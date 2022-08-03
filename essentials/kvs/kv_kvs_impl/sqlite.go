package kv_kvs_impl

import (
	"database/sql"
	"encoding/json"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewSqlite(name string, log esl.Logger, db *sql.DB) kv_kvs.Kvs {
	return &sqImpl{
		name:   name,
		logger: log,
		db:     db,
	}
}

type sqImpl struct {
	name   string
	logger esl.Logger
	db     *sql.DB
}

func (z *sqImpl) PutString(key string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) PutJson(key string, j json.RawMessage) error {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) PutJsonModel(key string, v interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) GetString(key string) (value string, err error) {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) GetJson(key string) (j json.RawMessage, err error) {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) GetJsonModel(key string, v interface{}) (err error) {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) ForEach(f func(key string, value []byte) error) error {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) ForEachRaw(f func(key []byte, value []byte) error) error {
	//TODO implement me
	panic("implement me")
}

func (z *sqImpl) ForEachModel(model interface{}, f func(key string, m interface{}) error) error {
	//TODO implement me
	panic("implement me")
}
