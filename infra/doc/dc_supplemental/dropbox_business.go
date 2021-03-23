package dc_supplemental

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sort"
)

type MsgDropboxBusiness struct {
	Title    app_msg.Message
	Overview app_msg.Message

	CommandHeaderName app_msg.Message
	CommandHeaderDesc app_msg.Message

	MemberTitle         app_msg.Message
	MemberInfoCommands  app_msg.Message
	MemberInfoOverview  app_msg.Message
	MemberBasicCommands app_msg.Message
	MemberBasicOverview app_msg.Message

	GroupTitle       app_msg.Message
	GroupMgmtTitle   app_msg.Message
	GroupMgmtInfo    app_msg.Message
	GroupUnusedTitle app_msg.Message
	GroupUnusedInfo  app_msg.Message

	FootnoteTitle app_msg.Message
	FootnoteInfo  app_msg.Message
}

var (
	MDropboxBusiness = app_msg.Apply(&MsgDropboxBusiness{}).(*MsgDropboxBusiness)
)

type DropboxBusinessCatalogue interface {
	Recipe(path string) rc_recipe.Spec
	RecipeTable(name string, ui app_ui.UI, paths []string)
	WarnUnmentioned()
}

func NewDbxCatalogue() DropboxBusinessCatalogue {
	return &dbxCat{
		mentioned: make(map[string]bool),
	}
}

type dbxCat struct {
	mentioned map[string]bool
}

func (z *dbxCat) RecipeTable(name string, ui app_ui.UI, paths []string) {
	ui.WithTable(name, func(t app_ui.Table) {
		t.Header(MDropboxBusiness.CommandHeaderName, MDropboxBusiness.CommandHeaderDesc)

		for _, p := range paths {
			c := z.Recipe(p)
			t.Row(c.CliNameRef(""), c.Title())
		}
	})
}

func (z *dbxCat) Recipe(path string) rc_recipe.Spec {
	z.mentioned[path] = true
	return app_catalogue.Current().RecipeSpec(path)
}

func (z *dbxCat) WarnUnmentioned() {
	businessRecipes := make([]string, 0)
	for _, r := range app_catalogue.Current().Recipes() {
		spec := rc_spec.New(r)
		if spec.ConnUseBusiness() && !spec.IsSecret() {
			businessRecipes = append(businessRecipes, spec.CliPath())
		}
	}
	l := esl.Default()
	sort.Strings(businessRecipes)
	for _, r := range businessRecipes {
		if mentioned, ok := z.mentioned[r]; !ok || !mentioned {
			l.Warn("Unmentioned Dropbox Business recipe found", esl.String("Path", r))
		}
	}
}

func NewDropboxBusiness() dc_section.Document {
	return &DropboxBusiness{
		cat: NewDbxCatalogue(),
	}
}

type DropboxBusiness struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusiness) DocId() dc_index.DocId {
	return dc_index.DocSupplementalDropboxBusiness
}

func (z DropboxBusiness) Sections() []dc_section.Section {
	return []dc_section.Section{
		&DropboxBusinessMember{cat: z.cat},
		&DropboxBusinessGroup{cat: z.cat},

		// footnote section must be placed at the end of the doc
		&DropboxBusinessFootnote{cat: z.cat},
	}
}

type DropboxBusinessMember struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessMember) Title() app_msg.Message {
	return MDropboxBusiness.MemberTitle
}

func (z DropboxBusinessMember) Body(ui app_ui.UI) {
	ui.SubHeader(MDropboxBusiness.MemberInfoCommands)
	ui.Info(MDropboxBusiness.MemberInfoOverview)

	z.cat.RecipeTable("member info commands", ui, []string{
		"member list",
		"member folder list",
		"member quota list",
		"member quota usage",
		"team activity user",
	})

	ui.SubHeader(MDropboxBusiness.MemberBasicCommands)
	ui.Info(MDropboxBusiness.MemberBasicOverview)

	z.cat.RecipeTable("member management commands", ui, []string{
		"member invite",
		"member delete",
		"member detach",
		"member reinvite",
		"member update email",
		"member update profile",
		"member update visible",
		"member update invisible",
		"member quota update",
	})
}

type DropboxBusinessGroup struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessGroup) Title() app_msg.Message {
	return MDropboxBusiness.GroupTitle
}

func (z DropboxBusinessGroup) Body(ui app_ui.UI) {
	ui.SubHeader(MDropboxBusiness.GroupMgmtTitle)
	ui.Info(MDropboxBusiness.GroupMgmtInfo)

	z.cat.RecipeTable("group management", ui, []string{
		"group add",
		"group delete",
		"group batch delete",
		"group list",
		"group rename",
	})

	ui.SubHeader(MDropboxBusiness.GroupUnusedTitle)
	ui.Info(MDropboxBusiness.GroupUnusedInfo)

	z.cat.RecipeTable("handle unused groups", ui, []string{
		"group list",
		"group folder list",
		"group batch delete",
	})
}

type DropboxBusinessFootnote struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessFootnote) Title() app_msg.Message {
	return MDropboxBusiness.FootnoteTitle
}

func (z DropboxBusinessFootnote) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.FootnoteInfo)
	z.cat.WarnUnmentioned()
}
