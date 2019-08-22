package uc_team_migration

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_url"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	app2 "github.com/watermint/toolbox/legacy/app"
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

func (z *Actors) Emails() []string {
	return []string{
		z.TeamAAdmin01,
		z.TeamAMember02,
		z.TeamAMember03,
		z.TeamAMember04,
		z.TeamBAdmin01,
		z.Individual01,
	}
}

func (z *Actors) Members() []string {
	return []string{
		z.TeamAMember02,
		z.TeamAMember03,
		z.TeamAMember04,
	}
}

func NewScenario(ctxExe *app2.ExecContext, actor *Actors) *Scenario {
	return &Scenario{
		ctxExec: ctxExe,
		actors:  actor,
	}
}

type Scenario struct {
	ctxExec       *app2.ExecContext
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
	z.ctxIndividual, err = api_auth_impl.Auth(z.ctxExec,
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
	removeSharedFolders := func(label string, ctx api_context.Context) error {
		svc := sv_sharedfolder.New(ctx)
		svf := sv_file.NewFiles(ctx)
		folders, err := svc.List()
		if err != nil {
			return err
		}
		z.log().Info("Removing shared folder(s)", zap.String("label", label))
		for _, folder := range folders {
			l := z.log().With(zap.String("label", label), zap.String("Id", folder.SharedFolderId), zap.String("Name", folder.Name))
			if strings.HasPrefix(strings.ToLower(folder.Name), strings.ToLower(testNamePrefix)) {
				if folder.AccessType == sv_sharedfolder_member.LevelOwner {
					l.Info("Removing shared folder")
					err = svc.Remove(folder)
					if err != nil {
						l.Warn("Unable to remove shared folder")
						continue
					}

				} else {
					l.Info("Leave from shared folder")
					err = svc.Leave(folder)
					if err != nil {
						l.Warn("Unable to leave from shared folder")
						continue
					}
				}

				if folder.PathLower != "" {
					path := mo_path.NewPathDisplay(folder.PathLower)
					f, err := svf.Resolve(path)
					if err != nil {
						l.Debug("Path not found", zap.Error(err))
						continue
					}

					_, err = svf.Remove(mo_path.NewPathDisplay(f.PathDisplay()))
					if err != nil {
						l.Warn("Unable to remove file/folder", zap.String("path", f.PathDisplay()), zap.Error(err))
					}
				}
			}
		}

		return nil
	}
	if err = removeSharedFolders("individual01", z.ctxIndividual); err != nil {
		z.log().Warn("Individual01: Unable to remove shared folder(s), or leave from shared folder(s)")
	}
	removeSharedFoldersTeam := func(label string, ctxFile, ctxMgmt api_context.Context) error {
		members, err := sv_member.New(ctxMgmt).List()
		if err != nil {
			return err
		}
		for _, member := range members {
			for _, email := range z.actors.Emails() {
				if member.Email == email {
					err = removeSharedFolders(
						fmt.Sprintf("%s(%s)", label, email),
						ctxFile.AsMemberId(member.TeamMemberId),
					)
					if err != nil {
						z.log().Warn(label, zap.Error(err))
					}
					break
				}
			}
		}
		return nil
	}
	if err = removeSharedFoldersTeam("Team A", z.ctxTeamAFile, z.ctxTeamAMgmt); err != nil {
		z.log().Warn("Team A: Error occurred on one or more member")
	}
	if err = removeSharedFoldersTeam("Team B", z.ctxTeamBFile, z.ctxTeamBMgmt); err != nil {
		z.log().Warn("Team B: Error occurred on one or more member")
	}

	// Reverse transfer members if a member already in Team B
	reverseTransfer := func() error {
		// List members at Team B
		svcA := sv_member.New(z.ctxTeamAMgmt)
		svcB := sv_member.New(z.ctxTeamBMgmt)
		members, err := svcB.List()
		if err != nil {
			return err
		}

		for _, member := range members {
			for _, t := range z.actors.Members() {
				if strings.ToLower(member.Email) == t {
					z.log().Info("Downgrading member", zap.String("email", t))
					switch member.Status {
					case "active":
						err = svcB.Remove(member, sv_member.Downgrade())
						if err != nil {
							z.log().Warn("Unable to downgrade member", zap.String("email", t))
							break
						}

					case "invited":
						err = svcB.Remove(member)
						if err != nil {
							z.log().Warn("Unable to downgrade member", zap.String("email", t))
							break
						}

					default:
						z.log().Error("Unable to handle unexpected member status", zap.String("email", t), zap.String("status", member.Status))
						break

					}
					z.log().Info("Inviting member", zap.String("email", t))
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
	z.log().Info("Reverse transfer accounts")
	if err = reverseTransfer(); err != nil {
		z.log().Warn("Reverse transfer failed")
	}

	return nil
}

const (
	// Access: Individual01(owner), MemberA02 (editor)
	individualSharedFolderName = testNamePrefix + "SF-Prj-Individual"

	// Access: MemberA03(owner), MemberA02(editor), Individual01(editor)
	teamOwnedSharedFolderName = testNamePrefix + "SF-Prj-Team"

	// Team folder:
	// Sales (Group: Sales)
	// +- Sales East (Group: Sales East)
	// +- Sales West (MemberA04)
	// Eng (Group: Eng)
	// +- Eng East (Individual01)
	// +- Eng West

	// Group:
	// Sales (A02, A03)
	// Sales-East (A04)
	// Eng (A04)

	groupSalesName     = testNamePrefix + "G-Sales"
	groupSalesEastName = testNamePrefix + "G-Sales-East"
	groupSalesWestName = testNamePrefix + "G-Sales-West"
	groupEngName       = testNamePrefix + "G-Eng"

	teamFolderSalesName       = testNamePrefix + "TF-Sales"
	nestedFolderSalesEastName = testNamePrefix + "NF-Sales-East"
	nestedFolderSalesWestName = testNamePrefix + "NF-Sales-West"
	teamFolderEngName         = testNamePrefix + "TF-Eng"
	nestedFolderEngEastName   = testNamePrefix + "NF-Eng-East"
	nestedFolderEngWestName   = testNamePrefix + "NF-Eng-West"

	dummyFileUrl  = "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
	dummyFileName = testNamePrefix + "bootstrap.min.css"
)

func (z *Scenario) Create() (err error) {
	// Invite members
	inviteMembers := func() error {
		svm := sv_member.New(z.ctxTeamAMgmt)
		for _, member := range z.actors.Members() {
			_, err := svm.Add(member)
			if err != nil {
				z.log().Warn("Unable to invite member", zap.String("email", member), zap.Error(err))
			}
		}

		testees := make(map[string]*mo_member.Member)
		members, err := svm.List()
		if err != nil {
			return err
		}

		for _, member := range members {
			for _, email := range z.actors.Members() {
				if strings.ToLower(member.Email) == email {
					testees[member.TeamMemberId] = member
					break
				}
			}
		}

		for _, testee := range testees {
			switch testee.Status {
			case "invited":
				z.ctxExec.Msg("usecase.team.migration.test.prompt.please_accept_invitation").WithData(struct {
					Email string
				}{
					Email: testee.Email,
				}).AskConfirm()

			case "active":
				z.log().Debug("Testee is active", zap.String("email", testee.Email))

			default:
				z.log().Error("Unexpected testee status", zap.String("email", testee.Email), zap.String("status", testee.Status))
				return errors.New("unexpected testee status")
			}
		}

		return nil
	}
	if err = inviteMembers(); err != nil {
		return err
	}

	// Place files on the folder
	placeFileOnTheFolder := func(path mo_path.Path, ctx api_context.Context) error {
		svu := sv_file_url.New(ctx)
		e, err := svu.Save(path.ChildPath(dummyFileName), dummyFileUrl)
		if err != nil {
			z.log().Warn("Failed to create dummy file", zap.Error(err))
			return err
		}
		z.log().Info("Dummy file created", zap.String("path", e.PathDisplay()))
		return nil
	}

	// Individual shared folder
	individualSharedFolder := func() error {
		z.log().Info("Create individual shared folder", zap.String("name", individualSharedFolderName))
		svs := sv_sharedfolder.New(z.ctxIndividual)
		sf, err := svs.Create(mo_path.NewPath("/" + individualSharedFolderName))
		if err != nil {
			z.log().Error("Unable to create shared folder", zap.Error(err))
			return err
		}

		err = placeFileOnTheFolder(mo_path.NewPathDisplay(sf.PathLower), z.ctxIndividual)
		if err != nil {
			return err
		}

		z.log().Info("Share individual shared folder to MemberA02", zap.String("name", individualSharedFolderName), zap.String("invite", z.actors.TeamAMember02))
		svm := sv_sharedfolder_member.New(z.ctxIndividual, sf)
		err = svm.Add(sv_sharedfolder_member.AddByEmail(z.actors.TeamAMember02,
			sv_sharedfolder_member.LevelEditor))
		if err != nil {
			z.log().Error("Unable to share shared folder", zap.Error(err))
			return err
		}

		return nil
	}
	if err = individualSharedFolder(); err != nil {
		return err
	}

	// Team owned shared folder
	teamOwnedSharedFolder := func() error {
		l := z.log().With(zap.String("name", teamOwnedSharedFolderName))
		l.Info("Create team owned shared folder")
		svm := sv_member.New(z.ctxTeamAMgmt)
		ma02, err := svm.ResolveByEmail(z.actors.TeamAMember02)
		if err != nil {
			l.Error("Unable to resolve", zap.Error(err))
			return err
		}
		ma03, err := svm.ResolveByEmail(z.actors.TeamAMember03)
		if err != nil {
			l.Error("Unable to resolve", zap.Error(err))
			return err
		}

		cta03 := z.ctxTeamAFile.AsMemberId(ma03.TeamMemberId)
		svs03 := sv_sharedfolder.New(cta03)
		sf, err := svs03.Create(mo_path.NewPath("/" + teamOwnedSharedFolderName))
		if err != nil {
			l.Error("Unable to create shared folder", zap.Error(err))
			return err
		}

		err = placeFileOnTheFolder(mo_path.NewPathDisplay(sf.PathLower), cta03)
		if err != nil {
			z.log().Warn("Unable to place file", zap.Error(err))
		}

		svm03 := sv_sharedfolder_member.New(cta03, sf)

		l.Info("Share team owned shared folder to Individual")
		err = svm03.Add(sv_sharedfolder_member.AddByEmail(z.actors.Individual01,
			sv_sharedfolder_member.LevelEditor))
		if err != nil {
			z.log().Error("Unable to share shared folder", zap.Error(err))
			return err
		}

		l.Info("Updating member policy to team only")
		sf, err = svs03.UpdatePolicy(sf.SharedFolderId, sv_sharedfolder.MemberPolicy("team"))
		if err != nil {
			z.log().Warn("Unable to change member policy", zap.Error(err))
		}

		l.Info("Share team owned shared folder to MemberA02")
		err = svm03.Add(sv_sharedfolder_member.AddByEmail(ma02.Email,
			sv_sharedfolder_member.LevelEditor))
		if err != nil {
			z.log().Error("Unable to share shared folder", zap.Error(err))
			return err
		}

		return nil
	}
	if err = teamOwnedSharedFolder(); err != nil {
		return err
	}

	// Create groups
	groupsByName := make(map[string]*mo_group.Group)
	createGroups := func() error {
		svg := sv_group.New(z.ctxTeamAMgmt)
		groupNames := []string{
			groupSalesName,
			groupSalesEastName,
			groupSalesWestName,
			groupEngName,
		}

		z.log().Info("Create groups")
		for _, gn := range groupNames {
			l := z.log().With(zap.String("groupName", gn))
			l.Info("Create group")

			g, err := svg.Create(gn, sv_group.CompanyManaged())
			if err != nil {
				l.Error("Unable to create group", zap.Error(err))
				return err
			}

			groupsByName[gn] = g
		}
		return nil
	}
	if err = createGroups(); err != nil {
		return err
	}

	// Add members to groups
	addMemberToGroups := func() error {
		// Sales (A02, A03)
		// Sales-East (A04)
		// Eng (A04)

		salesGroup := groupsByName[groupSalesName]
		salesEastGroup := groupsByName[groupSalesEastName]
		engGroup := groupsByName[groupEngName]

		svm := sv_member.New(z.ctxTeamAMgmt)
		ma02, err := svm.ResolveByEmail(z.actors.TeamAMember02)
		if err != nil {
			z.log().Error("Unable to resolve", zap.Error(err))
			return err
		}
		ma03, err := svm.ResolveByEmail(z.actors.TeamAMember03)
		if err != nil {
			z.log().Error("Unable to resolve", zap.Error(err))
			return err
		}
		ma04, err := svm.ResolveByEmail(z.actors.TeamAMember04)
		if err != nil {
			z.log().Error("Unable to resolve", zap.Error(err))
			return err
		}

		{
			l := z.log().With(zap.String("groupName", salesGroup.GroupName), zap.String("member", ma02.Email))
			l.Info("Adding member to group")
			_, err = sv_group_member.New(z.ctxTeamAMgmt, salesGroup).Add(sv_group_member.ByTeamMemberId(ma02.TeamMemberId))
			if err != nil {
				l.Error("Unable to add", zap.Error(err))
				return err
			}
		}
		{
			l := z.log().With(zap.String("groupName", salesGroup.GroupName), zap.String("member", ma03.Email))
			l.Info("Adding member to group")
			_, err = sv_group_member.New(z.ctxTeamAMgmt, salesGroup).Add(sv_group_member.ByTeamMemberId(ma03.TeamMemberId))
			if err != nil {
				l.Error("Unable to add", zap.Error(err))
				return err
			}
		}
		{
			l := z.log().With(zap.String("groupName", salesEastGroup.GroupName), zap.String("member", ma04.Email))
			l.Info("Adding member to group")
			_, err = sv_group_member.New(z.ctxTeamAMgmt, salesEastGroup).Add(sv_group_member.ByTeamMemberId(ma04.TeamMemberId))
			if err != nil {
				l.Error("Unable to add", zap.Error(err))
				return err
			}
		}
		{
			l := z.log().With(zap.String("groupName", engGroup.GroupName), zap.String("member", ma04.Email))
			l.Info("Adding member to group")
			_, err = sv_group_member.New(z.ctxTeamAMgmt, engGroup).Add(sv_group_member.ByTeamMemberId(ma04.TeamMemberId))
			if err != nil {
				l.Error("Unable to add", zap.Error(err))
				return err
			}
		}
		return nil
	}
	if err = addMemberToGroups(); err != nil {
		return err
	}

	// Create team folders
	teamFoldersByName := make(map[string]*mo_teamfolder.TeamFolder)
	createTeamFolders := func() error {
		svt := sv_teamfolder.New(z.ctxTeamAFile)
		folderNames := []string{
			teamFolderSalesName,
			teamFolderEngName,
		}

		z.log().Info("Create team folders")
		for i, fn := range folderNames {
			l := z.log().With(zap.String("name", fn))
			l.Info("Create team folder")

			var opt sv_teamfolder.CreateOption
			if i%2 == 0 {
				opt = sv_teamfolder.SyncDefault()
			} else {
				opt = sv_teamfolder.SyncNoSync()
			}
			f, err := svt.Create(fn, opt)
			if err != nil {
				l.Error("Unable to create team folder", zap.Error(err))
				return err
			}
			teamFoldersByName[fn] = f
		}
		return nil
	}

	if err = createTeamFolders(); err != nil {
		return err
	}

	// Create nested folders, and apply permissions
	createNestedFolders := func() error {
		adminA01, err := sv_profile.NewTeam(z.ctxTeamAFile).Admin()
		if err != nil {
			return err
		}
		cta := z.ctxTeamAFile.AsAdminId(adminA01.TeamMemberId)
		svs := sv_sharedfolder.New(cta)

		// Team folder: Sales
		// Sales (Group: Sales)
		// +- Sales East (Group: Sales East)
		// +- Sales West (MemberA04)
		// Eng (Group: Eng)
		// +- Eng East (Individual01)
		// +- Eng West

		// Sales Team folder
		salesFolder, _ := teamFoldersByName[teamFolderSalesName]
		salesGroup, _ := groupsByName[groupSalesName]
		salesEastGroup, _ := groupsByName[groupSalesEastName]

		{
			z.log().Info("Add group to Sales team folder", zap.String("group", salesGroup.GroupName))
			ssm := sv_sharedfolder_member.NewBySharedFolderId(cta, salesFolder.TeamFolderId)
			err = ssm.Add(sv_sharedfolder_member.AddByGroup(salesGroup, sv_sharedfolder_member.LevelEditor))
			if err != nil {
				z.log().Error("Unable to add group", zap.Error(err))
			}
		}

		{
			z.log().Info("Create nested folder", zap.String("folder", nestedFolderSalesEastName))
			salesEast, err := svs.Create(mo_path.NewPath("ns:" + salesFolder.TeamFolderId + "/" + nestedFolderSalesEastName))
			if err != nil {
				return err
			}
			z.log().Info("Add group to nested folder", zap.String("group", salesEastGroup.GroupName))
			ssm := sv_sharedfolder_member.New(cta, salesEast)
			err = ssm.Add(sv_sharedfolder_member.AddByGroup(salesEastGroup, sv_sharedfolder_member.LevelEditor))
			if err != nil {
				z.log().Error("Unable to add group", zap.Error(err))
			}
		}

		{
			z.log().Info("Create nested folder", zap.String("folder", nestedFolderSalesWestName))
			salesWest, err := svs.Create(mo_path.NewPath("ns:" + salesFolder.TeamFolderId + "/" + nestedFolderSalesWestName))
			if err != nil {
				return err
			}

			ma04, err := sv_member.New(z.ctxTeamAMgmt).ResolveByEmail(z.actors.TeamAMember04)
			if err != nil {
				z.log().Error("Unable to resolve", zap.Error(err))
				return err
			}

			z.log().Info("Add member A04 to nested folder", zap.String("member", ma04.Email))
			ssm := sv_sharedfolder_member.New(cta, salesWest)
			err = ssm.Add(sv_sharedfolder_member.AddByTeamMemberId(ma04.TeamMemberId, sv_sharedfolder_member.LevelEditor))
			if err != nil {
				z.log().Error("Unable to add group", zap.Error(err))
			}
		}

		// Eng team folder
		engFolder, _ := teamFoldersByName[teamFolderEngName]
		engGroup, _ := groupsByName[groupEngName]

		{
			z.log().Info("Add group to Eng team folder", zap.String("group", engGroup.GroupName))
			ssm := sv_sharedfolder_member.NewBySharedFolderId(cta, engFolder.TeamFolderId)
			err = ssm.Add(sv_sharedfolder_member.AddByGroup(engGroup, sv_sharedfolder_member.LevelEditor))
			if err != nil {
				z.log().Error("Unable to add group", zap.Error(err))
			}
		}

		{
			z.log().Info("Create nested folder", zap.String("folder", nestedFolderEngEastName))
			engEast, err := svs.Create(mo_path.NewPath("ns:" + engFolder.TeamFolderId + "/" + nestedFolderEngEastName))
			if err != nil {
				return err
			}

			individual01, err := sv_profile.NewProfile(z.ctxIndividual).Current()

			z.log().Info("Add member Individual01 to nested folder", zap.String("member", individual01.Email))
			ssm := sv_sharedfolder_member.New(cta, engEast)
			err = ssm.Add(sv_sharedfolder_member.AddByEmail(individual01.Email, sv_sharedfolder_member.LevelEditor))
			if err != nil {
				z.log().Error("Unable to add member", zap.Error(err))
			}
		}

		return nil
	}
	if err = createNestedFolders(); err != nil {
		return err
	}

	return nil
}
