package kv_storage_impl

import (
	"github.com/etcd-io/bbolt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/kvs/kv_transaction"
	"github.com/watermint/toolbox/infra/kvs/kv_transaction_impl"
	"go.uber.org/zap"
	"path/filepath"
	"time"
)

const (
	monitorInterval = 15 * time.Second
)

func New(ctl app_control.Control, name string) (kv_storage.Storage, error) {
	bw := &bboltWrapper{ctl: ctl}
	if err := bw.init(name); err != nil {
		return nil, err
	}
	return bw, nil
}

type bboltWrapper struct {
	ctl       app_control.Control
	name      string
	db        *bbolt.DB
	closed    bool
	lastStats bbolt.Stats
}

func (z *bboltWrapper) View(f func(tx kv_transaction.Transaction) error) error {
	return z.db.View(func(tx *bbolt.Tx) error {
		return f(kv_transaction_impl.New(z.ctl, tx))
	})
}

func (z *bboltWrapper) Update(f func(tx kv_transaction.Transaction) error) error {
	return z.db.Update(func(tx *bbolt.Tx) error {
		return f(kv_transaction_impl.New(z.ctl, tx))
	})
}

func (z *bboltWrapper) Batch(f func(tx kv_transaction.Transaction) error) error {
	return z.db.Batch(func(tx *bbolt.Tx) error {
		return f(kv_transaction_impl.New(z.ctl, tx))
	})
}

func (z *bboltWrapper) Close() {
	l := z.ctl.Log().With(zap.String("name", z.name))
	l.Debug("Closing database")
	z.closed = true
	err := z.db.Close()
	l.Debug("Database closed", zap.Error(err))
}

func (z *bboltWrapper) monitor() {
	z.lastStats = z.db.Stats()
	l := z.ctl.Log().With(zap.String("name", z.name))
	for {
		time.Sleep(monitorInterval)

		if z.closed {
			return
		}
		stats := z.db.Stats()
		diff := stats.Sub(&z.lastStats)
		z.lastStats = stats
		l.Debug("Database stats", zap.Any("stats", diff))
	}
}

func (z *bboltWrapper) init(name string) (err error) {
	l := z.ctl.Log().With(zap.String("name", name))
	path, err := z.ctl.Workspace().Descendant("kvs")
	if err != nil {
		l.Debug("Unable to create kvs folder", zap.Error(err))
		return err
	}
	path = filepath.Join(path, "name.db")

	l = l.With(zap.String("path", path))
	l.Debug("Open database")
	z.db, err = bbolt.Open(path, 0600, nil)
	if err != nil {
		l.Debug("Unable to open database", zap.Error(err))
		return err
	}
	z.name = name
	go z.monitor()
	return nil
}
