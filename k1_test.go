package k1_test

import (
	"testing"

	"github.com/amberpixels/k1"
	"github.com/expectto/be"
	"github.com/expectto/be/be_string"
)

// word is a tiny fmt.Stringer for exercising JoinStringers.
type word string

func (w word) String() string { return string(w) }

func TestJoinStringers(t *testing.T) {
	t.Run("joins multiple stringers with the separator", func(t *testing.T) {
		be.Expect(t, k1.JoinStringers([]word{"a", "b", "c"}, ", ")).To(be.Eq("a, b, c"))
	})

	t.Run("single element has no separator", func(t *testing.T) {
		be.Expect(t, k1.JoinStringers([]word{"solo"}, ", ")).To(be.Eq("solo"))
	})

	t.Run("empty and nil slices return empty string", func(t *testing.T) {
		be.Expect(t, k1.JoinStringers([]word{}, ", ")).To(be.Eq(""))
		be.Expect(t, k1.JoinStringers([]word(nil), ", ")).To(be.Eq(""))
	})

	t.Run("composes with be_string matchers", func(t *testing.T) {
		be.Expect(t, k1.JoinStringers([]word{"foo", "bar"}, "-")).
			To(be_string.ContainingSubstring("foo-bar"))
	})
}
