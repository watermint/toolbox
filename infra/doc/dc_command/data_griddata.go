package dc_command

import (
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewGridDataInput(spec rc_recipe.Spec) dc_section.Section {
	return &DocGridDataInput{
		spec: spec,
	}
}

type DocGridDataInput struct {
	spec           rc_recipe.Spec
	Header         app_msg.Message
	HeaderGridData app_msg.Message
}

func (z DocGridDataInput) Title() app_msg.Message {
	return z.Header
}

func (z DocGridDataInput) gridData(ui app_ui.UI, gd da_griddata.GridDataInputSpec) {
	ui.SubHeader(z.HeaderGridData.With("GridData", gd.Name()))
	ui.Break()
	ui.Info(gd.Desc())
}

func (z DocGridDataInput) Body(ui app_ui.UI) {
	for _, gd := range z.spec.GridDataInput() {
		z.gridData(ui, gd)
	}
}

func NewGridDataOutput(spec rc_recipe.Spec) dc_section.Section {
	return &DocGridDataOutput{
		spec: spec,
	}
}

type DocGridDataOutput struct {
	spec           rc_recipe.Spec
	Header         app_msg.Message
	HeaderGridData app_msg.Message
}

func (z DocGridDataOutput) Title() app_msg.Message {
	return z.Header
}

func (z DocGridDataOutput) gridData(ui app_ui.UI, gd da_griddata.GridDataOutputSpec) {
	ui.SubHeader(z.HeaderGridData.With("GridData", gd.Name()))
	ui.Break()
	ui.Info(gd.Desc())
}

func (z DocGridDataOutput) Body(ui app_ui.UI) {
	for _, gd := range z.spec.GridDataOutput() {
		z.gridData(ui, gd)
	}
}
