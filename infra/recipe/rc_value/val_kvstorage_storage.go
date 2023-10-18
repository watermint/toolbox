package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"os"
	"reflect"
)

func newValueKvStorageStorage(name string) rc_recipe.Value {
	v := &ValueKvStorageStorage{name: name}
	v.storage = kv_storage_impl.NewProxy(name, esl.Default())
	return v
}

type ValueKvStorageStorage struct {
	name     string
	filePath string
	storage  kv_storage.Storage
}

func (z *ValueKvStorageStorage) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.storage), nil
}

func (z *ValueKvStorageStorage) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*kv_storage.Storage)(nil)).Elem()) {
		return newValueKvStorageStorage(name)
	}
	return nil
}

func (z *ValueKvStorageStorage) Bind() interface{} {
	return nil
}

func (z *ValueKvStorageStorage) Init() (v interface{}) {
	return z.storage
}

func (z *ValueKvStorageStorage) ApplyPreset(v0 interface{}) {
	z.storage = v0.(kv_storage.Storage)
}

func (z *ValueKvStorageStorage) Apply() (v interface{}) {
	return z.storage
}

func (z *ValueKvStorageStorage) Debug() interface{} {
	return map[string]string{
		"name": z.name,
	}
}

func (z *ValueKvStorageStorage) Capture(ctl app_control.Control) (v interface{}, err error) {
	return
}

func (z *ValueKvStorageStorage) Restore(v es_json.Json, ctl app_control.Control) error {
	return nil
}

func (z *ValueKvStorageStorage) SpinUp(ctl app_control.Control) error {
	storage := z.storage.(kv_storage.Lifecycle)
	engine := kv_storage.KvsEngineBitcask

	switch {
	case ctl.Feature().Experiment(app.ExperimentKvsBadger):
		engine = kv_storage.KvsEngineBadger
	case ctl.Feature().Experiment(app.ExperimentKvsBadgerTurnstile):
		engine = kv_storage.KvsEngineBadgerTurnstile
	case ctl.Feature().Experiment(app.ExperimentKvsBitcask):
		engine = kv_storage.KvsEngineBitcask
	case ctl.Feature().Experiment(app.ExperimentKvsBitcaskTurnstile):
		engine = kv_storage.KvsEngineBitcaskTurnstile
	case ctl.Feature().Experiment(app.ExperimentKvsSqlite):
		engine = kv_storage.KvsEngineSqlite
	case ctl.Feature().Experiment(app.ExperimentKvsSqliteTurnstile):
		engine = kv_storage.KvsEngineSqliteTurnstile
	}
	if p, ok := storage.(kv_storage.Proxy); ok {
		p.SetEngine(engine)
	} else {
		ctl.Log().Error("Unable to set engine", esl.Any("engine", engine))
	}
	storage.SetLogger(ctl.Log())
	return storage.Open(ctl.Workspace().KVS())
}

func (z *ValueKvStorageStorage) SpinDown(ctl app_control.Control) error {
	l := ctl.Log()
	if lc, ok := z.storage.(kv_storage.Lifecycle); ok {
		l.Debug("Delete storage", esl.String("path", lc.Path()))
		if err := lc.Delete(); err != nil {
			l.Debug("Unable to delete storage", esl.Error(err))
			return err
		}
	} else {
		l.Debug("Skip deleting")
	}
	z.storage.Close()
	if lc, ok := z.storage.(kv_storage.Lifecycle); ok {
		l.Debug("Delete storage", esl.String("path", lc.Path()))
		if err := lc.Delete(); err != nil {
			l.Debug("Unable to delete storage", esl.Error(err))
			return err
		}
		l.Debug("Remove storage", esl.String("path", lc.Path()))
		if err := os.RemoveAll(lc.Path()); err != nil {
			l.Debug("Unable to remove storage", esl.Error(err))
			// fall through, just ignore the error
		}
		return nil
	} else {
		l.Debug("Skip deleting")
	}
	return nil
}
