package quick

import "github.com/amberpixels/k1/set"

// Append appends elements of b to the slice a.
// Append DOESN'T make NEW duplicates. (But it respects the old duplicates).
//
// Note: quick.Append is ~300x faster than append but the cost is ~3x more memory to be used.
// Also, keep it mind, quick.Append can load GC much more than simple append.
func Append[T comparable](a []T, b ...T) []T {
	m := len(b)
	if m == 0 {
		return a
	}

	// N stands for the number of elements in A. We don't care if duplicates are there. We respect them.
	n := len(a)

	// seen stands for a unique set of elements in A (will be used for lookup when appending Bs)
	seen := set.NewLookupCapped(n+m, a...)

	// 2) Allocate the result slice: len(a) + at most len(b)
	res := make([]T, n, n+m)
	copy(res, a)

	// 3) As we scan b, append only if not already in seen—and then mark seen
	for _, e := range b {
		if !seen.Has(e) {
			seen[e] = struct{}{}
			res = append(res, e)
		}
	}

	return res
}
