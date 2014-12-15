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

	Describe("BucketName()", func() {
		It("should be 'hare' by default", func() {
			Expect(record.BucketName()).To(Equal("hare"))
		})
	})

	Context("K", func() {

		var (
			sk string
			ik uint64
			bk []byte
		)

		BeforeEach(func() {
			sk, ik, bk = "12345678", 12345678, []byte{0, 0, 0, 0, 0, 188, 97, 78}
		})

		Describe("Key()", func() {
			It("is the 8-byte, big-endian representation of the 64-bit unsigned integer internal key", func() {
				record.K = ik
				Expect(record.Key()).To(Equal(bk))
			})
		})

		Describe("SetKey()", func() {
			It("sets the internal key to the 64-bit unsigned integer represented numerically by the passed string", func() {
				record.SetKey(sk)
				Expect(record.K).To(Equal(ik))
			})
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
