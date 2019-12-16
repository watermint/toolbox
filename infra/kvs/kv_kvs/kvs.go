package kv_kvs

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/kvs/kv_cursor"
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

	Nested(key string) (kvs Kvs, err error)
	DeleteNested(key string) error

	NextSequence() (uint64, error)

	Cursor() kv_cursor.Cursor
}
