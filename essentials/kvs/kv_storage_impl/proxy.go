package kv_storage_impl

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"sync"
)

func NewProxy(name string, logger esl.Logger) kv_storage.Lifecycle {
	return &proxyImpl{
		engine: kv_storage.KvsEngineBadger,
		name:   name,
		logger: logger,
	}
}

type proxyImpl struct {
	mutex   sync.Mutex
	engine  kv_storage.KvsEngine
	name    string
	storage kv_storage.Lifecycle
	logger  esl.Logger
}

func (z *proxyImpl) SetEngine(engine kv_storage.KvsEngine) {
	z.engine = engine
}

func (z *proxyImpl) Delete() error {
	return z.storage.Delete()
}

func (z *proxyImpl) Path() string {
	return z.storage.Path()
}

func (z *proxyImpl) SetLogger(logger esl.Logger) {
	z.logger = logger
}

func (z *proxyImpl) newStorage() kv_storage.Lifecycle {
	l := z.logger
	switch z.engine {
	case kv_storage.KvsEngineBitcask:
		l.Debug("Use Bitcask")
		return newBitcask(z.name, z.logger)
	case kv_storage.KvsEngineBitcaskTurnstile:
		l.Debug("Use Bitcask Turnstile")
		return newTurnstileStorage(newBitcask(z.name, z.logger))
	case kv_storage.KvsEngineSqlite:
		l.Debug("Use Sqlite")
		return newSqlite(z.name, z.logger)
	case kv_storage.KvsEngineSqliteTurnstile:
		l.Debug("Use Sqlite Turnstile")
		return newTurnstileStorage(newSqlite(z.name, z.logger))
	case kv_storage.KvsEngineBadger:
		l.Debug("Use Badger")
		return newBadger(z.name, z.logger)
	case kv_storage.KvsEngineBadgerTurnstile:
		l.Debug("Use Badger Turnstile")
		return newTurnstileStorage(newBadger(z.name, z.logger))

	default:
		l.Debug("Use default engine: Badger")
		// fallback
		return newBadger(z.name, z.logger)
	}
}

func (z *proxyImpl) Close() {
	z.storage.Close()
}

func (z *proxyImpl) View(f func(kvs kv_kvs.Kvs) error) error {
	return z.storage.View(f)
}

func (z *proxyImpl) Update(f func(kvs kv_kvs.Kvs) error) error {
	return z.storage.Update(f)
}

func (z *proxyImpl) Open(path string) error {
	z.storage = z.newStorage()

	return z.storage.Open(path)
}
