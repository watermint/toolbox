package kv_storage

import "github.com/watermint/toolbox/infra/kvs/kv_transaction"

// Storage interface.
type Storage interface {
	// Close KVS storage
	Close()

	// Read only transaction
	View(f func(tx kv_transaction.Transaction) error) error

	// Update transaction
	Update(f func(tx kv_transaction.Transaction) error) error

	// Read-write transaction
	Batch(f func(tx kv_transaction.Transaction) error) error
}
