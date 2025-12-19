package errs_test

import (
	"errors"
	"fmt"

	"github.com/amberpixels/k1/errs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UnwrapDeep", func() {
	It("should return nil for nil error", func() {
		result := errs.UnwrapDeep(nil)
		Expect(result).To(BeNil())
	})

	It("should return the same error for non-wrapped error", func() {
		rootErr := errors.New("root error")
		result := errs.UnwrapDeep(rootErr)
		Expect(result).To(Equal(rootErr))
	})

	It("should unwrap a single-level wrapped error", func() {
		rootErr := errors.New("root error")
		wrappedErr := fmt.Errorf("wrapped: %w", rootErr)

		result := errs.UnwrapDeep(wrappedErr)
		Expect(result).To(Equal(rootErr))
	})

	It("should unwrap a multi-level wrapped error chain", func() {
		rootErr := errors.New("database connection failed")
		err2 := fmt.Errorf("query failed: %w", rootErr)
		err3 := fmt.Errorf("operation failed: %w", err2)
		err4 := fmt.Errorf("handler failed: %w", err3)

		result := errs.UnwrapDeep(err4)
		Expect(result).To(Equal(rootErr))
	})

	It("should handle deeply nested error chains", func() {
		rootErr := errors.New("root cause")
		currentErr := rootErr

		// Create a chain of 10 wrapped errors
		for i := 1; i <= 10; i++ {
			currentErr = fmt.Errorf("level %d: %w", i, currentErr)
		}

		result := errs.UnwrapDeep(currentErr)
		Expect(result).To(Equal(rootErr))
	})
})
