package kv_storage

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_transaction"
)

// Storage interface.
type Storage interface {
	// Open KVS storage
	Open(ctl app_control.Control) error

	// Close KVS storage
	Close()

	// Read only transaction
	View(f func(tx kv_transaction.Transaction) error) error

	// Update transaction
	Update(f func(tx kv_transaction.Transaction) error) error

	// Read-write transaction
	Batch(f func(tx kv_transaction.Transaction) error) error
}
