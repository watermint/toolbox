package dc_web

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_readme"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
)

func NewCommands(mc app_msg_container.Container) dc_section.Document {
	return &homeCommands{
		mc: mc,
	}
}

type homeCommands struct {
	mc   app_msg_container.Container
	Desc app_msg.Message
}

func (z homeCommands) DocId() dc_index.DocId {
	return dc_index.DocWebCommandTableOfContent
}

func (z homeCommands) DocDesc() app_msg.Message {
	return z.Desc
}

func (z homeCommands) Sections() []dc_section.Section {
	return []dc_section.Section{
		dc_readme.NewCommand(true, dc_index.MediaWeb, z.mc),
	}
}
