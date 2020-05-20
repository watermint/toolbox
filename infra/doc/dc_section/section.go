package dc_section

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type Section interface {
	Title() app_msg.Message
	Body(ui app_ui.UI)
}

func Document(mc app_msg_container.Container, sections ...Section) string {
	return app_ui.MakeMarkdown(mc, func(ui app_ui.UI) {
		for _, section := range sections {
			ui.Header(section.Title())
			section.Body(ui)
			ui.Break()
		}
	})
}
