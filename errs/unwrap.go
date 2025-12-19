package errs

import "errors"

// UnwrapDeep recursively unwraps an error chain until it reaches the deepest/root error.
// It returns the root cause of the error chain, or nil if the input is nil.
//
// Example:
//
//	err1 := errors.New("database connection failed")
//	err2 := fmt.Errorf("query failed: %w", err1)
//	err3 := fmt.Errorf("operation failed: %w", err2)
//	root := errs.UnwrapDeep(err3) // returns err1
func UnwrapDeep(err error) error {
	if err == nil {
		return nil
	}

	for {
		child := errors.Unwrap(err)
		if child == nil {
			return err
		}
		err = child
	}
}
