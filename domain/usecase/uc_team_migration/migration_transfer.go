package uc_team_migration

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_util"
	"go.uber.org/zap"
	"strings"
)

func (z *migrationImpl) Transfer(ctx Context) (err error) {
	// Convert accounts into Basic, and invite from new team
	z.log().Info("Transfer: transfer accounts")
	transferAccounts := func() error {
		svmSrc := sv_member.New(z.ctxMgtSrc.NoRetryOnError())
		svmDst := sv_member.New(z.ctxMgtDst)
		failedMembers := make([]*mo_profile.Profile, 0)
		for retry := 0; retry < 1000; retry++ {
			failedMembers = make([]*mo_profile.Profile, 0)
			for _, member := range ctx.Members() {
				if member.TeamMemberId == ctx.AdminSrc().TeamMemberId {
					z.log().Debug("Skip admin", zap.String("teamMemberId", member.TeamMemberId), zap.String("email", member.Email))
					continue
				}

				z.log().Info("Transfer: transferring member", zap.String("email", member.Email))
				l := z.log().With(zap.String("teamMemberId", member.TeamMemberId), zap.String("email", member.Email))
				l.Debug("Transferring account")

				ms, err := svmSrc.Resolve(member.TeamMemberId)
				if err != nil {
					if strings.HasPrefix(api_util.ErrorSummary(err), "id_not_found") {
						md, err := svmDst.ResolveByEmail(member.Email)
						if err != nil {
							l.Debug("Assume user detached but not yet invited", zap.Error(err))
							_, err = svmDst.Add(member.Email)
							if err != nil {
								l.Warn("Unable to invite member", zap.Error(err))
								failedMembers = append(failedMembers, member)
								continue
							}
						} else {
							l.Debug("Skip: the user already transferred", zap.Any("member", md))
						}
						continue
					}
					l.Warn("Unable to resolve existing member", zap.Error(err))
					continue
				}
				err = svmSrc.Remove(ms, sv_member.Downgrade())
				if err != nil {
					l.Warn("Unable to downgrade existing member", zap.Error(err))
					failedMembers = append(failedMembers, member)
					continue
				}

				_, err = svmDst.Add(member.Email)
				if err != nil {
					l.Warn("Unable to downgrade existing member", zap.Error(err))
					failedMembers = append(failedMembers, member)
					continue
				}

				// TODO: add role if the member is an admin
			}

			for _, m := range failedMembers {
				z.log().Warn("Unable to transfer member", zap.Int("retry", retry), zap.Any("member", m))
			}
			if len(failedMembers) < 1 {
				break
			}
		}
		if len(failedMembers) > 0 {
			return errors.New("one or more members could not be transferred")
		}

		return nil
	}
	if err = transferAccounts(); err != nil {
		return err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}
