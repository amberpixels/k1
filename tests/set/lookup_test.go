package set_test

import (
	"testing"

	"github.com/amberpixels/k1/set"
	"github.com/expectto/be"
)

// TestNewLookupEmpty verifies an empty lookup reports no membership.
func TestNewLookupEmpty(t *testing.T) {
	l := set.NewLookup[string]()

	be.Expect(t, l).To(be.Not(be.HaveKey("a")))
	be.Expect(t, l).To(be.Not(be.HaveKey("b")))
	be.Expect(t, l).To(be.HaveLength(0))
}

// TestNewLookupAndFill verifies adding into an initially-empty lookup.
func TestNewLookupAndFill(t *testing.T) {
	l := set.NewLookup[string]()

	be.Expect(t, l).To(be.Not(be.HaveKey("a")))

	l.Add("a")

	be.Expect(t, l).To(be.HaveKey("a"))
	be.Expect(t, l).To(be.Not(be.HaveKey("b")))
	be.Expect(t, l).To(be.HaveLength(1))
}

// TestNewLookupAndHas tests that NewLookup correctly initializes a lookup and Has
// returns true for the initial keys.
func TestNewLookupAndHas(t *testing.T) {
	l := set.NewLookup("a", "b", "c")

	be.Expect(t, l).To(be.HaveLength(3))
	be.Expect(t, l).To(be.HaveKey("a"))
	be.Expect(t, l).To(be.HaveKey("b"))
	be.Expect(t, l).To(be.HaveKey("c"))
	be.Expect(t, l).To(be.Not(be.HaveKey("d")))
}

// TestHas exercises the Has method directly, asserting the returned bool with
// be.True/be.False.
func TestHas(t *testing.T) {
	l := set.NewLookup("a", "b", "c")

	be.Expect(t, l.Has("a")).To(be.True())
	be.Expect(t, l.Has("d")).To(be.False())
}

// TestNewLookupDeduplicates verifies that repeated initial keys collapse into one.
func TestNewLookupDeduplicates(t *testing.T) {
	l := set.NewLookup("a", "a", "b")

	be.Expect(t, l).To(be.HaveLength(2))
	be.Expect(t, l).To(be.HaveKey("a"))
	be.Expect(t, l).To(be.HaveKey("b"))
}

// TestAdd tests that Add inserts a new key and is idempotent.
func TestAdd(t *testing.T) {
	l := set.NewLookup("x")
	be.Require(t, l).To(be.HaveKey("x"))

	l.Add("y")
	be.Expect(t, l).To(be.HaveKey("y"))
	be.Expect(t, l).To(be.HaveLength(2))

	// Adding an existing key must not change the length.
	l.Add("y")
	be.Expect(t, l).To(be.HaveLength(2))
}

// TestDelete verifies Delete removes a key and is a no-op for missing keys.
func TestDelete(t *testing.T) {
	l := set.NewLookup("a", "b", "c")

	l.Delete("b")
	be.Expect(t, l).To(be.Not(be.HaveKey("b")))
	be.Expect(t, l).To(be.HaveKey("a"))
	be.Expect(t, l).To(be.HaveLength(2))

	// Deleting a non-existent key must be a safe no-op.
	l.Delete("zzz")
	be.Expect(t, l).To(be.HaveLength(2))
}

// TestClear verifies Clear removes every element.
func TestClear(t *testing.T) {
	l := set.NewLookup("a", "b", "c")
	be.Require(t, l).To(be.HaveLength(3))

	l.Clear()
	be.Expect(t, l).To(be.HaveLength(0))
	be.Expect(t, l).To(be.Not(be.HaveKey("a")))
	be.Expect(t, l).To(be.Not(be.HaveKey("b")))
	be.Expect(t, l).To(be.Not(be.HaveKey("c")))

	// A cleared lookup is still usable.
	l.Add("z")
	be.Expect(t, l).To(be.HaveKey("z"))
	be.Expect(t, l).To(be.HaveLength(1))
}

// TestNewLookupCapped verifies the capacity-hinted constructor behaves like NewLookup.
func TestNewLookupCapped(t *testing.T) {
	l := set.NewLookupCapped(10, "a", "b")

	be.Expect(t, l).To(be.HaveLength(2))
	be.Expect(t, l).To(be.HaveKey("a"))
	be.Expect(t, l).To(be.HaveKey("b"))
	be.Expect(t, l).To(be.Not(be.HaveKey("c")))

	l.Add("c")
	be.Expect(t, l).To(be.HaveKey("c"))
	be.Expect(t, l).To(be.HaveLength(3))
}

// TestNewLookupCappedZeroNoKeys verifies a zero-capacity capped lookup with no
// initial keys starts empty and remains usable.
func TestNewLookupCappedZeroNoKeys(t *testing.T) {
	l := set.NewLookupCapped[int](0)

	be.Expect(t, l).To(be.HaveLength(0))

	l.Add(42)
	be.Expect(t, l).To(be.HaveKey(42))
}

// TestLookupWithInts demonstrates the generic Lookup works for int types and the
// full Add/Has/Delete lifecycle.
func TestLookupWithInts(t *testing.T) {
	l := set.NewLookup(1, 2, 3)

	be.Expect(t, l).To(be.HaveLength(3))
	be.Expect(t, l).To(be.HaveKey(1))
	be.Expect(t, l).To(be.HaveKey(2))
	be.Expect(t, l).To(be.HaveKey(3))
	be.Expect(t, l).To(be.Not(be.HaveKey(4)))

	l.Add(4)
	be.Expect(t, l).To(be.HaveKey(4))

	l.Delete(1)
	be.Expect(t, l).To(be.Not(be.HaveKey(1)))
	be.Expect(t, l).To(be.HaveLength(3))
}
