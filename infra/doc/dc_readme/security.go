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
	HeaderTitle                   app_msg.Message
	HeaderInformationNotCollected app_msg.Message
	HeaderSensitiveData           app_msg.Message
	BodyInformationNotCollected   app_msg.Message
	BodySensitiveData             app_msg.Message
}

func (z SecurityDesc) Title() app_msg.Message {
	return z.HeaderTitle
}

func (z SecurityDesc) Body(ui app_ui.UI) {
	ui.SubHeader(z.HeaderInformationNotCollected)
	ui.Info(z.BodyInformationNotCollected)
	ui.Break()
	ui.SubHeader(z.HeaderSensitiveData)
	ui.Info(z.BodySensitiveData)
}
