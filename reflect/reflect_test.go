package abureflect_test

import (
	"reflect"

	abureflect "github.com/amberpixels/abu/reflect"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reflect", func() {
	Context("IndirectDeep", func() {
		It("should return the same value for non-pointer value", func() {
			str := "test"
			v := reflect.ValueOf(str)
			result := abureflect.IndirectDeep(v)

			Expect(result.Kind()).To(Equal(reflect.String))
			Expect(result.String()).To(Equal("test"))
		})

		It("should dereference a single pointer", func() {
			str := "test"
			strPtr := &str
			v := reflect.ValueOf(strPtr)
			result := abureflect.IndirectDeep(v)

			Expect(result.Kind()).To(Equal(reflect.String))
			Expect(result.String()).To(Equal("test"))
		})

		It("should dereference a double pointer", func() {
			str := "test"
			strPtr := &str
			strPtrPtr := &strPtr
			v := reflect.ValueOf(strPtrPtr)
			result := abureflect.IndirectDeep(v)

			Expect(result.Kind()).To(Equal(reflect.String))
			Expect(result.String()).To(Equal("test"))
		})
	})

	Context("LengthOf", func() {
		It("should return the correct length for a string", func() {
			str := "test"
			length, ok := abureflect.LengthOf(str)

			Expect(ok).To(BeTrue())
			Expect(length).To(Equal(4))
		})

		It("should return the correct length for a slice", func() {
			slice := []int{1, 2, 3}
			length, ok := abureflect.LengthOf(slice)

			Expect(ok).To(BeTrue())
			Expect(length).To(Equal(3))
		})

		It("should return the correct length for a map", func() {
			m := map[string]int{"a": 1, "b": 2}
			length, ok := abureflect.LengthOf(m)

			Expect(ok).To(BeTrue())
			Expect(length).To(Equal(2))
		})

		It("should return the correct length for a channel", func() {
			ch := make(chan int, 5)
			length, ok := abureflect.LengthOf(ch)

			Expect(ok).To(BeTrue())
			Expect(length).To(Equal(0))
		})

		It("should return false for an unsupported type", func() {
			i := 42
			length, ok := abureflect.LengthOf(i)

			Expect(ok).To(BeFalse())
			Expect(length).To(Equal(0))
		})

		It("should return false for nil", func() {
			length, ok := abureflect.LengthOf(nil)

			Expect(ok).To(BeFalse())
			Expect(length).To(Equal(0))
		})
	})
})
