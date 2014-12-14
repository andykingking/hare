package hare

import (
	"encoding/binary"
	"io"
)

type Key struct {
	Id    uint64 `capid:"0"`
	saved bool
}

type Keyed interface {
	IdBytes() []byte

	Load(r io.Reader)
	Save(w io.Writer)

	SetSaved(saved bool)
}

func (key *Key) IdBytes() []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, key.Id)
	return b
}

func (key *Key) SetSaved(saved bool) {
	key.saved = saved
}

func (key *Key) SetId(sId string) uint64 {
	key.Id = binary.LittleEndian.Uint64([]byte(sId))
	return key.Id
}
