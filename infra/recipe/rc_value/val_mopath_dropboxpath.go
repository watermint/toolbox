package rc_value

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueMoPathDropboxPath(name string) Value {
	v := &ValueMoPathDropboxPath{name: name}
	v.path = mo_path.NewDropboxPath("")
	return v
}

type ValueMoPathDropboxPath struct {
	name     string
	filePath string
	path     mo_path.DropboxPath
}

func (z *ValueMoPathDropboxPath) ValueText() string {
	return z.filePath
}

func (z *ValueMoPathDropboxPath) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	if t.Implements(reflect.TypeOf((*mo_path.Path)(nil)).Elem()) {
		return newValueMoPathDropboxPath(name)
	}
	return nil
}

func (z *ValueMoPathDropboxPath) Bind() interface{} {
	return &z.filePath
}

func (z *ValueMoPathDropboxPath) Init() (v interface{}) {
	return z.path
}

func (z *ValueMoPathDropboxPath) ApplyPreset(v0 interface{}) {
	z.path = v0.(mo_path.DropboxPath)
	z.filePath = z.path.Path()
}

func (z *ValueMoPathDropboxPath) Apply() (v interface{}) {
	z.path = mo_path.NewDropboxPath(z.filePath)
	return z.path
}

func (z *ValueMoPathDropboxPath) Debug() interface{} {
	return map[string]string{
		"path": z.filePath,
	}
}

func (z *ValueMoPathDropboxPath) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueMoPathDropboxPath) SpinDown(ctl app_control.Control) error {
	return nil
}
