package es_file_random

import (
	"io"
	"math/rand"
	"time"
)

func NewReaderWithSeed(seed int64, size uint64) io.ReadCloser {
	return &readerImpl{
		rand:      rand.New(rand.NewSource(seed)),
		size:      size,
		readBytes: 0,
	}
}

// Create new reader with pseudo random seed
func NewReader(size uint64) io.ReadCloser {
	return NewReaderWithSeed(time.Now().Unix(), size)
}

type readerImpl struct {
	rand      *rand.Rand
	size      uint64
	readBytes uint64
}

func (z *readerImpl) Close() error {
	return nil
}

func (z *readerImpl) Read(p []byte) (n int, err error) {
	bytesToRead := z.size - z.readBytes
	if bytesToRead <= 0 {
		return 0, io.EOF
	}

	n, err = z.rand.Read(p)
	newReadBytes := z.readBytes + uint64(n)
	z.readBytes += uint64(n)
	switch {
	case newReadBytes < z.size:
		return n, nil
	case newReadBytes == z.size:
		return n, io.EOF
	default:
		return int(bytesToRead), io.EOF
	}
}
