package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewRelease() dc_section.Section {
	return &Release{}
}

// Section of pointer to download
type Release struct {
	ReleaseHeader  app_msg.Message
	ReleaseBody    app_msg.Message
	HomebrewHeader app_msg.Message
	HomebrewDesc   app_msg.Message
}

func (z Release) Title() app_msg.Message {
	return z.ReleaseHeader
}

func (z Release) Body(ui app_ui.UI) {
	ui.Info(z.ReleaseBody)
	ui.Break()
	ui.SubHeader(z.HomebrewHeader)
	ui.Info(z.HomebrewDesc)
	ui.Code("brew tap watermint/toolbox\nbrew install toolbox")
}
