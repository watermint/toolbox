package kv_kvs_impl

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
)

func NewEmpty() kv_kvs.Kvs {
	return &emptyImpl{}
}

type emptyImpl struct {
}

func (z emptyImpl) PutString(key string, value string) error {
	return nil
}

func (z emptyImpl) PutBytes(key string, value []byte) error {
	return nil
}

func (z emptyImpl) PutJson(key string, j json.RawMessage) error {
	return nil
}

func (z emptyImpl) PutJsonModel(key string, v interface{}) error {
	return nil
}

func (z emptyImpl) PutRaw(key, value []byte) error {
	return nil
}

func (z emptyImpl) GetString(key string) (value string, err error) {
	return "", kv_kvs.ErrorNotFound
}

func (z emptyImpl) GetBytes(key string) (value []byte, err error) {
	return nil, kv_kvs.ErrorNotFound
}

func (z emptyImpl) GetJson(key string) (j json.RawMessage, err error) {
	return nil, kv_kvs.ErrorNotFound
}

func (z emptyImpl) GetJsonModel(key string, v interface{}) (err error) {
	return kv_kvs.ErrorNotFound
}

func (z emptyImpl) Delete(key string) error {
	return nil
}

func (z emptyImpl) ForEach(f func(key string, value []byte) error) error {
	return nil
}

func (z emptyImpl) ForEachRaw(f func(key []byte, value []byte) error) error {
	return nil
}

func (z emptyImpl) ForEachModel(model interface{}, f func(key string, m interface{}) error) error {
	return nil
}

func (z emptyImpl) Lock() error {
	return nil
}

func (z emptyImpl) Unlock() error {
	return nil
}
