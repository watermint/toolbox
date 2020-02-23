package kv_storage_impl

import (
	"errors"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"go.uber.org/zap"
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
	path = filepath.Join(path, ut_filepath.Escape(name))

	l = l.With(zap.String("path", path))
	l.Debug("Open database")
	opts := badger.DefaultOptions(path)
	opts.Logger = &badgerLogger{l: l}
	z.db, err = badger.Open(opts)
	if err != nil {
		l.Debug("Unable to open database", zap.Error(err))
		// Temporary workaround:
		// https://github.com/watermint/toolbox/issues/297
		if strings.Contains(err.Error(), "MapViewOfFile: Not enough memory resources are available to process this command") {
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
