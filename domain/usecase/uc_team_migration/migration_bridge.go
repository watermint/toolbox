package uc_team_migration

import (
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"go.uber.org/zap"
)

func (z *migrationImpl) Bridge(ctx Context) (err error) {
	// bridge shared folders
	z.log().Info("Bridge: shared folders")
	bridgeSharedFolders := func() error {
		failedFolders := make([]*mo_sharedfolder.SharedFolder, 0)
		folderTargets := ctx.SharedFolders()
		for _, namespace := range ctx.Namespaces() {
			// skip team folder
			if namespace.NamespaceType != "shared_folder" {
				continue
			}

			if f, e := folderTargets[namespace.NamespaceId]; e {
				owner, sameTeam := z.isTeamOwnedSharedFolder(ctx, namespace.NamespaceId)
				if !sameTeam {
					z.log().Debug("Skip non team owned shared folder", zap.String("namespaceId", namespace.NamespaceId), zap.String("name", namespace.Name))
					continue
				}
				if owner.TeamMemberId == ctx.AdminSrc().TeamMemberId {
					z.log().Debug("Skip admin owned shared folder", zap.String("namespaceId", namespace.NamespaceId), zap.String("name", namespace.Name))
					continue
				}

				l := z.log().With(zap.String("SharedFolderId", f.SharedFolderId), zap.String("SharedFolderName", f.Name), zap.String("dstAdminId", ctx.AdminDst().TeamMemberId))

				l.Info("Bridge shared folder")
				var ctxFileAsMember api_context.Context
				ctxFileAsMember = z.ctxFileSrc.AsMemberId(owner.TeamMemberId)

				if f.PolicyMember == "team" {
					l.Info("Update member policy from `team` to `anyone`")
					ssf := sv_sharedfolder.New(ctxFileAsMember)
					_, err := ssf.UpdatePolicy(f.SharedFolderId, sv_sharedfolder.MemberPolicy("anyone"))
					if err != nil {
						l.Warn("Unable to change member policy", zap.Error(err))
					}
				}

				// add
				svc := sv_sharedfolder_member.NewBySharedFolderId(ctxFileAsMember, namespace.NamespaceId)
				err = svc.Add(sv_sharedfolder_member.AddByEmail(ctx.AdminDst().Email, sv_sharedfolder_member.LevelEditor), sv_sharedfolder_member.AddQuiet())

				if err != nil {
					_, err2 := sv_member.New(z.ctxMgtSrc).ResolveByEmail(owner.Email)
					if err2 != nil {
						l.Debug("Skip bridge: assuming the owner already transferred to dest team", zap.String("namespaceId", namespace.NamespaceId), zap.String("name", namespace.Name))
						continue
					}
					l.Error("Unable to bridge shared folder permission", zap.Error(err))
					failedFolders = append(failedFolders, f)
				}

				// transfer ownership
				if ctx.DontTransferFolderOwnership() {
					l.Debug("Skip transfer ownership")
				} else {
					err = sv_sharedfolder.New(ctxFileAsMember).Transfer(f, sv_sharedfolder.ToAccountId(ctx.AdminDst().AccountId))
					if err != nil {
						l.Warn("Unable to transfer ownership to admin", zap.Error(err))
						return err
					}
				}
			}
		}

		for _, f := range failedFolders {
			z.log().Warn("Bridge failure folder", zap.Any("folder", f))
		}
		//if len(failedFolders) > 0 {
		//	return errors.New("bridge failed in one or more folders")
		//}
		return nil
	}
	if err = bridgeSharedFolders(); err != nil {
		return err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}
