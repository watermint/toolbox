package uc_team_migration

import (
	"errors"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/model/mo_namespace"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type Migration interface {
	Scope(opts ...ScopeOpt) (ctx Context, err error)

	// Preflight check
	Preflight(ctx Context) (err error)

	// Do entire migration process.
	Migrate(ctx Context) (err error)

	// Inspect team status.
	// Ensure both team allow externally sharing shared folders.
	Inspect(ctx Context) (err error)

	// Preserve members, groups, and sharing status.
	Preserve(ctx Context) (err error)

	// Bridge shared folders.
	// Share all shared folders to destination admin.
	Bridge(ctx Context) (err error)

	// Mount
	Mount(ctx Context) (err error)

	// Mirror team folders.
	Content(ctx Context) (err error)

	// Transfer members.
	// Convert accounts into Basic, and invite from destination team.
	Transfer(ctx Context) (err error)

	// Mirror permissions.
	// Create groups, invite members to shared folders or nested folders,
	// leave destination admin from bridged shared folders.
	Permissions(ctx Context) (err error)

	// Restore state.
	// Restore mount state.
	Restore(ctx Context) (err error)

	// Cleanup
	Cleanup(ctx Context) (err error)
}

// Migration context. Migration scope includes mutable states like permissions.
type Context interface {
	// Set admins
	SetAdmins(src, dst *mo_profile.Profile)

	// Add member who migrate to new team
	AddMember(member *mo_profile.Profile)

	// Add group which migrate to new team
	AddGroup(group *mo_group.Group)

	// Add group member for preserved state.
	// Given `member` must be listed on `Members` instead of call
	AddGroupMember(group *mo_group.Group, member *mo_group_member.Member)

	// Add team folder which migrate to new team
	AddTeamFolder(teamFolder *mo_teamfolder.TeamFolder)

	// Add shared folder which migrate to new team
	AddSharedFolder(sharedFolder *mo_sharedfolder.SharedFolder)

	// Add namespace member
	AddNamespaceMember(namespace *mo_namespace.Namespace, member mo_sharedfolder_member.Member)

	// Get src team admin
	AdminSrc() *mo_profile.Profile

	// Get dst team admin
	AdminDst() *mo_profile.Profile

	// Members to migrate
	Members() (members map[string]*mo_profile.Profile)

	// Groups to migrate
	Groups() (groups map[string]*mo_group.Group)

	// Group members to migrate
	GroupMembers(group *mo_group.Group) (members []*mo_group_member.Member)

	// Shared folders to migrate
	SharedFolders() (folders map[string]*mo_sharedfolder.SharedFolder)

	// Team folders to migrate
	TeamFolders() (folders map[string]*mo_teamfolder.TeamFolder)

	// Members of namespace
	NamespaceMembers(namespaceId string) (members []mo_sharedfolder_member.Member)

	// Set team folder migration context
	SetContextTeamFolder(ctx uc_teamfolder_mirror.Context)

	// Team folder migration context
	ContextTeamFolder() uc_teamfolder_mirror.Context

	// Whether migrate groups only related to members, or sharing
	GroupsOnlyRelated() bool
}

func newContext(ctxExec *app.ExecContext, groupOnlyRelated bool) Context {
	storage := &app_report.Factory{
		ExecContext: ctxExec,
		Path:        filepath.Join(ctxExec.JobsPath(), "state"),
	}
	if err := storage.Init(ctxExec); err != nil {
		ctxExec.Log().Warn("Unable to store state", zap.Error(err))
	}

	return &contextImpl{
		ctxExec:           ctxExec,
		storage:           storage,
		members:           make(map[string]*mo_profile.Profile),
		groups:            make(map[string]*mo_group.Group),
		groupMembers:      make(map[string][]*mo_group_member.Member),
		teamFolders:       make(map[string]*mo_teamfolder.TeamFolder),
		sharedFolders:     make(map[string]*mo_sharedfolder.SharedFolder),
		namespaceMember:   make(map[string][]mo_sharedfolder_member.Member),
		groupsOnlyRelated: groupOnlyRelated,
	}
}

type contextImpl struct {
	ctxExec           *app.ExecContext
	ctxTeamFolder     uc_teamfolder_mirror.Context
	storage           app_report.Report
	members           map[string]*mo_profile.Profile
	groups            map[string]*mo_group.Group
	groupMembers      map[string][]*mo_group_member.Member
	teamFolders       map[string]*mo_teamfolder.TeamFolder
	sharedFolders     map[string]*mo_sharedfolder.SharedFolder
	namespaceMember   map[string][]mo_sharedfolder_member.Member
	adminSrc          *mo_profile.Profile
	adminDst          *mo_profile.Profile
	groupsOnlyRelated bool
}

func (z *contextImpl) GroupsOnlyRelated() bool {
	return z.groupsOnlyRelated
}

func (z *contextImpl) SetAdmins(src, dst *mo_profile.Profile) {
	z.adminSrc = src
	z.adminDst = dst
}

func (z *contextImpl) AdminSrc() *mo_profile.Profile {
	return z.adminSrc
}

func (z *contextImpl) AdminDst() *mo_profile.Profile {
	return z.adminDst
}

func (z *contextImpl) AddMember(member *mo_profile.Profile) {
	z.storage.Report(member)
	z.members[member.TeamMemberId] = member
}

func (z *contextImpl) AddGroup(group *mo_group.Group) {
	z.storage.Report(group)
	z.groups[group.GroupId] = group
}

func (z *contextImpl) AddGroupMember(group *mo_group.Group, member *mo_group_member.Member) {
	z.storage.Report(mo_group_member.NewGroupMember(group, member))

	var members []*mo_group_member.Member
	if mem, e := z.groupMembers[group.GroupId]; !e {
		members = append(mem, member)
	} else {
		members = make([]*mo_group_member.Member, 0)
		members = append(members, member)
	}
	z.groupMembers[group.GroupId] = members
}

func (z *contextImpl) AddTeamFolder(teamFolder *mo_teamfolder.TeamFolder) {
	z.storage.Report(teamFolder)
	z.teamFolders[teamFolder.TeamFolderId] = teamFolder
}

func (z *contextImpl) AddSharedFolder(sharedFolder *mo_sharedfolder.SharedFolder) {
	z.storage.Report(sharedFolder)
	z.sharedFolders[sharedFolder.SharedFolderId] = sharedFolder
}

func (z *contextImpl) AddNamespaceMember(namespace *mo_namespace.Namespace, member mo_sharedfolder_member.Member) {
	z.storage.Report(mo_namespace.NewNamespaceMember(namespace, member))

	var members []mo_sharedfolder_member.Member
	if mem, e := z.namespaceMember[namespace.NamespaceId]; e {
		members = append(mem, member)
	} else {
		members = make([]mo_sharedfolder_member.Member, 0)
		members = append(members, member)
	}
	z.namespaceMember[namespace.NamespaceId] = members
}

func (z *contextImpl) Members() (members map[string]*mo_profile.Profile) {
	return z.members
}

func (z *contextImpl) Groups() (groups map[string]*mo_group.Group) {
	return z.groups
}

func (z *contextImpl) GroupMembers(group *mo_group.Group) (members []*mo_group_member.Member) {
	if members, e := z.groupMembers[group.GroupId]; e {
		return members
	} else {
		z.ctxExec.Log().Warn("Group members not found", zap.String("groupId", group.GroupId))
		return make([]*mo_group_member.Member, 0)
	}
}

func (z *contextImpl) SharedFolders() (folders map[string]*mo_sharedfolder.SharedFolder) {
	return z.sharedFolders
}

func (z *contextImpl) TeamFolders() (folders map[string]*mo_teamfolder.TeamFolder) {
	return z.teamFolders
}

func (z *contextImpl) NamespaceMembers(namespaceId string) (members []mo_sharedfolder_member.Member) {
	if members, e := z.namespaceMember[namespaceId]; e {
		return members
	} else {
		z.ctxExec.Log().Warn("Namespace members not found", zap.String("namespaceId", namespaceId))
		return make([]mo_sharedfolder_member.Member, 0)
	}
}

func (z *contextImpl) SetContextTeamFolder(ctx uc_teamfolder_mirror.Context) {
	z.ctxTeamFolder = ctx
}

func (z *contextImpl) ContextTeamFolder() uc_teamfolder_mirror.Context {
	return z.ctxTeamFolder
}

type ScopeOpt func(opt *scopeOpts) *scopeOpts
type scopeOpts struct {
	membersAllExceptAdmin    bool
	membersSpecifiedEmail    []string
	teamFoldersAll           bool
	teamFoldersSpecifiedName []string
	groupsOnlyRelated        bool
}

func MembersAllExceptAdmin() ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.membersAllExceptAdmin = true
		return opt
	}
}
func MembersSpecifiedEmail(members []string) ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.membersSpecifiedEmail = members
		return opt
	}
}
func TeamFoldersAll() ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.teamFoldersAll = true
		return opt
	}
}
func TeamFoldersSpecifiedName(name []string) ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.teamFoldersSpecifiedName = name
		return opt
	}
}
func GroupsOnlyRelated() ScopeOpt {
	return func(opt *scopeOpts) *scopeOpts {
		opt.groupsOnlyRelated = true
		return opt
	}
}

func New(ctxExec *app.ExecContext, ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst api_context.Context) Migration {
	return &migrationImpl{
		ctxExec:          ctxExec,
		ctxFileSrc:       ctxFileSrc,
		ctxMgtSrc:        ctxMgtSrc,
		ctxFileDst:       ctxFileDst,
		ctxMgtDst:        ctxMgtDst,
		teamFolderMirror: uc_teamfolder_mirror.New(ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst),
	}
}

type migrationImpl struct {
	ctxExec          *app.ExecContext
	ctxFileSrc       api_context.Context
	ctxFileDst       api_context.Context
	ctxMgtSrc        api_context.Context
	ctxMgtDst        api_context.Context
	teamFolderMirror uc_teamfolder_mirror.TeamFolder
}

func (z *migrationImpl) log() *zap.Logger {
	return z.ctxExec.Log()
}

func (z *migrationImpl) Scope(opts ...ScopeOpt) (ctx Context, err error) {
	so := &scopeOpts{
		membersSpecifiedEmail:    make([]string, 0),
		teamFoldersSpecifiedName: make([]string, 0),
	}
	for _, o := range opts {
		o(so)
	}

	// Prepare migration context
	ctx = newContext(z.ctxExec, so.groupsOnlyRelated)

	// validation
	if so.membersAllExceptAdmin && len(so.membersSpecifiedEmail) > 0 {
		z.log().Warn("Conflicted option `membersAllExceptAdmin` and `membersSpecifiedEmail`")
		return nil, errors.New("conflicted option")
	}
	if !so.membersAllExceptAdmin && len(so.membersSpecifiedEmail) < 1 {
		z.log().Warn("Please specify `memberAllExceptAdmin` or `membersSpecifiedEmail`")
		return nil, errors.New("not enough options")
	}
	if so.teamFoldersAll && len(so.teamFoldersSpecifiedName) > 0 {
		z.log().Warn("Conflicted option `teamFoldersAll` and `teamFoldersSpecifiedName`")
		return nil, errors.New("conflicted option")
	}
	if !so.teamFoldersAll && len(so.teamFoldersSpecifiedName) < 1 {
		z.log().Warn("Please specify `teamFoldersAll` or `teamFoldersSpecifiedName`")
		return nil, errors.New("not enough options")
	}

	// Identify admins
	identifyAdmins := func() error {
		adminSrc, err := sv_profile.NewTeam(z.ctxMgtSrc).Admin()
		if err != nil {
			return err
		}
		adminDst, err := sv_profile.NewTeam(z.ctxMgtDst).Admin()
		if err != nil {
			return err
		}
		z.log().Debug("Admins identified",
			zap.String("srcId", adminSrc.TeamMemberId),
			zap.String("srcEmail", adminSrc.Email),
			zap.String("dstId", adminDst.TeamMemberId),
			zap.String("dstEmail", adminDst.Email),
		)
		ctx.SetAdmins(adminSrc, adminDst)
		return nil
	}
	if err = identifyAdmins(); err != nil {
		return nil, err
	}

	// Define scope of members
	allMembers, err := sv_member.New(z.ctxMgtSrc).List()
	if err != nil {
		return nil, err
	}
	if so.membersAllExceptAdmin {
		for _, member := range allMembers {
			ctx.AddMember(member.Profile())
		}
	} else if len(so.membersSpecifiedEmail) > 0 {
		err = nil
		for _, email := range so.membersSpecifiedEmail {
			found := false
			emailLower := strings.ToLower(email)
			for _, member := range allMembers {
				if strings.ToLower(member.Email) == emailLower {
					ctx.AddMember(member.Profile())
					found = true
					break
				}
			}
			if !found {
				z.log().Warn("Member not found for email address", zap.String("email", email))
				err = errors.New("member not found")
			}
		}
		if err != nil {
			return nil, err
		}
	}
	if len(ctx.Members()) < 1 {
		z.log().Warn("No members found")
		return nil, errors.New("no member to migrate")
	}
	z.log().Debug("Members to migrate", zap.Int("count", len(ctx.Members())))

	// Define scope of team folders
	allFolders, err := sv_teamfolder.New(z.ctxFileSrc).List()
	if err != nil {
		return nil, err
	}
	if so.teamFoldersAll {
		for _, folder := range allFolders {
			ctx.AddTeamFolder(folder)
		}
	} else if len(so.teamFoldersSpecifiedName) > 0 {
		err = nil
		for _, name := range so.teamFoldersSpecifiedName {
			found := false
			nameLower := strings.ToLower(name)
			for _, folder := range allFolders {
				if strings.ToLower(folder.Name) == nameLower {
					ctx.AddTeamFolder(folder)
					found = true
					break
				}
			}
			if !found {
				z.log().Warn("Team folder not found for name", zap.String("name", name))
				err = errors.New("team folder not found")
			}
		}
		if err != nil {
			return nil, err
		}
	}
	z.log().Debug("Team folders to migrate", zap.Int("count", len(ctx.TeamFolders())))

	// Team folder mirror
	prepTeamFolderMirror := func() error {
		names := make([]string, 0)
		for _, f := range ctx.TeamFolders() {
			names = append(names, f.Name)
		}
		ctxTf, err := z.teamFolderMirror.PartialScope(names)
		if err != nil {
			return err
		}
		ctx.SetContextTeamFolder(ctxTf)
		return nil
	}
	if err = prepTeamFolderMirror(); err != nil {
		return nil, err
	}

	return ctx, nil
}

func (z *migrationImpl) Preflight(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Migrate(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Inspect(ctx Context) (err error) {
	// Inspect team information.
	inspectTeams := func() error {
		var inspectErr error
		inspectErr = nil
		infoSrc, err := sv_team.New(z.ctxMgtSrc).Info()
		if err != nil {
			return err
		}
		infoDst, err := sv_team.New(z.ctxMgtDst).Info()
		if err != nil {
			return err
		}
		z.log().Debug("Team info",
			zap.String("srcId", infoSrc.TeamId),
			zap.String("srcName", infoSrc.Name),
			zap.Int("srcLicenses", infoSrc.NumLicensedUsers),
			zap.Int("srcProvisioned", infoSrc.NumProvisionedUsers),
			zap.String("srcPolicySharedFolderMember", infoSrc.PolicySharedFolderMember),
			zap.String("srcPolicySharedFolderJoin", infoSrc.PolicySharedFolderJoin),
			zap.String("dstId", infoDst.TeamId),
			zap.String("dstName", infoDst.Name),
			zap.Int("dstLicenses", infoDst.NumLicensedUsers),
			zap.Int("dstProvisioned", infoDst.NumProvisionedUsers),
			zap.String("dstPolicySharedFolderMember", infoDst.PolicySharedFolderMember),
			zap.String("dstPolicySharedFolderJoin", infoDst.PolicySharedFolderJoin),
		)
		if infoSrc.TeamId == infoDst.TeamId {
			z.log().Warn("Source and destination team are the same team.")
			inspectErr = errors.New("source and destination teams are the same team")
		}
		if infoSrc.PolicySharedFolderMember != "anyone" {
			z.log().Warn("Source team: Shared folder member policy must be `anyone` during migration")
			inspectErr = errors.New("invalid sharing policy")
		}
		if infoSrc.PolicySharedFolderJoin != "from_anyone" {
			z.log().Warn("Source team: Shared folder join policy must be `from_anyone` during migration")
			inspectErr = errors.New("invalid sharing policy")
		}
		if infoDst.PolicySharedFolderMember != "anyone" {
			z.log().Warn("Dest team: Shared folder member policy must be `anyone` during migration")
			inspectErr = errors.New("invalid sharing policy")
		}
		if infoDst.PolicySharedFolderJoin != "from_anyone" {
			z.log().Warn("Dest team: Shared folder join policy must be `from_anyone` during migration")
			inspectErr = errors.New("invalid sharing policy")
		}

		return inspectErr
	}
	if err = inspectTeams(); err != nil {
		return err
	}

	// Inspect team folder mirror
	if err = z.teamFolderMirror.Inspect(ctx.ContextTeamFolder()); err != nil {
		return err
	}

	return nil
}

func (z *migrationImpl) Preserve(ctx Context) (err error) {
	// Preserve group
	preserveGroups := func() error {
		groups, err := sv_group.New(z.ctxMgtSrc).List()
		if err != nil {
			return err
		}

		// fetch group members
		for _, group := range groups {
			members, err := sv_group_member.New(z.ctxMgtSrc, group).List()
			if err != nil {
				return err
			}

			// filter into members who migrate to dest team
			membersDst := make(map[string]*mo_group_member.Member)
			for _, m := range ctx.Members() {
				for _, gm := range members {
					if gm.TeamMemberId == m.TeamMemberId {
						membersDst[gm.TeamMemberId] = gm
						break
					}
				}
			}

			migrate := false
			if ctx.GroupsOnlyRelated() {
				// ensure any of group member is in migration target members
				if len(membersDst) > 0 {
					migrate = true
				}
			} else {
				migrate = true
			}

			if migrate {
				// preserve group
				ctx.AddGroup(group)

				// preserve group members
				for _, m := range membersDst {
					ctx.AddGroupMember(group, m)
				}
			}
		}

		return nil
	}
	if err = preserveGroups(); err != nil {
		return err
	}

	// Preserve shared folders
	preserveSharedFolders := func() error {
		// fetch all shared folders of migrating members
		for _, member := range ctx.Members() {
			ctxFileOfMember := z.ctxFileSrc.AsMemberId(member.TeamMemberId)
			folders, err := sv_sharedfolder.New(ctxFileOfMember).List()
			if err != nil {
				return err
			}
			for _, folder := range folders {
				ctx.AddSharedFolder(folder)
			}
		}
		return nil
	}
	if err = preserveSharedFolders(); err != nil {
		return err
	}

	// Preserve team folder permissions
	panic("implement me")

	// Verify any of group already created in dest team or not
	verifyDestGroups := func() error {
		groupsDst, err := sv_group.New(z.ctxMgtDst).List()
		if err != nil {
			return err
		}

		var verifyErr error
		verifyErr = nil
		for _, groupSrc := range ctx.Groups() {
			nameSrcLower := strings.ToLower(groupSrc.GroupName)
			for _, groupDst := range groupsDst {
				if strings.ToLower(groupDst.GroupName) == nameSrcLower {
					z.log().Warn("Group already exists in dest team", zap.String("GroupName", groupSrc.GroupName))
					verifyErr = errors.New("group already exists")
				}
			}
		}
		return verifyErr
	}
	if err = verifyDestGroups(); err != nil {
		return err
	}

	return nil
}

func (z *migrationImpl) Bridge(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Mount(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Content(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Transfer(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Permissions(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Restore(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Cleanup(ctx Context) (err error) {
	panic("implement me")
}
