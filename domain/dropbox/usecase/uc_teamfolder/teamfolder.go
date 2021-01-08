package uc_teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type AccessType string

const (
	AccessTypeOwner           AccessType = "owner"
	AccessTypeEditor          AccessType = "editor"
	AccessTypeViewer          AccessType = "viewer"
	AccessTypeViewerNoComment AccessType = "viewer_no_comment"
)

const (
	DefaultAdminWorkGroupName = "watermint-toolbox-admin"
)

var (
	ErrorUnableToIdentifyFolder = errors.New("unable to identify folder")
)

type TeamContent interface {
	// Get or create a team folder with the name.
	GetOrCreateTeamFolder(name string) (teamfolder TeamFolder, err error)

	// Get a team folder with the name.
	GetTeamFolder(name string) (teamfolder TeamFolder, err error)
}

type TeamFolder interface {
	// Add a member to the folder.
	// Nothing happens if a user already have access to the folder with given access type.
	// Returns an  error when failed to add a member to the folder.
	MemberAddUser(path mo_path.DropboxPath, accessType AccessType, memberEmail string) (err error)

	// Add a group to the folder.
	// Nothing happens if a group already have access to the folder with given access type.
	// Returns an  error when failed to add a group to the folder.
	MemberAddGroup(path mo_path.DropboxPath, accessType AccessType, groupName string) (err error)

	// Remove a member from the folder.
	// Nothing happens if a user is not in the folder.
	MemberRemoveUser(path mo_path.DropboxPath, memberEmail string) (err error)

	// Remove a group from the folder.
	// Nothing happens if a group is not in the folder.
	MemberRemoveGroup(path mo_path.DropboxPath, groupName string) (err error)

	// Update access inheritance setting for the folder.
	UpdateInheritance(path mo_path.DropboxPath, inherit bool) (folder *mo_sharedfolder.SharedFolder, err error)
}

func New(ctx dbx_context.Context, adminGroupName string) (tc TeamContent, err error) {
	admin, err := sv_profile.NewTeam(ctx).Admin()
	if err != nil {
		return nil, err
	}

	tc = &teamContentImpl{
		ctx:            ctx,
		stf:            sv_teamfolder.NewCached(ctx),
		sg:             sv_group.NewCached(ctx),
		adminGroupName: adminGroupName,
		admin:          admin,
	}
	return
}

type teamContentImpl struct {
	ctx            dbx_context.Context
	stf            sv_teamfolder.TeamFolder
	sg             sv_group.Group
	adminGroupName string
	admin          *mo_profile.Profile
}

func (z teamContentImpl) newTeamFolder(tf *mo_teamfolder.TeamFolder) TeamFolder {
	return &teamFolderImpl{
		ctx:              z.ctx,
		stf:              z.stf,
		sg:               z.sg,
		adminGroupName:   z.adminGroupName,
		admin:            z.admin,
		tf:               tf,
		cacheNamespaceId: make(map[string]string),
	}
}

func (z teamContentImpl) createTeamFolder(name string) (teamfolder TeamFolder, err error) {
	l := z.ctx.Log().With(esl.String("name", name))
	l.Debug("Create a team folder")
	tf, err := z.stf.Create(name)
	if err != nil {
		l.Debug("Unable to create a team folder", esl.Error(err))
		return nil, err
	}
	l.Debug("The team folder created", esl.Any("teamFolder", tf))
	return z.newTeamFolder(tf), nil
}

func (z teamContentImpl) GetTeamFolder(name string) (teamfolder TeamFolder, err error) {
	l := z.ctx.Log().With(esl.String("name", name))
	l.Debug("Resolve team folder by name")
	tf, err := z.stf.ResolveByName(name)
	if err != nil {
		l.Debug("Unable to resolve the team folder", esl.Error(err))
		return nil, err
	}
	return z.newTeamFolder(tf), nil
}

func (z teamContentImpl) GetOrCreateTeamFolder(name string) (teamfolder TeamFolder, err error) {
	l := z.ctx.Log().With(esl.String("name", name))
	l.Debug("Resolve team folder by name")
	tf, err := z.stf.ResolveByName(name)
	if sv_teamfolder.IsNotFound(err) {
		l.Debug("The team folder is not found. Try create a team folder")
		return z.createTeamFolder(name)
	} else if err != nil {
		l.Debug("Unable to resolve the team folder", esl.Error(err))
		return nil, err
	} else {
		l.Debug("The team folder found", esl.Any("teamFolder", tf))
		return z.newTeamFolder(tf), nil
	}
}

type teamFolderImpl struct {
	ctx              dbx_context.Context
	stf              sv_teamfolder.TeamFolder
	sg               sv_group.Group
	adminGroupName   string
	admin            *mo_profile.Profile
	tf               *mo_teamfolder.TeamFolder
	cacheAdminGroup  *mo_group.Group
	cacheNamespaceId map[string]string // path -> namespaceId
}

func (z *teamFolderImpl) logger() esl.Logger {
	return z.ctx.Log().With(esl.String("teamFolderId", z.tf.TeamFolderId), esl.String("teamFolderName", z.tf.Name))
}

func (z *teamFolderImpl) ctxAdmin() dbx_context.Context {
	return z.ctx.AsAdminId(z.admin.TeamMemberId).WithPath(dbx_context.Namespace(z.tf.TeamFolderId))
}

func (z *teamFolderImpl) ctxAdminTeamFolder() dbx_context.Context {
	return z.ctxAdmin().WithPath(dbx_context.Namespace(z.tf.TeamFolderId))
}

func (z *teamFolderImpl) adminGroup() (group *mo_group.Group, err error) {
	if z.cacheAdminGroup != nil {
		return z.cacheAdminGroup, nil
	}
	l := z.logger()
	l.Debug("Resolve the admin group")
	z.cacheAdminGroup, err = z.sg.ResolveByName(z.adminGroupName)
	if err != nil {
		l.Debug("Unable to resolve the admin group, try create an admin group")
		z.cacheAdminGroup, err = z.sg.Create(z.adminGroupName)
		if err != nil {
			l.Debug("Unable to create the group", esl.Error(err))
			return nil, err
		}
		l.Debug("The admin group created", esl.Any("group", z.cacheAdminGroup))
		return z.cacheAdminGroup, nil
	}

	l.Debug("The admin group resolved", esl.Any("group", z.cacheAdminGroup))
	return
}

func (z *teamFolderImpl) namespaceIdForPath(path mo_path.DropboxPath, createIfNotExist bool) (namespaceId string, err error) {
	l := z.logger().With(esl.String("path", path.Path()))
	if path.IsRoot() {
		l.Debug("The path is root")
		return z.tf.TeamFolderId, nil
	}

	if nsId, ok := z.cacheNamespaceId[path.Path()]; ok {
		l.Debug("Namespace retrieved from the cache", esl.String("nsId", nsId))
		return nsId, nil
	}

	nestedMeta, err := sv_file.NewFiles(z.ctxAdminTeamFolder()).Resolve(path)
	if err != nil {
		l.Debug("Unable to resolve nested folder", esl.Error(err))
		return "", err
	}
	nestedNsId := nestedMeta.Concrete().SharedFolderId
	if nestedNsId != "" {
		l.Debug("Nested folder found", esl.Any("nestedFolderMeta", nestedMeta))
		z.cacheNamespaceId[path.Path()] = nestedNsId
		return "", ErrorUnableToIdentifyFolder
	}

	if !createIfNotExist {
		l.Debug("Unable to find the folder")
		return "", ErrorUnableToIdentifyFolder
	}

	l.Debug("Try create nested folder")
	nested, err := sv_sharedfolder.New(z.ctxAdminTeamFolder()).Create(path)
	de := dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Debug("The nested folder created", esl.Any("nested", nested))
		z.cacheNamespaceId[path.Path()] = nested.SharedFolderId
		return nested.SharedFolderId, nil

	case de.BadPath().IsAlreadyShared():
		l.Debug("Retry resolve")
		nestedMeta, err = sv_file.NewFiles(z.ctxAdminTeamFolder()).Resolve(path)
		if err != nil {
			l.Debug("Unable to resolve nested folder", esl.Error(err))
			return "", err
		}
		nestedNsId := nestedMeta.Concrete().SharedFolderId
		z.cacheNamespaceId[path.Path()] = nestedNsId
		return nestedNsId, nil

	default:
		l.Debug("Unable to create a nested folder", esl.Error(err))
		return "", err
	}
}

func (z *teamFolderImpl) MemberAddUser(path mo_path.DropboxPath, accessType AccessType, memberEmail string) (err error) {
	l := z.logger().With(esl.String("path", path.Path()), esl.String("accessType", string(accessType)), esl.String("memberEmail", memberEmail))
	l.Debug("Add an user")
	nsId, err := z.namespaceIdForPath(path, true)
	if err != nil {
		l.Debug("namespace lookup failed", esl.Error(err))
		return err
	}

	err = sv_sharedfolder_member.NewBySharedFolderId(z.ctxAdminTeamFolder(), nsId).
		Add(sv_sharedfolder_member.AddByEmail(memberEmail, string(accessType)))
	if err != nil {
		l.Debug("Unable to add a member", esl.Error(err))
		return err
	}

	l.Debug("The member is successfully added")
	return nil
}

func (z *teamFolderImpl) MemberAddGroup(path mo_path.DropboxPath, accessType AccessType, groupName string) (err error) {
	l := z.logger().With(esl.String("path", path.Path()), esl.String("accessType", string(accessType)), esl.String("groupName", groupName))
	l.Debug("Add a group")
	nsId, err := z.namespaceIdForPath(path, true)
	if err != nil {
		l.Debug("namespace lookup failed", esl.Error(err))
		return err
	}

	group, err := z.sg.ResolveByName(groupName)
	if err != nil {
		l.Debug("Unable to lookup the group", esl.Error(err))
		return err
	}

	err = sv_sharedfolder_member.NewBySharedFolderId(z.ctxAdminTeamFolder(), nsId).
		Add(sv_sharedfolder_member.AddByGroupId(group.GroupId, string(accessType)))
	if err != nil {
		l.Debug("Unable to add a group", esl.Error(err))
		return err
	}

	l.Debug("The group is successfully added")
	return nil
}

func (z *teamFolderImpl) MemberRemoveUser(path mo_path.DropboxPath, memberEmail string) (err error) {
	l := z.logger().With(esl.String("path", path.Path()), esl.String("memberEmail", memberEmail))
	l.Debug("Remove an user")
	nsId, err := z.namespaceIdForPath(path, false)
	if err != nil {
		l.Debug("namespace lookup failed", esl.Error(err))
		return err
	}

	err = sv_sharedfolder_member.NewBySharedFolderId(z.ctxAdminTeamFolder(), nsId).
		Remove(sv_sharedfolder_member.RemoveByEmail(memberEmail))
	if err != nil {
		l.Debug("Unable to add a member", esl.Error(err))
		return err
	}

	l.Debug("The member is successfully added")
	return nil
}

func (z *teamFolderImpl) MemberRemoveGroup(path mo_path.DropboxPath, groupName string) (err error) {
	l := z.logger().With(esl.String("path", path.Path()), esl.String("groupName", groupName))
	l.Debug("Remove a group")
	nsId, err := z.namespaceIdForPath(path, false)
	if err != nil {
		l.Debug("namespace lookup failed", esl.Error(err))
		return err
	}

	group, err := z.sg.ResolveByName(groupName)
	if err != nil {
		l.Debug("Unable to lookup the group", esl.Error(err))
		return err
	}

	err = sv_sharedfolder_member.NewBySharedFolderId(z.ctxAdminTeamFolder(), nsId).
		Remove(sv_sharedfolder_member.RemoveByGroupId(group.GroupId))
	if err != nil {
		l.Debug("Unable to add a group", esl.Error(err))
		return err
	}

	l.Debug("The group is successfully added")
	return nil
}

func (z *teamFolderImpl) UpdateInheritance(path mo_path.DropboxPath, inherit bool) (folder *mo_sharedfolder.SharedFolder, err error) {
	panic("implement me")
}
