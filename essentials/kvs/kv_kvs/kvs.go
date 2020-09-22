package kv_kvs

import (
	"encoding/json"
)

// Key-value store interface.
type Kvs interface {
	PutString(key string, value string) error
	PutBytes(key string, value []byte) error
	PutJson(key string, j json.RawMessage) error
	PutJsonModel(key string, v interface{}) error
	PutRaw(key, value []byte) error

	GetString(key string) (value string, err error)
	GetBytes(key string) (value []byte, err error)
	GetJson(key string) (j json.RawMessage, err error)
	GetJsonModel(key string, v interface{}) (err error)

	Delete(key string) error
	ForEach(f func(key string, value []byte) error) error
	ForEachRaw(f func(key, value []byte) error) error
	ForEachModel(model interface{}, f func(key string, m interface{}) error) error

	Lock() error
	Unlock() error
}
