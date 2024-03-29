package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"os"
	"path/filepath"
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
	return es_reflect.Key(z.storage), nil
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
	engine := ctl.Feature().KvsEngine()
	if p, ok := storage.(kv_storage.Proxy); ok {
		p.SetEngine(engine)
	} else {
		ctl.Log().Error("Unable to set engine", esl.Any("engine", engine))
	}
	storage.SetLogger(ctl.Log())
	storagePath := filepath.Join(ctl.Workspace().KVS(), es_filepath.Escape(z.name))
	return storage.Open(storagePath)
}

func (z *ValueKvStorageStorage) SpinDown(ctl app_control.Control) error {
	l := ctl.Log()
	l.Debug("Close storage")
	z.storage.Close()

	if lc, ok := z.storage.(kv_storage.Lifecycle); ok {
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
