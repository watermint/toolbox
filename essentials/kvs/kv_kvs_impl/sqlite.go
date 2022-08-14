package kv_kvs_impl

import (
	"database/sql"
	"encoding/json"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/log/esl"
	"reflect"
)

const (
	KvsTableName     = "kvs"
	KvsTableSchema   = "CREATE TABLE " + KvsTableName + " (k TEXT PRIMARY KEY NOT NULL, v TEXT)"
	KvsStmtPutString = "INSERT OR REPLACE INTO " + KvsTableName + "(K, V) VALUES(?, ?)"
	KvsStmtGetString = "SELECT v FROM " + KvsTableName + " WHERE k = ?"
	KvsStmtForEach   = "SELECT k, v FROM " + KvsTableName + " ORDER BY k"
	KvsStmtDelete    = "DELETE FROM " + KvsTableName + " WHERE k = ?"
)

func NewSqlite(name string, log esl.Logger, db *sql.DB) (kv_kvs.Kvs, error) {
	putStmt, err := db.Prepare(KvsStmtPutString)
	if err != nil {
		return nil, err
	}
	getStmt, err := db.Prepare(KvsStmtGetString)
	if err != nil {
		return nil, err
	}
	forStmt, err := db.Prepare(KvsStmtForEach)
	if err != nil {
		return nil, err
	}
	delStmt, err := db.Prepare(KvsStmtDelete)
	if err != nil {
		return nil, err
	}

	return &sqImpl{
		putStmt: putStmt,
		getStmt: getStmt,
		forStmt: forStmt,
		delStmt: delStmt,
		name:    name,
		logger:  log,
		db:      db,
	}, nil
}

type sqImpl struct {
	putStmt *sql.Stmt
	getStmt *sql.Stmt
	forStmt *sql.Stmt
	delStmt *sql.Stmt
	name    string
	logger  esl.Logger
	db      *sql.DB
}

func (z *sqImpl) PutString(key string, value string) error {
	_, err := z.putStmt.Exec(key, value)
	return err
}

func (z *sqImpl) PutJson(key string, j json.RawMessage) error {
	return z.PutString(key, string(j))
}

func (z *sqImpl) PutJsonModel(key string, v interface{}) error {
	s, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return z.PutString(key, string(s))
}

func (z *sqImpl) GetString(key string) (value string, err error) {
	r, err := z.getStmt.Query(key)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = r.Close()
	}()
	if !r.Next() {
		return "", kv_kvs.ErrorNotFound
	}
	err = r.Scan(&value)
	return value, err
}

func (z *sqImpl) GetJson(key string) (j json.RawMessage, err error) {
	s, err := z.GetString(key)
	return json.RawMessage(s), err
}

func (z *sqImpl) GetJsonModel(key string, v interface{}) (err error) {
	s, err := z.GetString(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(s), v)
}

func (z *sqImpl) Delete(key string) error {
	_, err := z.delStmt.Exec(key)
	return err
}

func (z *sqImpl) ForEach(f func(key string, value []byte) error) error {
	h := func(rows *sql.Rows) error {
		var k, v string
		if sErr := rows.Scan(&k, &v); sErr != nil {
			return sErr
		}
		return f(k, []byte(v))
	}
	r, err := z.forStmt.Query()
	if err != nil {
		return err
	}
	for r.Next() {
		if err := h(r); err != nil {
			return err
		}
	}
	return nil
}

func (z *sqImpl) ForEachRaw(f func(key []byte, value []byte) error) error {
	return z.ForEach(func(key string, value []byte) error {
		return f([]byte(key), value)
	})
}

func (z *sqImpl) ForEachModel(model interface{}, f func(key string, m interface{}) error) error {
	mt := reflect.ValueOf(model).Elem().Type()
	return z.ForEach(func(key string, value []byte) error {
		m := reflect.New(mt).Interface()
		if err := json.Unmarshal(value, m); err != nil {
			return err
		}
		return f(key, m)
	})
}
