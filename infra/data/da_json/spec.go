package da_json

import (
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type JsonInputSpec interface {
	Name() string
	Desc() app_msg.Message
	Doc(ui app_ui.UI) *dc_recipe.DocJsonInput
}

func NewJsonSpec(name string, recipe interface{}) JsonInputSpec {
	return &jsSpec{
		name:   name,
		recipe: recipe,
	}
}

type jsSpec struct {
	name   string
	recipe interface{}
}

func (z jsSpec) Name() string {
	return z.name
}

func (z jsSpec) Desc() app_msg.Message {
	return app_msg.CreateMessage(es_reflect.Key(app.Pkg, z.recipe) + ".json_input." + strcase.ToSnake(z.name) + ".desc")
}

func (z jsSpec) Doc(ui app_ui.UI) *dc_recipe.DocJsonInput {
	return &dc_recipe.DocJsonInput{
		Name: z.name,
		Desc: ui.Text(z.Desc()),
	}
}
