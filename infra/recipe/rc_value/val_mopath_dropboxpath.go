package rc_value

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueMoPathDropboxPath(name string) rc_recipe.Value {
	v := &ValueMoPathDropboxPath{name: name}
	v.path = mo_path.NewDropboxPath("")
	return v
}

type ValueMoPathDropboxPath struct {
	name     string
	filePath string
	path     mo_path.DropboxPath
}

func (z *ValueMoPathDropboxPath) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.path), nil
}

func (z *ValueMoPathDropboxPath) ValueText() string {
	return z.filePath
}

func (z *ValueMoPathDropboxPath) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*mo_path.DropboxPath)(nil)).Elem()) {
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
	l := esl.Default()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.filePath)
	if err != nil {
		p = z.filePath
		l.Debug("Unable to format", esl.String("path", z.filePath), esl.Error(err))
	}
	z.path = mo_path.NewDropboxPath(p)
	return z.path
}

func (z *ValueMoPathDropboxPath) Debug() interface{} {
	return map[string]string{
		"path": z.filePath,
	}
}

func (z *ValueMoPathDropboxPath) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.filePath, nil
}

func (z *ValueMoPathDropboxPath) Restore(v es_json.Json, ctl app_control.Control) error {
	if w, found := v.String(); found {
		z.filePath = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueMoPathDropboxPath) SpinUp(ctl app_control.Control) error {
	if !z.path.IsValid() {
		return ErrorMissingRequiredOption
	}
	return nil
}

func (z *ValueMoPathDropboxPath) SpinDown(ctl app_control.Control) error {
	return nil
}
