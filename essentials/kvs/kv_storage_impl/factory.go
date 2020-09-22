package kv_storage_impl

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"sync"
)

type Storage interface {
	kv_storage.Storage

	// Open KVS storage
	Open(path string) error

	SetLogger(logger esl.Logger)
}

func NewFactory(basePath string, logger esl.Logger) kv_storage.Factory {
	return &factoryImpl{
		basePath: basePath,
		logger:   logger,
		storages: make([]Storage, 0),
	}
}

type factoryImpl struct {
	basePath string
	logger   esl.Logger
	storages []Storage
	mutex    sync.Mutex
}

func (z *factoryImpl) log() esl.Logger {
	return z.logger.With(esl.String("BasePath", z.basePath))
}

func (z *factoryImpl) New(name string) (kv_storage.Storage, error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.log().With(esl.String("name", name))
	l.Debug("Create new storage")

	sto := newProxy(name, z.logger)
	err := sto.Open(z.basePath)
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
	z.storages = make([]Storage, 0)
	l.Debug("Closed")
}
