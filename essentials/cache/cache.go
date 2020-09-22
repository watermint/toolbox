package cache

import "github.com/watermint/toolbox/essentials/kvs/kv_kvs"

type Cache interface {
	// Create new kvs
	NewKvs(name string) kv_kvs.Kvs

	// Evict this cache
	Evict()

	// Flush and close this cache
	Close()
}
