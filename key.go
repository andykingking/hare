package hare

import (
	"github.com/spaolacci/murmur3"
)

type Key struct {
	Digest uint64
	Index  uint64
}

func NewKey(index uint64) *Key {
	var key = &Key{Index: index}
	var h64 = murmur3.New64()
	h64.Write(uint64Bytes(key.Index))
	key.Digest = h64.Sum64()
	return key
}

func uint64Bytes(v uint64) []byte {
	// little-endian encode uint64
	var b = make([]byte, 8)
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	return b
}
