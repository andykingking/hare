package hare_test

import (
	. "github.com/captainpete/hare"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Record", func() {

	var (
		record *Record
	)

	BeforeEach(func() {
		record = &Record{}
	})

	Describe("Key()", func() {
		It("is the keys bytes", func() {
			key := NewKey(12345678)
			record.K = key
			Expect(record.Key()).To(Equal(key.Bytes()))
		})
	})

	Describe("SetKey()", func() {
		It("sets the key using a string", func() {
			sKey := "12345678"
			record.SetKey(sKey)
			k, _ := StringToKey(sKey)
			Expect(record.K).To(Equal(k))
		})
	})

	Describe("BucketName()", func() {
		It("should be 'hare' by default", func() {
			Expect(record.BucketName()).To(Equal("hare"))
		})
	})

	Context("saved", func() {

		Describe("Saved()", func() {
			It("returns false by default", func() {
				Expect(record.Saved()).To(BeFalse())
			})
		})

		Describe("SetSaved()", func() {
			It("sets the value of saved subsequently returned by Saved()", func() {
				record.SetSaved(true)
				Expect(record.Saved()).To(BeTrue())
				record.SetSaved(false)
				Expect(record.Saved()).To(BeFalse())
			})
		})

	})

})
