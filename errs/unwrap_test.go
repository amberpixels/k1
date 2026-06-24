package errs_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/amberpixels/k1/errs"
	"github.com/expectto/be"
)

func TestUnwrapDeep_NilError(t *testing.T) {
	result := errs.UnwrapDeep(nil)

	be.Expect(t, result).To(be.Nil())
}

func TestUnwrapDeep_NonWrappedError(t *testing.T) {
	rootErr := errors.New("root error")

	result := errs.UnwrapDeep(rootErr)

	be.Expect(t, result).To(be.NotNil())
	be.Expect(t, result).To(be.Eq(rootErr))
}

func TestUnwrapDeep_SingleLevelWrappedError(t *testing.T) {
	rootErr := errors.New("root error")
	wrappedErr := fmt.Errorf("wrapped: %w", rootErr)

	result := errs.UnwrapDeep(wrappedErr)

	be.Expect(t, result).To(be.Eq(rootErr))
	be.Expect(t, result).To(be.MatchError(rootErr))
}

func TestUnwrapDeep_MultiLevelWrappedError(t *testing.T) {
	rootErr := errors.New("database connection failed")
	err2 := fmt.Errorf("query failed: %w", rootErr)
	err3 := fmt.Errorf("operation failed: %w", err2)
	err4 := fmt.Errorf("handler failed: %w", err3)

	result := errs.UnwrapDeep(err4)

	be.Expect(t, result).To(be.Eq(rootErr))
}

func TestUnwrapDeep_DeeplyNestedErrorChain(t *testing.T) {
	rootErr := errors.New("root cause")
	currentErr := rootErr

	// Create a chain of 10 wrapped errors.
	for i := 1; i <= 10; i++ {
		currentErr = fmt.Errorf("level %d: %w", i, currentErr)
	}

	result := errs.UnwrapDeep(currentErr)

	be.Expect(t, result).To(be.Eq(rootErr))
}
