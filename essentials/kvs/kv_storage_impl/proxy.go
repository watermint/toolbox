package kv_storage_impl

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func newProxy(name string, logger esl.Logger) Storage {
	return &proxyImpl{
		name:   name,
		logger: logger,
	}
}

type proxyImpl struct {
	name    string
	storage Storage
	logger  esl.Logger
}

func (z *proxyImpl) SetLogger(logger esl.Logger) {
	z.logger = logger
}

func (z *proxyImpl) Kvs() kv_kvs.Kvs {
	return z.storage.Kvs()
}

func (z *proxyImpl) newStorage() Storage {
	return InternalNewBitcask(z.name, z.logger)
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

func (z *proxyImpl) Open(path string) error {
	z.storage = z.newStorage()
	return z.storage.Open(path)
}
