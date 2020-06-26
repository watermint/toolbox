package kv_storage_impl

import (
	"errors"
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_badger"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type MsgStorage struct {
	ErrorThisCommandMayNotWorkOnWin32 app_msg.Message
}

var (
	MStorage = app_msg.Apply(&MsgStorage{}).(*MsgStorage)
)

func NewWithPath(ctl app_control.Control, path string) (kv_storage.Storage, error) {
	name := filepath.Base(path)
	bw := &badgerWrapper{ctl: ctl, name: name}
	if err := bw.openWithPath(name, path); err != nil {
		return nil, err
	}
	return bw, nil
}

func New(name string) kv_storage.Storage {
	bw := &badgerWrapper{name: name}
	return bw
}

type badgerWrapper struct {
	ctl    app_control.Control
	path   string
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
	l := z.ctl.Log().With(esl.String("name", z.name))
	l.Debug("Closing database")
	z.closed = true
	err := z.db.Close()
	l.Debug("Database closed", esl.Error(err))

	// #323 : remove storage data on close unless debug option enabled
	if z.ctl.Feature().IsDebug() {
		l.Debug("Skip removing database")
		return
	}

	l.Debug("Trying to clean up database", esl.String("path", z.path))
	if err := os.RemoveAll(z.path); err != nil {
		l.Debug("Unable to delete database", esl.Error(err))
	} else {
		l.Debug("The database removed")
	}
}

func (z *badgerWrapper) openWithPath(name, path string) (err error) {
	l := z.ctl.Log().With(esl.String("name", name))
	z.name = name
	z.path = path

	l = l.With(esl.String("path", path))
	l.Debug("Open database")
	opts := badger.DefaultOptions(path)
	opts = opts.WithLogger(lgw_badger.New(l))
	opts = opts.WithMaxBfCacheSize(4 * 1_048_576)
	opts = opts.WithMaxCacheSize(16 * 1_048_576)
	opts = opts.WithMaxTableSize(16 * 1_048_576)
	opts = opts.WithNumCompactors(1)
	opts = opts.WithNumMemtables(1)
	if z.ctl.Feature().BudgetMemory() == app_budget.BudgetLow {
		opts = opts.WithInMemory(false)
		opts = opts.WithKeepL0InMemory(false)
	}
	opts = opts.WithTableLoadingMode(options.FileIO)

	// Use lesser ValueLogFileSize for Windows 32bit environment
	if app.IsWindows() && runtime.GOARCH == "386" {
		opts = opts.WithValueLogFileSize(2 << 20)
	}

	z.db, err = badger.Open(opts)
	if err != nil {
		l.Debug("Unable to open database", esl.Error(err))
		// Temporary workaround:
		// https://github.com/watermint/toolbox/issues/297
		// Win 64 bit / GOARCH=386 : MapViewOfFile: Not enough memory resources are available to process this command
		// Win 32 bit / GOARCH=386 : MapViewOfFile: The parameter is incorrect.
		// This look like: https://github.com/dgraph-io/badger/issues/1072
		if strings.Contains(err.Error(), "MapViewOfFile") {
			l.Debug("Memory map error", esl.Error(err))
			if app.IsWindows() && runtime.GOARCH == "386" {
				z.ctl.UI().Error(MStorage.ErrorThisCommandMayNotWorkOnWin32)
				return errors.New("this command may not work on 32 bit windows")
			}
		}
		return err
	}
	return nil
}

func (z *badgerWrapper) init(name string) (err error) {
	l := z.ctl.Log().With(esl.String("name", name))
	kvsBasePath, err := z.ctl.Workspace().Descendant("kvs")
	if err != nil {
		l.Debug("Unable to create kvs folder", esl.Error(err))
		return err
	}
	path := filepath.Join(kvsBasePath, es_filepath.Escape(name))

	return z.openWithPath(name, path)
}
