package cast_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	cast "github.com/amberpixels/k1/cast"
	"github.com/expectto/be"
	"github.com/expectto/be/be_math"
	"github.com/expectto/be/be_reflected"
)

type (
	customString string
	customBytes  []byte
	customInt    int
	customFloat  float64
	customBool   bool
	customTime   time.Time
)

func TestAsString(t *testing.T) {
	be.Expect(t, cast.AsString("something")).To(be.Eq("something"))
	be.Expect(t, cast.AsString("")).To(be.Eq(""))
	be.Expect(t, cast.AsString([]byte("foobar"))).To(be.Eq("foobar"))
	be.Expect(t, cast.AsString([]byte{})).To(be.Eq(""))
	be.Expect(t, cast.AsString(json.RawMessage(`{"foo":"bar"}`))).To(be.Eq(`{"foo":"bar"}`))

	msg := json.RawMessage(`{"foo":"bar"}`)
	be.Expect(t, cast.AsString(&msg)).To(be.Eq(`{"foo":"bar"}`))

	// custom string type goes through the reflect fallback
	be.Expect(t, cast.AsString(customString("foobar"))).To(be.Eq("foobar"))
	// custom []byte type goes through the reflect fallback
	be.Expect(t, cast.AsString(customBytes("byte_data"))).To(be.Eq("byte_data"))
	// string under a pointer
	be.Expect(t, cast.AsString(new(string))).To(be.Eq(""))

	// the result is genuinely a string
	be.Expect(t, cast.AsString("hello")).To(be_reflected.AsString())
}

func TestAsStringPanics(t *testing.T) {
	be.Expect(t, func() { cast.AsString(123) }).To(be.Panic())
}

func TestAsBytes(t *testing.T) {
	be.Expect(t, cast.AsBytes([]byte("byte_data"))).To(be.Eq([]byte("byte_data")))
	be.Expect(t, cast.AsBytes("example")).To(be.Eq([]byte("example")))
	be.Expect(t, cast.AsBytes(json.RawMessage(`x`))).To(be.Eq([]byte("x")))

	msg := json.RawMessage(`y`)
	be.Expect(t, cast.AsBytes(&msg)).To(be.Eq([]byte("y")))

	// custom string type via reflect fallback
	be.Expect(t, cast.AsBytes(customString("foobar"))).To(be.Eq([]byte("foobar")))
	// custom []byte type via reflect fallback
	be.Expect(t, cast.AsBytes(customBytes("foobar"))).To(be.Eq([]byte("foobar")))

	be.Expect(t, cast.AsBytes("data")).To(be_reflected.AsBytes())
}

func TestAsBytesPanics(t *testing.T) {
	be.Expect(t, func() { cast.AsBytes(123) }).To(be.Panic())
}

func TestAsBool(t *testing.T) {
	be.Expect(t, cast.AsBool(true)).To(be.Eq(true))
	be.Expect(t, cast.AsBool(false)).To(be.Eq(false))

	b := true
	be.Expect(t, cast.AsBool(&b)).To(be.Eq(true))

	// custom bool via reflect fallback
	be.Expect(t, cast.AsBool(customBool(true))).To(be.Eq(true))
}

func TestAsBoolPanics(t *testing.T) {
	be.Expect(t, func() { cast.AsBool("nope") }).To(be.Panic())
}

func TestAsInt(t *testing.T) {
	be.Expect(t, cast.AsInt(42)).To(be.Eq(42))
	be.Expect(t, cast.AsInt(int8(8))).To(be.Eq(8))
	be.Expect(t, cast.AsInt(int16(16))).To(be.Eq(16))
	be.Expect(t, cast.AsInt(int32(32))).To(be.Eq(32))
	be.Expect(t, cast.AsInt(int64(64))).To(be.Eq(64))
	be.Expect(t, cast.AsInt(uint(1))).To(be.Eq(1))
	be.Expect(t, cast.AsInt(uint8(2))).To(be.Eq(2))
	be.Expect(t, cast.AsInt(uint16(3))).To(be.Eq(3))
	be.Expect(t, cast.AsInt(uint32(4))).To(be.Eq(4))
	be.Expect(t, cast.AsInt(uint64(5))).To(be.Eq(5))

	// integral floats convert cleanly
	be.Expect(t, cast.AsInt(42.0)).To(be.Eq(42))
	be.Expect(t, cast.AsInt(float32(7.0))).To(be.Eq(7))

	// the result is a real int and matches math matchers
	be.Expect(t, cast.AsInt(50)).To(be.All(be_reflected.AsInteger(), be_math.GreaterThan(10)))
}

func TestAsIntPointers(t *testing.T) {
	i := 42
	be.Expect(t, cast.AsInt(&i)).To(be.Eq(42))
	i8 := int8(8)
	be.Expect(t, cast.AsInt(&i8)).To(be.Eq(8))
	i16 := int16(16)
	be.Expect(t, cast.AsInt(&i16)).To(be.Eq(16))
	i32 := int32(32)
	be.Expect(t, cast.AsInt(&i32)).To(be.Eq(32))
	i64 := int64(64)
	be.Expect(t, cast.AsInt(&i64)).To(be.Eq(64))
	u := uint(1)
	be.Expect(t, cast.AsInt(&u)).To(be.Eq(1))
	u8 := uint8(2)
	be.Expect(t, cast.AsInt(&u8)).To(be.Eq(2))
	u16 := uint16(3)
	be.Expect(t, cast.AsInt(&u16)).To(be.Eq(3))
	u32 := uint32(4)
	be.Expect(t, cast.AsInt(&u32)).To(be.Eq(4))
	u64 := uint64(5)
	be.Expect(t, cast.AsInt(&u64)).To(be.Eq(5))
	f64 := 9.0
	be.Expect(t, cast.AsInt(&f64)).To(be.Eq(9))
	f32 := float32(11.0)
	be.Expect(t, cast.AsInt(&f32)).To(be.Eq(11))
}

func TestAsIntReflectFallback(t *testing.T) {
	// custom int / uint / float types go through the reflect fallback
	be.Expect(t, cast.AsInt(customInt(123))).To(be.Eq(123))
	be.Expect(t, cast.AsInt(customFloat(5.0))).To(be.Eq(5))

	type customUint uint
	be.Expect(t, cast.AsInt(customUint(7))).To(be.Eq(7))
}

func TestAsIntPanics(t *testing.T) {
	// non-integral float64
	be.Expect(t, func() { cast.AsInt(42.5) }).To(be.Panic())
	// non-integral *float64
	f := 1.5
	be.Expect(t, func() { cast.AsInt(&f) }).To(be.Panic())
	// non-integral float32
	be.Expect(t, func() { cast.AsInt(float32(1.5)) }).To(be.Panic())
	f32 := float32(2.5)
	be.Expect(t, func() { cast.AsInt(&f32) }).To(be.Panic())
	// non-integral custom float via reflect fallback
	be.Expect(t, func() { cast.AsInt(customFloat(3.5)) }).To(be.Panic())
	// not a number at all
	be.Expect(t, func() { cast.AsInt("nope") }).To(be.Panic())
}

func TestAsFloat(t *testing.T) {
	be.Expect(t, cast.AsFloat(3.14)).To(be.Eq(3.14))
	be.Expect(t, cast.AsFloat(float32(2.5))).To(be.Eq(2.5))
	be.Expect(t, cast.AsFloat(42)).To(be.Eq(42.0))
	be.Expect(t, cast.AsFloat(int8(8))).To(be.Eq(8.0))
	be.Expect(t, cast.AsFloat(int16(16))).To(be.Eq(16.0))
	be.Expect(t, cast.AsFloat(int32(32))).To(be.Eq(32.0))
	be.Expect(t, cast.AsFloat(int64(64))).To(be.Eq(64.0))
	be.Expect(t, cast.AsFloat(uint(1))).To(be.Eq(1.0))
	be.Expect(t, cast.AsFloat(uint8(2))).To(be.Eq(2.0))
	be.Expect(t, cast.AsFloat(uint16(3))).To(be.Eq(3.0))
	be.Expect(t, cast.AsFloat(uint32(4))).To(be.Eq(4.0))
	be.Expect(t, cast.AsFloat(uint64(5))).To(be.Eq(5.0))

	be.Expect(t, cast.AsFloat(3.14)).To(be_reflected.AsFloat())
}

func TestAsFloatPointers(t *testing.T) {
	f64 := 3.14
	be.Expect(t, cast.AsFloat(&f64)).To(be.Eq(3.14))
	f32 := float32(2.5)
	be.Expect(t, cast.AsFloat(&f32)).To(be.Eq(2.5))
	i := 42
	be.Expect(t, cast.AsFloat(&i)).To(be.Eq(42.0))
	i8 := int8(8)
	be.Expect(t, cast.AsFloat(&i8)).To(be.Eq(8.0))
	i16 := int16(16)
	be.Expect(t, cast.AsFloat(&i16)).To(be.Eq(16.0))
	i32 := int32(32)
	be.Expect(t, cast.AsFloat(&i32)).To(be.Eq(32.0))
	i64 := int64(64)
	be.Expect(t, cast.AsFloat(&i64)).To(be.Eq(64.0))
	u := uint(1)
	be.Expect(t, cast.AsFloat(&u)).To(be.Eq(1.0))
	u8 := uint8(2)
	be.Expect(t, cast.AsFloat(&u8)).To(be.Eq(2.0))
	u16 := uint16(3)
	be.Expect(t, cast.AsFloat(&u16)).To(be.Eq(3.0))
	u32 := uint32(4)
	be.Expect(t, cast.AsFloat(&u32)).To(be.Eq(4.0))
	u64 := uint64(5)
	be.Expect(t, cast.AsFloat(&u64)).To(be.Eq(5.0))
}

func TestAsFloatReflectFallback(t *testing.T) {
	be.Expect(t, cast.AsFloat(customFloat(1.5))).To(be.Eq(1.5))
	be.Expect(t, cast.AsFloat(customInt(3))).To(be.Eq(3.0))

	type customUint uint
	be.Expect(t, cast.AsFloat(customUint(7))).To(be.Eq(7.0))
}

func TestAsFloatPanics(t *testing.T) {
	be.Expect(t, func() { cast.AsFloat("nope") }).To(be.Panic())
}

func TestAsKind(t *testing.T) {
	be.Expect(t, cast.AsKind(reflect.Int)).To(be.Eq(reflect.Int))

	k := reflect.String
	be.Expect(t, cast.AsKind(&k)).To(be.Eq(reflect.String))
}

func TestAsKindPanics(t *testing.T) {
	be.Expect(t, func() { cast.AsKind("nope") }).To(be.Panic())
}

func TestAsSliceOfAny(t *testing.T) {
	be.Expect(t, cast.AsSliceOfAny([]any{1, "two", 3.0})).To(be.HaveLength(3))

	// non-[]any slice goes through the reflect fallback
	out := cast.AsSliceOfAny([]string{"a", "b", "c"})
	be.Expect(t, out).To(be.HaveLength(3))
	be.Expect(t, out).To(be.Eq([]any{"a", "b", "c"}))

	// pointer to slice
	s := []int{1, 2}
	be.Expect(t, cast.AsSliceOfAny(&s)).To(be.Eq([]any{1, 2}))
}

func TestAsSliceOfAnyPanics(t *testing.T) {
	be.Expect(t, func() { cast.AsSliceOfAny(123) }).To(be.Panic())
}

func TestAsStrings(t *testing.T) {
	be.Expect(t, cast.AsStrings([]string{"a", "b"})).To(be.Eq([]string{"a", "b"}))

	// custom string slice goes through the reflect fallback
	be.Expect(t, cast.AsStrings([]customString{"x", "y"})).To(be.Eq([]string{"x", "y"}))

	// pointer to slice
	s := []string{"p", "q"}
	be.Expect(t, cast.AsStrings(&s)).To(be.Eq([]string{"p", "q"}))

	be.Expect(t, cast.AsStrings([]string{"a"})).To(be_reflected.AsSliceOf[string]())
}

func TestAsStringsPanics(t *testing.T) {
	// not a slice at all
	be.Expect(t, func() { cast.AsStrings(123) }).To(be.Panic())
	// slice whose elements are not strings
	be.Expect(t, func() { cast.AsStrings([]struct{}{{}}) }).To(be.Panic())
	// []int is rune-convertible to string but is NOT a slice of strings.
	be.Expect(t, func() { cast.AsStrings([]int{1, 2}) }).To(be.Panic())
}

func TestAsTime(t *testing.T) {
	now := time.Now()
	be.Expect(t, cast.AsTime(now)).To(be.Eq(now))
	be.Expect(t, cast.AsTime(&now)).To(be.Eq(now))

	// A custom type whose underlying type is time.Time converts correctly.
	be.Expect(t, cast.AsTime(customTime(now))).To(be.Eq(now))
}

func TestAsTimePanics(t *testing.T) {
	be.Expect(t, func() { cast.AsTime("nope") }).To(be.Panic())
}
