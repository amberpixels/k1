package reflectish_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestReflect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reflect Suite")
}
