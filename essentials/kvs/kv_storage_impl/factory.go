package kv_storage_impl

import (
	"github.com/watermint/essentials/eformat/euuid"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
	"path/filepath"
	"sync"
)

func NewFactory(basePath string, logger esl.Logger) kv_storage.Factory {
	return &factoryImpl{
		basePath: basePath,
		logger:   logger,
		storages: make([]kv_storage.Lifecycle, 0),
	}
}

type factoryImpl struct {
	basePath string
	logger   esl.Logger
	storages []kv_storage.Lifecycle
	mutex    sync.Mutex
}

func (z *factoryImpl) log() esl.Logger {
	return z.logger.With(esl.String("BasePath", z.basePath))
}

func (z *factoryImpl) New(name string) (kv_storage.Lifecycle, error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.log().With(esl.String("name", name))
	l.Debug("Create new storage")
	kvPath := filepath.Join(z.basePath, euuid.NewV4().String())
	if err := os.MkdirAll(kvPath, 0755); err != nil {
		l.Debug("Unable to create a directory", esl.Error(err))
		return nil, err
	}

	sto := NewProxy(name, z.logger)
	err := sto.Open(kvPath)
	if err != nil {
		l.Debug("Unable to open the storage", esl.Error(err))
		return nil, err
	}
	z.storages = append(z.storages, sto)

	l.Debug("Storage created")
	return sto, nil
}

func (z *factoryImpl) Close() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.log()
	l.Debug("Closing all storages")

	for _, sto := range z.storages {
		l.Debug("Closing a storage")
		sto.Close()
	}
	z.storages = make([]kv_storage.Lifecycle, 0)
	l.Debug("Closed")
}
