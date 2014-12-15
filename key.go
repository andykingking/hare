package hare

import (
	"encoding/binary"
	"errors"
	"strconv"
)

// Converts a string containing digits
// to the equivalent numeric key.
func StringToKey(key string) (uint64, error) {
	n, err := strconv.ParseUint(key, 10, 64)
	if n < 0 {
		return 0, errors.New("Key must be positive")
	}
	if err != nil {
		return 0, err
	}
	return n, nil
}

// Returns the 8-byte big-endian encoding
// from the unsigned, 64-bit integer key value.
func KeyToBytes(key uint64) []byte {
	bKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bKey, uint64(key))
	return bKey
}
