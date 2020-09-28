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
	"reflect"
)

func newValueKvStorageStorage(name string) rc_recipe.Value {
	v := &ValueKvStorageStorage{name: name}
	v.storage = kv_storage_impl.NewStorage(name, esl.Default())
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

func (z *ValueKvStorageStorage) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
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
	storage.SetLogger(ctl.Log())
	return storage.Open(ctl.Workspace().KVS())
}

func (z *ValueKvStorageStorage) SpinDown(ctl app_control.Control) error {
	z.storage.Close()
	return nil
}
