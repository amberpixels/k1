package ptr

// Deref returns values under the pointer or zero value in case of nil
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
