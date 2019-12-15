package kv_transaction

import "github.com/watermint/toolbox/infra/kvs/kv_kvs"

type Transaction interface {
	Kvs(name string) (kvs kv_kvs.Kvs, err error)
	ForEach(f func(name string, kvs kv_kvs.Kvs) error) error
}
