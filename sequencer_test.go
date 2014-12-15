package hare

import (
	"bytes"
	"fmt"
	cv "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSequencerBucketName(t *testing.T) {
	seq := &Sequencer{}
	expectedText := "seq"

	cv.Convey("Given a sequencer", t, func() {
		cv.Convey(fmt.Sprintf("It should have the BucketName '%s'", expectedText), func() {
			cv.So(seq.BucketName(), cv.ShouldEqual, expectedText)
		})
	})
}

func TestSequencerKey(t *testing.T) {
	seq := &Sequencer{}
	expectedKey := []byte{0, 0, 0, 0, 0, 0, 0, 1}

	cv.Convey("Given a sequencer", t, func() {
		cv.Convey(fmt.Sprintf("It should have the Key '%v'", expectedKey), func() {
			cv.So(seq.Key(), cv.ShouldResemble, expectedKey)
		})
	})
}

func TestSequencerNext(t *testing.T) {
	seq := &Sequencer{}

	cv.Convey("Given a sequencer", t, func() {
		cv.Convey("It should start with the Index 0", func() {
			cv.So(seq.Index, cv.ShouldEqual, 0)
		})
		cv.Convey("It should increment the Index with Next", func() {
			for i := 1; i <= 3; i++ {
				seq.Next()
				cv.So(seq.Index, cv.ShouldEqual, i)
			}
		})
	})
}

func TestSequencerSetKey(t *testing.T) {
	seq := &Sequencer{}

	cv.Convey("Given a sequencer", t, func() {
		cv.Convey("It should return an error when setting the key", func() {
			var err error = seq.SetKey("123")
			cv.So(err, cv.ShouldNotBeEmpty)
		})
	})
}

func TestSequencerSerialisation(t *testing.T) {
	var buf bytes.Buffer
	seq1 := &Sequencer{Index: 42}
	seq2 := &Sequencer{}

	cv.Convey("Given a sequencer", t, func() {
		cv.Convey("It should serialise to/from a buffer", func() {
			seq1.Save(&buf)
			seq2.Load(&buf)
			cv.So(seq2.Index, cv.ShouldEqual, 42)
		})
	})
}
