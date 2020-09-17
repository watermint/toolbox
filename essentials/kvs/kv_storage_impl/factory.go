package kv_storage_impl

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"path/filepath"
	"sync"
)

func NewStorage(name string) kv_storage.Storage {
	return newProxy(name)
}

func NewStorageWithPath(ctl app_control.Control, path string) (kv_storage.Storage, error) {
	s := newProxy(filepath.Base(path))
	err := s.OpenWithPath(ctl, path)
	return s, err
}

func NewFactory(ctl app_control.Control) kv_storage.Factory {
	return &factoryImpl{
		ctl:      ctl,
		storages: make([]Storage, 0),
	}
}

type factoryImpl struct {
	ctl      app_control.Control
	storages []Storage
	mutex    sync.Mutex
}

func (z *factoryImpl) New(name string) (kv_storage.Storage, error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.ctl.Log().With(esl.String("name", name))
	l.Debug("Create new storage")

	sto := newProxy(name)
	err := sto.Open(z.ctl)
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

	l := z.ctl.Log()
	l.Debug("Closing all storages")

	for _, sto := range z.storages {
		l.Debug("Closing a storage")
		sto.Close()
	}
	z.storages = make([]Storage, 0)
	l.Debug("Closed")
}

type Storage interface {
	kv_storage.Storage

	// Open KVS storage
	Open(ctl app_control.Control) error

	OpenWithPath(ctl app_control.Control, path string) error
}

func newProxy(name string) Storage {
	return &proxyImpl{
		name: name,
	}
}

type proxyImpl struct {
	name    string
	storage Storage
}

func (z *proxyImpl) newStorage(ctl app_control.Control) Storage {
	return InternalNewBitcask(z.name)
}

func (z *proxyImpl) Open(ctl app_control.Control) error {
	z.storage = z.newStorage(ctl)
	return z.storage.Open(ctl)
}

func (z *proxyImpl) Close() {
	z.storage.Close()
}

func (z *proxyImpl) View(f func(kvs kv_kvs.Kvs) error) error {
	return z.storage.View(f)
}

func (z *proxyImpl) Update(f func(kvs kv_kvs.Kvs) error) error {
	return z.storage.Update(f)
}

func (z *proxyImpl) OpenWithPath(ctl app_control.Control, path string) error {
	z.storage = z.newStorage(ctl)
	return z.storage.OpenWithPath(ctl, path)
}
