package uc_sharedfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/strings/es_mailaddr"
)

var (
	ErrorSharedFolderNotFound = errors.New("shared folder not found")
	ErrorNotSharedFolder      = errors.New("the folder is not a shared folder")
	ErrorNotGroupOrEmail      = errors.New("not a group or an email address")
)

type Resolver interface {
	Resolve(path mo_path.DropboxPath) (sf *mo_sharedfolder.SharedFolder, err error)

	GroupOrEmail(groupOrEmail string) (s *mo_member.MemberSelector, err error)
	GroupOrEmailForAddMember(groupOrEmail, accessLevel string) (m sv_sharedfolder_member.MemberAddOption, err error)
	GroupOrEmailForRemoveMember(groupOrEmail string) (m sv_sharedfolder_member.MemberRemoveOption, err error)
}

func NewResolver(ctx dbx_client.Client) Resolver {
	return &resImpl{
		ctx: ctx,
		svg: sv_group.NewCached(ctx),
	}
}

type resImpl struct {
	ctx dbx_client.Client
	svg sv_group.Group
}

func (z resImpl) GroupOrEmailForAddMember(groupOrEmail, accessLevel string) (m sv_sharedfolder_member.MemberAddOption, err error) {
	s, err := z.GroupOrEmail(groupOrEmail)
	switch {
	case s.Email != "":
		return sv_sharedfolder_member.AddByEmail(s.Email, accessLevel), nil
	case s.DropboxId != "":
		return sv_sharedfolder_member.AddByTeamMemberId(s.DropboxId, accessLevel), nil
	default:
		return nil, ErrorNotGroupOrEmail
	}
}

func (z resImpl) GroupOrEmailForRemoveMember(groupOrEmail string) (m sv_sharedfolder_member.MemberRemoveOption, err error) {
	s, err := z.GroupOrEmail(groupOrEmail)
	switch {
	case s.Email != "":
		return sv_sharedfolder_member.RemoveByEmail(s.Email), nil
	case s.DropboxId != "":
		return sv_sharedfolder_member.RemoveByGroupId(s.DropboxId), nil
	default:
		return nil, ErrorNotGroupOrEmail
	}
}

func (z resImpl) GroupOrEmail(groupOrEmail string) (s *mo_member.MemberSelector, err error) {
	group, err := z.svg.ResolveByName(groupOrEmail)
	if err == nil && group != nil {
		return mo_member.NewMemberSelectorDropboxId(group.GroupId), nil
	}
	if es_mailaddr.IsEmailAddr(groupOrEmail) {
		return mo_member.NewMemberSelectorEmail(groupOrEmail), nil
	}
	return nil, ErrorNotGroupOrEmail
}

func (z resImpl) Resolve(path mo_path.DropboxPath) (sf *mo_sharedfolder.SharedFolder, err error) {
	f1, err := sv_file.NewFiles(z.ctx).Resolve(path)
	if err != nil {
		return nil, err
	}
	f2, ok := f1.Folder()
	if !ok {
		return nil, ErrorSharedFolderNotFound
	}
	if f2.EntrySharedFolderId == "" {
		return nil, ErrorNotSharedFolder
	}
	sf, err = sv_sharedfolder.New(z.ctx).Resolve(f2.EntrySharedFolderId)
	if err != nil {
		return nil, err
	}
	return sf, nil
}
