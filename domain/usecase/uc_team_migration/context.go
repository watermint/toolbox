package uc_team_migration

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/app/app_report/app_report_json"
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

// Migration context. Migration scope includes mutable states like permissions.
type Context interface {
	// Set admins
	SetAdmins(src, dst *mo_profile.Profile)

	// Add member who migrate to new team
	AddMember(member *mo_profile.Profile)

	// Add group which migrate to new team
	AddGroup(group *mo_group.Group)

	// Add group in destination team
	AddDestGroup(group *mo_group.Group)

	// Add group member for preserved state.
	// Given `member` must be listed on `Members` instead of call
	AddGroupMember(group *mo_group.Group, member *mo_group_member.Member)

	// Add team folder which migrate to new team
	AddTeamFolder(teamFolder *mo_teamfolder.TeamFolder)

	// Add shared folder which migrate to new team
	AddSharedFolder(sharedFolder *mo_sharedfolder.SharedFolder)

	// Add namespace
	AddNamespace(namespace *mo_namespace.Namespace)

	// Add namespace detail
	AddNamespaceDetail(sharedFolder *mo_sharedfolder.SharedFolder)

	// Add namespace member
	AddNamespaceMember(namespace *mo_namespace.Namespace, member mo_sharedfolder_member.Member)

	// Filter by migration target members
	FilterByMigrationTarget(members []*mo_profile.Profile) (filtered []*mo_profile.Profile)

	// Get src team admin
	AdminSrc() *mo_profile.Profile

	// Get dst team admin
	AdminDst() *mo_profile.Profile

	// Members to migrate
	Members() (members map[string]*mo_profile.Profile)

	// Groups to migrate
	Groups() (groups map[string]*mo_group.Group)

	// Groups created in destination team
	DestGroups() (groups map[string]*mo_group.Group)

	// Group members to migrate
	GroupMembers(group *mo_group.Group) (members []*mo_group_member.Member)

	// Shared folders to migrate
	SharedFolders() (folders map[string]*mo_sharedfolder.SharedFolder)

	// Team folders to migrate
	TeamFolders() (folders map[string]*mo_teamfolder.TeamFolder)

	// Namespaces of source team
	Namespaces() (namespaces map[string]*mo_namespace.Namespace)

	// Namespace details of source team
	NamespaceDetails() (details map[string]*mo_sharedfolder.SharedFolder)

	// Members of namespace
	NamespaceMembers(namespaceId string) (members []mo_sharedfolder_member.Member)

	// Set team folder migration context
	SetContextTeamFolder(ctx uc_teamfolder_mirror.Context)

	// Team folder migration context
	ContextTeamFolder() uc_teamfolder_mirror.Context

	// Set whether migrate groups only related to members, or sharing
	SetGroupsOnlyRelated(onlyRelated bool)

	// Whether migrate groups only related to members, or sharing
	GroupsOnlyRelated() bool
}

const (
	storageTagMembers          = "src_members"
	storageTagDestGroups       = "dst_groups"
	storageTagGroups           = "src_groups"
	storageTagGroupMembers     = "src_group_members"
	storageTagTeamFolders      = "src_team_folders"
	storageTagNamespaces       = "src_namespaces"
	storageTagNamespaceDetails = "src_namespace_details"
	storageTagSharedFolders    = "src_shared_folders"
	storageTagNamespaceMembers = "src_namespace_members"
	storageTagOptions          = "options"
)

func newContext(ctxExec *app.ExecContext) Context {
	storageTags := []string{
		storageTagMembers,
		storageTagDestGroups,
		storageTagGroups,
		storageTagGroupMembers,
		storageTagTeamFolders,
		storageTagNamespaces,
		storageTagNamespaceDetails,
		storageTagSharedFolders,
		storageTagNamespaceMembers,
		storageTagOptions,
	}
	storages := make(map[string]app_report.Report)
	for _, tag := range storageTags {
		s := &app_report_json.JsonReport{
			ReportPath: filepath.Join(ctxExec.JobsPath(), "state", tag),
		}
		if err := s.Init(ctxExec); err != nil {
			ctxExec.Log().Warn("Unable to store state", zap.String("tag", tag), zap.Error(err))
		}
		storages[tag] = s
	}

	return &contextImpl{
		ctxExec:          ctxExec,
		storages:         storages,
		members:          make(map[string]*mo_profile.Profile),
		destGroups:       make(map[string]*mo_group.Group),
		groups:           make(map[string]*mo_group.Group),
		groupMembers:     make(map[string][]*mo_group_member.Member),
		teamFolders:      make(map[string]*mo_teamfolder.TeamFolder),
		namespaces:       make(map[string]*mo_namespace.Namespace),
		namespaceDetails: make(map[string]*mo_sharedfolder.SharedFolder),
		sharedFolders:    make(map[string]*mo_sharedfolder.SharedFolder),
		namespaceMember:  make(map[string][]mo_sharedfolder_member.Member),
		contextOpts:      &contextOpts{},
	}
}

type contextOpts struct {
	groupsOnlyRelated bool                `json:"groups_only_related"`
	adminSrc          *mo_profile.Profile `json:"admin_src"`
	adminDst          *mo_profile.Profile `json:"admin_dst"`
}

type contextImpl struct {
	ctxExec          *app.ExecContext
	ctxTeamFolder    uc_teamfolder_mirror.Context
	storages         map[string]app_report.Report
	members          map[string]*mo_profile.Profile
	destGroups       map[string]*mo_group.Group
	groups           map[string]*mo_group.Group
	groupMembers     map[string][]*mo_group_member.Member
	teamFolders      map[string]*mo_teamfolder.TeamFolder
	namespaces       map[string]*mo_namespace.Namespace
	namespaceDetails map[string]*mo_sharedfolder.SharedFolder
	sharedFolders    map[string]*mo_sharedfolder.SharedFolder
	namespaceMember  map[string][]mo_sharedfolder_member.Member
	contextOpts      *contextOpts
}

func (z *contextImpl) SetGroupsOnlyRelated(onlyRelated bool) {
	z.contextOpts.groupsOnlyRelated = onlyRelated
	if s, e := z.storages[storageTagOptions]; e {
		s.Report(z.contextOpts)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
}

func (z *contextImpl) AddNamespaceDetail(sharedFolder *mo_sharedfolder.SharedFolder) {
	if s, e := z.storages[storageTagNamespaceDetails]; e {
		s.Report(sharedFolder)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.namespaceDetails[sharedFolder.SharedFolderId] = sharedFolder
}

func (z *contextImpl) NamespaceDetails() (details map[string]*mo_sharedfolder.SharedFolder) {
	return z.namespaceDetails
}

func (z *contextImpl) AddDestGroup(group *mo_group.Group) {
	if s, e := z.storages[storageTagDestGroups]; e {
		s.Report(group)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.destGroups[group.GroupId] = group
}

func (z *contextImpl) DestGroups() (groups map[string]*mo_group.Group) {
	return z.destGroups
}

func (z *contextImpl) FilterByMigrationTarget(members []*mo_profile.Profile) (filtered []*mo_profile.Profile) {
	filtered = make([]*mo_profile.Profile, 0)

	for _, member := range members {
		for _, marked := range z.members {
			if member.TeamMemberId == marked.TeamMemberId {
				filtered = append(filtered, member)
				break
			}
		}
	}
	return filtered
}

func (z *contextImpl) AddNamespace(namespace *mo_namespace.Namespace) {
	if s, e := z.storages[storageTagNamespaces]; e {
		s.Report(namespace)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.namespaces[namespace.NamespaceId] = namespace
}

func (z *contextImpl) Namespaces() (namespaces map[string]*mo_namespace.Namespace) {
	return z.namespaces
}

func (z *contextImpl) GroupsOnlyRelated() bool {
	return z.contextOpts.groupsOnlyRelated
}

func (z *contextImpl) SetAdmins(src, dst *mo_profile.Profile) {
	z.contextOpts.adminSrc = src
	z.contextOpts.adminDst = dst
	if s, e := z.storages[storageTagOptions]; e {
		s.Report(z.contextOpts)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
}

func (z *contextImpl) AdminSrc() *mo_profile.Profile {
	return z.contextOpts.adminSrc
}

func (z *contextImpl) AdminDst() *mo_profile.Profile {
	return z.contextOpts.adminDst
}

func (z *contextImpl) AddMember(member *mo_profile.Profile) {
	if s, e := z.storages[storageTagMembers]; e {
		s.Report(member)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.members[member.TeamMemberId] = member
}

func (z *contextImpl) AddGroup(group *mo_group.Group) {
	if s, e := z.storages[storageTagGroups]; e {
		s.Report(group)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.groups[group.GroupId] = group
}

func (z *contextImpl) AddGroupMember(group *mo_group.Group, member *mo_group_member.Member) {
	if s, e := z.storages[storageTagGroupMembers]; e {
		s.Report(mo_group_member.NewGroupMember(group, member))
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}

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
	if s, e := z.storages[storageTagTeamFolders]; e {
		s.Report(teamFolder)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.teamFolders[teamFolder.TeamFolderId] = teamFolder
}

func (z *contextImpl) AddSharedFolder(sharedFolder *mo_sharedfolder.SharedFolder) {
	if s, e := z.storages[storageTagSharedFolders]; e {
		s.Report(sharedFolder)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}

	z.sharedFolders[sharedFolder.SharedFolderId] = sharedFolder
}

func (z *contextImpl) AddNamespaceMember(namespace *mo_namespace.Namespace, member mo_sharedfolder_member.Member) {
	if s, e := z.storages[storageTagNamespaceMembers]; e {
		s.Report(mo_namespace.NewNamespaceMember(namespace, member))
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}

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
