package dc_command

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewTextInput(spec rc_recipe.Spec) dc_section.Section {
	return &DocTextInput{
		spec: spec,
	}
}

type DocTextInput struct {
	spec        rc_recipe.Spec
	Header      app_msg.Message
	FieldHeader app_msg.Message
}

func (z DocTextInput) Title() app_msg.Message {
	return z.Header
}

func (z DocTextInput) Body(ui app_ui.UI) {
	for _, text := range z.spec.TextInput() {
		ui.SubHeader(z.FieldHeader.With("Name", text.Name()))
		ui.Info(text.Desc())
	}
}
