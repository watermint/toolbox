package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_json"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueDaJsonInput(recipe interface{}, name string) rc_recipe.Value {
	return &ValueDaJsonInput{
		recipe:  recipe,
		jsInput: da_json.NewInput(name, recipe),
		name:    name,
		path:    "",
	}
}

type ValueDaJsonInput struct {
	recipe  interface{}
	jsInput da_json.JsonInput
	name    string
	path    string
}

func (z ValueDaJsonInput) ValueText() string {
	return ""
}

func (z *ValueDaJsonInput) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*da_json.JsonInput)(nil)).Elem()) {
		return newValueDaJsonInput(recipe, name)
	}
	return nil
}

func (z *ValueDaJsonInput) Bind() interface{} {
	return &z.path
}

func (z *ValueDaJsonInput) Init() (v interface{}) {
	return z.jsInput
}

func (z *ValueDaJsonInput) ApplyPreset(v0 interface{}) {
	z.jsInput = v0.(da_json.JsonInput)
	if z.jsInput.FilePath() != "" {
		z.path = z.jsInput.FilePath()
	}
}

func (z *ValueDaJsonInput) Apply() (v interface{}) {
	l := esl.Default()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.path)
	if err != nil {
		p = z.path
		l.Debug("Unable to format", esl.String("path", z.path), esl.Error(err))
	}

	if p != "" {
		z.jsInput.SetFilePath(p)
	}
	return z.jsInput
}

func (z *ValueDaJsonInput) Debug() interface{} {
	return z.jsInput.Debug()
}

func (z *ValueDaJsonInput) Capture(ctl app_control.Control) (v interface{}, err error) {
	filePath := z.path

	if z.path == "" {
		filePath = z.jsInput.FilePath()
	}
	return captureFile(ctl, filePath, func(path string) {
		z.path = path
		z.jsInput.SetFilePath(path)
	})
}

func (z *ValueDaJsonInput) Restore(v es_json.Json, ctl app_control.Control) error {
	return restoreFile(v, ctl, func(path string) {
		z.jsInput.SetFilePath(path)
	})
}

func (z *ValueDaJsonInput) SpinUp(ctl app_control.Control) error {
	if z.jsInput.FilePath() == "" {
		return ErrorMissingRequiredOption
	} else {
		return z.jsInput.Open(ctl)
	}
}

func (z *ValueDaJsonInput) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDaJsonInput) Spec() (typeName string, typeAttr interface{}) {
	return z.jsInput.Spec().Name(), nil
}

func (z *ValueDaJsonInput) JsonInput() (js da_json.JsonInput, valid bool) {
	return z.jsInput, true
}
