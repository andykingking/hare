package hare_test

import (
	. "github.com/captainpete/hare"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Key", func() {

	Describe("StringToKey()", func() {
		It("sets the internal key to the 64-bit unsigned integer represented numerically by the passed string", func() {
			var kk uint64 = 12345
			Expect(StringToKey("012345")).To(Equal(kk))
		})
	})

	Describe("KeyToBytes()", func() {
		It("is the 8-byte, big-endian representation of the 64-bit unsigned integer internal key", func() {
			var kk uint64 = 12345678
			bk := []byte{0, 0, 0, 0, 0, 188, 97, 78}

			Expect(KeyToBytes(kk)).To(Equal(bk))
		})
	})

})
