package kv_transaction_impl

import (
	"github.com/etcd-io/bbolt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/infra/kvs/kv_transaction"
	"go.uber.org/zap"
)

func New(ctl app_control.Control, tx *bbolt.Tx) kv_transaction.Transaction {
	return &bboltWrapper{
		ctl: ctl,
		tx:  tx,
	}
}

type bboltWrapper struct {
	ctl app_control.Control
	tx  *bbolt.Tx
}

func (z *bboltWrapper) Kvs(name string) (kvs kv_kvs.Kvs, err error) {
	l := z.ctl.Log().With(zap.String("name", name))
	b := z.tx.Bucket([]byte(name))
	if b != nil {
		return kv_kvs_impl.New(z.ctl, name, b), nil
	}

	bk, err := z.tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		l.Debug("Unable to create bucket", zap.Error(err))
		return nil, err
	}
	return kv_kvs_impl.New(z.ctl, name, bk), nil
}

func (z *bboltWrapper) ForEach(f func(name string, kvs kv_kvs.Kvs) error) error {
	return z.tx.ForEach(func(bn []byte, b *bbolt.Bucket) error {
		return f(string(bn), kv_kvs_impl.New(z.ctl, string(bn), b))
	})
}
