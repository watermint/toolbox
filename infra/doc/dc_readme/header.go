package dc_readme

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewHeader(forPublish bool) dc_section.Section {
	return &Header{
		publish: forPublish,
	}
}

type Header struct {
	// True when the doc is for publish on the GitHub project root.
	publish     bool
	HeaderTitle app_msg.Message
	HeaderBody  app_msg.Message
}

func (z Header) Title() app_msg.Message {
	return z.HeaderTitle.With("AppName", app.Name)
}

func (z Header) Body(ui app_ui.UI) {
	if z.publish {
		ui.Info(app_msg.Raw(app.ProjectStatusBadge))
		ui.Info(app_msg.Raw(app.ProjectLogo))
		ui.Break()
	}
	ui.Info(z.HeaderBody)
}
