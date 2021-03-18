package sc_random

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	mathrand "math/rand"
)

func MustGetSecureRandomString(size int) string {
	r, err := GetSecureRandomString(size)
	if err != nil {
		panic(err)
	}
	return r
}

// size: length of the string
func GetSecureRandomString(size int) (string, error) {
	if size < 1 {
		return "", errors.New(fmt.Sprintf("Size must greater than 1, given size was %d", size))
	}
	seq := make([]byte, size)
	_, err := rand.Read(seq)
	if err != nil {
		return "", err
	}
	encoded := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(seq)
	return encoded[:size], nil
}

func MustGetPseudoRandomString(r *mathrand.Rand, size int) string {
	if size < 1 {
		panic("Size must grater than 1")
	}
	seq := make([]byte, size)
	_, err := r.Read(seq)
	if err != nil {
		return ""
	}
	encoded := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(seq)
	return encoded[:size]
}
