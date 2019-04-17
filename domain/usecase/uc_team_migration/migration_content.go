package uc_team_migration

import (
	"github.com/watermint/toolbox/domain/model/mo_device"
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/domain/service/sv_device"
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"go.uber.org/zap"
)

func (z *migrationImpl) Content(ctx Context) (err error) {
	// Detach desktop clients of migration target end users to prevent content inconsistency
	unlinkDesktopClients := func() error {
		devices, err := sv_device.New(z.ctxFileSrc).List()
		if err != nil {
			z.log().Error("Unable to retrieve list of devices of source team", zap.Error(err))
			return err
		}
		sourceMembers := make(map[string]*mo_profile.Profile)
		for _, member := range ctx.Members() {
			sourceMembers[member.TeamMemberId] = member
		}
		retryDevices := make([]*mo_device.Desktop, 0)
		for _, device := range devices {
			l := z.log().With(zap.String("sessionId", device.SessionId()), zap.String("tag", device.EntryTag()))
			d, e := device.Desktop()
			if !e {
				l.Debug("Skip non desktop sessions")
				continue
			}
			if m, e := sourceMembers[device.EntryTeamMemberId()]; e {
				l.Info("Unlink Desktop session", zap.String("member", m.Email), zap.String("platform", d.Platform), zap.String("updated", d.Updated))
				err = sv_device.New(z.ctxFileSrc.NoRetryOnError()).Revoke(d, sv_device.DeleteOnUnlink())
				if err != nil {
					l.Warn("Unable to unlink desktop session, retry later", zap.Error(err))
					retryDevices = append(retryDevices, d)
				}
			}
		}

		if len(retryDevices) > 0 {
			maxRetry := 3
			for retryCount := 0; retryCount < maxRetry; retryCount++ {
				moreRetryDevices := make([]*mo_device.Desktop, 0)
				l := z.log().With(zap.Int("retry", retryCount+1))
				l.Info("Retry", zap.Int("numDevices", len(retryDevices)))
				for _, d := range retryDevices {
					err = sv_device.New(z.ctxFileSrc.NoRetryOnError()).Revoke(d, sv_device.DeleteOnUnlink())
					if err != nil {
						l.Warn("Unable to unlink desktop session, retry later", zap.Error(err))
						moreRetryDevices = append(moreRetryDevices, d)
					}
				}
				if len(moreRetryDevices) < 1 {
					break
				}
				retryDevices = moreRetryDevices
			}
			if len(retryDevices) > 0 {
				for _, d := range retryDevices {
					z.log().Warn("Unable to unlink device",
						zap.String("sessionId", d.SessionId()),
						zap.String("teamMemberId", d.EntryTeamMemberId()),
						zap.String("hostname", d.HostName),
						zap.String("platform", d.Platform),
						zap.String("clientType", d.ClientType),
						zap.String("clientVersion", d.ClientVersion),
					)
				}
			}
		}

		return nil
	}
	if !ctx.KeepDesktopSessions() {
		z.log().Info("Content: unlink desktop clients of members to prevent inconsistency")
		if err = unlinkDesktopClients(); err != nil {
			return err
		}
	}

	// Mirror team folders
	z.log().Info("Content: mirroring team folder contents")
	if err = z.teamFolderMirror.Mirror(ctx.ContextTeamFolder(), uc_teamfolder_mirror.SkipVerify()); err != nil {
		return err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return err
	}

	return nil
}
