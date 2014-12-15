package hare

import (
	"encoding/binary"
	"strconv"
)

type Key struct {
	uint64 `capid:"0"`
}

func NewKey(n uint64) Key {
	return Key{n}
}

// Converts a string containing digits
// to the equivalent numeric key.
func StringToKey(key string) (Key, error) {
	n, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return NewKey(0), err
	}
	return NewKey(n), nil
}

// Returns the 8-byte big-endian encoding
// from the unsigned, 64-bit integer key value.
func (key Key) Bytes() []byte {
	bKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bKey, key.uint64)
	return bKey
}
