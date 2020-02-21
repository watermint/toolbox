package kv_storage_impl

import (
	"github.com/dgraph-io/badger"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"go.uber.org/zap"
	"path/filepath"
)

func New(name string) kv_storage.Storage {
	bw := &badgerWrapper{name: name}
	return bw
}

type badgerWrapper struct {
	ctl    app_control.Control
	name   string
	db     *badger.DB
	closed bool
}

func (z *badgerWrapper) Open(ctl app_control.Control) error {
	z.ctl = ctl
	return z.init(z.name)
}

func (z *badgerWrapper) View(f func(kv kv_kvs.Kvs) error) error {
	return z.db.View(func(tx *badger.Txn) error {
		return f(kv_kvs_impl.New(z.ctl, z.db, tx))
	})
}

func (z *badgerWrapper) Update(f func(kv kv_kvs.Kvs) error) error {
	return z.db.Update(func(tx *badger.Txn) error {
		return f(kv_kvs_impl.New(z.ctl, z.db, tx))
	})
}

func (z *badgerWrapper) Close() {
	l := z.ctl.Log().With(zap.String("name", z.name))
	l.Debug("Closing database")
	z.closed = true
	err := z.db.Close()
	l.Debug("Database closed", zap.Error(err))
}

func (z *badgerWrapper) init(name string) (err error) {
	l := z.ctl.Log().With(zap.String("name", name))
	path, err := z.ctl.Workspace().Descendant("kvs")
	if err != nil {
		l.Debug("Unable to create kvs folder", zap.Error(err))
		return err
	}
	path = filepath.Join(path, ut_filepath.Escape(name)+".db")

	l = l.With(zap.String("path", path))
	l.Debug("Open database")
	z.db, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		l.Debug("Unable to open database", zap.Error(err))
		return err
	}
	z.name = name
	return nil
}
