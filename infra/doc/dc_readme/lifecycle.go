package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewLifecycle() dc_section.Section {
	return &secLifecycle{}
}

type secLifecycle struct {
	SecTitle               app_msg.Message
	HeadMaintenancePolicy  app_msg.Message
	DescMaintenancePolicy  app_msg.Message
	HeadSpecChange         app_msg.Message
	DescSpecChange         app_msg.Message
	HeadReleaseStatusCheck app_msg.Message
	DescReleaseStatusCheck app_msg.Message
}

func (z secLifecycle) Title() app_msg.Message {
	return z.SecTitle
}

func (z secLifecycle) Body(ui app_ui.UI) {
	ui.SubHeader(z.HeadMaintenancePolicy)
	ui.Break()
	ui.Info(z.DescMaintenancePolicy)
	ui.Break()

	ui.SubHeader(z.HeadSpecChange)
	ui.Break()
	ui.Info(z.DescSpecChange)
	ui.Break()

	ui.SubHeader(z.HeadReleaseStatusCheck)
	ui.Break()
	ui.Info(z.DescReleaseStatusCheck)
}
