package kv_storage_impl

import (
	"errors"
	"git.mills.io/prologic/bitcask"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"path/filepath"
)

func newBitcask(name string, log esl.Logger) kv_storage.Lifecycle {
	bc := &bcWrapper{
		name:   name,
		logger: log,
	}
	return bc
}

type bcWrapper struct {
	path   string
	name   string
	db     *bitcask.Bitcask
	logger esl.Logger
	kvs    kv_kvs.Kvs
}

func (z *bcWrapper) Delete() error {
	return z.db.DeleteAll()
}

func (z *bcWrapper) Path() string {
	return z.path
}

func (z *bcWrapper) SetLogger(logger esl.Logger) {
	z.logger = logger
}

func (z *bcWrapper) OpenWithPath(path string) error {
	z.name = filepath.Base(path)
	return z.openWithPath(path)
}

func (z *bcWrapper) log() esl.Logger {
	return z.logger.With(esl.String("name", z.name), esl.String("path", z.path))
}

func (z *bcWrapper) openWithPath(path string) (err error) {
	z.path = path
	l := z.log()
	l.Debug("open")

	z.db, err = bitcask.Open(path,

		// 64bytes (default) -> 64kB (file path may up to 32kb in UNIX path or UNC path)
		bitcask.WithMaxKeySize(2<<16),

		// 64kB (default) -> 256kB
		bitcask.WithMaxValueSize(2<<18),
	)

	switch {
	case err == nil:
		l.Debug("Database open")

	case errors.Is(err, bitcask.ErrDatabaseLocked):
		l.Debug("Database locked", esl.Error(err))
		return kv_storage.ErrorStorageLocked

	default:
		l.Debug("Unable to open the database", esl.Error(err))
		return err
	}

	z.kvs = kv_kvs_impl.NewBitcask(z.name, z.logger, z.db)

	return nil
}

func (z *bcWrapper) Open(path string) error {
	kvsPath := filepath.Join(path, es_filepath.Escape(z.name))
	return z.openWithPath(kvsPath)
}

func (z *bcWrapper) Close() {
	l := z.log()
	l.Debug("Close database")
	if err := z.db.Sync(); err != nil {
		l.Debug("Unable to sync the database", esl.Error(err))
	}
	err := z.db.Close()
	if err != nil {
		l.Debug("There is an error on close", esl.Error(err))
	}
}

func (z *bcWrapper) View(f func(kvs kv_kvs.Kvs) error) error {
	return f(z.kvs)
}

func (z *bcWrapper) Update(f func(kvs kv_kvs.Kvs) error) error {
	return f(z.kvs)
}
