package dc_readme

import (
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewUsage() dc_section.Section {
	return &Usage{}
}

type Usage struct {
	UsageHeader app_msg.Message
	UsageBody   app_msg.Message
}

func (z Usage) Title() app_msg.Message {
	return z.UsageHeader
}

func (z Usage) Body(ui app_ui.UI) {
	ui.Info(z.UsageBody)
	ui.Break()

	bodyUsage := app_ui.MakeConsoleDemo(ui.Messages(), func(cui app_ui.UI) {
		cat := app_catalogue.Current()
		rg := cat.RootGroup()
		rg.PrintUsage(cui, "./tbx", "xx.x.xxx")
	})
	ui.Code("% ./tbx\n" + bodyUsage)
}
