package da_text

import (
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type TextInputSpec interface {
	Name() string
	Desc() app_msg.Message
	Doc(ui app_ui.UI) *dc_recipe.DocTextInput
}

func NewInputSpec(name string, recipe interface{}) TextInputSpec {
	return &inSpec{
		name:   name,
		recipe: recipe,
	}
}

type inSpec struct {
	name   string
	recipe interface{}
}

func (z inSpec) Name() string {
	return z.name
}

func (z inSpec) Desc() app_msg.Message {
	return app_msg.CreateMessage(es_reflect.Key(app.Pkg, z.recipe) + ".text_input." + es_case.ToLowerSnakeCase(z.name) + ".desc")
}

func (z inSpec) Doc(ui app_ui.UI) *dc_recipe.DocTextInput {
	return &dc_recipe.DocTextInput{
		Name: z.name,
		Desc: ui.Text(z.Desc()),
	}
}
