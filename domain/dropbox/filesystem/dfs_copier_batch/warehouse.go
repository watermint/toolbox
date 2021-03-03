package dfs_copier_batch

import (
	"errors"
	"github.com/watermint/toolbox/essentials/io/es_block"
	"github.com/watermint/toolbox/essentials/log/esl"
	"go.uber.org/atomic"
	"os"
	"sync"
	"time"
)

type BlockWarehouseReceipt int64

type BlockWarehouseCallback func(bw BlockWarehouse, path string, offset int64, isLastBlock bool, receipt BlockWarehouseReceipt)
type BlockWarehouseFailureCallback func(bw BlockWarehouse, path string, err error)

const (
	blockWarehouseReceiptZeroBlock BlockWarehouseReceipt = -1
)

type BlockWarehouse interface {
	// Startup warehouse worker
	Startup()

	// Shutdown warehouse worker
	Shutdown()

	// Load file as blocks, then callback the func r.
	// Callback f in case of the error during read data.
	// This func will not block the operation.
	Load(path string, r BlockWarehouseCallback, f BlockWarehouseFailureCallback)

	// Receive the block data with the receipt. Then, release data from the warehouse.
	Receive(receipt BlockWarehouseReceipt) ([]byte, error)

	// Check status for the path
	Status(path string) (found bool)
}

func NewBlockWarehouse(l esl.Logger, blockSize int, warehouseSize int) BlockWarehouse {
	return &bwImpl{
		l:                  l,
		backlogCount:       sync.WaitGroup{},
		backlogMutex:       sync.Mutex{},
		backlogPaths:       make([]string, 0),
		backlogCallbacks:   make(map[string]BlockWarehouseCallback),
		backlogFailure:     make(map[string]BlockWarehouseFailureCallback),
		blockMutex:         sync.Mutex{},
		block:              make(map[BlockWarehouseReceipt][]byte),
		blocKIds:           atomic.NewInt64(1),
		warehouseShutdown:  false,
		warehouseSize:      warehouseSize,
		warehouseBlockSize: blockSize,
	}
}

type bwImpl struct {
	l            esl.Logger
	backlogCount sync.WaitGroup

	backlogMutex     sync.Mutex
	backlogPaths     []string
	backlogCallbacks map[string]BlockWarehouseCallback        // path -> callback
	backlogFailure   map[string]BlockWarehouseFailureCallback // path -> failure callback

	blockMutex sync.Mutex
	block      map[BlockWarehouseReceipt][]byte
	blocKIds   *atomic.Int64

	warehouseShutdown  bool
	warehouseSize      int
	warehouseBlockSize int
	//warehouseBlockSem  semaphore.Weighted
}

func (z *bwImpl) Status(path string) (found bool) {
	z.backlogMutex.Lock()
	_, found = z.backlogCallbacks[path]
	z.backlogMutex.Unlock()
	return
}

func (z *bwImpl) Startup() {
	go z.blockLoader()
}

func (z *bwImpl) Shutdown() {
	z.backlogCount.Wait()
	//z.warehouseShutdown = true
}

// Store data & retrieve the receipt. This func will block the operation
func (z *bwImpl) getReceipt(data []byte) BlockWarehouseReceipt {
	l := z.l
	if len(data) < 1 {
		l.Debug("Return zero block receipt")
		return blockWarehouseReceiptZeroBlock
	}

	for {
		z.blockMutex.Lock()
		if x := len(z.block); x < z.warehouseSize {
			var receipt = BlockWarehouseReceipt(z.blocKIds.Inc())
			z.block[receipt] = data
			z.blockMutex.Unlock()
			l.Debug("Warehouse has a room to store", esl.Int("currentWarehouseSize", x), esl.Int64("receipt", int64(receipt)))
			return receipt
		}
		z.blockMutex.Unlock()

		time.Sleep(100 * time.Millisecond)
	}
}

func (z *bwImpl) loadBlock(path string, callback BlockWarehouseCallback, failure BlockWarehouseFailureCallback) {
	l := z.l.With(esl.String("path", path))
	defer func() {
		l.Debug("Release backlog")
		z.backlogCount.Done()
	}()

	fi, osErr := os.Lstat(path)
	if osErr != nil {
		l.Debug("Unable to retrieve file info", esl.Error(osErr))
		failure(z, path, osErr)
		return
	}

	// immediately finish if file size == 0
	if fi.Size() < 1 {
		l.Debug("Finis zero size file")
		callback(z, path, 0, true, blockWarehouseReceiptZeroBlock)
		return
	}

	f, osErr := os.Open(path)
	if osErr != nil {
		l.Debug("Unable to open file", esl.Error(osErr))
		failure(z, path, osErr)
		return
	}
	defer func() {
		_ = f.Close()
	}()

	var offset int64 = 0
	blockReader := es_block.NewReader(f, z.warehouseBlockSize)

	for {
		l.Debug("Reading data", esl.Int64("offset", offset))
		data, isEof, rdErr := blockReader.ReadBlock()
		if isEof {
			l.Debug("Reached to EOF")
			callback(z, path, offset, true, z.getReceipt(data))
			return
		}

		if rdErr != nil {
			failure(z, path, rdErr)
			return
		}

		callback(z, path, offset, false, z.getReceipt(data))
		offset += int64(len(data))
	}
}

func (z *bwImpl) blockLoader() {
	l := z.l
	for !z.warehouseShutdown {
		var backlogPath string = ""
		var backlogFound bool
		var backlogCallback BlockWarehouseCallback
		var backlogFailure BlockWarehouseFailureCallback

		z.backlogMutex.Lock()
		if len(z.backlogPaths) > 1 {
			backlogPath = z.backlogPaths[0]
			z.backlogPaths = z.backlogPaths[1:]

			backlogCallback, backlogFound = z.backlogCallbacks[backlogPath]
			delete(z.backlogCallbacks, backlogPath)
			if !backlogFound {
				l.Warn("Backlog callback not found", esl.String("path", backlogPath))
				panic("backlog callback not found")
			}

			backlogFailure, backlogFound = z.backlogFailure[backlogPath]
			delete(z.backlogFailure, backlogPath)
			if !backlogFound {
				l.Warn("Backlog failure callback not found", esl.String("path", backlogPath))
				panic("Backlog failure callback not found")
			}
		}
		z.backlogMutex.Unlock()

		if backlogPath != "" {
			z.loadBlock(backlogPath, backlogCallback, backlogFailure)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
	l.Debug("Block loader shutdown")
}

func (z *bwImpl) Load(path string, r BlockWarehouseCallback, f BlockWarehouseFailureCallback) {
	z.l.Debug("Load path", esl.String("path", path))
	z.backlogCount.Add(1)
	z.backlogMutex.Lock()
	z.backlogPaths = append(z.backlogPaths, path)
	z.backlogCallbacks[path] = r
	z.backlogFailure[path] = f
	z.backlogMutex.Unlock()
}

func (z *bwImpl) Receive(receipt BlockWarehouseReceipt) ([]byte, error) {
	z.blockMutex.Lock()
	defer func() {
		z.blockMutex.Unlock()
	}()

	data, found := z.block[receipt]
	if found {
		z.l.Debug("Release a receipt", esl.Int64("receipt", int64(receipt)), esl.Int("backlogSize", len(z.block)))
		delete(z.block, receipt)
		return data, nil
	} else {
		z.l.Debug("Unable to find a block for the receipt", esl.Int64("receipt", int64(receipt)))
		return nil, errors.New("no block found for the receipt")
	}
}
