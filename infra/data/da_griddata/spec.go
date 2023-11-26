package da_griddata

import (
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/islet/estring/ecase"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type GridDataInputSpec interface {
	Name() string
	Desc() app_msg.Message
	Doc(ui app_ui.UI) *dc_recipe.DocGridDataInput
}

type GridDataOutputSpec interface {
	Name() string
	Desc() app_msg.Message
	Doc(ui app_ui.UI) *dc_recipe.DocGridDataOutput
}

func NewInputSpec(recipe interface{}, name string) GridDataInputSpec {
	return &rdSpec{
		name:   name,
		recipe: recipe,
	}
}

type rdSpec struct {
	name   string
	recipe interface{}
}

func (z rdSpec) Name() string {
	return z.name
}

func (z rdSpec) Desc() app_msg.Message {
	return app_msg.CreateMessage(es_reflect.Key(app.Pkg, z.recipe) + ".grid_data_input." + ecase.ToLowerSnakeCase(z.name) + ".desc")
}

func (z rdSpec) Doc(ui app_ui.UI) *dc_recipe.DocGridDataInput {
	return &dc_recipe.DocGridDataInput{
		Name: z.Name(),
		Desc: ui.Text(z.Desc()),
	}
}

func NewOutputSpec(recipe interface{}, name string) GridDataOutputSpec {
	return &wrSpec{
		name:   name,
		recipe: recipe,
	}
}

type wrSpec struct {
	name   string
	recipe interface{}
}

func (z wrSpec) Name() string {
	return z.name
}

func (z wrSpec) Desc() app_msg.Message {
	return app_msg.CreateMessage(es_reflect.Key(app.Pkg, z.recipe) + ".grid_data_output." + ecase.ToLowerSnakeCase(z.name) + ".desc")
}

func (z wrSpec) Doc(ui app_ui.UI) *dc_recipe.DocGridDataOutput {
	return &dc_recipe.DocGridDataOutput{
		Name: z.Name(),
		Desc: ui.Text(z.Desc()),
	}
}
