package cast_test

import (
	"encoding/json"
	"testing"

	cast "github.com/amberpixels/k1/cast"
	"github.com/expectto/be"
)

func TestIsStringStrict(t *testing.T) {
	// strict mode (no opts) accepts only actual strings
	be.Expect(t, cast.IsString("something")).To(be.Eq(true))
	be.Expect(t, cast.IsString("")).To(be.Eq(true))
	be.Expect(t, cast.IsString([]byte("foobar"))).To(be.Eq(false))
	be.Expect(t, cast.IsString([]byte{})).To(be.Eq(false))

	hw := "hello world"
	be.Expect(t, cast.IsString(&hw)).To(be.Eq(false))

	// nil is never a string
	be.Expect(t, cast.IsString(nil)).To(be.Eq(false))

	// explicit Strict() option
	be.Expect(t, cast.IsString("x", cast.Strict())).To(be.Eq(true))
	be.Expect(t, cast.IsString([]byte("x"), cast.Strict())).To(be.Eq(false))
}

func TestIsStringAllowAll(t *testing.T) {
	be.Expect(t, cast.IsString(customString("hello world"), cast.AllowAll())).To(be.Eq(true))

	hw := customString("hello world")
	be.Expect(t, cast.IsString(&hw, cast.AllowAll())).To(be.Eq(true))

	ptr := &hw
	be.Expect(t, cast.IsString(&ptr, cast.AllowAll())).To(be.Eq(true))

	be.Expect(t, cast.IsString([]byte("hello"), cast.AllowAll())).To(be.Eq(true))

	cb := customBytes("hello")
	be.Expect(t, cast.IsString(cb, cast.AllowAll())).To(be.Eq(true))
	be.Expect(t, cast.IsString(&cb, cast.AllowAll())).To(be.Eq(true))
	cbp := &cb
	be.Expect(t, cast.IsString(&cbp, cast.AllowAll())).To(be.Eq(true))

	// a genuinely non-stringish value is still rejected under AllowAll
	be.Expect(t, cast.IsString(123, cast.AllowAll())).To(be.Eq(false))
}

func TestIsStringAllowPointers(t *testing.T) {
	be.Expect(t, cast.IsString(new(string), cast.AllowPointers())).To(be.Eq(true))
	be.Expect(t, cast.IsString(new(string), cast.AllowDeepPointers())).To(be.Eq(true))

	s := "hello"
	be.Expect(t, cast.IsString(&s, cast.AllowPointers())).To(be.Eq(true))
	be.Expect(t, cast.IsString(&s, cast.AllowDeepPointers())).To(be.Eq(true))

	// double pointer: only deep pointers reach the underlying string
	ss := &s
	be.Expect(t, cast.IsString(&ss, cast.AllowDeepPointers())).To(be.Eq(true))
	be.Expect(t, cast.IsString(&ss, cast.AllowPointers())).To(be.Eq(false))

	// not a string under the pointer
	be.Expect(t, cast.IsString(new(int), cast.AllowPointers())).To(be.Eq(false))
}

func TestIsStringAllowBytesConversion(t *testing.T) {
	// bytes / json.RawMessage accepted via the type-cast fast path
	be.Expect(t, cast.IsString([]byte("x"), cast.AllowBytesConversion())).To(be.Eq(true))
	be.Expect(t, cast.IsString(json.RawMessage(`x`), cast.AllowBytesConversion())).To(be.Eq(true))

	// pointer to bytes only when pointers are also allowed
	b := []byte("x")
	be.Expect(t, cast.IsString(&b, cast.AllowBytesConversion())).To(be.Eq(false))
	be.Expect(t, cast.IsString(&b, cast.AllowBytesConversion(), cast.AllowPointers())).To(be.Eq(true))

	rm := json.RawMessage(`x`)
	be.Expect(t, cast.IsString(&rm, cast.AllowBytesConversion(), cast.AllowPointers())).To(be.Eq(true))
}

func TestIsStringAllowCustomTypes(t *testing.T) {
	// custom string type accepted only with AllowCustomTypes
	be.Expect(t, cast.IsString(customString("x"))).To(be.Eq(false))
	be.Expect(t, cast.IsString(customString("x"), cast.AllowCustomTypes())).To(be.Eq(true))

	// custom []byte requires both custom-types and bytes-conversion
	be.Expect(t, cast.IsString(customBytes("x"), cast.AllowCustomTypes())).To(be.Eq(false))
	be.Expect(t,
		cast.IsString(customBytes("x"), cast.AllowCustomTypes(), cast.AllowBytesConversion()),
	).To(be.Eq(true))
}

func TestConfigureIsStringConfig(t *testing.T) {
	// Change the global default so that IsString accepts custom string types
	// without per-call options, then restore strict mode.
	cast.ConfigureIsStringConfig(cast.AllowCustomTypes())
	t.Cleanup(func() { cast.ConfigureIsStringConfig() })

	be.Expect(t, cast.IsString(customString("x"))).To(be.Eq(true))

	cast.ConfigureIsStringConfig() // back to strict
	be.Expect(t, cast.IsString(customString("x"))).To(be.Eq(false))
}
