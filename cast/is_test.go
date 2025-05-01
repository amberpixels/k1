package abucast_test

import (
	abucast "github.com/amberpixels/abu/cast"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Is", func() {
	Context("IsNil", func() {
		It("should return true for nil", func() {
			Expect(abucast.IsNil(nil)).To(BeTrue())
		})
		It("should return true for typed nil", func() {
			var i *int
			Expect(abucast.IsNil(i)).To(BeTrue())
		})
		It("should return true for interface nil", func() {
			var i interface{}
			Expect(abucast.IsNil(i)).To(BeTrue())
		})

		It("should return false for non-nil pointer", func() {
			Expect(abucast.IsNil(&struct{}{})).To(BeFalse())
		})
		It("should return false for non-nil map", func() {
			Expect(abucast.IsNil(map[string]int{})).To(BeFalse())
		})
		It("should return false for non-nil func", func() {
			Expect(abucast.IsNil(func() {})).To(BeFalse())
		})

		It("should return false for non-nil digit", func() {
			Expect(abucast.IsNil(0)).To(BeFalse())
		})
		It("should return false for non-nil string", func() {
			Expect(abucast.IsNil("")).To(BeFalse())
		})
	})

	Context("IsStringish", func() {
		When("considered stringish", func() {
			It("should return true for string", func() {
				Expect(abucast.IsStringish("something")).To(BeTrue())
			})
			It("should return true for empty string", func() {
				Expect(abucast.IsStringish("")).To(BeTrue())
			})
			It("should return true for []byte", func() {
				Expect(abucast.IsStringish([]byte("foobar"))).To(BeTrue())
			})
			It("should return true for empty []byte", func() {
				Expect(abucast.IsStringish([]byte{})).To(BeTrue())
			})
		})

		When("considered not stringish", func() {
			It("should return false for nil", func() {
				Expect(abucast.IsStringish(nil)).To(BeFalse())
			})
			It("should return false for int", func() {
				Expect(abucast.IsStringish(123)).To(BeFalse())
			})
			It("should return false for float", func() {
				Expect(abucast.IsStringish(123.456)).To(BeFalse())
			})
			It("should return false for bool", func() {
				Expect(abucast.IsStringish(true)).To(BeFalse())
			})
			It("should return false for complex", func() {
				Expect(abucast.IsStringish(1 + 2i)).To(BeFalse())
			})
			It("should return false for struct", func() {
				Expect(abucast.IsStringish(struct{}{})).To(BeFalse())
			})
			It("should return false for map", func() {
				Expect(abucast.IsStringish(map[string]int{})).To(BeFalse())
			})
			It("should return false for func", func() {
				Expect(abucast.IsStringish(func() {})).To(BeFalse())
			})
		})
	})

	Context("IsInt", func() {
		It("should return true for int", func() {
			Expect(abucast.IsInt(123)).To(BeTrue())
		})
		It("should return true for int8", func() {
			Expect(abucast.IsInt(int8(123))).To(BeTrue())
		})
		It("should return true for int16", func() {
			Expect(abucast.IsInt(int16(123))).To(BeTrue())
		})
		It("should return true for int32", func() {
			Expect(abucast.IsInt(int32(123))).To(BeTrue())
		})
		It("should return true for int64", func() {
			Expect(abucast.IsInt(int64(123))).To(BeTrue())
		})
		It("should return true for uint", func() {
			Expect(abucast.IsInt(uint(123))).To(BeTrue())
		})

		It("should return false for float", func() {
			Expect(abucast.IsInt(123.456)).To(BeFalse())
		})
		It("should return false for string", func() {
			Expect(abucast.IsInt("123")).To(BeFalse())
		})
		It("should return false for bool", func() {
			Expect(abucast.IsInt(true)).To(BeFalse())
		})
		It("should return false for nil", func() {
			Expect(abucast.IsInt(nil)).To(BeFalse())
		})
		It("should return false for struct", func() {
			Expect(abucast.IsInt(struct{}{})).To(BeFalse())
		})
	})
})
