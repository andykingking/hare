package hare

import (
	"io"
)

type Recorded interface {
	BucketName() string

	Key() []byte
	SetKey(key string) error

	Saved() bool
	SetSaved(saved bool)

	Load(r io.Reader)
	Save(w io.Writer)
}

type Record struct {
	K     uint64 `capid:"skip", json:"key"`
	saved bool
}

func (record *Record) BucketName() string {
	return "hare"
}

func (record *Record) Key() []byte {
	return KeyToBytes(record.K)
}

func (record *Record) SetKey(key string) error {
	k, err := StringToKey(key)
	if err != nil {
		return err
	}

	record.K = k
	return nil
}

func (record *Record) Saved() bool {
	return record.saved
}

func (record *Record) SetSaved(saved bool) {
	record.saved = saved
}
