package dc_command

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewInstall() dc_section.Section {
	return &Install{}
}

type Install struct {
	Header      app_msg.Message
	Instruction app_msg.Message
}

func (z Install) Title() app_msg.Message {
	return z.Header
}

func (z Install) Body(ui app_ui.UI) {
	ui.Info(z.Instruction)
}
