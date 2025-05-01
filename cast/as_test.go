package abucast_test

import (
	"encoding/json"

	abucast "github.com/amberpixels/abu/cast"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("As", func() {
	Context("AsString", func() {
		It("should return string for string", func() {
			Expect(abucast.AsString("something")).To(Equal("something"))
		})

		It("should return string for empty string", func() {
			Expect(abucast.AsString("")).To(Equal(""))
		})

		It("should return string for []byte", func() {
			Expect(abucast.AsString([]byte("foobar"))).To(Equal("foobar"))
		})

		It("should return string for empty []byte", func() {
			Expect(abucast.AsString([]byte{})).To(Equal(""))
		})

		It("should return string for CustomString", func() {
			type CustomString string
			Expect(abucast.AsString(CustomString("foobar"))).To(Equal("foobar"))
		})

		It("should return string for json.RawMessage", func() {
			Expect(abucast.AsString(json.RawMessage(`{"foo":"bar"}`))).To(Equal(`{"foo":"bar"}`))
		})

		It("should return string for *json.RawMessage", func() {
			msg := json.RawMessage(`{"foo":"bar"}`)
			Expect(abucast.AsString(&msg)).To(Equal(`{"foo":"bar"}`))
		})

		It("should return string for a string under the pointer", func() {
			Expect(abucast.AsString(new(string))).To(Equal(""))
		})

		It("should panic for non-stringish", func() {
			Expect(func() { abucast.AsString(123) }).To(Panic())
		})
	})
})
