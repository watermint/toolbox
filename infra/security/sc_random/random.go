package sc_random

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

func MustGenerateRandomString(size int) string {
	r, err := GenerateRandomString(size)
	if err != nil {
		panic(err)
	}
	return r
}

// size: length of the string
func GenerateRandomString(size int) (string, error) {
	if size < 1 {
		return "", errors.New(fmt.Sprintf("Size must greater than 1, given size was %d", size))
	}
	seq := make([]byte, size)
	_, err := rand.Read(seq)
	if err != nil {
		return "", err
	}
	encoded := base64.RawStdEncoding.EncodeToString(seq)
	return encoded[:size], nil
}
