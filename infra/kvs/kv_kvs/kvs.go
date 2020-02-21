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

	GetString(key string) (value string, err error)
	GetBytes(key string) (value []byte, err error)
	GetJson(key string) (j json.RawMessage, err error)
	GetJsonModel(key string, v interface{}) (err error)

	Delete(key string) error
	ForEach(func(key string, value []byte) error) error

	NextSequence(name string) (uint64, error)
}
