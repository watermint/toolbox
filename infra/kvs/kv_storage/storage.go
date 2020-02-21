package kv_storage

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
)

// Storage interface.
type Storage interface {
	// Open KVS storage
	Open(ctl app_control.Control) error

	// Close KVS storage
	Close()

	// Read only transaction
	View(f func(kvs kv_kvs.Kvs) error) error

	// Read-write transaction
	Update(f func(kvs kv_kvs.Kvs) error) error
}
