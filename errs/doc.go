// Package errs provides error handling utilities that extend Go's standard "errors" library.
//
// It offers:
//   - UnwrapDeep: recursively unwraps error chains until the root/deepest error is reached.
//
// Usage:
//
//	import "github.com/amberpixels/k1/errs"
//
//	// Get the root cause of a wrapped error chain
//	rootErr := errs.UnwrapDeep(wrappedError)
//
//	// Example with error wrapping:
//	err1 := errors.New("root cause")
//	err2 := fmt.Errorf("layer 2: %w", err1)
//	err3 := fmt.Errorf("layer 3: %w", err2)
//	root := errs.UnwrapDeep(err3) // returns err1
//
// Package errs is intended as a lightweight helper for deep error chain traversal.
package errs
