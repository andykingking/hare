package hare

import (
	"encoding/binary"
	"errors"
	"strconv"
)

// TODO: serialisation of uint64 is broken.
// Update cap'n'proto, change to uint64 in future.
type kInt uint64

type Key struct {
	kInt `capid:"0"`
}

func NewKey(n kInt) Key {
	return Key{n}
}

// Converts a string containing digits
// to the equivalent numeric key.
func StringToKey(key string) (Key, error) {
	n, err := strconv.ParseInt(key, 10, 64)
	if n < 0 {
		return NewKey(0), errors.New("Key must be positive")
	}
	if err != nil {
		return NewKey(0), err
	}
	return NewKey(kInt(n)), nil
}

// Returns the 8-byte big-endian encoding
// from the unsigned, 64-bit integer key value.
func (key Key) Bytes() []byte {
	bKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bKey, uint64(key.kInt))
	return bKey
}
