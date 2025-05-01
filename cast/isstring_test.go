package abucast_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	abucast "github.com/amberpixels/abu/cast"
)

var _ = Describe("IsString", func() {
	When("in strict mode", func() {
		It("should return true for string", func() {
			Expect(abucast.IsString("something")).To(BeTrue())
		})
		It("should return true for empty string", func() {
			Expect(abucast.IsString("")).To(BeTrue())
		})
		It("should return false for []byte", func() {
			Expect(abucast.IsString([]byte("foobar"))).To(BeFalse())
		})
		It("should return false for empty []byte", func() {
			Expect(abucast.IsString([]byte{})).To(BeFalse())
		})
		It("should return false for *string", func() {
			helloWorld := "hello world"
			Expect(abucast.IsString(&helloWorld)).To(BeFalse())
		})
	})

	When("allowing all", func() {
		type CustomString string
		type CustomBytes []byte
		It("should return true for custom strings", func() {
			Expect(abucast.IsString(CustomString("hello world"), abucast.AllowAll())).To(BeTrue())
		})
		It("should return true for custom strings under pointer", func() {
			helloWorld := CustomString("hello world")
			ptrHelloWorld := &helloWorld
			Expect(abucast.IsString(&helloWorld, abucast.AllowAll())).To(BeTrue())
			Expect(abucast.IsString(&ptrHelloWorld, abucast.AllowAll())).To(BeTrue())
		})
		It("should return true for bytes", func() {
			Expect(abucast.IsString([]byte("hello world"), abucast.AllowAll())).To(BeTrue())
		})
		It("should return true for custom bytes", func() {
			helloWorld := CustomBytes("hello world")
			Expect(abucast.IsString(helloWorld, abucast.AllowAll())).To(BeTrue())
			Expect(abucast.IsString(&helloWorld, abucast.AllowAll())).To(BeTrue())

			ptrHelloWorld := &helloWorld
			Expect(abucast.IsString(&ptrHelloWorld, abucast.AllowAll())).To(BeTrue())
		})
	})

	When("allowing pointers", func() {
		It("should return true for string under the pointer", func() {
			Expect(abucast.IsString(new(string), abucast.AllowPointers())).To(BeTrue())
			Expect(abucast.IsString(new(string), abucast.AllowDeepPointers())).To(BeTrue())

			s := "hello"
			Expect(abucast.IsString(&s, abucast.AllowPointers())).To(BeTrue())
			Expect(abucast.IsString(&s, abucast.AllowDeepPointers())).To(BeTrue())
			ss := &s
			Expect(abucast.IsString(&ss, abucast.AllowDeepPointers())).To(BeTrue())
			Expect(abucast.IsString(&ss, abucast.AllowPointers())).To(BeFalse())
		})

		It("should return false for not-a-string under the pointer", func() {
			Expect(abucast.IsString(new(int), abucast.AllowPointers())).To(BeFalse())
		})
	})

})
