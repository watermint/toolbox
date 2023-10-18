package kv_storage_impl

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewTurnstileStorage(storage kv_storage.Lifecycle) kv_storage.Lifecycle {
	return &turnstileImpl{
		storage: storage,
	}
}

type turnstileImpl struct {
	storage kv_storage.Lifecycle
}

func (z *turnstileImpl) Close() {
	z.storage.Close()
}

func (z *turnstileImpl) View(f func(kvs kv_kvs.Kvs) error) error {
	return z.storage.View(f)
}

func (z *turnstileImpl) Update(f func(kvs kv_kvs.Kvs) error) error {
	return z.storage.Update(f)
}

func (z *turnstileImpl) Open(path string) error {
	err := z.storage.Open(path)
	if err != nil {
		return err
	}
	return nil
}

func (z *turnstileImpl) SetLogger(logger esl.Logger) {
	z.storage.SetLogger(logger)
}

func (z *turnstileImpl) Path() string {
	return z.storage.Path()
}

func (z *turnstileImpl) Delete() error {
	return z.storage.Delete()
}
