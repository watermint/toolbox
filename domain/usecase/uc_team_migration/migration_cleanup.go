package uc_team_migration

import (
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"go.uber.org/zap"
	"strings"
)

func (z *migrationImpl) Cleanup(ctx Context) (err error) {

	// Clean up bridge permission
	z.log().Info("Cleanup: clean up permissions of shared folders")
	cleanupPermissionSharedFolder := func() error {
		for _, folder := range ctx.SharedFolders() {
			l := z.log().With(zap.String("name", folder.Name))
			if folder.IsTeamFolder || folder.IsInsideTeamFolder {
				l.Debug("Skip team folder & nested folder")
				continue
			}
			owner, sameTeam := z.isTeamOwnedSharedFolder(ctx, folder.SharedFolderId)
			if !sameTeam {
				l.Debug("Skip non team owned folder")
				continue
			}
			if owner.TeamMemberId == ctx.AdminSrc().TeamMemberId ||
				owner.Email == ctx.AdminSrc().Email {
				l.Debug("Skip shared folder which owned by src admin")
				continue
			}
			ownerMember, err := sv_member.New(z.ctxMgtDst).ResolveByEmail(owner.Email)
			if err != nil {
				l.Error("Unable to resolve folder owner user", zap.String("email", owner.Email), zap.Error(err))
				return err
			}

			members := ctx.NamespaceMembers(folder.SharedFolderId)
			isAdminExist := false
			for _, member := range members {
				if u, e := member.User(); e {
					if u.TeamMemberId == ctx.AdminSrc().TeamMemberId {
						isAdminExist = true
					}
				}
			}

			if isAdminExist {
				l.Debug("Admin exists on the folder. Keep admin permission")
			} else {
				l.Info("Clean up permission of the folder", zap.String("admin", ctx.AdminDst().Email))
				ctf := z.ctxFileDst.AsMemberId(ownerMember.TeamMemberId)
				svm := sv_sharedfolder_member.NewBySharedFolderId(ctf, folder.SharedFolderId)
				err = svm.Remove(sv_sharedfolder_member.RemoveByTeamMemberId(ctx.AdminDst().TeamMemberId))
				if err != nil {
					if strings.HasPrefix(api_util.ErrorSummary(err), "access_error/not_a_member") {
						l.Debug("Unable to remove admin due to original owner not yet activated", zap.Error(err))
					} else {
						l.Warn("Unable to remove admin from the folder", zap.String("admin", ctx.AdminDst().Email), zap.Error(err))
					}
				}
			}
		}
		return nil
	}
	if err = cleanupPermissionSharedFolder(); err != nil {
		return nil
	}

	return nil
}
