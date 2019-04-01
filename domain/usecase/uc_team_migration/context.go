package uc_team_migration

import (
	"encoding/json"
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
	"io/ioutil"
	"os"
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

	// Store state
	StoreState() error
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
	ctx := &contextImpl{
		ctxExec:             ctxExec,
		MapMembers:          make(map[string]*mo_profile.Profile),
		MapDestGroups:       make(map[string]*mo_group.Group),
		MapGroups:           make(map[string]*mo_group.Group),
		MapGroupMembers:     make(map[string][]*mo_group_member.Member),
		MapTeamFolders:      make(map[string]*mo_teamfolder.TeamFolder),
		MapNamespaces:       make(map[string]*mo_namespace.Namespace),
		MapNamespaceDetails: make(map[string]*mo_sharedfolder.SharedFolder),
		MapSharedFolders:    make(map[string]*mo_sharedfolder.SharedFolder),
		MapNamespaceMember:  make(map[string][]mo_sharedfolder_member.Member),
		ContextOpts:         &contextOpts{},
	}
	ctx.init(ctxExec)
	return ctx
}

type contextOpts struct {
	GroupsOnlyRelated bool                `json:"groups_only_related"`
	AdminSrc          *mo_profile.Profile `json:"admin_src"`
	AdminDst          *mo_profile.Profile `json:"admin_dst"`
}

type contextImpl struct {
	ctxExec            *app.ExecContext                           `json:"-"`
	storages           map[string]app_report.Report               `json:"-"`
	storagePath        string                                     `json:"-"`
	ctxTeamFolder      uc_teamfolder_mirror.Context               `json:"-"`
	MapNamespaceMember map[string][]mo_sharedfolder_member.Member `json:"-"`

	MapMembers          map[string]*mo_profile.Profile           `json:"members"`
	MapDestGroups       map[string]*mo_group.Group               `json:"dest_groups"`
	MapGroups           map[string]*mo_group.Group               `json:"groups"`
	MapGroupMembers     map[string][]*mo_group_member.Member     `json:"group_members"`
	MapTeamFolders      map[string]*mo_teamfolder.TeamFolder     `json:"team_folders"`
	MapNamespaces       map[string]*mo_namespace.Namespace       `json:"namespaces"`
	MapNamespaceDetails map[string]*mo_sharedfolder.SharedFolder `json:"namespace_details"`
	MapSharedFolders    map[string]*mo_sharedfolder.SharedFolder `json:"shared_folders"`
	ContextOpts         *contextOpts                             `json:"context_opts"`
}

func (z *contextImpl) init(ec *app.ExecContext) {
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
	z.storagePath = filepath.Join(ec.JobsPath(), "state")
	storages := make(map[string]app_report.Report)
	for _, tag := range storageTags {
		s := &app_report_json.JsonReport{
			ReportPath: filepath.Join(z.storagePath, tag),
		}
		if err := s.Init(ec); err != nil {
			ec.Log().Warn("Unable to store state", zap.String("tag", tag), zap.Error(err))
		}
		storages[tag] = s
	}
	z.storages = storages
	z.ctxExec = ec
}

func (z *contextImpl) StoreState() error {
	{
		if _, err := os.Stat(z.storagePath); os.IsNotExist(err) {
			err := os.MkdirAll(z.storagePath, 0755)
			if err != nil {
				z.ctxExec.Log().Error("unable to create state folder", zap.Error(err), zap.String("path", z.storagePath))
				return err
			}
		}
	}

	{
		b, err := json.Marshal(z)
		if err != nil {
			z.ctxExec.Log().Error("unable to marshal context", zap.Error(err))
			return err
		}
		err = ioutil.WriteFile(filepath.Join(z.storagePath, "context.json"), b, 0644)
		if err != nil {
			z.ctxExec.Log().Error("unable to store context", zap.Error(err))
			return err
		}
	}

	// namespace member
	{
		bnm, err := json.Marshal(z.MapNamespaceMember)
		if err != nil {
			z.ctxExec.Log().Error("unable to marshal context", zap.Error(err))
			return err
		}
		err = ioutil.WriteFile(filepath.Join(z.storagePath, "namespace_member.json"), bnm, 0644)
		if err != nil {
			z.ctxExec.Log().Error("unable to store context", zap.Error(err))
			return err
		}
	}

	// team folder
	{
		tb, err := uc_teamfolder_mirror.MarshalContext(z.ctxTeamFolder)
		if err != nil {
			z.ctxExec.Log().Error("unable to marshal team folder mirror context", zap.Error(err))
			return err
		}
		err = ioutil.WriteFile(filepath.Join(z.storagePath, "teamfolder_content.json"), tb, 0644)
		if err != nil {
			z.ctxExec.Log().Error("unable to store team folder mirror context", zap.Error(err))
			return err
		}

		z.ctxExec.Log().Info("Context preserved", zap.String("path", z.storagePath))
	}
	return nil
}

func (z *contextImpl) SetGroupsOnlyRelated(onlyRelated bool) {
	z.ContextOpts.GroupsOnlyRelated = onlyRelated
	if s, e := z.storages[storageTagOptions]; e {
		s.Report(z.ContextOpts)
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
	z.MapNamespaceDetails[sharedFolder.SharedFolderId] = sharedFolder
}

func (z *contextImpl) NamespaceDetails() (details map[string]*mo_sharedfolder.SharedFolder) {
	return z.MapNamespaceDetails
}

func (z *contextImpl) AddDestGroup(group *mo_group.Group) {
	if s, e := z.storages[storageTagDestGroups]; e {
		s.Report(group)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.MapDestGroups[group.GroupId] = group
}

func (z *contextImpl) DestGroups() (groups map[string]*mo_group.Group) {
	return z.MapDestGroups
}

func (z *contextImpl) FilterByMigrationTarget(members []*mo_profile.Profile) (filtered []*mo_profile.Profile) {
	filtered = make([]*mo_profile.Profile, 0)

	for _, member := range members {
		for _, marked := range z.MapMembers {
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
	z.MapNamespaces[namespace.NamespaceId] = namespace
}

func (z *contextImpl) Namespaces() (namespaces map[string]*mo_namespace.Namespace) {
	return z.MapNamespaces
}

func (z *contextImpl) GroupsOnlyRelated() bool {
	return z.ContextOpts.GroupsOnlyRelated
}

func (z *contextImpl) SetAdmins(src, dst *mo_profile.Profile) {
	z.ContextOpts.AdminSrc = src
	z.ContextOpts.AdminDst = dst
	if s, e := z.storages[storageTagOptions]; e {
		s.Report(z.ContextOpts)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
}

func (z *contextImpl) AdminSrc() *mo_profile.Profile {
	return z.ContextOpts.AdminSrc
}

func (z *contextImpl) AdminDst() *mo_profile.Profile {
	return z.ContextOpts.AdminDst
}

func (z *contextImpl) AddMember(member *mo_profile.Profile) {
	if s, e := z.storages[storageTagMembers]; e {
		s.Report(member)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.MapMembers[member.TeamMemberId] = member
}

func (z *contextImpl) AddGroup(group *mo_group.Group) {
	if s, e := z.storages[storageTagGroups]; e {
		s.Report(group)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.MapGroups[group.GroupId] = group
}

func (z *contextImpl) AddGroupMember(group *mo_group.Group, member *mo_group_member.Member) {
	if s, e := z.storages[storageTagGroupMembers]; e {
		s.Report(mo_group_member.NewGroupMember(group, member))
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}

	var members []*mo_group_member.Member
	if mem, e := z.MapGroupMembers[group.GroupId]; !e {
		members = append(mem, member)
	} else {
		members = make([]*mo_group_member.Member, 0)
		members = append(members, member)
	}
	z.MapGroupMembers[group.GroupId] = members
}

func (z *contextImpl) AddTeamFolder(teamFolder *mo_teamfolder.TeamFolder) {
	if s, e := z.storages[storageTagTeamFolders]; e {
		s.Report(teamFolder)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}
	z.MapTeamFolders[teamFolder.TeamFolderId] = teamFolder
}

func (z *contextImpl) AddSharedFolder(sharedFolder *mo_sharedfolder.SharedFolder) {
	if s, e := z.storages[storageTagSharedFolders]; e {
		s.Report(sharedFolder)
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}

	z.MapSharedFolders[sharedFolder.SharedFolderId] = sharedFolder
}

func (z *contextImpl) AddNamespaceMember(namespace *mo_namespace.Namespace, member mo_sharedfolder_member.Member) {
	if s, e := z.storages[storageTagNamespaceMembers]; e {
		s.Report(mo_namespace.NewNamespaceMember(namespace, member))
	} else {
		z.ctxExec.Log().Error("unable to find storage")
	}

	var members []mo_sharedfolder_member.Member
	if mem, e := z.MapNamespaceMember[namespace.NamespaceId]; e {
		members = append(mem, member)
	} else {
		members = make([]mo_sharedfolder_member.Member, 0)
		members = append(members, member)
	}
	z.MapNamespaceMember[namespace.NamespaceId] = members
}

func (z *contextImpl) Members() (members map[string]*mo_profile.Profile) {
	return z.MapMembers
}

func (z *contextImpl) Groups() (groups map[string]*mo_group.Group) {
	return z.MapGroups
}

func (z *contextImpl) GroupMembers(group *mo_group.Group) (members []*mo_group_member.Member) {
	if members, e := z.MapGroupMembers[group.GroupId]; e {
		return members
	} else {
		z.ctxExec.Log().Warn("Group members not found", zap.String("groupId", group.GroupId))
		return make([]*mo_group_member.Member, 0)
	}
}

func (z *contextImpl) SharedFolders() (folders map[string]*mo_sharedfolder.SharedFolder) {
	return z.MapSharedFolders
}

func (z *contextImpl) TeamFolders() (folders map[string]*mo_teamfolder.TeamFolder) {
	return z.MapTeamFolders
}

func (z *contextImpl) NamespaceMembers(namespaceId string) (members []mo_sharedfolder_member.Member) {
	if members, e := z.MapNamespaceMember[namespaceId]; e {
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
