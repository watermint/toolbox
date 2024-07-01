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

var (
	SkipDropboxBusinessCommandDoc = false
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

	LegalHoldTitle app_msg.Message
	LegalHoldInfo  app_msg.Message

	InsightTitle app_msg.Message
	InsightInfo  app_msg.Message
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
	if SkipDropboxBusinessCommandDoc {
		return []dc_section.Section{}
	}
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
		&DropboxBusinessLegalHold{cat: z.cat},

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
		"dropbox team member list",
		"dropbox team member feature",
		"dropbox team member folder list",
		"dropbox team member quota list",
		"dropbox team member quota usage",
		"dropbox team activity user",
	})

	ui.SubHeader(MDropboxBusiness.MemberBasicCommands)
	ui.Info(MDropboxBusiness.MemberBasicOverview)

	z.cat.RecipeTable("member management commands", ui, []string{
		"dropbox team member batch invite",
		"dropbox team member batch delete",
		"dropbox team member batch detach",
		"dropbox team member batch reinvite",
		"dropbox team member update batch email",
		"dropbox team member update batch profile",
		"dropbox team member update batch visible",
		"dropbox team member update batch invisible",
		"dropbox team member quota batch update",
	})

	ui.SubHeader(MDropboxBusiness.MemberProfileTitle)
	ui.Info(MDropboxBusiness.MemberProfileInfo)
	z.cat.RecipeTable("member profile commands", ui, []string{
		"dropbox team member update batch email",
		"dropbox team member update batch profile",
	})

	ui.SubHeader(MDropboxBusiness.MemberQuotaTitle)
	ui.Info(MDropboxBusiness.MemberQuotaInfo)
	z.cat.RecipeTable("member quota control", ui, []string{
		"dropbox team member quota list",
		"dropbox team member quota usage",
		"dropbox team member quota batch update",
	})

	ui.SubHeader(MDropboxBusiness.MemberSuspendTitle)
	ui.Info(MDropboxBusiness.MemberSuspendInfo)

	z.cat.RecipeTable("member suspend commands", ui, []string{
		"dropbox team member suspend",
		"dropbox team member unsuspend",
		"dropbox team member batch suspend",
		"dropbox team member batch unsuspend",
	})

	ui.SubHeader(MDropboxBusiness.MemberDirectoryRestrictionTitle)
	ui.Info(MDropboxBusiness.MemberDirectoryRestrictionInfo)
	z.cat.RecipeTable("directory restriction", ui, []string{
		"dropbox team member update batch visible",
		"dropbox team member update batch invisible",
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
		"dropbox team group add",
		"dropbox team group batch add",
		"dropbox team group batch delete",
		"dropbox team group delete",
		"dropbox team group list",
		"dropbox team group rename",
		"dropbox team group update type",
	})

	ui.SubHeader(MDropboxBusiness.GroupMemberTitle)
	ui.Info(MDropboxBusiness.GroupMemberInfo)

	z.cat.RecipeTable("group member management", ui, []string{
		"dropbox team group member add",
		"dropbox team group member delete",
		"dropbox team group member list",
		"dropbox team group member batch add",
		"dropbox team group member batch delete",
		"dropbox team group member batch update",
	})

	ui.SubHeader(MDropboxBusiness.GroupUnusedTitle)
	ui.Info(MDropboxBusiness.GroupUnusedInfo)

	z.cat.RecipeTable("handle unused groups", ui, []string{
		"dropbox team group list",
		"dropbox team group folder list",
		"dropbox team group batch delete",
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
	z.cat.RecipeTable("dropbox team folder operation", ui, []string{
		"dropbox team teamfolder add",
		"dropbox team teamfolder archive",
		"dropbox team teamfolder batch archive",
		"dropbox team teamfolder batch permdelete",
		"dropbox team teamfolder batch replication",
		"dropbox team teamfolder file size",
		"dropbox team teamfolder list",
		"dropbox team teamfolder permdelete",
		"dropbox team teamfolder policy list",
		"dropbox team teamfolder sync setting list",
		"dropbox team teamfolder sync setting update",
	})

	ui.SubHeader(MDropboxBusiness.ContentTeamFolderPermissionTitle)
	ui.Info(MDropboxBusiness.ContentTeamFolderPermissionInfo)
	z.cat.RecipeTable("dropbox team folder permission", ui, []string{
		"dropbox team teamfolder member list",
		"dropbox team teamfolder member add",
		"dropbox team teamfolder member delete",
	})

	ui.SubHeader(MDropboxBusiness.ContentTeamContentTitle)
	ui.Info(MDropboxBusiness.ContentTeamContentInfo)
	z.cat.RecipeTable("dropbox team content", ui, []string{
		"dropbox team content member list",
		"dropbox team content member size",
		"dropbox team content mount list",
		"dropbox team content policy list",
	})

	ui.SubHeader(MDropboxBusiness.ContentNamespaceTitle)
	z.cat.RecipeTable("dropbox team namespace", ui, []string{
		"dropbox team namespace list",
		"dropbox team namespace summary",
		"dropbox team namespace file list",
		"dropbox team namespace file size",
		"dropbox team namespace member list",
	})

	ui.SubHeader(MDropboxBusiness.ContentFileRequestTitle)
	z.cat.RecipeTable("dropbox team file request", ui, []string{
		"dropbox team filerequest list",
	})

	ui.SubHeader(MDropboxBusiness.ContentMemberFileTitle)
	z.cat.RecipeTable("member file commands", ui, []string{
		"dropbox team member file permdelete",
	})

	ui.SubHeader(MDropboxBusiness.InsightTitle)
	ui.Info(MDropboxBusiness.InsightInfo)
	z.cat.RecipeTable("insight commands", ui, []string{
		"dropbox team insight scan",
		"dropbox team insight scanretry",
		"dropbox team insight summarize",
		"dropbox team insight report teamfoldermember",
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
		"dropbox team device list",
		"dropbox team device unlink",
		"dropbox team linkedapp list",
		"dropbox team backup device status",
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
	z.cat.RecipeTable("dropbox team shared link commands", ui, []string{
		"dropbox team sharedlink list",
		"dropbox team sharedlink cap expiry",
		"dropbox team sharedlink cap visibility",
		"dropbox team sharedlink update expiry",
		"dropbox team sharedlink update password",
		"dropbox team sharedlink update visibility",
		"dropbox team sharedlink delete links",
		"dropbox team sharedlink delete member",
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
		"dropbox team member file lock all release",
		"dropbox team member file lock list",
		"dropbox team member file lock release",
	})

	ui.SubHeader(MDropboxBusiness.FileLockTeamFolderTitle)
	z.cat.RecipeTable("dropbox team teamfolder file lock management", ui, []string{
		"dropbox team teamfolder file list",
		"dropbox team teamfolder file lock all release",
		"dropbox team teamfolder file lock list",
		"dropbox team teamfolder file lock release",
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
		"dropbox team activity batch user",
		"dropbox team activity daily event",
		"dropbox team activity event",
		"dropbox team activity user",
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
		"dropbox team member list",
		"dropbox team member clear externalid",
		"dropbox team member update batch externalid",
		"dropbox team group list",
		"dropbox team group clear externalid",
	})

	ui.SubHeader(MDropboxBusiness.UsecaseDataMigrationTitle)
	ui.Info(MDropboxBusiness.UsecaseDataMigrationInfo)
	z.cat.RecipeTable("data migration commands", ui, []string{
		"dropbox team member folder replication",
		"dropbox team member replication",
		"dropbox team teamfolder partial replication",
		"dropbox team teamfolder replication",
	})

	ui.SubHeader(MDropboxBusiness.UsecaseTeamInfoTitle)
	z.cat.RecipeTable("information commands", ui, []string{
		"dropbox team feature",
		"dropbox team filesystem",
		"dropbox team info",
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
		"dropbox team content legacypaper count",
		"dropbox team content legacypaper list",
		"dropbox team content legacypaper export",
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
	z.cat.RecipeTable("dropbox team admin commands", ui, []string{
		"dropbox team admin list",
		"dropbox team admin role add",
		"dropbox team admin role clear",
		"dropbox team admin role delete",
		"dropbox team admin role list",
		"dropbox team admin group role add",
		"dropbox team admin group role delete",
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
	z.cat.RecipeTable("dropbox team runas commands", ui, []string{
		"dropbox team runas file list",
		"dropbox team runas file batch copy",
		"dropbox team runas file sync batch up",
		"dropbox team runas sharedfolder list",
		"dropbox team runas sharedfolder isolate",
		"dropbox team runas sharedfolder mount add",
		"dropbox team runas sharedfolder mount delete",
		"dropbox team runas sharedfolder mount list",
		"dropbox team runas sharedfolder mount mountable",
		"dropbox team runas sharedfolder batch leave",
		"dropbox team runas sharedfolder batch share",
		"dropbox team runas sharedfolder batch unshare",
		"dropbox team runas sharedfolder member batch add",
		"dropbox team runas sharedfolder member batch delete",
	})
}

type DropboxBusinessLegalHold struct {
	cat DropboxBusinessCatalogue
}

func (z DropboxBusinessLegalHold) Title() app_msg.Message {
	return MDropboxBusiness.LegalHoldTitle
}

func (z DropboxBusinessLegalHold) Body(ui app_ui.UI) {
	ui.Info(MDropboxBusiness.LegalHoldInfo)
	z.cat.RecipeTable("dropbox team legalhold commands", ui, []string{
		"dropbox team legalhold add",
		"dropbox team legalhold list",
		"dropbox team legalhold member batch update",
		"dropbox team legalhold member list",
		"dropbox team legalhold release",
		"dropbox team legalhold revision list",
		"dropbox team legalhold update desc",
		"dropbox team legalhold update name",
	})
}
