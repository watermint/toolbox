package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewSecurity() dc_section.Document {
	return &docSecurity{}
}

type docSecurity struct {
	Desc app_msg.Message
}

func (z docSecurity) DocId() dc_index.DocId {
	return dc_index.DocRootSecurityAndPrivacy
}

func (z docSecurity) DocDesc() app_msg.Message {
	return z.Desc
}

func (z docSecurity) Sections() []dc_section.Section {
	return []dc_section.Section{
		NewSecuritySection(),
	}
}

func NewSecuritySection() dc_section.Section {
	return &SecurityDesc{}
}

type SecurityDesc struct {
	HeaderTitle          app_msg.Message
	BodyOverview         app_msg.Message
	HeaderDataProtection app_msg.Message
	BodyDataProtection   app_msg.Message
	HeaderUse            app_msg.Message
	BodyUse              app_msg.Message
	HeaderSharing        app_msg.Message
	BodySharing          app_msg.Message
}

func (z SecurityDesc) Title() app_msg.Message {
	return z.HeaderTitle
}

func (z SecurityDesc) Body(ui app_ui.UI) {
	ui.Info(z.BodyOverview)

	ui.SubHeader(z.HeaderDataProtection)
	ui.Info(z.BodyDataProtection)

	ui.SubHeader(z.HeaderUse)
	ui.Info(z.BodyUse)

	ui.SubHeader(z.HeaderSharing)
	ui.Info(z.BodySharing)
}
