package kv_kvs_impl

import (
	"encoding/json"
	"github.com/dgraph-io/badger"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"go.uber.org/zap"
)

func New(ctl app_control.Control, db *badger.DB, tx *badger.Txn) kv_kvs.Kvs {
	return &badgerWrapper{
		ctl: ctl,
		db:  db,
		tx:  tx,
	}
}

type badgerWrapper struct {
	ctl app_control.Control
	db  *badger.DB
	tx  *badger.Txn
}

func (z *badgerWrapper) NextSequence(name string) (uint64, error) {
	l := z.ctl.Log().With(zap.String("name", name))
	seq, err := z.db.GetSequence([]byte(name), 100)
	if err != nil {
		l.Debug("Unable to get seq", zap.Error(err))
		return 0, err
	}
	defer seq.Release()
	s, err := seq.Next()
	if err != nil {
		l.Debug("Unable to generate seq", zap.Error(err))
		return 0, err
	}
	return s, nil
}

func (z *badgerWrapper) PutString(key string, value string) error {
	l := z.ctl.Log()
	err := z.tx.Set([]byte(key), []byte(value))
	if err != nil {
		l.Debug("Unable to put key/value", zap.String("key", key))
		return err
	}
	return nil
}

func (z *badgerWrapper) PutBytes(key string, value []byte) error {
	l := z.ctl.Log()
	err := z.tx.Set([]byte(key), value)
	if err != nil {
		l.Debug("Unable to put key/value", zap.String("key", key))
		return err
	}
	return nil
}

func (z *badgerWrapper) PutJson(key string, j json.RawMessage) error {
	l := z.ctl.Log()
	err := z.tx.Set([]byte(key), j)
	if err != nil {
		l.Debug("Unable to put key/value", zap.String("key", key))
		return err
	}
	return nil
}

func (z *badgerWrapper) PutJsonModel(key string, v interface{}) error {
	l := z.ctl.Log()
	b, err := json.Marshal(v)
	if err != nil {
		l.Debug("Unable to marshal value", zap.Error(err))
		return err
	}
	err = z.tx.Set([]byte(key), b)
	if err != nil {
		l.Debug("Unable to put key/value", zap.String("key", key))
		return err
	}
	return nil
}

func (z *badgerWrapper) GetString(key string) (value string, err error) {
	r, err := z.tx.Get([]byte(key))
	if err != nil {
		return "", err
	}
	v, err := r.ValueCopy(nil)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func (z *badgerWrapper) GetBytes(key string) (value []byte, err error) {
	r, err := z.tx.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	return r.ValueCopy(nil)
}

func (z *badgerWrapper) GetJson(key string) (j json.RawMessage, err error) {
	r, err := z.tx.Get([]byte(key))
	if err != nil {
		return nil, err
	}
	return r.ValueCopy(nil)
}

func (z *badgerWrapper) GetJsonModel(key string, v interface{}) (err error) {
	r, err := z.GetJson(key)
	if err = json.Unmarshal(r, v); err != nil {
		return err
	}
	return nil
}

func (z *badgerWrapper) Delete(key string) error {
	return z.tx.Delete([]byte(key))
}

func (z *badgerWrapper) ForEach(f func(key string, value []byte) error) error {
	opts := badger.DefaultIteratorOptions
	opts.PrefetchSize = 100
	it := z.tx.NewIterator(opts)
	defer it.Close()
	for it.Rewind(); it.Valid(); it.Next() {
		i := it.Item()
		v, err := i.ValueCopy(nil)
		if err != nil {
			return err
		}
		if err := f(string(i.Key()), v); err != nil {
			return err
		}
	}
	return nil
}
