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
	DocDesc  app_msg.Message
	Title    app_msg.Message
	Overview app_msg.Message

	CommandHeaderName app_msg.Message
	CommandHeaderDesc app_msg.Message

	MemberTitle                     app_msg.Message
	MemberInfoCommands              app_msg.Message
	MemberInfoOverview              app_msg.Message
	MemberBasicCommands             app_msg.Message
	MemberBasicOverview             app_msg.Message
	MemberQuotaTitle                app_msg.Message
	MemberQuotaInfo                 app_msg.Message
	MemberDirectoryRestrictionTitle app_msg.Message
	MemberDirectoryRestrictionInfo  app_msg.Message
	MemberProfileTitle              app_msg.Message
	MemberProfileInfo               app_msg.Message
	MemberSuspendTitle              app_msg.Message
	MemberSuspendInfo               app_msg.Message

	GroupTitle       app_msg.Message
	GroupMgmtTitle   app_msg.Message
	GroupMgmtInfo    app_msg.Message
	GroupMemberTitle app_msg.Message
	GroupMemberInfo  app_msg.Message
	GroupUnusedTitle app_msg.Message
	GroupUnusedInfo  app_msg.Message

	ContentTitle                     app_msg.Message
	ContentInfo                      app_msg.Message
	ContentAboutNamespace            app_msg.Message
	ContentNamespaceTitle            app_msg.Message
	ContentTeamFolderOperationTitle  app_msg.Message
	ContentTeamFolderOperationInfo   app_msg.Message
	ContentTeamContentTitle          app_msg.Message
	ContentTeamContentInfo           app_msg.Message
	ContentTeamFolderPermissionTitle app_msg.Message
	ContentTeamFolderPermissionInfo  app_msg.Message
	ContentFileRequestTitle          app_msg.Message
	ContentMemberFileTitle           app_msg.Message

	ConnectTitle app_msg.Message
	ConnectInfo  app_msg.Message

	SharedLinkTitle               app_msg.Message
	SharedLinkInfo                app_msg.Message
	SharedLinkCapVsUpdateTitle    app_msg.Message
	SharedLinkCapVsUpdateDesc     app_msg.Message
	SharedLinkWithJqListTitle     app_msg.Message
	SharedLinkWithJqListExample   app_msg.Message
	SharedLinkWithJqDeleteTitle   app_msg.Message
	SharedLinkWithJqDeleteExample app_msg.Message

	ActivitiesTitle app_msg.Message
	ActivitiesInfo  app_msg.Message

	FileLockTitle           app_msg.Message
	FileLockInfo            app_msg.Message
	FileLockMemberTitle     app_msg.Message
	FileLockTeamFolderTitle app_msg.Message

	UsecaseTitle              app_msg.Message
	UsecaseExternalIdTitle    app_msg.Message
	UsecaseExternalIdInfo     app_msg.Message
	UsecaseDataMigrationTitle app_msg.Message
	UsecaseDataMigrationInfo  app_msg.Message
	UsecaseTeamInfoTitle      app_msg.Message

	FootnoteTitle app_msg.Message
	FootnoteInfo  app_msg.Message

	PaperTitle       app_msg.Message
	LegacyPaperTitle app_msg.Message
	LegacyPaperInfo  app_msg.Message

	TeamAdminTitle app_msg.Message
	TeamAdminInfo  app_msg.Message

	RunAsTitle app_msg.Message
	RunAsInfo  app_msg.Message
}

var (
	MDropboxBusiness = app_msg.Apply(&MsgDropboxBusiness{}).(*MsgDropboxBusiness)
)

type DropboxBusinessCatalogue interface {
	Recipe(path string) rc_recipe.Spec
	RecipeTable(name string, ui app_ui.UI, paths []string)
	// WarnUnmentioned returns true if there is one or more unmentioned commands found.
	WarnUnmentioned() bool
}

func NewDbxCatalogue(media dc_index.MediaType) DropboxBusinessCatalogue {
	return &dbxCat{
		media:     media,
		mentioned: make(map[string]bool),
	}
}

type dbxCat struct {
	media     dc_index.MediaType
	mentioned map[string]bool
}

func (z *dbxCat) RecipeTable(name string, ui app_ui.UI, paths []string) {
	lg := ui.Messages().Lang()

	ui.WithTable(name, func(t app_ui.Table) {
		t.Header(MDropboxBusiness.CommandHeaderName, MDropboxBusiness.CommandHeaderDesc)

		for _, p := range paths {
			spec := z.Recipe(p)

			relPath := ""
			if z.media == dc_index.MediaWeb {
				relPath = dc_index.DocName(z.media, dc_index.DocManualCommand, lg, dc_index.RefPath(true))
			}

			t.Row(spec.CliNameRef(z.media, lg, relPath), spec.Title())
		}
	})
}

func (z *dbxCat) Recipe(path string) rc_recipe.Spec {
	z.mentioned[path] = true
	return app_catalogue.Current().RecipeSpec(path)
}

func (z *dbxCat) WarnUnmentioned() bool {
	businessRecipes := make([]string, 0)
	for _, r := range app_catalogue.Current().Recipes() {
		spec := rc_spec.New(r)
		if spec.ConnUseBusiness() && !spec.IsSecret() {
			businessRecipes = append(businessRecipes, spec.CliPath())
		}
	}
	l := esl.Default()
	sort.Strings(businessRecipes)
	warn := false
	for _, r := range businessRecipes {
		if mentioned, ok := z.mentioned[r]; !ok || !mentioned {
			l.Warn("Unmentioned Dropbox Business recipe found", esl.String("Path", r))
			warn = true
		}
	}
	return warn
}

func NewDropboxBusiness(media dc_index.MediaType) dc_section.Document {
	return &DropboxBusiness{
		cat: NewDbxCatalogue(media),
	}
}

type DropboxBusiness struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusiness) DocDesc() app_msg.Message {
	return MDropboxBusiness.DocDesc
}

func (z DropboxBusiness) DocId() dc_index.DocId {
	return dc_index.DocSupplementalDropboxBusiness
}

func (z DropboxBusiness) Sections() []dc_section.Section {
	return []dc_section.Section{
		&DropboxBusinessMember{cat: z.cat},
		&DropboxBusinessGroup{cat: z.cat},
		&DropboxBusinessContent{cat: z.cat},
		&DropboxBusinessSharedLink{cat: z.cat},
		&DropboxBusinessFileLock{cat: z.cat},
		&DropboxBusinessActivities{cat: z.cat},
		&DropboxBusinessConnect{cat: z.cat},
		&DropboxBusinessUsecase{cat: z.cat},
		&DropboxBusinessPaper{cat: z.cat},
		&DropboxBusinessTeamAdmin{cat: z.cat},
		&DropboxBusinessRunAs{cat: z.cat},

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
		"member feature",
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

	ui.SubHeader(MDropboxBusiness.MemberProfileTitle)
	ui.Info(MDropboxBusiness.MemberProfileInfo)
	z.cat.RecipeTable("member profile commands", ui, []string{
		"member update email",
		"member update profile",
	})

	ui.SubHeader(MDropboxBusiness.MemberQuotaTitle)
	ui.Info(MDropboxBusiness.MemberQuotaInfo)
	z.cat.RecipeTable("member quota control", ui, []string{
		"member quota list",
		"member quota usage",
		"member quota update",
	})

	ui.SubHeader(MDropboxBusiness.MemberSuspendTitle)
	ui.Info(MDropboxBusiness.MemberSuspendInfo)

	z.cat.RecipeTable("member suspend commands", ui, []string{
		"member suspend",
		"member unsuspend",
		"member batch suspend",
		"member batch unsuspend",
	})

	ui.SubHeader(MDropboxBusiness.MemberDirectoryRestrictionTitle)
	ui.Info(MDropboxBusiness.MemberDirectoryRestrictionInfo)
	z.cat.RecipeTable("directory restriction", ui, []string{
		"member update visible",
		"member update invisible",
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
		"group batch add",
		"group batch delete",
		"group list",
		"group rename",
	})

	ui.SubHeader(MDropboxBusiness.GroupMemberTitle)
	ui.Info(MDropboxBusiness.GroupMemberInfo)

	z.cat.RecipeTable("group member management", ui, []string{
		"group member add",
		"group member delete",
		"group member list",
		"group member batch add",
		"group member batch delete",
		"group member batch update",
	})

	ui.SubHeader(MDropboxBusiness.GroupUnusedTitle)
	ui.Info(MDropboxBusiness.GroupUnusedInfo)

	z.cat.RecipeTable("handle unused groups", ui, []string{
		"group list",
		"group folder list",
		"group batch delete",
	})
}

type DropboxBusinessContent struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessContent) Title() app_msg.Message {
	return MDropboxBusiness.ContentTitle
}

func (z DropboxBusinessContent) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.ContentInfo)
	ui.Info(MDropboxBusiness.ContentAboutNamespace)

	ui.SubHeader(MDropboxBusiness.ContentTeamFolderOperationTitle)
	ui.Info(MDropboxBusiness.ContentTeamFolderOperationInfo)
	z.cat.RecipeTable("team folder operation", ui, []string{
		"teamfolder list",
		"teamfolder policy list",
		"teamfolder file size",
		"teamfolder add",
		"teamfolder archive",
		"teamfolder permdelete",
		"teamfolder batch archive",
		"teamfolder batch permdelete",
		"teamfolder batch replication",
	})

	ui.SubHeader(MDropboxBusiness.ContentTeamFolderPermissionTitle)
	ui.Info(MDropboxBusiness.ContentTeamFolderPermissionInfo)
	z.cat.RecipeTable("team folder permission", ui, []string{
		"teamfolder member list",
		"teamfolder member add",
		"teamfolder member delete",
	})

	ui.SubHeader(MDropboxBusiness.ContentTeamContentTitle)
	ui.Info(MDropboxBusiness.ContentTeamContentInfo)
	z.cat.RecipeTable("team content", ui, []string{
		"team content member list",
		"team content member size",
		"team content mount list",
		"team content policy list",
	})

	ui.SubHeader(MDropboxBusiness.ContentNamespaceTitle)
	z.cat.RecipeTable("team namespace", ui, []string{
		"team namespace list",
		"team namespace file list",
		"team namespace file size",
		"team namespace member list",
	})

	ui.SubHeader(MDropboxBusiness.ContentFileRequestTitle)
	z.cat.RecipeTable("team file request", ui, []string{
		"team filerequest list",
	})

	ui.SubHeader(MDropboxBusiness.ContentMemberFileTitle)
	z.cat.RecipeTable("member file commands", ui, []string{
		"member file permdelete",
	})
}

type DropboxBusinessConnect struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessConnect) Title() app_msg.Message {
	return MDropboxBusiness.ConnectTitle
}

func (z DropboxBusinessConnect) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.ConnectInfo)
	z.cat.RecipeTable("connected applications and devices commands", ui, []string{
		"team device list",
		"team device unlink",
		"team linkedapp list",
	})
}

type DropboxBusinessSharedLink struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessSharedLink) Title() app_msg.Message {
	return MDropboxBusiness.SharedLinkTitle
}

func (z DropboxBusinessSharedLink) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.SharedLinkInfo)
	z.cat.RecipeTable("team shared link commands", ui, []string{
		"team sharedlink list",
		"team sharedlink cap expiry",
		"team sharedlink cap visibility",
		"team sharedlink update expiry",
		"team sharedlink update password",
		"team sharedlink update visibility",
		"team sharedlink delete links",
		"team sharedlink delete member",
	})

	ui.SubHeader(MDropboxBusiness.SharedLinkCapVsUpdateTitle)
	ui.Info(MDropboxBusiness.SharedLinkCapVsUpdateDesc)

	ui.SubHeader(MDropboxBusiness.SharedLinkWithJqListTitle)
	ui.Info(MDropboxBusiness.SharedLinkWithJqListExample)

	ui.SubHeader(MDropboxBusiness.SharedLinkWithJqDeleteTitle)
	ui.Info(MDropboxBusiness.SharedLinkWithJqDeleteExample)
}

type DropboxBusinessFileLock struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessFileLock) Title() app_msg.Message {
	return MDropboxBusiness.FileLockTitle
}

func (z DropboxBusinessFileLock) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.FileLockInfo)
	ui.SubHeader(MDropboxBusiness.FileLockMemberTitle)
	z.cat.RecipeTable("member file lock management", ui, []string{
		"member file lock all release",
		"member file lock list",
		"member file lock release",
	})

	ui.SubHeader(MDropboxBusiness.FileLockTeamFolderTitle)
	z.cat.RecipeTable("teamfolder file lock management", ui, []string{
		"teamfolder file list",
		"teamfolder file lock all release",
		"teamfolder file lock list",
		"teamfolder file lock release",
	})
}

type DropboxBusinessActivities struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessActivities) Title() app_msg.Message {
	return MDropboxBusiness.ActivitiesTitle
}

func (z DropboxBusinessActivities) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.ActivitiesInfo)
	z.cat.RecipeTable("activities commands", ui, []string{
		"team activity batch user",
		"team activity daily event",
		"team activity event",
		"team activity user",
	})
}

type DropboxBusinessUsecase struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessUsecase) Title() app_msg.Message {
	return MDropboxBusiness.UsecaseTitle
}

func (z DropboxBusinessUsecase) Body(ui app_ui.UI) {
	ui.SubHeader(MDropboxBusiness.UsecaseExternalIdTitle)
	ui.Info(MDropboxBusiness.UsecaseExternalIdInfo)
	z.cat.RecipeTable("external id commands", ui, []string{
		"member list",
		"member clear externalid",
		"member update externalid",
	})

	ui.SubHeader(MDropboxBusiness.UsecaseDataMigrationTitle)
	ui.Info(MDropboxBusiness.UsecaseDataMigrationInfo)
	z.cat.RecipeTable("data migration commands", ui, []string{
		"member folder replication",
		"member replication",
		"teamfolder partial replication",
		"teamfolder replication",
	})

	ui.SubHeader(MDropboxBusiness.UsecaseTeamInfoTitle)
	z.cat.RecipeTable("information commands", ui, []string{
		"team feature",
		"team info",
	})
}

type DropboxBusinessPaper struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessPaper) Title() app_msg.Message {
	return MDropboxBusiness.PaperTitle
}

func (z DropboxBusinessPaper) Body(ui app_ui.UI) {
	ui.SubHeader(MDropboxBusiness.LegacyPaperTitle)
	ui.Info(MDropboxBusiness.LegacyPaperInfo)
	z.cat.RecipeTable("legacy paper commands", ui, []string{
		"team content legacypaper count",
		"team content legacypaper list",
		"team content legacypaper export",
	})
}

type DropboxBusinessTeamAdmin struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessTeamAdmin) Title() app_msg.Message {
	return MDropboxBusiness.TeamAdminTitle
}

func (z DropboxBusinessTeamAdmin) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.TeamAdminInfo)
	z.cat.RecipeTable("team admin commands", ui, []string{
		"team admin list",
		"team admin role add",
		"team admin role clear",
		"team admin role delete",
		"team admin role list",
		"team admin group role add",
		"team admin group role delete",
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
	if z.cat.WarnUnmentioned() {
		panic("Unmentioned Dropbox Business command found")
	}
}

type DropboxBusinessRunAs struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessRunAs) Title() app_msg.Message {
	return MDropboxBusiness.RunAsTitle
}

func (z DropboxBusinessRunAs) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.RunAsInfo)
	z.cat.RecipeTable("team runas commands", ui, []string{
		"team runas file batch copy",
		"team runas file sync batch up",
		"team runas sharedfolder batch share",
		"team runas sharedfolder batch unshare",
		"team runas sharedfolder member batch add",
		"team runas sharedfolder member batch delete",
	})
}
