package sq_sharedfolder

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type Mount struct {
	SharedFolderId string `json:"shared_folder_id"`
	UserEmail      string `json:"user_email"`
	MountPoint     string `json:"mount_point"`
}

func (z *Mount) Do(biz service.Business) error {
	l := biz.Log().With(zap.Any("task", z))

	l.Debug("Resolve user")
	user, err := biz.Member().ResolveByEmail(z.UserEmail)
	if err != nil {
		l.Debug("Unable to resolve user", zap.Error(err))
		return err
	}
	l = l.With(zap.Any("user", user))

	l.Debug("Resolve shared folder")
	sf, err := biz.SharedFolderAsMember(user.TeamMemberId).Resolve(z.SharedFolderId)
	if err != nil {
		l.Debug("Unable to resolve shared folder for user", zap.Error(err))
		return err
	}

	if sf.PathLower == "" {
		l.Debug("Mount")
		sf, err = biz.SharedFolderMountAsMember(user.TeamMemberId).Mount(sf)
		if err != nil {
			l.Debug("Unable to mount shared folder", zap.Error(err))
			return err
		}
	}

	if filepath.Dir(sf.PathLower) == "/" && sf.PathLower != strings.ToLower(z.MountPoint) {
		l.Debug("Move mount point")
		entry, err := biz.RelocationAsMember(user.TeamMemberId).Move(mo_path.NewPath(sf.PathLower), mo_path.NewPath(z.MountPoint))
		if err != nil {
			l.Debug("Unable to relocate", zap.Error(err))
			return err
		}
		l.Debug("Moved", zap.Any("entry", entry))
	}

	return nil
}
