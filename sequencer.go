package hare

import (
	"errors"
)

type Sequencer struct {
	Record `capid:"skip"`
	Mark   uint64 `capid:"0"`
}

func (seq *Sequencer) BucketName() []byte {
	return []byte("sequencers")
}

func (seq *Sequencer) Key() []byte {
	return []byte{0, 0, 0, 0, 0, 0, 0, 1}
}

func (seq *Sequencer) SetKey(sId string) error {
	return errors.New("Can't set key for sequence")
}

func (seq *Sequencer) NextKey() uint64 {
	seq.Mark++
	return seq.Mark
}
