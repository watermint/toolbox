package dc_command

import (
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type DocJsonInput struct {
	spec        rc_recipe.Spec
	Header      app_msg.Message
	FieldHeader app_msg.Message
}

func (z DocJsonInput) Title() app_msg.Message {
	return z.Header
}

func (z DocJsonInput) Body(ui app_ui.UI) {
	for _, text := range z.spec.JsonInput() {
		ui.SubHeader(z.FieldHeader.With("Name", text.Name()))
		ui.Info(text.Desc())
	}
}
