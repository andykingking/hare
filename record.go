package hare

import (
	"encoding/binary"
	"io"
	"strconv"
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
	bKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bKey, record.K)
	return bKey
}

func (record *Record) SetKey(key string) (err error) {
	var k uint64

	k, err = strconv.ParseUint(key, 10, 64)
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
