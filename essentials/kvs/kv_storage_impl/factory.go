package kv_storage_impl

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"path/filepath"
)

func New(name string) kv_storage.Storage {
	return newProxy(name)
}

func NewWithPath(ctl app_control.Control, path string) (kv_storage.Storage, error) {
	s := newProxy(filepath.Base(path))
	err := s.OpenWithPath(ctl, path)
	return s, err
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
	switch {
	case ctl.Feature().Experiment(app.ExperimentKvsStorageUseBadger):
		return InternalNewBadger(z.name)
	default:
		return InternalNewBitcask(z.name)
	}
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
