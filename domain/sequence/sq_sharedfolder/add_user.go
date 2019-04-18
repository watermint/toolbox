package sq_sharedfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"go.uber.org/zap"
)

type AddUser struct {
	SharedFolderId string `json:"shared_folder_id"`
	UserEmail      string `json:"user_email"`
	AccessLevel    string `json:"access_level"`
}

func (z *AddUser) Do(biz service.Business) error {
	l := biz.Log().With(zap.Any("task", z))

	l.Debug("Resolve shared folder members")
	members, err := biz.SharedFolderMemberAsAdmin(z.SharedFolderId).List()
	if err != nil {
		l.Debug("Unable to list member", zap.Error(err))
		return err
	}

	l.Debug("Lookup owner, or editor")
	var owner, editor *mo_sharedfolder_member.User = nil, nil

	for _, m := range members {
		if u, e := m.User(); e {
			switch {
			case u.SameTeam && u.AccessType() == mo_sharedfolder_member.AccessTypeOwner:
				owner = u
				break
			case u.SameTeam && u.AccessType() == mo_sharedfolder_member.AccessTypeEditor:
				editor = u
			}
		}
	}

	var svm sv_sharedfolder_member.Member
	switch {
	case owner != nil:
		l = l.With(zap.Any("owner", owner))
		svm = biz.SharedFolderMemberAsMember(z.SharedFolderId, owner.TeamMemberId)

	case editor != nil:
		l = l.With(zap.Any("editor", editor))
		svm = biz.SharedFolderMemberAsMember(z.SharedFolderId, editor.TeamMemberId)

	default:
		l.Debug("Both owner & editor not found")
		return errors.New("both owner and editor not found in same team")
	}

	l.Debug("Add group")
	err = svm.Add(
		sv_sharedfolder_member.AddByEmail(z.UserEmail, z.AccessLevel),
		sv_sharedfolder_member.AddQuiet(),
	)

	if err != nil {
		l.Debug("Unable to add user as member", zap.Error(err))
		return err
	}

	l.Debug("Success")
	return nil
}
