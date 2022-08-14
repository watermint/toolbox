package kv_storage

import (
	"errors"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/log/esl"
)

var (
	ErrorStorageLocked = errors.New("storage locked")
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
	New(name string) (Lifecycle, error)

	// Close all storages which created thru this factory.
	Close()
}

const (
	KvsEngineBitCask KvsEngine = iota
	KvsEngineSqlite
)

type KvsEngine int

type Proxy interface {
	// SetEngine apply new engine type from next call of Open
	SetEngine(engine KvsEngine)
}

// Lifecycle Storage with lifecycle control
type Lifecycle interface {
	Storage

	// Open KVS storage. Returns ErrorStorageLocked if the storage
	// used by the another process.
	Open(path string) error

	// Update logger
	SetLogger(logger esl.Logger)

	// KVS path
	Path() string

	// Delete this KVS
	Delete() error
}
