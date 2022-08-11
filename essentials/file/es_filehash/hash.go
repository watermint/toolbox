package es_filehash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"hash"
	"io"
	"os"
)

var (
	ErrorHashNotFound = errors.New("hash not found")
)

type Hash interface {
	MD5(filepath string) (digest string, err error)
	SHA256(filepath string) (digest string, err error)
	SHA1(filepath string) (digest string, err error)
}

func NewHash(l esl.Logger) Hash {
	return &hashImpl{l: l}
}

type hashImpl struct {
	l esl.Logger
}

func (z hashImpl) sum(filepath string, algorithm hash.Hash, sumLength int) (digest string, err error) {
	l := z.l.With(esl.String("file", filepath))
	f, err := os.Open(filepath)
	if err != nil {
		l.Debug("Unable to open file", esl.Error(err))
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(algorithm, f); err != nil {
		l.Debug("Unable to calculate or read file", esl.Error(err))
		return "", err
	}
	dh := algorithm.Sum(nil)[:sumLength]
	return hex.EncodeToString(dh), nil
}

func (z hashImpl) MD5(filepath string) (digest string, err error) {
	return z.sum(filepath, md5.New(), 16)
}

func (z hashImpl) SHA1(filepath string) (digest string, err error) {
	return z.sum(filepath, sha1.New(), 20)
}

func (z hashImpl) SHA256(filepath string) (digest string, err error) {
	return z.sum(filepath, sha256.New(), 32)
}
