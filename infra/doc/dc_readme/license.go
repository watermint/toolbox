package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewLicense() dc_section.Section {
	return &License{}
}

type License struct {
	HeaderTitle        app_msg.Message
	BodyLicense        app_msg.Message
	BodyLicenseRemarks app_msg.Message
	BodyLicenseQuote   app_msg.Message
}

func (z License) Title() app_msg.Message {
	return z.HeaderTitle
}

func (z License) Body(ui app_ui.UI) {
	ui.Info(z.BodyLicense)
	ui.Break()
	ui.Info(z.BodyLicenseRemarks)
	ui.Quote(z.BodyLicenseQuote)
}
