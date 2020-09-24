package cache

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Cache interface {
	// Get KVS of this cache. Note: Kvs may not actually store data.
	// It's depends on the cache policy. Caller must avoid rely on this Kvs as
	// a regular storage.
	Kvs() kv_kvs.Kvs
}

type Lifecycle interface {
	Cache

	// Open cache
	Open() error

	// Evict this cache if required (e.g. timeout)
	EvictIfRequired()

	// Flush and close this cache
	Close()
}

type noCache struct {
}

func (z noCache) Kvs() kv_kvs.Kvs {
	return kv_kvs_impl.NewEmpty()
}

func NoCache() Cache {
	return &noCache{}
}

func newCache(info Info, infoPath, basePath string, factory kv_storage.Factory, logger esl.Logger, timeout time.Duration) Lifecycle {
	return &cacheImpl{
		info:     info,
		infoPath: infoPath,
		basePath: basePath,
		factory:  factory,
		logger:   logger,
		timeout:  timeout,
	}
}

type cacheImpl struct {
	info       Info
	infoPath   string
	basePath   string
	factory    kv_storage.Factory
	sto        kv_storage.Lifecycle
	logger     esl.Logger
	cacheMutex sync.Mutex
	timeout    time.Duration
}

func (z *cacheImpl) noLockOpen() (err error) {
	l := z.log()
	l.Debug("Open database")
	z.sto, err = z.factory.New(z.info.DatabaseName())
	l.Debug("Open database: done", esl.Error(err))
	return
}

func (z *cacheImpl) updateInfo() (err error) {
	l := z.log()

	z.info.Last = time.Now()
	z.info.Expire = z.info.Last.Add(z.timeout)

	l.Debug("Updating info", esl.Any("info", z.info))

	infoData, err := json.Marshal(&z.info)
	if err != nil {
		l.Debug("Unable to marshal", esl.Error(err))
		return err
	}

	err = ioutil.WriteFile(z.infoPath, infoData, 0600)
	l.Debug("Info updated", esl.Error(err))
	return
}

func (z *cacheImpl) Open() (err error) {
	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	l := z.log()
	if z.info.Version != CurrentVersion {
		l.Debug("Unsupported cache version", esl.Error(err))
		return nil
	}

	err = z.noLockOpen()
	if err != nil {
		l.Debug("Unable to open", esl.Error(err))
		if z.info.Expire.Before(time.Now()) {
			dbPath := filepath.Join(z.basePath, z.info.DatabaseName())
			l.Debug("The database could not open, probably locked. Try removing database as file", esl.String("dbName", dbPath))
			osErr := os.RemoveAll(dbPath)
			if osErr != nil {
				l.Debug("Unable to remove database file", esl.Error(osErr))
				return
			}

			l.Debug("Database file removed, try re-opening database")
			err = z.noLockOpen()
			if err != nil {
				l.Debug("Database could not be opened", esl.Error(err))
				return err
			}
		}
		return err
	}

	if z.info.Expire.Before(time.Now()) {
		l.Debug("Evict required, evict", esl.Time("expire", z.info.Expire))
		if err := z.sto.Delete(); err != nil {
			l.Debug("Unable to evict data", esl.Error(err))
			z.sto.Close()
			z.sto = nil
			return err
		}
	}

	l.Debug("Updating info")
	return z.updateInfo()
}

func (z *cacheImpl) log() esl.Logger {
	return z.logger.With(esl.String("namespace", z.info.Namespace),
		esl.String("name", z.info.Name))
}

func (z *cacheImpl) Kvs() kv_kvs.Kvs {
	if z.sto == nil {
		return kv_kvs_impl.NewEmpty()
	}
	return z.sto.Kvs()
}

func (z *cacheImpl) EvictIfRequired() {
	l := z.log()
	l.Debug("EvictIfRequired cache")
	if z.info.Expire.After(time.Now()) {
		l.Debug("Evict not required", esl.Time("expire", z.info.Expire))
		return
	}

	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	if z.sto == nil {
		if err := z.noLockOpen(); err != nil {
			dbPath := filepath.Join(z.basePath, z.info.DatabaseName())
			l.Debug("The database could not open, probably locked. Try removing database as file", esl.String("dbName", dbPath))
			osErr := os.RemoveAll(dbPath)
			l.Debug("Removed", esl.Error(osErr))
			return
		}

		l.Debug("Database found, delete all entries")
		if err := z.sto.Delete(); err != nil {
			l.Debug("Unable to open database", esl.Error(err))
		}
		return
	}

	z.sto.Close()
	if sto, ok := z.sto.(kv_storage.Lifecycle); ok {
		l.Debug("Try removing cache", esl.String("path", sto.Path()))
		err := sto.Delete()
		l.Debug("Removed", esl.Error(err))
	}
}

func (z *cacheImpl) Close() {
	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	l := z.log()
	l.Debug("Closing cache")
	if z.sto != nil {
		z.sto.Close()
		z.sto = nil
	}

	err := z.updateInfo()
	if err != nil {
		l.Debug("Unable to update info", esl.Error(err))
	}
}
