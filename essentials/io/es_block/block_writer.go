package es_block

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
	"sync"
)

var (
	ErrorInvalidDataBlockSize = errors.New("invalid data block size")
)

// BlockWriterFactory Creates the stateful BlockWriter instance for write data in blocks.
// BlockWriter requests subsequent blocks if the instance can handle multiple blocks.
type BlockWriterFactory interface {
	// Open a file with the block size.
	Open(path string, fileSize int64, req BlockRequester, fin BlockFinisher, be BlockError) BlockWriter

	// Wait for finish all writers
	Wait()
}

type BlockWriter interface {
	WriteBlock(data []byte, offset int64)

	// Abort abort operation and clean up resources, then send error report
	Abort(offset int64, err error)

	// Path file path to write
	Path() string

	// Wait for finish
	Wait()
}

type BlockRequester func(w BlockWriter, offset, blockSize int64)
type BlockFinisher func(w BlockWriter, size int64)
type BlockError func(w BlockWriter, offset int64, err error)

func NewWriterFactory(l esl.Logger, batchSize int, blockSize int) BlockWriterFactory {
	if batchSize < 1 || blockSize < 1 {
		l.Error("Invalid batchSize or blockSize",
			esl.Int("batchSize", batchSize),
			esl.Int("blockSize", blockSize))
		panic("Invalid batchSize or blockSize")
	}
	return &bwfImpl{
		l:         l,
		mutex:     sync.Mutex{},
		batchSize: batchSize,
		blockSize: blockSize,
		writers:   make(map[string]BlockWriter, 0),
		wg:        sync.WaitGroup{},
	}
}

type bwfImpl struct {
	l         esl.Logger
	mutex     sync.Mutex
	batchSize int
	blockSize int
	writers   map[string]BlockWriter
	wg        sync.WaitGroup
}

func (z *bwfImpl) Wait() {
	z.wg.Wait()

	for _, w := range z.writers {
		w.Wait()
	}
}

func (z *bwfImpl) releaseWriter(w BlockWriter) {
	l := z.l
	l.Debug("Release writer", esl.String("path", w.Path()))
	z.mutex.Lock()
	if _, ok := z.writers[w.Path()]; ok {
		z.wg.Done()
		delete(z.writers, w.Path())
	}
	z.mutex.Unlock()
}

func (z *bwfImpl) Open(path string, fileSize int64, req BlockRequester, fin BlockFinisher, be BlockError) BlockWriter {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	z.wg.Add(1)
	bw := &bwImpl{
		l:            z.l,
		batchSize:    z.batchSize,
		path:         path,
		blockSize:    z.blockSize,
		blockCache:   make(map[int64][]byte),
		posWritten:   0,
		posRequested: 0,
		fileSize:     fileSize,
		lastErr:      nil,
		mutex:        sync.Mutex{},
		handleReq:    req,
		handleFin: func(w BlockWriter, size int64) {
			z.releaseWriter(w)
			fin(w, size)
		},
		handleErr: func(w BlockWriter, offset int64, err error) {
			z.releaseWriter(w)
			be(w, offset, err)
		},
		wg: sync.WaitGroup{},
	}
	bw.start()
	z.writers[path] = bw

	return bw
}

type bwImpl struct {
	l            esl.Logger
	batchSize    int
	path         string
	blockSize    int
	blockCache   map[int64][]byte
	posWritten   int64
	posRequested int64
	fileSize     int64
	lastErr      error
	mutex        sync.Mutex
	handleReq    BlockRequester
	handleFin    BlockFinisher
	handleErr    BlockError
	wg           sync.WaitGroup
}

func (z *bwImpl) Abort(offset int64, err error) {
	z.noLockReportError(offset, err)
}

func (z *bwImpl) Wait() {
	z.wg.Wait()
}

func (z *bwImpl) Path() string {
	return z.path
}

func (z *bwImpl) start() {
	l := z.l.With(esl.String("path", z.path), esl.Int64("fileSize", z.fileSize))
	z.mutex.Lock()
	defer z.mutex.Unlock()

	z.wg.Add(1)
	f, err := os.Create(z.path)
	if err != nil {
		l.Debug("Unable to create the file", esl.Error(err))
		z.noLockReportError(0, err)
		z.wg.Done()
		return
	}
	_ = f.Close()
	z.noLockRequestBlocks()
}

func (z *bwImpl) noLockReportError(offset int64, err error) {
	l := z.l.With(esl.String("path", z.path), esl.Int64("fileSize", z.fileSize))
	rmErr := os.Remove(z.path)
	l.Debug("Removed the file", esl.Error(rmErr))
	z.lastErr = err
	z.handleErr(z, offset, err)
	z.wg.Done()
}

func (z *bwImpl) noLockReportFinish() {
	z.handleFin(z, z.fileSize) // notify to the caller
	z.wg.Done()
}

func (z *bwImpl) noLockRequestBlocks() {
	l := z.l.With(esl.String("path", z.path), esl.Int64("fileSize", z.fileSize))

	reqBlocks := max(0, z.batchSize-len(z.blockCache))
	l.Debug("Request blocks", esl.Int("reqBlocks", reqBlocks))

	for i := 0; i < reqBlocks; i++ {
		if z.fileSize < z.posRequested {
			l.Debug("Request reached to the end of the file")
			break
		}
		var reqBlockSize = int64(z.blockSize)

		if z.fileSize < reqBlockSize+z.posRequested {
			reqBlockSize = z.fileSize - z.posRequested
		}

		l.Debug("Request block", esl.Int64("reqOffset", z.posRequested), esl.Int64("reqBlockSize", reqBlockSize))
		go z.handleReq(z, z.posRequested, reqBlockSize)

		z.posRequested += int64(z.blockSize)
	}
}

func (z *bwImpl) noLockFlushBlock(data []byte, offset int64) (isEOF bool, err error) {
	l := z.l.With(esl.String("path", z.path), esl.Int64("fileSize", z.fileSize), esl.Int64("offset", offset))
	newPosition := int64(z.blockSize) + z.posWritten
	if z.fileSize < newPosition {
		isEOF = true
		newPosition = z.fileSize
		l.Debug("The final block")
	} else if len(data) != z.blockSize {
		l.Debug("Invalid data size", esl.Int("dataSize", len(data)))
		return false, ErrorInvalidDataBlockSize
	}

	l.Debug("Open file for append")
	f, err := os.OpenFile(z.path, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		l.Debug("Unable to open the file", esl.Error(err))
		return false, err
	}
	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write(data)
	if err != nil {
		l.Debug("Unable to write the block", esl.Error(err))
		return false, err
	}

	l.Debug("Increment position", esl.Int64("oldPos", z.posWritten), esl.Int64("newPos", newPosition))
	z.posWritten = newPosition

	if nextBloc, ok := z.blockCache[z.posWritten]; ok {
		l.Debug("Write next block", esl.Int64("nextBlock", z.posWritten))
		delete(z.blockCache, z.posWritten)
		return z.noLockFlushBlock(nextBloc, z.posWritten)
	}

	return isEOF, nil
}

func (z *bwImpl) WriteBlock(data []byte, offset int64) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.l.With(esl.String("path", z.path), esl.Int64("offset", offset), esl.Int64("fileSize", z.fileSize))

	if z.lastErr != nil {
		l.Debug("Skip write due to the last error")
		return
	}

	if offset != z.posWritten {
		l.Debug("sparse block found, cache it")
		z.blockCache[offset] = data
		return
	}

	isEOF, fbErr := z.noLockFlushBlock(data, offset)
	if fbErr != nil {
		// shutdown
		z.lastErr = fbErr
		z.noLockReportError(offset, fbErr) // notify to the caller
		return
	}

	if isEOF {
		l.Debug("Reached to EOF")
		z.noLockReportFinish()
		return
	}

	// request other blocks
	z.noLockRequestBlocks()
}
