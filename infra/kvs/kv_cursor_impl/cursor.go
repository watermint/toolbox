package kv_cursor_impl

import (
	"github.com/etcd-io/bbolt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_cursor"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
)

func New(ctl app_control.Control, kvs kv_kvs.Kvs, cursor *bbolt.Cursor) kv_cursor.Cursor {
	return &bboltImpl{
		ctl:    ctl,
		kvs:    kvs,
		cursor: cursor,
	}
}

type bboltImpl struct {
	ctl    app_control.Control
	kvs    kv_kvs.Kvs
	cursor *bbolt.Cursor
}

func (z *bboltImpl) First() (key string, value []byte, exist bool) {
	k, v := z.cursor.First()
	if k == nil {
		return "", nil, false
	}
	return string(k), v, true
}

func (z *bboltImpl) Last() (key string, value []byte, exist bool) {
	k, v := z.cursor.Last()
	if k == nil {
		return "", nil, false
	}
	return string(k), v, true
}

func (z *bboltImpl) Next() (key string, value []byte, exist bool) {
	k, v := z.cursor.Next()
	if k == nil {
		return "", nil, false
	}
	return string(k), v, true
}

func (z *bboltImpl) Prev() (key string, value []byte, exist bool) {
	k, v := z.cursor.Prev()
	if k == nil {
		return "", nil, false
	}
	return string(k), v, true
}

func (z *bboltImpl) Seek(seek string) (key string, value []byte, exist bool) {
	k, v := z.cursor.Seek([]byte(seek))
	if k == nil {
		return "", nil, false
	}
	return string(k), v, true
}

func (z *bboltImpl) Delete() error {
	return z.cursor.Delete()
}
