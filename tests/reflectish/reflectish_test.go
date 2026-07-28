package reflectish_test

import (
	"reflect"
	"testing"

	"github.com/amberpixels/k1/reflectish"
	"github.com/expectto/be"
	"github.com/expectto/be/be_reflected"
)

func TestIndirectDeep(t *testing.T) {
	t.Run("returns the same value for a non-pointer value", func(t *testing.T) {
		str := "test"
		result := reflectish.IndirectDeep(reflect.ValueOf(str))

		be.Expect(t, result.Kind()).To(be.Eq(reflect.String))
		// The indirected value is a string under reflection.
		be.Expect(t, result.Interface()).To(be_reflected.AsString())
		be.Expect(t, result.String()).To(be.Eq("test"))
	})

	t.Run("dereferences a single pointer", func(t *testing.T) {
		str := "test"
		strPtr := &str
		result := reflectish.IndirectDeep(reflect.ValueOf(strPtr))

		be.Expect(t, result.Kind()).To(be.Eq(reflect.String))
		be.Expect(t, result.String()).To(be.Eq("test"))
	})

	t.Run("dereferences a double pointer", func(t *testing.T) {
		str := "test"
		strPtr := &str
		strPtrPtr := &strPtr
		result := reflectish.IndirectDeep(reflect.ValueOf(strPtrPtr))

		be.Expect(t, result.Kind()).To(be.Eq(reflect.String))
		be.Expect(t, result.String()).To(be.Eq("test"))
	})
}

func TestInterface(t *testing.T) {
	t.Run("returns the underlying value for a valid value", func(t *testing.T) {
		be.Expect(t, reflectish.Interface(reflect.ValueOf("test"))).To(be.Eq("test"))
		be.Expect(t, reflectish.Interface(reflect.ValueOf("test"))).To(be_reflected.AsString())

		be.Expect(t, reflectish.Interface(reflect.ValueOf(42))).To(be.Eq(42))
		be.Expect(t, reflectish.Interface(reflect.ValueOf(42))).To(be_reflected.AsInteger())
	})

	t.Run("returns nil for the invalid (zero) value", func(t *testing.T) {
		be.Expect(t, reflectish.Interface(reflect.Value{})).To(be.Nil())
	})

	t.Run("returns nil when composed with IndirectDeep over a nil pointer", func(t *testing.T) {
		var p *int
		be.Expect(t, reflectish.Interface(reflectish.IndirectDeep(reflect.ValueOf(p)))).To(be.Nil())
	})

	t.Run("returns the pointed-at value when composed with IndirectDeep", func(t *testing.T) {
		n := 7
		p := &n
		be.Expect(t, reflectish.Interface(reflectish.IndirectDeep(reflect.ValueOf(p)))).To(be.Eq(7))
	})
}

func TestIndirectInterface(t *testing.T) {
	t.Run("returns the value as-is for a non-pointer", func(t *testing.T) {
		be.Expect(t, reflectish.IndirectInterface(reflect.ValueOf("test"))).To(be.Eq("test"))
	})

	t.Run("dereferences through nested pointers", func(t *testing.T) {
		n := 7
		p := &n
		pp := &p
		be.Expect(t, reflectish.IndirectInterface(reflect.ValueOf(pp))).To(be.Eq(7))
		be.Expect(t, reflectish.IndirectInterface(reflect.ValueOf(pp))).To(be_reflected.AsInteger())
	})

	t.Run("returns nil for a nil pointer", func(t *testing.T) {
		var p *int
		be.Expect(t, reflectish.IndirectInterface(reflect.ValueOf(p))).To(be.Nil())
	})
}

func TestLengthOf(t *testing.T) {
	t.Run("returns the correct length for a string", func(t *testing.T) {
		length, ok := reflectish.LengthOf("test")

		be.Expect(t, ok).To(be.True())
		be.Expect(t, length).To(be.Eq(4))
	})

	t.Run("returns the correct length for a slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		length, ok := reflectish.LengthOf(slice)

		be.Expect(t, ok).To(be.True())
		be.Expect(t, length).To(be.Eq(3))
		// Cross-check via the native HaveLength matcher on the slice itself.
		be.Expect(t, slice).To(be.HaveLength(3))
	})

	t.Run("returns the correct length for a map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		length, ok := reflectish.LengthOf(m)

		be.Expect(t, ok).To(be.True())
		be.Expect(t, length).To(be.Eq(2))
		be.Expect(t, m).To(be.HaveLength(2))
	})

	t.Run("returns the correct length for a channel", func(t *testing.T) {
		ch := make(chan int, 5)
		length, ok := reflectish.LengthOf(ch)

		be.Expect(t, ok).To(be.True())
		be.Expect(t, length).To(be.Eq(0))
	})

	t.Run("returns false for an unsupported type", func(t *testing.T) {
		length, ok := reflectish.LengthOf(42)

		be.Expect(t, ok).To(be.False())
		be.Expect(t, length).To(be.Eq(0))
	})

	t.Run("returns false for nil", func(t *testing.T) {
		length, ok := reflectish.LengthOf(nil)

		be.Expect(t, ok).To(be.False())
		be.Expect(t, length).To(be.Eq(0))
	})
}
