package kv_kvs_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"sync"
)

func NewTurnstile(kvs kv_kvs.Kvs) kv_kvs.Kvs {
	return &turnstileImpl{
		m:   sync.Mutex{},
		kvs: kvs,
	}
}

type turnstileImpl struct {
	m   sync.Mutex
	kvs kv_kvs.Kvs
}

func (z *turnstileImpl) PutString(key string, value string) error {
	z.m.Lock()
	defer z.m.Unlock()
	return z.kvs.PutString(key, value)
}

func (z *turnstileImpl) PutJson(key string, j json.RawMessage) error {
	z.m.Lock()
	defer z.m.Unlock()
	return z.kvs.PutJson(key, j)
}

func (z *turnstileImpl) PutJsonModel(key string, v interface{}) error {
	z.m.Lock()
	defer z.m.Unlock()
	return z.PutJsonModel(key, v)
}

func (z *turnstileImpl) GetString(key string) (value string, err error) {
	z.m.Lock()
	defer z.m.Unlock()
	return z.GetString(key)
}

func (z *turnstileImpl) GetJson(key string) (j json.RawMessage, err error) {
	z.m.Lock()
	defer z.m.Unlock()
	return z.GetJson(key)
}

func (z *turnstileImpl) GetJsonModel(key string, v interface{}) (err error) {
	z.m.Lock()
	defer z.m.Unlock()
	return z.GetJsonModel(key, v)
}

func (z *turnstileImpl) Delete(key string) error {
	z.m.Lock()
	defer z.m.Unlock()
	return z.Delete(key)
}

func (z *turnstileImpl) ForEach(f func(key string, value []byte) error) error {
	z.m.Lock()
	defer z.m.Unlock()
	return z.ForEach(f)
}

func (z *turnstileImpl) ForEachRaw(f func(key []byte, value []byte) error) error {
	z.m.Lock()
	defer z.m.Unlock()
	return z.ForEachRaw(f)
}

func (z *turnstileImpl) ForEachModel(model interface{}, f func(key string, m interface{}) error) error {
	z.m.Lock()
	defer z.m.Unlock()
	return z.ForEachModel(model, f)
}
