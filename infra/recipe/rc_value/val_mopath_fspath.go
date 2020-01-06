package rc_value

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueMoPathFileSystemPath(name string) rc_recipe.Value {
	v := &ValueMoPathFileSystemPath{name: name}
	v.path = mo_path.NewFileSystemPath("")
	return v
}

type ValueMoPathFileSystemPath struct {
	name     string
	filePath string
	path     mo_path.FileSystemPath
}

func (z ValueMoPathFileSystemPath) ValueText() string {
	return z.filePath
}

func (z *ValueMoPathFileSystemPath) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*mo_path.Path)(nil)).Elem()) {
		return newValueMoPathFileSystemPath(name)
	}
	return nil
}

func (z *ValueMoPathFileSystemPath) Bind() interface{} {
	return &z.filePath
}

func (z *ValueMoPathFileSystemPath) Init() (v interface{}) {
	return z.path
}

func (z *ValueMoPathFileSystemPath) ApplyPreset(v0 interface{}) {
	z.path = v0.(mo_path.FileSystemPath)
	z.filePath = z.path.Path()
}

func (z *ValueMoPathFileSystemPath) Apply() (v interface{}) {
	z.path = mo_path.NewFileSystemPath(z.filePath)
	return z.path
}

func (z *ValueMoPathFileSystemPath) Debug() interface{} {
	return map[string]string{
		"path": z.filePath,
	}
}

func (z *ValueMoPathFileSystemPath) SpinUp(ctl app_control.Control) error {
	return nil
}

func (z *ValueMoPathFileSystemPath) SpinDown(ctl app_control.Control) error {
	return nil
}
