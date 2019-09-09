package service

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_relocation"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/infra/api/api_context"
	"go.uber.org/zap"
)

type Business interface {
	// Purge caches
	Purge()
	Admin() *mo_profile.Profile
	Log() *zap.Logger

	Group() sv_group.Group
	GroupMember(groupId string) sv_group_member.GroupMember

	Member() sv_member.Member
	FilesAsMember(teamMemberId string) sv_file.Files
	RelocationAsMember(teamMemberId string) sv_file_relocation.Relocation
	SharedFolderAsMember(teamMemberId string) sv_sharedfolder.SharedFolder
	SharedFolderMountAsMember(teamMemberId string) sv_sharedfolder_mount.Mount
	SharedFolderMemberAsAdmin(sharedFolderId string) sv_sharedfolder_member.Member
	SharedFolderMemberAsMember(sharedFolderId string, teamMemberId string) sv_sharedfolder_member.Member
}

func New(ctxMgmt, ctxFile api_context.Context) (biz Business, err error) {
	svc := &businessImpl{
		ctxBusinessManagement: ctxMgmt,
		ctxBusinessFile:       ctxFile,
	}
	svc.Purge()
	adminMgmt, err := sv_profile.NewTeam(ctxMgmt).Admin()
	if err != nil {
		ctxMgmt.Log().Debug("Unable to determine admin", zap.Error(err))
		return nil, err
	}
	adminFile, err := sv_profile.NewTeam(ctxFile).Admin()
	if err != nil {
		ctxFile.Log().Debug("Unable to determine admin", zap.Error(err))
		return nil, err
	}
	if adminMgmt.TeamMemberId != adminFile.TeamMemberId {
		ctxMgmt.Log().Debug("Admins not match", zap.Any("adminMgmt", adminMgmt), zap.Any("adminFile", adminFile))
		return nil, errors.New("admins not match")
	}
	svc.adminProfile = adminFile
	return svc, nil
}

type businessImpl struct {
	ctxBusinessManagement api_context.Context
	ctxBusinessFile       api_context.Context
	group                 map[string]sv_group.Group
	member                map[string]sv_member.Member
	sharedFolderMember    map[string]sv_sharedfolder_member.Member
	adminProfile          *mo_profile.Profile
}

func (z *businessImpl) Purge() {
	z.group = make(map[string]sv_group.Group)
	z.member = make(map[string]sv_member.Member)
	z.sharedFolderMember = make(map[string]sv_sharedfolder_member.Member)
}

func (z *businessImpl) RelocationAsMember(teamMemberId string) sv_file_relocation.Relocation {
	ctx := z.ctxBusinessFile.AsMemberId(teamMemberId)
	return sv_file_relocation.New(ctx)
}

func (z *businessImpl) FilesAsMember(teamMemberId string) sv_file.Files {
	ctx := z.ctxBusinessFile.AsMemberId(teamMemberId)
	return sv_file.NewFiles(ctx)
}

func (z *businessImpl) SharedFolderAsMember(teamMemberId string) sv_sharedfolder.SharedFolder {
	ctx := z.ctxBusinessFile.AsMemberId(teamMemberId)
	return sv_sharedfolder.New(ctx)
}

func (z *businessImpl) SharedFolderMountAsMember(teamMemberId string) sv_sharedfolder_mount.Mount {
	ctx := z.ctxBusinessFile.AsMemberId(teamMemberId)
	return sv_sharedfolder_mount.New(ctx)
}

func (z *businessImpl) Admin() (admin *mo_profile.Profile) {
	return z.adminProfile
}

func (z *businessImpl) SharedFolderMemberAsMember(sharedFolderId string, teamMemberId string) sv_sharedfolder_member.Member {
	ctx := z.ctxBusinessFile.AsMemberId(teamMemberId)
	hash := ctx.Hash() + sharedFolderId
	if s, e := z.sharedFolderMember[hash]; e {
		return s
	}
	s := sv_sharedfolder_member.NewCached(ctx, sharedFolderId)
	z.sharedFolderMember[hash] = s
	return s
}

func (z *businessImpl) SharedFolderMemberAsAdmin(sharedFolderId string) sv_sharedfolder_member.Member {
	ctx := z.ctxBusinessFile.AsAdminId(z.adminProfile.TeamMemberId)
	hash := ctx.Hash() + sharedFolderId
	if s, e := z.sharedFolderMember[hash]; e {
		return s
	}
	s := sv_sharedfolder_member.NewCached(ctx, sharedFolderId)
	z.sharedFolderMember[hash] = s
	return s
}

func (z *businessImpl) Log() *zap.Logger {
	return z.ctxBusinessManagement.Log()
}

func (z *businessImpl) Member() sv_member.Member {
	hash := z.ctxBusinessManagement.Hash()
	if m, e := z.member[hash]; e {
		return m
	}
	m := sv_member.New(z.ctxBusinessManagement)
	z.member[hash] = m
	return m
}

func (z *businessImpl) GroupMember(groupId string) sv_group_member.GroupMember {
	return sv_group_member.NewByGroupId(z.ctxBusinessManagement, groupId)
}

func (z *businessImpl) Group() sv_group.Group {
	ch := z.ctxBusinessManagement.Hash()
	if g, e := z.group[ch]; e {
		return g
	}
	g := sv_group.NewCached(z.ctxBusinessManagement)
	z.group[ch] = g
	return g
}
