package cast_test

import (
	"testing"
	"time"

	cast "github.com/amberpixels/k1/cast"
	"github.com/expectto/be"
)

func TestIsNil(t *testing.T) {
	be.Expect(t, cast.IsNil(nil)).To(be.Eq(true))

	var ip *int
	be.Expect(t, cast.IsNil(ip)).To(be.Eq(true))

	var iface any
	be.Expect(t, cast.IsNil(iface)).To(be.Eq(true))

	var m map[string]int
	be.Expect(t, cast.IsNil(m)).To(be.Eq(true))

	var s []int
	be.Expect(t, cast.IsNil(s)).To(be.Eq(true))

	var fn func()
	be.Expect(t, cast.IsNil(fn)).To(be.Eq(true))

	var ch chan int
	be.Expect(t, cast.IsNil(ch)).To(be.Eq(true))

	// non-nil cases
	be.Expect(t, cast.IsNil(&struct{}{})).To(be.Eq(false))
	be.Expect(t, cast.IsNil(map[string]int{})).To(be.Eq(false))
	be.Expect(t, cast.IsNil(func() {})).To(be.Eq(false))
	be.Expect(t, cast.IsNil(0)).To(be.Eq(false))
	be.Expect(t, cast.IsNil("")).To(be.Eq(false))
}

func TestIsStringish(t *testing.T) {
	// stringish values
	be.Expect(t, cast.IsStringish("something")).To(be.Eq(true))
	be.Expect(t, cast.IsStringish("")).To(be.Eq(true))
	be.Expect(t, cast.IsStringish([]byte("foobar"))).To(be.Eq(true))
	be.Expect(t, cast.IsStringish([]byte{})).To(be.Eq(true))
	be.Expect(t, cast.IsStringish(customString("x"))).To(be.Eq(true))

	// not stringish
	be.Expect(t, cast.IsStringish(nil)).To(be.Eq(false))
	be.Expect(t, cast.IsStringish(123)).To(be.Eq(false))
	be.Expect(t, cast.IsStringish(123.456)).To(be.Eq(false))
	be.Expect(t, cast.IsStringish(true)).To(be.Eq(false))
	be.Expect(t, cast.IsStringish(1+2i)).To(be.Eq(false))
	be.Expect(t, cast.IsStringish(struct{}{})).To(be.Eq(false))
	be.Expect(t, cast.IsStringish(map[string]int{})).To(be.Eq(false))
	be.Expect(t, cast.IsStringish(func() {})).To(be.Eq(false))
}

func TestIsInt(t *testing.T) {
	be.Expect(t, cast.IsInt(123)).To(be.Eq(true))
	be.Expect(t, cast.IsInt(int8(1))).To(be.Eq(true))
	be.Expect(t, cast.IsInt(int16(1))).To(be.Eq(true))
	be.Expect(t, cast.IsInt(int32(1))).To(be.Eq(true))
	be.Expect(t, cast.IsInt(int64(1))).To(be.Eq(true))
	be.Expect(t, cast.IsInt(uint(1))).To(be.Eq(true))
	// integral floats are considered ints
	be.Expect(t, cast.IsInt(42.0)).To(be.Eq(true))

	// not ints
	be.Expect(t, cast.IsInt(123.456)).To(be.Eq(false))
	be.Expect(t, cast.IsInt("123")).To(be.Eq(false))
	be.Expect(t, cast.IsInt(true)).To(be.Eq(false))
	be.Expect(t, cast.IsInt(nil)).To(be.Eq(false))
	be.Expect(t, cast.IsInt(struct{}{})).To(be.Eq(false))
}

func TestIsStrings(t *testing.T) {
	be.Expect(t, cast.IsStrings([]string{"a", "b"})).To(be.True())
	be.Expect(t, cast.IsStrings([]customString{"a"})).To(be.True())

	be.Expect(t, cast.IsStrings("not a slice")).To(be.False())
	be.Expect(t, cast.IsStrings(nil)).To(be.False())

	// []int is rune-convertible to string but is NOT a slice of strings.
	be.Expect(t, cast.IsStrings([]int{1, 2})).To(be.False())
}

func TestIsTime(t *testing.T) {
	be.Expect(t, cast.IsTime(time.Now())).To(be.Eq(true))

	n := time.Now()
	be.Expect(t, cast.IsTime(&n)).To(be.Eq(true))

	be.Expect(t, cast.IsTime("not a time")).To(be.Eq(false))
	be.Expect(t, cast.IsTime(123)).To(be.Eq(false))
}
