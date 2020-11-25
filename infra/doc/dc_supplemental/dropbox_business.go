package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgDropboxBusiness struct {
	Title    app_msg.Message
	Overview app_msg.Message

	MemberTitle             app_msg.Message
	MemberCommandHeaderName app_msg.Message
	MemberCommandHeaderDesc app_msg.Message
	MemberInfoCommands      app_msg.Message
	MemberInfoOverview      app_msg.Message
	MemberBasicCommands     app_msg.Message
	MemberBasicOverview     app_msg.Message
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
	cat := app_catalogue.Current()
	//memberClearExternalid := cat.RecipeSpec("member clear externalid")
	memberList := cat.RecipeSpec("member list")
	memberDelete := cat.RecipeSpec("member delete")
	memberDetach := cat.RecipeSpec("member detach")
	memberInvite := cat.RecipeSpec("member invite")
	memberFolderList := cat.RecipeSpec("member folder list")
	memberUpdateEmail := cat.RecipeSpec("member update email")
	memberUpdateProfile := cat.RecipeSpec("member update profile")
	memberQuotaUpdate := cat.RecipeSpec("member quota update")
	memberQuotaList := cat.RecipeSpec("member quota list")
	memberQuotaUsage := cat.RecipeSpec("member quota usage")
	teamActivityUser := cat.RecipeSpec("team activity user")

	ui.SubHeader(MDropboxBusiness.MemberInfoCommands)
	ui.Info(MDropboxBusiness.MemberInfoOverview)

	dropboxBusinessCommandTable(ui, "member info commands", []rc_recipe.Spec{
		memberList,
		memberFolderList,
		memberQuotaList,
		memberQuotaUsage,
		teamActivityUser,
	})

	ui.SubHeader(MDropboxBusiness.MemberBasicCommands)
	ui.Info(MDropboxBusiness.MemberBasicOverview)

	dropboxBusinessCommandTable(ui, "management commands", []rc_recipe.Spec{
		memberInvite,
		memberDelete,
		memberDetach,
		memberUpdateEmail,
		memberUpdateProfile,
		memberQuotaUpdate,
	})

}

func dropboxBusinessCommandTable(ui app_ui.UI, name string, commands []rc_recipe.Spec) {
	t := ui.InfoTable(name)
	t.Header(MDropboxBusiness.MemberCommandHeaderName, MDropboxBusiness.MemberCommandHeaderDesc)

	for _, c := range commands {
		t.Row(c.CliNameRef(""), c.Title())
	}
	t.Flush()
}
