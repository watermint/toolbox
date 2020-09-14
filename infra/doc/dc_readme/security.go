package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewSecurity() dc_section.Section {
	return app_msg.Apply(&Security{}).(dc_section.Section)
}

type Security struct {
	HeaderTitle                   app_msg.Message
	HeaderInformationNotCollected app_msg.Message
	HeaderSensitiveData           app_msg.Message
	BodyInformationNotCollected   app_msg.Message
	BodySensitiveData             app_msg.Message
}

func (z Security) Title() app_msg.Message {
	return z.HeaderTitle
}

func (z Security) Body(ui app_ui.UI) {
	ui.SubHeader(z.HeaderInformationNotCollected)
	ui.Info(z.BodyInformationNotCollected)
	ui.Break()
	ui.SubHeader(z.HeaderSensitiveData)
	ui.Info(z.BodySensitiveData)
}
