package es_block

import "github.com/watermint/toolbox/essentials/log/esl"

type cacheImpl struct {
	l         esl.Logger
	blockSize int
}

func (z *cacheImpl) FileBlocks(path string) (offsets []int64, err error) {
	return fileBlocks(z.l, path, z.blockSize)
}

func (z *cacheImpl) ReadBlock(path string, offset int64) (data []byte, isLastBlock bool, err error) {
	panic("implement me")
}
