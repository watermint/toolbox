package rc_value

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueMoPathPath(name string) Value {
	v := &ValueMoPathPath{name: name}
	v.path = mo_path.NewDropboxPath("")
	return v
}

type ValueMoPathPath struct {
	name     string
	filePath string
	path     mo_path.DropboxPath
}

func (z *ValueMoPathPath) ValueText() string {
	return z.filePath
}

func (z *ValueMoPathPath) Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	if t.Implements(reflect.TypeOf((*mo_path.Path)(nil)).Elem()) {
		return newValueMoPathPath(name)
	}
	return nil
}

func (z *ValueMoPathPath) Bind() interface{} {
	return &z.filePath
}

func (z *ValueMoPathPath) Init() (v interface{}) {
	return z.path
}

func (z *ValueMoPathPath) Apply() (v interface{}) {
	z.path = mo_path.NewDropboxPath(z.filePath)
	return z.path
}

func (z *ValueMoPathPath) Debug() interface{} {
	return map[string]string{
		"path": z.filePath,
	}
}

func (z *ValueMoPathPath) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueMoPathPath) SpinDown(ctl app_control.Control) error {
	return nil
}
