package kv_kvs_impl

import (
	"encoding/json"
	"git.mills.io/prologic/bitcask"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/log/esl"
	"reflect"
)

func NewBitcask(name string, log esl.Logger, db *bitcask.Bitcask) kv_kvs.Kvs {
	return &bcImpl{
		logger: log,
		name:   name,
		db:     db,
	}
}

type bcImpl struct {
	name   string
	logger esl.Logger
	db     *bitcask.Bitcask
}

func (z *bcImpl) Lock() error {
	return z.db.Lock()
}

func (z *bcImpl) Unlock() error {
	return z.db.Unlock()
}

func (z *bcImpl) log() esl.Logger {
	return z.logger.With(esl.String("name", z.name))
}

func (z *bcImpl) op(opName string, key string, f func() error) error {
	if len(key) < 1 {
		return kv_kvs.ErrorInvalidKey
	}
	l := z.log().With(esl.String("opName", opName))
	if err := f(); err != nil {
		switch err {
		case bitcask.ErrKeyTooLarge, bitcask.ErrEmptyKey:
			l.Debug("Invalid key", esl.Error(err))
			return kv_kvs.ErrorInvalidKey

		default:
			l.Debug("Op failed", esl.Error(err))
			return err
		}
	}
	return nil
}

func (z *bcImpl) PutString(key string, value string) error {
	return z.op("PutString", key, func() error {
		return z.db.Put([]byte(key), []byte(value))
	})
}

func (z *bcImpl) PutBytes(key string, value []byte) error {
	return z.op("PutBytes", key, func() error {
		return z.db.Put([]byte(key), value)
	})
}

func (z *bcImpl) PutJson(key string, j json.RawMessage) error {
	return z.op("PutJson", key, func() error {
		return z.db.Put([]byte(key), j)
	})
}

func (z *bcImpl) PutJsonModel(key string, v interface{}) error {
	l := z.log()
	b, err := json.Marshal(v)
	if err != nil {
		l.Debug("Unable to marshal value", esl.Error(err))
		return err
	}
	return z.op("PutJsonModel", key, func() error {
		return z.db.Put([]byte(key), b)
	})
}

func (z *bcImpl) PutRaw(key, value []byte) error {
	return z.op("PutRaw", string(key), func() error {
		return z.db.Put(key, value)
	})
}

func (z *bcImpl) getOp(opName string, key string, unmarshal func(v []byte) error) (err error) {
	l := z.log().With(esl.String("opName", opName), esl.String("key", key))
	v, err := z.db.Get([]byte(key))
	if err == bitcask.ErrKeyNotFound {
		l.Debug("Key not found", esl.Error(err))
		return kv_kvs.ErrorNotFound
	}
	if err != nil {
		l.Debug("Get failed", esl.Error(err))
		return err
	}
	if err := unmarshal(v); err != nil {
		l.Debug("Unmarshal failed", esl.Error(err))
	}
	return nil
}

func (z *bcImpl) GetString(key string) (value string, err error) {
	err = z.getOp("GetString", key, func(v []byte) error {
		value = string(v)
		return nil
	})
	return
}

func (z *bcImpl) GetBytes(key string) (value []byte, err error) {
	err = z.getOp("GetBytes", key, func(v []byte) error {
		value = v
		return nil
	})
	return
}

func (z *bcImpl) GetJson(key string) (j json.RawMessage, err error) {
	err = z.getOp("GetJson", key, func(v []byte) error {
		j = v
		return nil
	})
	return
}

func (z *bcImpl) GetJsonModel(key string, v interface{}) (err error) {
	err = z.getOp("GetBytes", key, func(v0 []byte) error {
		return json.Unmarshal(v0, v)
	})
	return
}

func (z *bcImpl) Delete(key string) error {
	return z.op("Delete", key, func() error {
		return z.db.Delete([]byte(key))
	})
}

func (z *bcImpl) ForEach(f func(key string, value []byte) error) error {
	l := z.log()
	return z.db.Fold(func(key []byte) error {
		value, err := z.db.Get(key)
		if err != nil {
			l.Debug("Unable to get a value", esl.Error(err))
			return err
		}
		return f(string(key), value)
	})
}

func (z *bcImpl) ForEachRaw(f func(key []byte, value []byte) error) error {
	l := z.log()
	return z.db.Fold(func(key []byte) error {
		value, err := z.db.Get(key)
		if err != nil {
			l.Debug("Unable to get a value", esl.Error(err))
			return err
		}
		return f(key, value)
	})
}

func (z *bcImpl) ForEachModel(model interface{}, f func(key string, m interface{}) error) error {
	l := z.log()
	mt := reflect.ValueOf(model).Elem().Type()
	return z.db.Fold(func(key []byte) error {
		value, err := z.db.Get(key)
		if err != nil {
			l.Debug("Unable to get a value", esl.Error(err))
			return err
		}
		m := reflect.New(mt).Interface()
		if err := json.Unmarshal(value, m); err != nil {
			l.Debug("Unable to unmarshal", esl.Error(err))
			return err
		}
		return f(string(key), m)
	})
}
