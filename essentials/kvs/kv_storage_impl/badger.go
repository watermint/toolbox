package kv_storage_impl

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_badger"
)

func InternalNewBadger(name string, logger esl.Logger) kv_storage.Lifecycle {
	return &badgerWrapper{
		name:   name,
		logger: logger,
	}
}

type badgerWrapper struct {
	path   string
	name   string
	db     *badger.DB
	logger esl.Logger
}

func (z *badgerWrapper) openWithPath(path string) error {
	z.path = path
	z.logger = z.logger.With(esl.String("path", path), esl.String("name", z.name))
	opts := badger.DefaultOptions(path).
		WithLogger(lgw_badger.NewLogWrapper(z.logger)).
		WithBlockCacheSize(32 << 20).   // 32MiB (default: 256MiB)
		WithValueLogFileSize(32 << 20). // 32MiB (default: 1GiB)
		WithIndexCacheSize(8 << 20).    // 8MiB
		WithCompression(options.ZSTD)

	db, err := badger.Open(opts)
	if err != nil {
		return err
	}
	z.db = db
	return nil
}

func (z *badgerWrapper) Close() {
	l := z.logger
	if err := z.db.Close(); err != nil {
		l.Debug("Unable to close db", esl.Error(err))
	}
}

func (z *badgerWrapper) View(f func(kvs kv_kvs.Kvs) error) error {
	return z.db.View(func(tx *badger.Txn) error {
		return f(kv_kvs_impl.NewBadger(z.name, z.logger, z.db, tx))
	})
}

func (z *badgerWrapper) Update(f func(kvs kv_kvs.Kvs) error) error {
	return z.db.Update(func(tx *badger.Txn) error {
		return f(kv_kvs_impl.NewBadger(z.name, z.logger, z.db, tx))
	})
}

func (z *badgerWrapper) Open(path string) error {
	return z.openWithPath(path)
}

func (z *badgerWrapper) SetLogger(logger esl.Logger) {
	z.logger = logger.With(esl.String("name", z.name), esl.String("path", z.path))
}

func (z *badgerWrapper) Path() string {
	return z.path
}

func (z *badgerWrapper) Delete() error {
	return z.db.DropAll()
}
