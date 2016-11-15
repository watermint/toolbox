package util

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

// size: length of the string
func GenerateRandomString(size int) (string, error) {
	if size < 1 {
		return "", errors.New(fmt.Sprintf("Size must greater than 1, given size was %d", size))
	}
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	encoded := base64.URLEncoding.EncodeToString(bytes)
	return encoded[:size], nil
}
