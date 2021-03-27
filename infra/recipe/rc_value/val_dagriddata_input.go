package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueDaGridDataInput(recipe interface{}, name string) rc_recipe.Value {
	return &ValueDaGridDataInput{
		recipe:  recipe,
		gdInput: da_griddata.NewInput(recipe, name),
		name:    name,
		path:    "",
	}
}

type ValueDaGridDataInput struct {
	recipe  interface{}
	gdInput da_griddata.GridDataInput
	name    string
	path    string
}

func (z *ValueDaGridDataInput) ValueText() string {
	return ""
}

func (z *ValueDaGridDataInput) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*da_griddata.GridDataInput)(nil)).Elem()) {
		return newValueDaGridDataInput(recipe, name)
	}
	return nil
}

func (z *ValueDaGridDataInput) Bind() interface{} {
	return &z.path
}

func (z *ValueDaGridDataInput) Init() (v interface{}) {
	return z.gdInput
}

func (z *ValueDaGridDataInput) ApplyPreset(v0 interface{}) {
	z.gdInput = v0.(da_griddata.GridDataInput)
	if z.gdInput.FilePath() != "" {
		z.path = z.gdInput.FilePath()
	}
}

func (z *ValueDaGridDataInput) Apply() (v interface{}) {
	l := esl.Default()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.path)
	if err != nil {
		p = z.path
		l.Debug("Unable to format", esl.String("path", z.path), esl.Error(err))
	}

	if p != "" {
		z.gdInput.SetFilePath(p)
	}
	return z.gdInput
}

func (z *ValueDaGridDataInput) Debug() interface{} {
	return z.gdInput.Debug()
}

func (z *ValueDaGridDataInput) Capture(ctl app_control.Control) (v interface{}, err error) {
	filePath := z.path

	if z.path == "" {
		filePath = z.gdInput.FilePath()
	}
	return captureFile(ctl, filePath, func(path string) {
		z.path = path
		z.gdInput.SetFilePath(path)
	})
}

func (z *ValueDaGridDataInput) Restore(v es_json.Json, ctl app_control.Control) error {
	return restoreFile(v, ctl, func(path string) {
		z.gdInput.SetFilePath(path)
	})
}

func (z *ValueDaGridDataInput) SpinUp(ctl app_control.Control) (err error) {
	if z.gdInput.FilePath() == "" {
		return ErrorMissingRequiredOption
	} else {
		return z.gdInput.Open(ctl)
	}
}

func (z *ValueDaGridDataInput) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDaGridDataInput) Spec() (typeName string, typeAttr interface{}) {
	return z.gdInput.Spec().Name(), nil
}

func (z *ValueDaGridDataInput) GridDataInput() (gd da_griddata.GridDataInput, valid bool) {
	return z.gdInput, true
}
