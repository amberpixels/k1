package ptr_test

import (
	"testing"

	"github.com/amberpixels/k1/ptr"
	"github.com/expectto/be"
)

func TestDeref(t *testing.T) {
	t.Run("returns the value behind a non-nil pointer", func(t *testing.T) {
		n := 42
		be.Expect(t, ptr.Deref(&n)).To(be.Eq(42))

		s := "hello"
		be.Expect(t, ptr.Deref(&s)).To(be.Eq("hello"))
	})

	t.Run("returns the zero value for a nil pointer", func(t *testing.T) {
		var pi *int
		be.Expect(t, ptr.Deref(pi)).To(be.Eq(0))

		var ps *string
		be.Expect(t, ptr.Deref(ps)).To(be.Eq(""))

		var pb *bool
		be.Expect(t, ptr.Deref(pb)).To(be.Eq(false))
	})
}

func TestClone(t *testing.T) {
	t.Run("returns a new pointer with the same value", func(t *testing.T) {
		n := 42
		c := ptr.Clone(&n)

		// same value...
		be.Expect(t, ptr.Deref(c)).To(be.Eq(42))
		// ...but a distinct pointer (mutating the clone must not touch the original)
		*c = 100
		be.Expect(t, n).To(be.Eq(42))
		be.Expect(t, ptr.Deref(c)).To(be.Eq(100))
	})

	t.Run("nil in, nil out", func(t *testing.T) {
		var p *int
		be.Expect(t, ptr.Clone(p)).To(be.Nil())
	})
}
