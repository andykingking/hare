package hare

import (
	"errors"
)

type Sequencer struct {
	Record `capid:"skip"`
	Index  Key `capid:"0"`
}

func (seq *Sequencer) BucketName() string {
	return "seq"
}

func (seq *Sequencer) Key() []byte {
	return []byte{0, 0, 0, 0, 0, 0, 0, 1}
}

func (seq *Sequencer) SetKey(sId string) error {
	return errors.New("Can't set key for sequence")
}

func (seq *Sequencer) NextKey() Key {
	seq.Index.kInt++
	return seq.Index
}
