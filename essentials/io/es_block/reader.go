package es_block

import "io"

type Reader interface {
	// Read a block. isEof will be true when it's the last block of the source.
	// EOF will not return error.
	ReadBlock() (block []byte, isEof bool, err error)
}

func NewReader(r io.Reader, blockSize int) Reader {
	return &readerImpl{
		r:         r,
		blockSize: blockSize,
	}
}

type readerImpl struct {
	r         io.Reader
	blockSize int
}

func (z *readerImpl) ReadBlock() (block []byte, eof bool, err error) {
	dataOffset := 0
	dataRead := 0
	block = make([]byte, z.blockSize)

	for dataRead < z.blockSize {
		readBuf := make([]byte, z.blockSize-dataOffset)
		readSize, rdErr := z.r.Read(readBuf)
		dataRead += readSize
		switch rdErr {
		case nil:
			copy(block[dataOffset:], readBuf[:readSize])
			dataOffset += readSize

		case io.EOF:
			copy(block[dataOffset:], readBuf[:readSize])
			return block[:dataRead], true, nil

		default:
			return nil, false, rdErr
		}
	}

	return block[:dataRead], false, nil
}
