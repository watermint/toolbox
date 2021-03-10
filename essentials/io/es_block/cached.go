package es_block

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"sync"
)

type cacheImpl struct {
	l                esl.Logger
	blockSize        int
	cacheSize        int
	cacheMutex       sync.Mutex
	cacheData        map[string][]byte // cache key: offset-path => data
	cacheIsLastBlock map[string]bool
	cacheHit         int64
	cacheMiss        int64
}

func (z *cacheImpl) FileBlocks(path string) (offsets []int64, err error) {
	return fileBlocks(z.l, path, z.blockSize)
}

//func (z *cacheImpl) ReadBlock(path string, offset int64) (data []byte, isLastBlock bool, err error) {
//	l := z.l.With(esl.String("path", path), esl.Int64("offset", offset))
//	cacheKey := fmt.Sprintf("%x-%s", offset, path)
//	z.cacheMutex.Lock()
//	defer z.cacheMutex.Unlock()
//
//	if d, ok := z.cacheData[cacheKey]; ok {
//		isLastBlock = z.cacheIsLastBlock[cacheKey]
//		z.cacheHit++
//		delete(z.cacheData, cacheKey)
//		delete(z.cacheIsLastBlock, cacheKey)
//		return d, isLastBlock, nil
//	}
//
//	cacheCapacity := z.cacheSize - len(z.cacheData)
//
//}
