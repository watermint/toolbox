package uc_team_migration

import (
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
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"go.uber.org/zap"
	"path/filepath"
)

type Migration interface {
	Scope(opts ...ScopeOpt) (ctx Context, err error)

	// Preflight check
	Preflight(ctx Context) (err error)

	// Do entire migration process.
	Migrate(ctx Context) (err error)

	// Inspect team status.
	// Ensure both team allow externally sharing shared folders.
	// Ensure which contents are not migrated by this migration.
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
}

func newContext(ctxExec *app.ExecContext) Context {
	storage := &app_report.Factory{
		ExecContext: ctxExec,
		Path:        filepath.Join(ctxExec.JobsPath(), "state"),
	}
	if err := storage.Init(ctxExec); err != nil {
		ctxExec.Log().Warn("Unable to store state", zap.Error(err))
	}

	return &contextImpl{
		ctxExec:         ctxExec,
		storage:         storage,
		members:         make(map[string]*mo_profile.Profile),
		groups:          make(map[string]*mo_group.Group),
		groupMembers:    make(map[string][]*mo_group_member.Member),
		teamFolders:     make(map[string]*mo_teamfolder.TeamFolder),
		sharedFolders:   make(map[string]*mo_sharedfolder.SharedFolder),
		namespaceMember: make(map[string][]mo_sharedfolder_member.Member),
	}
}

type contextImpl struct {
	ctxExec         *app.ExecContext
	ctxTeamFolder   uc_teamfolder_mirror.Context
	storage         app_report.Report
	members         map[string]*mo_profile.Profile
	groups          map[string]*mo_group.Group
	groupMembers    map[string][]*mo_group_member.Member
	teamFolders     map[string]*mo_teamfolder.TeamFolder
	sharedFolders   map[string]*mo_sharedfolder.SharedFolder
	namespaceMember map[string][]mo_sharedfolder_member.Member
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

func (z *migrationImpl) Scope(opts ...ScopeOpt) (ctx Context, err error) {
	panic("implement me")
}

func (z *migrationImpl) Preflight(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Migrate(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Inspect(ctx Context) (err error) {
	panic("implement me")
}

func (z *migrationImpl) Preserve(ctx Context) (err error) {
	panic("implement me")
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
