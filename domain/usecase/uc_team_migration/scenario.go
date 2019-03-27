package uc_team_migration

import (
	"errors"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"go.uber.org/zap"
	"strings"
)

const (
	peerNameActorTeamAAdmin01 = "test-migration-a01"
	peerNameActorTeamBAdmin01 = "test-migration-b01"
	peerNameActorIndividual01 = "test-migration-i01"

	// Prefix for team folders, shared folders, and groups
	testNamePrefix = "UCTM"
)

type Actors struct {
	TeamAAdmin01  string `json:"team_a_admin_01"`
	TeamAMember02 string `json:"team_a_member_02"`
	TeamAMember03 string `json:"team_a_member_03"`
	TeamAMember04 string `json:"team_a_member_04"`
	TeamBAdmin01  string `json:"team_b_admin_01"`
	Individual01  string `json:"individual_01"`
}

func NewScenario(ctxExe *app.ExecContext, actor *Actors) *Scenario {
	return &Scenario{
		ctxExec: ctxExe,
		actors:  actor,
	}
}

type Scenario struct {
	ctxExec       *app.ExecContext
	ctxTeamAFile  api_context.Context
	ctxTeamAMgmt  api_context.Context
	ctxTeamBFile  api_context.Context
	ctxTeamBMgmt  api_context.Context
	ctxIndividual api_context.Context
	actors        *Actors
}

func (z *Scenario) log() *zap.Logger {
	return z.ctxExec.Log()
}

func (z *Scenario) Auth() (err error) {
	z.log().Info("Auth: Team A File", zap.String("admin", z.actors.TeamAAdmin01))
	z.ctxTeamAFile, err = api_auth_impl.Auth(z.ctxExec,
		api_auth_impl.PeerName(peerNameActorTeamAAdmin01),
		api_auth_impl.BusinessFile(),
	)
	if err != nil {
		return err
	}

	z.log().Info("Auth: Team A Management", zap.String("admin", z.actors.TeamAAdmin01))
	z.ctxTeamAMgmt, err = api_auth_impl.Auth(z.ctxExec,
		api_auth_impl.PeerName(peerNameActorTeamAAdmin01),
		api_auth_impl.BusinessManagement(),
	)
	if err != nil {
		return err
	}

	z.log().Info("Auth: Team B File", zap.String("admin", z.actors.TeamBAdmin01))
	z.ctxTeamBFile, err = api_auth_impl.Auth(z.ctxExec,
		api_auth_impl.PeerName(peerNameActorTeamBAdmin01),
		api_auth_impl.BusinessFile(),
	)
	if err != nil {
		return err
	}

	z.log().Info("Auth: Team B Management", zap.String("admin", z.actors.TeamBAdmin01))
	z.ctxTeamBMgmt, err = api_auth_impl.Auth(z.ctxExec,
		api_auth_impl.PeerName(peerNameActorTeamBAdmin01),
		api_auth_impl.BusinessManagement(),
	)
	if err != nil {
		return err
	}

	z.log().Info("Auth: Individual full", zap.String("admin", z.actors.Individual01))
	z.ctxTeamAFile, err = api_auth_impl.Auth(z.ctxExec,
		api_auth_impl.PeerName(peerNameActorIndividual01),
		api_auth_impl.Full(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (z *Scenario) Cleanup() (err error) {
	// Display team information for confirmation
	displayTeamInfo := func(label string, ctx api_context.Context) error {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			return err
		}
		z.log().Info(label, zap.String("TeamId", info.TeamId), zap.String("Name", info.Name))
		return nil
	}
	displayTeamAdmin := func(label string, expectedEmail string, ctx api_context.Context) error {
		admin, err := sv_profile.NewTeam(ctx).Admin()
		if err != nil {
			return err
		}
		if strings.ToLower(expectedEmail) != strings.ToLower(admin.Email) {
			z.log().Warn("Admin Email address didn't match to given actor list", zap.String("found", admin.Email), zap.String("expected", expectedEmail))
			return errors.New("invalid admin & token")
		}
		z.log().Info(label, zap.String("TeamMemberId", admin.TeamMemberId), zap.String("email", admin.Email))
		return nil
	}
	if err = displayTeamAdmin("Team A: Admin (file token)", z.actors.TeamAAdmin01, z.ctxTeamAFile); err != nil {
		return err
	}
	if err = displayTeamAdmin("Team A: Admin (management token)", z.actors.TeamAAdmin01, z.ctxTeamAMgmt); err != nil {
		return err
	}
	if err = displayTeamInfo("Team A: Info", z.ctxTeamAMgmt); err != nil {
		return err
	}
	if err = displayTeamAdmin("Team B: Admin (file token)", z.actors.TeamBAdmin01, z.ctxTeamBFile); err != nil {
		return err
	}
	if err = displayTeamAdmin("Team B: Admin (management token)", z.actors.TeamBAdmin01, z.ctxTeamBMgmt); err != nil {
		return err
	}
	if err = displayTeamInfo("Team B: Info", z.ctxTeamBMgmt); err != nil {
		return err
	}
	displayIndividual := func() error {
		account, err := sv_profile.NewProfile(z.ctxIndividual).Current()
		if err != nil {
			return err
		}
		if strings.ToLower(account.Email) != strings.ToLower(z.actors.Individual01) {
			z.log().Warn("Individual01: Email address didn't match to given actor list")
			return errors.New("invalid individual01 token")
		}
		return nil
	}
	if err = displayIndividual(); err != nil {
		return err
	}

	z.log().Warn("Caution: Please do not run on production environment")
	if !z.ctxExec.Msg("usecase.team.migration.test.cleanup.confirmation").AskConfirm() {
		return
	}

	// Remove team folders (file token)
	removeTeamFolders := func(ctx api_context.Context) error {
		svc := sv_teamfolder.New(ctx)
		folders, err := svc.List()
		if err != nil {
			return err
		}
		for _, f := range folders {
			if strings.HasPrefix(strings.ToLower(f.Name), strings.ToLower(testNamePrefix)) {
				z.log().Info("Archive team folder", zap.String("name", f.Name))
				if _, err := svc.Archive(f); err != nil {
					z.log().Warn("Unable to archive", zap.String("name", f.Name), zap.Error(err))
					// continue
				}
				if err := svc.PermDelete(f); err != nil {
					z.log().Warn("Unable to perm delete", zap.String("name", f.Name), zap.Error(err))
				}
			}
		}
		return nil
	}

	z.log().Info("Team A: Clean up team folders")
	if err = removeTeamFolders(z.ctxTeamAFile); err != nil {
		z.log().Warn("Team A: Unable to clean up team folder(s)")
	}
	z.log().Info("Team B: Clean up team folders")
	if err = removeTeamFolders(z.ctxTeamBFile); err != nil {
		z.log().Warn("Team B: Unable to clean up team folder(s)")
	}

	// Clean up groups (mgmt token)
	removeGroups := func(ctx api_context.Context) error {
		svc := sv_group.New(ctx)
		groups, err := svc.List()
		if err != nil {
			return err
		}
		for _, g := range groups {
			if strings.HasPrefix(strings.ToLower(g.GroupName), strings.ToLower(uc_teamfolder_mirror.MirrorGroupNamePrefix)) ||
				strings.HasPrefix(strings.ToLower(g.GroupName), strings.ToLower(testNamePrefix)) {
				z.log().Info("Remove group", zap.String("groupName", g.GroupName))
				if err := svc.Remove(g.GroupId); err != nil {
					z.log().Warn("Unable to group", zap.Error(err))
				}
			}
		}
		return nil
	}

	z.log().Info("Team A: Clean up groups")
	if err = removeGroups(z.ctxTeamAMgmt); err != nil {
		z.log().Warn("Team A: Unable to clean up group(s)")
	}
	z.log().Info("Team B: Clean up groups")
	if err = removeGroups(z.ctxTeamBMgmt); err != nil {
		z.log().Warn("Team B: Unable to clean up group(s)")
	}

	// Remove shared folders
	//removeSharedFolders := func() error {
	//
	//}

	// Reverse transfer members if a member already in Team B
	reverseTransfer := func() error {
		// List members at Team B
		svcA := sv_member.New(z.ctxTeamAMgmt)
		svcB := sv_member.New(z.ctxTeamBMgmt)
		members, err := svcB.List()
		if err != nil {
			return err
		}

		targetEmails := make([]string, 0)
		targetEmails = append(targetEmails, strings.ToLower(z.actors.TeamAMember02))
		targetEmails = append(targetEmails, strings.ToLower(z.actors.TeamAMember03))
		targetEmails = append(targetEmails, strings.ToLower(z.actors.TeamAMember04))

		for _, member := range members {
			for _, t := range targetEmails {
				if strings.ToLower(member.Email) == t {
					z.log().Info("Downgrading member", zap.String("email", t))
					err = svcB.Remove(member, sv_member.Downgrade())
					if err != nil {
						z.log().Warn("Unable to downgrade member", zap.String("email", t))
						break
					}
					_, err = svcA.Add(t)
					if err != nil {
						z.log().Warn("Unable to invite member", zap.String("email", t))
					}

					break
				}
			}
		}
		return nil
	}
	if err = reverseTransfer(); err != nil {
		z.log().Warn("Reverse transfer failed")
	}

	return nil
}
