package sq_sharedfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"go.uber.org/zap"
)

type TransferOwnership struct {
	SharedFolderId string `json:"shared_folder_id"`
	FromUserEmail  string `json:"from_user_email"`
	ToUserEmail    string `json:"to_user_email"`
}

func (z *TransferOwnership) Do(src, dst service.Business) error {
	l := src.Log().With(zap.Any("task", z))

	l.Debug("Resolve from user")
	fromUser, err := src.Member().ResolveByEmail(z.FromUserEmail)
	if err != nil {
		l.Debug("Not found in src team", zap.Error(err))
		// fallback to dst team, the user might moved to dst team before the operation
		fromUser, err = dst.Member().ResolveByEmail(z.FromUserEmail)
		if err != nil {
			l.Debug("From user not found in both src and dst", zap.Error(err))
			return err
		}
	}

	l.Debug("Resolve to user")
	toUser, err := dst.Member().ResolveByEmail(z.ToUserEmail)
	if err != nil {
		l.Debug("To user not found", zap.Error(err))
		// Do not fallback to resolve in src. Because the user not yet moved
		return err
	}

	l.Debug("From and to users", zap.Any("fromUser", fromUser), zap.Any("toUser", toUser))

	sf, err := src.SharedFolderAsMember(fromUser.TeamMemberId).Resolve(z.SharedFolderId)
	if err != nil {
		l.Debug("Unable to find shared folder", zap.Error(err))
		return err
	}

	if sf.AccessType != mo_sharedfolder_member.AccessTypeOwner {
		l.Debug("From user is no longer an owner of the folder", zap.Any("folder", sf))
		return errors.New("from user is no longer an owner of the folder")
	}

	members, err := src.SharedFolderMemberAsMember(sf.SharedFolderId, fromUser.TeamMemberId).List()
	if err != nil {
		l.Debug("Unable to retrieve list of members", zap.Error(err))
		return err
	}

	found := false
	for _, m := range members {
		if u, e := m.User(); e {
			l.Debug("To user found", zap.Any("user", u))
			found = true
			break
		}
	}

	if !found {
		l.Debug("To user does not have an access to the shared folder")
		return errors.New("to user does not have an access to the shared folder")
	}

	l.Debug("Transfer")
	err = src.SharedFolderAsMember(fromUser.TeamMemberId).Transfer(sf, sv_sharedfolder.ToAccountId(toUser.AccountId))
	if err != nil {
		l.Debug("Unable to transfer ownership", zap.Error(err))
		return err
	}

	l.Debug("Success")

	return nil
}
