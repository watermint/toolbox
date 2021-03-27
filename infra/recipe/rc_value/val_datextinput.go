package rc_value

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"reflect"
)

func newValueDaTextInput(recipe interface{}, name string) rc_recipe.Value {
	return &ValueDaTextInput{
		recipe:  recipe,
		txInput: da_text.NewTextInput(name, recipe),
		name:    name,
		path:    "",
	}
}

type ValueDaTextInput struct {
	recipe  interface{}
	txInput da_text.TextInput
	name    string
	path    string
}

func (z *ValueDaTextInput) ValueText() string {
	return ""
}

func (z *ValueDaTextInput) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*da_text.TextInput)(nil)).Elem()) {
		return newValueDaTextInput(recipe, name)
	}
	return nil
}

func (z *ValueDaTextInput) Bind() interface{} {
	return &z.path
}

func (z *ValueDaTextInput) Init() (v interface{}) {
	return z.txInput
}

func (z *ValueDaTextInput) ApplyPreset(v0 interface{}) {
	z.txInput = v0.(da_text.TextInput)
	if z.txInput.FilePath() != "" {
		z.path = z.txInput.FilePath()
	}
}

func (z *ValueDaTextInput) Apply() (v interface{}) {
	l := esl.Default()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.path)
	if err != nil {
		p = z.path
		l.Debug("Unable to format", esl.String("path", z.path), esl.Error(err))
	}

	if p != "" {
		z.txInput.SetFilePath(p)
	}
	return z.txInput
}

func (z *ValueDaTextInput) Debug() interface{} {
	return z.txInput.Debug()
}

func (z *ValueDaTextInput) Capture(ctl app_control.Control) (v interface{}, err error) {
	filePath := z.path

	if z.path == "" {
		filePath = z.txInput.FilePath()
	}
	return captureFile(ctl, filePath, func(path string) {
		z.path = path
		z.txInput.SetFilePath(path)
	})
}

func (z *ValueDaTextInput) Restore(v es_json.Json, ctl app_control.Control) error {
	return restoreFile(v, ctl, func(path string) {
		z.txInput.SetFilePath(path)
	})
}

func (z *ValueDaTextInput) SpinUp(ctl app_control.Control) error {
	if z.txInput.FilePath() == "" {
		return ErrorMissingRequiredOption
	} else {
		return z.txInput.Open(ctl)
	}
}

func (z *ValueDaTextInput) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDaTextInput) Spec() (typeName string, typeAttr interface{}) {
	return z.txInput.Spec().Name(), nil
}

func (z *ValueDaTextInput) TextInput() (tx da_text.TextInput, valid bool) {
	return z.txInput, true
}
