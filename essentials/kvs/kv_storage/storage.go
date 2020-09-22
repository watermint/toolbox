package kv_storage

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
)

// Storage interface.
type Storage interface {
	// Close KVS storage
	Close()

	// Read only transaction
	View(f func(kvs kv_kvs.Kvs) error) error

	// Read-write transaction
	Update(f func(kvs kv_kvs.Kvs) error) error

	// Use direct operation
	Kvs() kv_kvs.Kvs
}

type Factory interface {
	// Create new storage
	New(name string) (Storage, error)

	// Close all storages which created thru this factory.
	Close()
}
