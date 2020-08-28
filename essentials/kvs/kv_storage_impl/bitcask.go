package kv_storage_impl

import (
	"github.com/prologic/bitcask"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"path/filepath"
)

func InternalNewBitcask(name string) Storage {
	bc := &bcWrapper{
		name: name,
	}
	return bc
}

type bcWrapper struct {
	ctl  app_control.Control
	path string
	name string
	db   *bitcask.Bitcask
	kvs  kv_kvs.Kvs
}

func (z *bcWrapper) OpenWithPath(ctl app_control.Control, path string) error {
	z.name = filepath.Base(path)
	z.ctl = ctl
	return z.openWithPath(path)
}

func (z *bcWrapper) log() esl.Logger {
	return z.ctl.Log().With(esl.String("name", z.name), esl.String("path", z.path))
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

	if err != nil {
		l.Debug("Unable to open the database", esl.Error(err))
		return err
	}
	z.kvs = kv_kvs_impl.NewBitcask(z.name, z.ctl, z.db)

	return nil
}

func (z *bcWrapper) Open(ctl app_control.Control) error {
	z.ctl = ctl
	path := filepath.Join(z.ctl.Workspace().KVS(), es_filepath.Escape(z.name))
	return z.openWithPath(path)
}

func (z *bcWrapper) Close() {
	l := z.log()
	l.Debug("Close database")
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
