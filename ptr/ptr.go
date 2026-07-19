package ptr

// Deref is the safe equivalent of * operator.
// It returns values under the pointer or zero value in case of nil.
func Deref[T any](v *T) T {
	if v == nil {
		var empty T
		return empty
	}

	return *v
}

// Clone returns a shallow copy of the value behind p.
// Nil in, nil out.
func Clone[T any](p *T) *T {
	if p == nil {
		return nil
	}
	c := *p
	return &c
}

// Equal reports whether two pointers hold equal values.
// Two nils are equal; a nil never equals a non-nil.
func Equal[T comparable](a, b *T) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}
