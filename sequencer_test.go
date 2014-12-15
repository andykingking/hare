package hare_test

import (
	. "github.com/captainpete/hare"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
)

var _ = Describe("Sequencer", func() {

	var (
		seq *Sequencer
	)

	BeforeEach(func() {
		seq = &Sequencer{}
	})

	Describe("BucketName()", func() {
		It("should be 'seq'", func() {
			Expect(seq.BucketName()).To(Equal("seq"))
		})
	})

	Describe("Key()", func() {
		It("should be a constant", func() {
			seqKey := []byte{0, 0, 0, 0, 0, 0, 0, 1}
			Expect(seq.Key()).To(Equal(seqKey))
		})
	})

	Describe("Index", func() {
		It("has a default value of 0", func() {
			var e int64 = 0
			Expect(seq.Index).To(Equal(e))
		})
	})

	Describe("Next()", func() {
		It("increments and returns Index", func() {
			seq.Index = 8
			n := seq.Next()
			Expect(n).To(Equal(uint64(9)))
			Expect(seq.Index).To(Equal(int64(9)))
		})
	})

	Describe("SetKey()", func() {
		It("returns an error", func() {
			var err error = seq.SetKey("123")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).NotTo(BeEmpty())
		})
	})

	Describe("serialisation", func() {
		It("saves and loads to a buffer", func() {
			seq.Index = 42000
			seqCopy := &Sequencer{}

			var buf bytes.Buffer
			seq.Save(&buf)
			seqCopy.Load(&buf)

			Expect(seqCopy.Index).To(Equal(seq.Index))
		})
	})

})
