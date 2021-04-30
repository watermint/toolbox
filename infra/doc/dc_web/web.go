package dc_web

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
)

func WebDocuments(mc app_msg_container.Container) []dc_section.Document {
	return []dc_section.Document{
		Home(),
		NewCommands(mc),
	}
}
