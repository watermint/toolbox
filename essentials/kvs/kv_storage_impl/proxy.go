package kv_storage_impl

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs_impl"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewProxy(name string, logger esl.Logger) kv_storage.Lifecycle {
	return &proxyImpl{
		engine: kv_storage.KvsEngineBitCask,
		name:   name,
		logger: logger,
	}
}

type proxyImpl struct {
	engine  kv_storage.KvsEngine
	name    string
	storage kv_storage.Lifecycle
	kvs     kv_kvs.Kvs
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

func (z *proxyImpl) Kvs() kv_kvs.Kvs {
	return z.kvs
}

func (z *proxyImpl) newStorage() kv_storage.Lifecycle {
	switch z.engine {
	case kv_storage.KvsEngineBitCask, kv_storage.KvsEngineBitcaskTurnstile:
		return InternalNewBitcask(z.name, z.logger)
	case kv_storage.KvsEngineSqlite, kv_storage.KvsEngineSqliteTurnstile:
		return InternalNewSqlite(z.name, z.logger)
	default:
		// fallback
		return InternalNewBitcask(z.name, z.logger)
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
	switch z.engine {
	case kv_storage.KvsEngineSqliteTurnstile, kv_storage.KvsEngineBitcaskTurnstile:
		z.kvs = kv_kvs_impl.NewTurnstile(z.storage.Kvs())
	default:
		z.kvs = z.Kvs()
	}
	return z.storage.Open(path)
}
