package hare

import (
	"errors"
)

type Sequencer struct {
	Record `capid:"skip"`

	// TODO: serialisation of uint64 is broken.
	// Update cap'n'proto, change to uint64 in future.
	Index int64 `capid:"0"`
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
	seq.Index++
	return Key{uint64(seq.Index)}
}
