package es_block

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
)

type BlockReader interface {
	// FileBlocks Returns offsets of the file.
	FileBlocks(path string) (offsets []int64, err error)

	// ReadBlock Returns block data.
	ReadBlock(path string, offset int64) (data []byte, isLastBlock bool, err error)
}

func NewPlainReader(l esl.Logger, blockSize int) BlockReader {
	return &bfsImpl{
		l:         l,
		blockSize: blockSize,
	}
}

func fileBlocks(l esl.Logger, path string, blockSize int) (offsets []int64, err error) {
	l = l.With(esl.String("path", path))
	info, err := os.Lstat(path)
	if err != nil {
		l.Debug("Unable to retrieve file info", esl.Error(err))
		return nil, err
	}

	offsets = make([]int64, 0)
	var offset int64 = 0
	for offset < info.Size() {
		offsets = append(offsets, offset)
		offset += int64(blockSize)
	}
	l.Debug("File blocks", esl.Int("numBlocks", len(offsets)))
	return offsets, nil
}

type bfsImpl struct {
	l         esl.Logger
	blockSize int
}

func (z bfsImpl) FileBlocks(path string) (offsets []int64, err error) {
	return fileBlocks(z.l, path, z.blockSize)
}

func (z bfsImpl) ReadBlock(path string, offset int64) (data []byte, isLastBlock bool, err error) {
	l := z.l.With(esl.String("path", path), esl.Int64("offset", offset))
	info, err := os.Lstat(path)
	if err != nil {
		l.Debug("Unable to retrieve file info", esl.Error(err))
		return nil, false, err
	}

	f, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open the file", esl.Error(err))
		return nil, false, err
	}

	if _, err = f.Seek(offset, io.SeekStart); err != nil {
		l.Debug("Unable to seek to the position", esl.Error(err))
		return nil, false, err
	}

	data, isLastBlock, err = NewReader(f, z.blockSize).ReadBlock()

	// adjust last block
	if info.Size()-offset <= int64(z.blockSize) {
		isLastBlock = true
	}
	l.Debug("Read block", esl.Int("readSize", len(data)), esl.Bool("isLastBlock", isLastBlock), esl.Error(err))
	return
}
