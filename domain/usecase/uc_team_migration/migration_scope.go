package uc_team_migration

import (
	"errors"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"go.uber.org/zap"
	"strings"
)

func (z *migrationImpl) Scope(opts ...ScopeOpt) (ctx Context, err error) {
	so := &scopeOpts{
		membersSpecifiedEmail:    make([]string, 0),
		teamFoldersSpecifiedName: make([]string, 0),
	}
	for _, o := range opts {
		o(so)
	}

	z.log().Info("Define scope")

	// Prepare migration context
	ctx = newContext(z.ctxExec)
	ctx.SetGroupsOnlyRelated(so.groupsOnlyRelated)
	ctx.SetKeepDesktopSessions(so.keepDesktopSessions)
	ctx.SetDontTransferFolderOwnership(so.dontTransferFolderOwnership)

	// validation
	if so.membersAllExceptAdmin && len(so.membersSpecifiedEmail) > 0 {
		z.log().Warn("Conflicted option `membersAllExceptAdmin` and `membersSpecifiedEmail`")
		return nil, errors.New("conflicted option")
	}
	if !so.membersAllExceptAdmin && len(so.membersSpecifiedEmail) < 1 {
		z.log().Warn("Please specify `memberAllExceptAdmin` or `membersSpecifiedEmail`")
		return nil, errors.New("not enough options")
	}
	if so.teamFoldersAll && len(so.teamFoldersSpecifiedName) > 0 {
		z.log().Warn("Conflicted option `teamFoldersAll` and `teamFoldersSpecifiedName`")
		return nil, errors.New("conflicted option")
	}
	if !so.teamFoldersAll && len(so.teamFoldersSpecifiedName) < 1 {
		z.log().Warn("Please specify `teamFoldersAll` or `teamFoldersSpecifiedName`")
		return nil, errors.New("not enough options")
	}

	// Identify admins
	identifyAdmins := func() error {
		adminSrc, err := sv_profile.NewTeam(z.ctxMgtSrc).Admin()
		if err != nil {
			return err
		}
		adminDst, err := sv_profile.NewTeam(z.ctxMgtDst).Admin()
		if err != nil {
			return err
		}
		z.log().Debug("Admins identified",
			zap.String("srcId", adminSrc.TeamMemberId),
			zap.String("srcEmail", adminSrc.Email),
			zap.String("dstId", adminDst.TeamMemberId),
			zap.String("dstEmail", adminDst.Email),
		)
		ctx.SetAdmins(adminSrc, adminDst)
		return nil
	}
	if err = identifyAdmins(); err != nil {
		return nil, err
	}

	// Define scope of members
	z.log().Info("Define scope: members")
	allMembers, err := sv_member.New(z.ctxMgtSrc).List()
	if err != nil {
		return nil, err
	}
	if so.membersAllExceptAdmin {
		for _, member := range allMembers {
			if ctx.AdminSrc().TeamMemberId != member.TeamMemberId {
				ctx.AddMember(member.Profile())
			} else {
				z.log().Debug("Skip admin", zap.String("teamMemberId", member.TeamMemberId), zap.String("email", member.Email))
			}
		}
	} else if len(so.membersSpecifiedEmail) > 0 {
		err = nil
		for _, email := range so.membersSpecifiedEmail {
			found := false
			emailLower := strings.ToLower(email)
			for _, member := range allMembers {
				if strings.ToLower(member.Email) == emailLower {
					ctx.AddMember(member.Profile())
					found = true
					break
				}
			}
			if !found {
				z.log().Warn("Member not found for email address", zap.String("email", email))
				err = errors.New("member not found")
			}
		}
		if err != nil {
			return nil, err
		}
	}
	if len(ctx.Members()) < 1 {
		z.log().Warn("No members found")
		return nil, errors.New("no member to migrate")
	}
	z.log().Debug("Members to migrate", zap.Int("count", len(ctx.Members())))

	// Define scope of team folders
	z.log().Info("Define scope: team folders")
	allFolders, err := sv_teamfolder.New(z.ctxFileSrc).List()
	if err != nil {
		return nil, err
	}
	if so.teamFoldersAll {
		for _, folder := range allFolders {
			if folder.Status == "active" {
				ctx.AddTeamFolder(folder)
			} else {
				z.log().Warn("Skip mirroring non active team folder", zap.String("name", folder.Name))
			}
		}
	} else if len(so.teamFoldersSpecifiedName) > 0 {
		err = nil
		for _, name := range so.teamFoldersSpecifiedName {
			found := false
			nameLower := strings.ToLower(name)
			for _, folder := range allFolders {
				if strings.ToLower(folder.Name) == nameLower {
					ctx.AddTeamFolder(folder)
					found = true
					break
				}
			}
			if !found {
				z.log().Warn("Team folder not found for name", zap.String("name", name))
				err = errors.New("team folder not found")
			}
		}
		if err != nil {
			return nil, err
		}
	}
	z.log().Debug("Team folders to migrate", zap.Int("count", len(ctx.TeamFolders())))

	// Team folder mirror
	z.log().Info("Define scope: mirroring content of team folders")
	prepTeamFolderMirror := func() error {
		names := make([]string, 0)
		for _, f := range ctx.TeamFolders() {
			names = append(names, f.Name)
		}
		ctxTf, err := z.teamFolderMirror.PartialScope(names)
		if err != nil {
			return err
		}
		ctx.SetContextTeamFolder(ctxTf)
		return nil
	}
	if err = prepTeamFolderMirror(); err != nil {
		return nil, err
	}

	// Store context
	if err = ctx.StoreState(); err != nil {
		return nil, err
	}

	return ctx, nil
}
