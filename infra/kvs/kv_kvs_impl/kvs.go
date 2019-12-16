package kv_kvs_impl

import (
	"encoding/json"
	"errors"
	"github.com/etcd-io/bbolt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_cursor"
	"github.com/watermint/toolbox/infra/kvs/kv_cursor_impl"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"go.uber.org/zap"
)

func New(ctl app_control.Control, name string, bucket *bbolt.Bucket) kv_kvs.Kvs {
	return &bboltWrapper{
		ctl:    ctl,
		name:   name,
		bucket: bucket,
	}
}

type bboltWrapper struct {
	ctl    app_control.Control
	name   string
	bucket *bbolt.Bucket
}

func (z *bboltWrapper) NextSequence() (uint64, error) {
	return z.bucket.NextSequence()
}

func (z *bboltWrapper) Nested(key string) (kvs kv_kvs.Kvs, err error) {
	l := z.ctl.Log()
	b := z.bucket.Bucket([]byte(key))
	if b != nil {
		return New(z.ctl, key, b), nil
	}

	nb, err := z.bucket.CreateBucketIfNotExists([]byte(key))
	if err != nil {
		l.Debug("Unable to create nested bucket", zap.String("name", z.name), zap.String("key", key))
		return nil, err
	}
	return New(z.ctl, key, nb), nil
}

func (z *bboltWrapper) DeleteNested(key string) error {
	return z.bucket.DeleteBucket([]byte(key))
}

func (z *bboltWrapper) PutString(key string, value string) error {
	l := z.ctl.Log()
	err := z.bucket.Put([]byte(key), []byte(value))
	if err != nil {
		l.Debug("Unable to put key/value", zap.String("name", z.name), zap.String("key", key))
		return err
	}
	return nil
}

func (z *bboltWrapper) PutBytes(key string, value []byte) error {
	l := z.ctl.Log()
	err := z.bucket.Put([]byte(key), value)
	if err != nil {
		l.Debug("Unable to put key/value", zap.String("name", z.name), zap.String("key", key))
		return err
	}
	return nil
}

func (z *bboltWrapper) PutJson(key string, j json.RawMessage) error {
	l := z.ctl.Log()
	err := z.bucket.Put([]byte(key), j)
	if err != nil {
		l.Debug("Unable to put key/value", zap.String("name", z.name), zap.String("key", key))
		return err
	}
	return nil
}

func (z *bboltWrapper) PutJsonModel(key string, v interface{}) error {
	l := z.ctl.Log()
	b, err := json.Marshal(v)
	if err != nil {
		l.Debug("Unable to marshal value", zap.Error(err))
		return err
	}
	err = z.bucket.Put([]byte(key), b)
	if err != nil {
		l.Debug("Unable to put key/value", zap.String("name", z.name), zap.String("key", key))
		return err
	}
	return nil
}

func (z *bboltWrapper) GetString(key string) (value string, err error) {
	r := z.bucket.Get([]byte(key))
	if r == nil {
		return "", errors.New("key not found in this kvs")
	}
	return string(r), nil
}

func (z *bboltWrapper) GetBytes(key string) (value []byte, err error) {
	r := z.bucket.Get([]byte(key))
	if r == nil {
		return nil, errors.New("key not found in this kvs")
	}
	return r, nil
}

func (z *bboltWrapper) GetJson(key string) (j json.RawMessage, err error) {
	r := z.bucket.Get([]byte(key))
	if r == nil {
		return nil, errors.New("key not found in this kvs")
	}
	return r, nil
}

func (z *bboltWrapper) GetJsonModel(key string, v interface{}) (err error) {
	r := z.bucket.Get([]byte(key))
	if r == nil {
		return errors.New("key not found in this kvs")
	}
	if err = json.Unmarshal(r, v); err != nil {
		return err
	}
	return nil
}

func (z *bboltWrapper) Delete(key string) error {
	return z.bucket.Delete([]byte(key))
}

func (z *bboltWrapper) ForEach(f func(key string, value []byte) error) error {
	return z.bucket.ForEach(func(k, v []byte) error {
		return f(string(k), v)
	})
}

func (z *bboltWrapper) Cursor() kv_cursor.Cursor {
	return kv_cursor_impl.New(z.ctl, z, z.bucket.Cursor())
}
