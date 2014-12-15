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
		It("has a default key value of 0", func() {
			Expect(seq.Index).To(Equal(NewKey(0)))
		})
	})

	Describe("NextKey()", func() {
		It("increments and returns Index", func() {
			seq.Index = NewKey(8)
			next := NewKey(9)
			res := seq.NextKey()
			Expect(res).To(Equal(next))
			Expect(seq.Index).To(Equal(next))
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
			seq.Index = NewKey(42000)
			seqCopy := &Sequencer{}

			var buf bytes.Buffer
			seq.Save(&buf)
			seqCopy.Load(&buf)

			Expect(seqCopy.Index).To(Equal(seq.Index))
		})
	})

})
