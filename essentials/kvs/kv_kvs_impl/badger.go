package kv_kvs_impl

import (
	"encoding/json"
	"errors"
	"git.mills.io/prologic/bitcask"
	"github.com/dgraph-io/badger/v4"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewBadger(name string, log esl.Logger, db *badger.DB, tx *badger.Txn) kv_kvs.Kvs {
	return &badgerImpl{
		logger: log,
		name:   name,
		db:     db,
		tx:     tx,
	}
}

type badgerImpl struct {
	name   string
	logger esl.Logger
	db     *badger.DB
	tx     *badger.Txn
}

func (z *badgerImpl) log() esl.Logger {
	return z.logger.With(esl.String("name", z.name))
}

func (z *badgerImpl) opWrite(opName string, key string, f func() error) error {
	if len(key) < 1 {
		return kv_kvs.ErrorInvalidKey
	}
	l := z.log().With(esl.String("opName", opName))
	if err := f(); err != nil {
		switch {
		case errors.Is(err, bitcask.ErrKeyTooLarge),
			errors.Is(err, bitcask.ErrEmptyKey):
			l.Debug("Invalid key", esl.Error(err))
			return kv_kvs.ErrorInvalidKey

		default:
			l.Debug("Op failed", esl.Error(err))
			return err
		}
	}
	return nil
}

func (z *badgerImpl) opRead(opName string, key string, unmarshal func(v []byte) error) (err error) {
	l := z.log().With(esl.String("opName", opName), esl.String("key", key))
	v, err := z.tx.Get([]byte(key))
	if errors.Is(err, badger.ErrKeyNotFound) {
		l.Debug("Key not found", esl.Error(err))
		return kv_kvs.ErrorNotFound
	}
	if err != nil {
		l.Debug("Get failed", esl.Error(err))
		return err
	}
	var b []byte
	if b, err = v.ValueCopy(nil); err != nil {
		l.Debug("Unable to copy value", esl.Error(err))
		return err
	}
	if err := unmarshal(b); err != nil {
		l.Debug("Unmarshal failed", esl.Error(err))
	}
	return nil
}

func (z *badgerImpl) PutString(key string, value string) error {
	return z.opWrite("PutString", key, func() error {
		return z.tx.Set([]byte(key), []byte(value))
	})
}

func (z *badgerImpl) PutJson(key string, j json.RawMessage) error {
	return z.opWrite("PutJson", key, func() error {
		return z.tx.Set([]byte(key), j)
	})
}

func (z *badgerImpl) PutJsonModel(key string, v interface{}) error {
	l := z.log()
	b, err := json.Marshal(v)
	if err != nil {
		l.Debug("Unable to marshal value", esl.Error(err))
		return err
	}
	return z.opWrite("PutJsonModel", key, func() error {
		return z.tx.Set([]byte(key), b)
	})
}

func (z *badgerImpl) GetString(key string) (value string, err error) {
	err = z.opRead("GetString", key, func(v []byte) error {
		value = string(v)
		return nil
	})
	return
}

func (z *badgerImpl) GetJson(key string) (j json.RawMessage, err error) {
	err = z.opRead("GetJson", key, func(v []byte) error {
		j = v
		return nil
	})
	return
}

func (z *badgerImpl) GetJsonModel(key string, v interface{}) (err error) {
	err = z.opRead("GetJsonModel", key, func(b []byte) error {
		return json.Unmarshal(b, v)
	})
	return
}

func (z *badgerImpl) Delete(key string) error {
	return z.opWrite("Delete", key, func() error {
		return z.tx.Delete([]byte(key))
	})
}

func (z *badgerImpl) ForEach(f func(key string, value []byte) error) error {
	l := z.log()
	itr := z.tx.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	for itr.Rewind(); itr.Valid(); itr.Next() {
		item := itr.Item()
		k := item.Key()
		v, err := item.ValueCopy(nil)
		if err != nil {
			l.Debug("Unable to copy value", esl.Error(err))
			return err
		}
		if err := f(string(k), v); err != nil {
			l.Debug("Callback failed", esl.Error(err))
			return err
		}
	}
	return nil
}

func (z *badgerImpl) ForEachRaw(f func(key []byte, value []byte) error) error {
	l := z.log()
	itr := z.tx.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	for itr.Rewind(); itr.Valid(); itr.Next() {
		item := itr.Item()
		k := item.Key()
		v, err := item.ValueCopy(nil)
		if err != nil {
			l.Debug("Unable to copy value", esl.Error(err))
			return err
		}
		if err := f(k, v); err != nil {
			l.Debug("Callback failed", esl.Error(err))
			return err
		}
	}
	return nil
}

func (z *badgerImpl) ForEachModel(model interface{}, f func(key string, m interface{}) error) error {
	l := z.log()
	itr := z.tx.NewIterator(badger.DefaultIteratorOptions)
	defer itr.Close()
	for itr.Rewind(); itr.Valid(); itr.Next() {
		item := itr.Item()
		k := item.Key()
		v, err := item.ValueCopy(nil)
		if err != nil {
			l.Debug("Unable to copy value", esl.Error(err))
			return err
		}
		if err := json.Unmarshal(v, model); err != nil {
			l.Debug("Unable to unmarshal value", esl.Error(err))
			return err
		}
		if err := f(string(k), model); err != nil {
			l.Debug("Callback failed", esl.Error(err))
			return err
		}
	}
	return nil
}
