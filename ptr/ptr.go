package ptr

// Deref returns values under the pointer or zero value in case of nil
func Deref[T any](v *T) T {
	if v == nil {
		var empty T
		return empty
	}

	return *v
}
