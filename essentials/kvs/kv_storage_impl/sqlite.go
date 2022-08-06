package kv_storage_impl

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
	"path/filepath"
)

func InternalNewSqlite(name string, logger esl.Logger) kv_storage.Lifecycle {
	return &sqWrapper{
		name:   name,
		logger: logger,
	}
}

type sqWrapper struct {
	path   string
	dbPath string
	name   string
	logger esl.Logger
	db     *sql.DB
	kvs    kv_kvs.Kvs
}

func (z *sqWrapper) log() esl.Logger {
	if z.logger != nil {
		return z.logger.With(esl.String("name", z.name))
	} else {
		return esl.Default().With(esl.String("name", z.name))
	}
}

func (z *sqWrapper) Close() {
	z.log().Debug("Close database")
	if err := z.db.Close(); err != nil {
		z.log().Debug("Unable to close the database", esl.Error(err))
	}
	z.db = nil
	z.dbPath = ""
}

func (z *sqWrapper) View(f func(kvs kv_kvs.Kvs) error) error {
	return f(z.kvs)
}

func (z *sqWrapper) Update(f func(kvs kv_kvs.Kvs) error) error {
	return f(z.kvs)
}

func (z *sqWrapper) Kvs() kv_kvs.Kvs {
	return z.kvs
}

func (z *sqWrapper) Open(path string) (err error) {
	l := z.log()
	z.dbPath = filepath.Join(path, es_filepath.Escape(z.name))
	z.db, err = sql.Open("sqlite3", z.dbPath)
	if err != nil {
		l.Debug("Unable to open the database", esl.Error(err), esl.String("path", z.dbPath))
		return err
	}
	if err := z.db.Ping(); err != nil {
		l.Debug("Ping failed", esl.Error(err))
		return err
	}
	_, err = z.db.Exec(kv_kvs_impl.KvsTableSchema)
	if err != nil {
		l.Debug("Unable to create the table", esl.Error(err))
		return err
	}
	z.kvs, err = kv_kvs_impl.NewSqlite(z.name, z.log(), z.db)
	if err != nil {
		_ = z.db.Close()
	}
	return err
}

func (z *sqWrapper) SetLogger(logger esl.Logger) {
	z.logger = logger
}

func (z *sqWrapper) Path() string {
	return z.path
}

func (z *sqWrapper) Delete() error {
	if err := z.db.Close(); err != nil {
		return err
	}
	return os.RemoveAll(z.dbPath)
}
