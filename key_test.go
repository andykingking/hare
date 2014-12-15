package hare_test

import (
	. "github.com/captainpete/hare"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Key", func() {

	BeforeEach(func() {
	})

	Describe("StringToKey()", func() {
		It("sets the internal key to the 64-bit unsigned integer represented numerically by the passed string", func() {
			kk := NewKey(12345)
			Expect(StringToKey("012345")).To(Equal(kk))
		})
	})

	Describe("Bytes()", func() {
		It("is the 8-byte, big-endian representation of the 64-bit unsigned integer internal key", func() {
			kk := NewKey(12345678)
			bk := []byte{0, 0, 0, 0, 0, 188, 97, 78}

			Expect(kk.Bytes()).To(Equal(bk))
		})
	})

})
