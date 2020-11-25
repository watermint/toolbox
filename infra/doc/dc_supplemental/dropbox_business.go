package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgDropboxBusiness struct {
	Title    app_msg.Message
	Overview app_msg.Message

	MemberTitle         app_msg.Message
	MemberBasicCommands app_msg.Message
}

var (
	MDropboxBusiness = app_msg.Apply(&MsgDropboxBusiness{}).(*MsgDropboxBusiness)
)

type DropboxBusiness struct {
}

func (z DropboxBusiness) DocId() dc_index.DocId {
	return dc_index.DocSupplementalDropboxBusiness
}

func (z DropboxBusiness) Sections() []dc_section.Section {
	return []dc_section.Section{
		&DropboxBusinessMember{},
	}
}

type DropboxBusinessMember struct {
}

func (z DropboxBusinessMember) Title() app_msg.Message {
	return MDropboxBusiness.MemberTitle
}

func (z DropboxBusinessMember) Body(ui app_ui.UI) {
	ui.SubHeader(MDropboxBusiness.MemberBasicCommands)
}
