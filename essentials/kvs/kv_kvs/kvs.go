package kv_kvs

import (
	"encoding/json"
	"errors"
)

// Key-value store interface.
type Kvs interface {
	PutString(key string, value string) error
	PutJson(key string, j json.RawMessage) error
	PutJsonModel(key string, v interface{}) error

	GetString(key string) (value string, err error)
	GetJson(key string) (j json.RawMessage, err error)
	GetJsonModel(key string, v interface{}) (err error)

	Delete(key string) error
	ForEach(f func(key string, value []byte) error) error
	ForEachRaw(f func(key, value []byte) error) error
	ForEachModel(model interface{}, f func(key string, m interface{}) error) error
}

var (
	ErrorInvalidKey = errors.New("invalid key")
	ErrorNotFound   = errors.New("key not found")
)
