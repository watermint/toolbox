package kv_storage_impl

import (
	"errors"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"go.uber.org/zap"
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
	l := z.ctl.Log().With(zap.String("name", z.name))
	l.Debug("Closing database")
	z.closed = true
	err := z.db.Close()
	l.Debug("Database closed", zap.Error(err))

	// #323 : remove storage data on close unless debug option enabled
	if z.ctl.IsDebug() {
		l.Debug("Skip removing database")
		return
	}

	l.Debug("Trying to clean up database", zap.String("path", z.path))
	if err := os.RemoveAll(z.path); err != nil {
		l.Debug("Unable to delete database", zap.Error(err))
	} else {
		l.Debug("The database removed")
	}
}

func (z *badgerWrapper) init(name string) (err error) {
	l := z.ctl.Log().With(zap.String("name", name))
	kvsBasePath, err := z.ctl.Workspace().Descendant("kvs")
	if err != nil {
		l.Debug("Unable to create kvs folder", zap.Error(err))
		return err
	}
	z.path = filepath.Join(kvsBasePath, ut_filepath.Escape(name))

	l = l.With(zap.String("path", z.path))
	l.Debug("Open database")
	opts := badger.DefaultOptions(z.path)
	opts = opts.WithLogger(&badgerLogger{l.WithOptions(zap.AddCallerSkip(1))})
	opts = opts.WithMaxCacheSize(32 * 1_048_576) // 32MB
	opts = opts.WithNumCompactors(1)
	opts = opts.WithTableLoadingMode(options.FileIO)

	// Use lesser ValueLogFileSize for Windows 32bit environment
	if app.IsWindows() && runtime.GOARCH == "386" {
		opts = opts.WithValueLogFileSize(2 << 20)
		opts = opts.WithNumMemtables(2)
	}

	z.db, err = badger.Open(opts)
	if err != nil {
		l.Debug("Unable to open database", zap.Error(err))
		// Temporary workaround:
		// https://github.com/watermint/toolbox/issues/297
		// Win 64 bit / GOARCH=386 : MapViewOfFile: Not enough memory resources are available to process this command
		// Win 32 bit / GOARCH=386 : MapViewOfFile: The parameter is incorrect.
		// This look like: https://github.com/dgraph-io/badger/issues/1072
		if strings.Contains(err.Error(), "MapViewOfFile") {
			l.Debug("Memory map error", zap.Error(err))
			if app.IsWindows() && runtime.GOARCH == "386" {
				z.ctl.UI().Error(MStorage.ErrorThisCommandMayNotWorkOnWin32)
				return errors.New("this command may not work on 32 bit windows")
			}
		}
		return err
	}
	z.name = name
	return nil
}

type badgerLogger struct {
	l *zap.Logger
}

func (z *badgerLogger) Errorf(f string, p ...interface{}) {
	z.l.Warn(fmt.Sprintf(f, p...), zap.String("level", "error"))
}

func (z *badgerLogger) Warningf(f string, p ...interface{}) {
	z.l.Debug(fmt.Sprintf(f, p...), zap.String("level", "warn"))
}

func (z *badgerLogger) Infof(f string, p ...interface{}) {
	z.l.Debug(fmt.Sprintf(f, p...), zap.String("level", "info"))
}

func (z *badgerLogger) Debugf(f string, p ...interface{}) {
	z.l.Debug(fmt.Sprintf(f, p...), zap.String("level", "debug"))
}
