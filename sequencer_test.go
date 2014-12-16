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
			Expect(seq.BucketName()).To(Equal([]byte("sequencers")))
		})
	})

	Describe("Key()", func() {
		It("should be a constant", func() {
			seqKey := []byte{0, 0, 0, 0, 0, 0, 0, 1}
			Expect(seq.Key()).To(Equal(seqKey))
		})
	})

	Describe("Mark", func() {
		It("has a default key value of 0", func() {
			Expect(seq.Mark).To(Equal(uint64(0)))
		})
	})

	Describe("NextKey()", func() {
		It("increments and returns Mark", func() {
			seq.Mark = 8
			res := seq.NextKey()
			Expect(res).To(Equal(uint64(9)))
			Expect(seq.Mark).To(Equal(uint64(9)))
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
			seq.Mark = 42000
			seqCopy := &Sequencer{}

			var buf bytes.Buffer
			seq.Save(&buf)
			seqCopy.Load(&buf)

			Expect(seqCopy.Mark).To(Equal(seq.Mark))
		})
	})

})
