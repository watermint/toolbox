package cmd_teamfolder

import (
	"flag"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_file/compare"
	"github.com/watermint/toolbox/model/dbx_file/copy_ref"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/model/dbx_group/group_members"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_sharing"
	"github.com/watermint/toolbox/model/dbx_teamfolder"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
	"strings"
	"time"
)

type CmdTeamTeamFolderMirror struct {
	*cmd.SimpleCommandlet
	optFromAccount    string
	optToAccount      string
	optVerify         bool
	optAllTeamFolders bool
	report            report.Factory
	toTeamFolders     map[string]*dbx_teamfolder.TeamFolder
	fromTeamFolders   map[string]*dbx_teamfolder.TeamFolder
	toTeamAdminId     string
	fromTeamAdminId   string
	toTempGroupId     string
	fromFileApi       *dbx_api.Context
	toFileApi         *dbx_api.Context
	toMgmtApi         *dbx_api.Context
}

func (CmdTeamTeamFolderMirror) Name() string {
	return "mirror"
}

func (CmdTeamTeamFolderMirror) Desc() string {
	return "cmd.team.teamfolder.mirror.desc"
}

func (CmdTeamTeamFolderMirror) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamTeamFolderMirror) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descFromAccount := z.ExecContext.Msg("cmd.team.teamfolder.mirror.flag.from_account").Text()
	f.StringVar(&z.optFromAccount, "from-account", "", descFromAccount)

	descToAccount := z.ExecContext.Msg("cmd.team.teamfolder.mirror.flag.to_account").Text()
	f.StringVar(&z.optToAccount, "to-account", "", descToAccount)

	descVerify := z.ExecContext.Msg("cmd.team.teamfolder.mirror.flag.verify").Text()
	f.BoolVar(&z.optVerify, "verify", false, descVerify)

	descAll := z.ExecContext.Msg("cmd.team.teamfolder.mirror.flag.all").Text()
	f.BoolVar(&z.optAllTeamFolders, "all", false, descAll)
}

func (z *CmdTeamTeamFolderMirror) Exec(args []string) {
	if z.optFromAccount == "" ||
		z.optToAccount == "" {

		z.ExecContext.Msg("cmd.team.teamfolder.mirror.err.not_enough_params").TellError()
		return
	}
	if len(args) < 1 && !z.optAllTeamFolders {
		z.ExecContext.Msg("cmd.team.teamfolder.mirror.err.not_enough_arguments").TellError()
		return
	}
	var err error

	// Ask for FROM account authentication
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.prompt.ask_from_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optFromAccount,
	}).Tell()
	auFrom := dbx_auth.NewAuth(z.ExecContext, z.optFromAccount)
	z.fromFileApi, err = auFrom.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	// Ask for TO account authentication
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.prompt.ask_to_file_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optToAccount,
	}).Tell()
	auTo := dbx_auth.NewAuth(z.ExecContext, z.optToAccount)
	z.toFileApi, err = auTo.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	// Ask for TO account authentication
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.prompt.ask_to_mgmt_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optToAccount,
	}).Tell()
	z.toMgmtApi, err = auTo.Auth(dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}

	// Identify FROM team admin
	var fromAdminEmail, toAdminEmail string
	z.fromTeamAdminId, fromAdminEmail, err = z.identifyAdmin(z.fromFileApi)
	if err != nil {
		return
	}
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.identified_from_team_admin").WithData(struct {
		Alias        string
		TeamMemberId string
		Email        string
	}{
		Alias:        z.optFromAccount,
		TeamMemberId: z.fromTeamAdminId,
		Email:        fromAdminEmail,
	}).Tell()
	z.ExecContext.Log().Debug("from team admin", zap.String("teamMemberId", z.fromTeamAdminId), zap.String("email", fromAdminEmail))

	// Identify TO team admin
	z.toTeamAdminId, toAdminEmail, err = z.identifyAdmin(z.toFileApi)
	if err != nil {
		return
	}
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.identified_to_team_admin").WithData(struct {
		Alias        string
		TeamMemberId string
		Email        string
	}{
		Alias:        z.optFromAccount,
		TeamMemberId: z.fromTeamAdminId,
		Email:        fromAdminEmail,
	}).Tell()
	z.ExecContext.Log().Debug("to team admin", zap.String("teamMemberId", z.toTeamAdminId), zap.String("email", toAdminEmail))

	// Create temporary group for mirroring
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.create_tmp_group").WithData(struct {
		Alias string
	}{
		Alias: z.optToAccount,
	}).Tell()
	z.ExecContext.Log().Debug("create temporary group for mirroring")
	err = z.createTempGroup()
	if err != nil {
		return
	}

	// Adding admin user into temporary group
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.add_admin_to_tmp_group").WithData(struct {
		Email string
	}{
		Email: toAdminEmail,
	}).Tell()
	z.ExecContext.Log().Debug("adding admin user into temporary group")
	err = z.addAdminIntoTempGroup()
	if err != nil {
		// clean up temp group
		z.removeTempGroup()
		return
	}

	z.report.Init(z.ExecContext)
	z.report.Close()

	z.fromTeamFolders = z.listTeamFolders(z.fromFileApi)
	z.toTeamFolders = z.listTeamFolders(z.toFileApi)

	if z.optAllTeamFolders {
		for n, _ := range z.fromTeamFolders {
			z.mirrorTeamFolder(n)
		}
	} else {
		for _, n := range args {
			z.mirrorTeamFolder(n)
		}
	}

	// clean up
	z.removeTempGroup()
}

func (z *CmdTeamTeamFolderMirror) removeTempGroup() bool {
	remove := dbx_group.Remove{
		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			z.Log().Error("unable to clean up temporary group", zap.String("group_id", z.toTempGroupId), zap.Any("error", annotation))
			return true
		},
		OnSuccess: func() {
			z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.tmp_group_removed")
		},
	}
	return remove.Remove(z.toMgmtApi, z.toTempGroupId)
}

func (z *CmdTeamTeamFolderMirror) createTempGroup() error {
	groupName := fmt.Sprintf("%s-teamfolder-mirror-%x", app.AppName, time.Now().Unix())
	z.Log().Debug("temporary group name", zap.String("groupName", groupName))

	c := dbx_group.Create{
		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			z.Log().Warn("unable to create temporary group", zap.Any("error", annotation))
			return true
		},
		OnSuccess: func(group dbx_group.Group) {
			z.Log().Debug("group created", zap.String("group_id", group.GroupId))
			z.toTempGroupId = group.GroupId
			z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.tmp_group_created").WithData(struct {
				Name  string
				Alias string
			}{
				Name:  group.GroupName,
				Alias: z.optToAccount,
			}).Tell()
		},
	}
	return c.Create(z.toMgmtApi, groupName, dbx_group.ManagementTypeCompany)
}

func (z *CmdTeamTeamFolderMirror) addAdminIntoTempGroup() error {
	z.Log().Debug("adding admin", zap.String("group_id", z.toTempGroupId), zap.String("adminId", z.toTeamAdminId))
	add := group_members.Add{
		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			z.Log().Warn("unable to add admin into temporary group", zap.Any("error", annotation))
			return true
		},
		OnSuccess: func(group dbx_group.Group) {
			z.Log().Debug("group is ready", zap.String("group_id", group.GroupId))
		},
	}
	return add.AddMembers(z.toMgmtApi, z.toTempGroupId, []string{z.toTeamAdminId})
}

func (z *CmdTeamTeamFolderMirror) mirrorTeamFolder(name string) {
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.mirroring_team_folder").WithData(struct {
		Name string
	}{
		Name: name,
	}).Text()

	var err error
	ftf, e := z.fromTeamFolders[strings.ToLower(name)]
	if !e {
		// TODO: Report error, a team folder not found in from side
		return
	}
	ttf, e := z.toTeamFolders[strings.ToLower(name)]
	if !e {
		ttf, err = z.createTeamFolder(ftf.Name, z.toFileApi)
		if err != nil {
			return
		}
	}
	err = z.addTempGroupToTeamFolder(ttf)
	if err != nil {
		return
	}

	m := copy_ref.Mirror{
		FromAsMemberId:   z.fromTeamAdminId,
		FromApi:          z.fromFileApi,
		FromPath:         "/",
		FromAccountAlias: z.optFromAccount,
		FromNamespaceId:  ftf.TeamFolderId,
		ToAsMemberId:     z.toTeamAdminId,
		ToApi:            z.toFileApi,
		ToPath:           "/",
		ToNamespaceId:    ttf.TeamFolderId,
		ToAccountAlias:   z.optToAccount,
		ExecContext:      z.ExecContext,
	}
	m.MirrorAncestors()

	if z.optVerify {
		ba := compare.BetweenAccounts{
			ExecContext:       z.ExecContext,
			LeftAsMemberId:    z.fromTeamAdminId,
			LeftAccountAlias:  z.optFromAccount,
			LeftPath:          "/",
			LeftPathRoot:      dbx_api.NewPathRootNamespace(ttf.TeamFolderId),
			LeftApi:           z.fromFileApi,
			RightAsMemberId:   z.toTeamAdminId,
			RightAccountAlias: z.optToAccount,
			RightPath:         "/",
			RightPathRoot:     dbx_api.NewPathRootNamespace(ftf.TeamFolderId),
			RightApi:          z.toFileApi,
			OnDiff: func(diff compare.Diff) {
				z.report.Report(diff)
			},
		}
		ba.Compare()
	}
}

func (z *CmdTeamTeamFolderMirror) addTempGroupToTeamFolder(tf *dbx_teamfolder.TeamFolder) error {
	add := dbx_sharing.AddMembers{
		AsAdminId: z.toTeamAdminId,
		Context:   z.toFileApi,
		Quiet:     true,
	}
	return add.AddGroups(tf.TeamFolderId, []string{z.toTempGroupId}, dbx_sharing.AccessLevelEditor)
}

func (z *CmdTeamTeamFolderMirror) identifyAdmin(c *dbx_api.Context) (teamMemberId string, email string, err error) {
	admin, _, err := dbx_profile.AuthenticatedAdmin(c)
	if err != nil {
		return "", "", err
	} else {
		return admin.TeamMemberId, admin.Email, nil
	}
}

func (z *CmdTeamTeamFolderMirror) createTeamFolder(name string, acTo *dbx_api.Context) (tf *dbx_teamfolder.TeamFolder, err error) {
	// TODO: show progress
	cr := dbx_teamfolder.Create{
		OnError: z.DefaultErrorHandler,
		OnSuccess: func(teamFolder dbx_teamfolder.TeamFolder) {
			// TODO: show progress
			tf = &teamFolder
		},
	}
	err = cr.Create(acTo, name)
	if err != nil {
		switch e := err.(type) {
		case dbx_api.ApiError:
			tag := e.ErrorTag
			switch {
			case strings.HasPrefix(tag, "invalid_folder_name"),
				strings.HasPrefix(tag, "folder_name_reserved"):
				// TODO: show some err
				return

			case strings.HasPrefix(tag, "folder_name_already_used"):
				// ignore & proceed
				z.Log().Debug("ok") //TODO: detailed log
				return

			default:
				// TODO: show some err
				return
			}

		default:
			// TODO: log or show err
			return
		}
	}

	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.team_folder_created_on_to_team").WithData(struct {
		Name  string
		Alias string
	}{
		Name:  name,
		Alias: z.optToAccount,
	}).Tell()
	return
}

func (z *CmdTeamTeamFolderMirror) listTeamFolders(c *dbx_api.Context) map[string]*dbx_teamfolder.TeamFolder {
	folders := make(map[string]*dbx_teamfolder.TeamFolder)

	l := dbx_teamfolder.ListTeamFolder{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(teamFolder *dbx_teamfolder.TeamFolder) bool {
			// potentially unsafe for chars like Turkish `i/ı`
			folders[strings.ToLower(teamFolder.Name)] = teamFolder
			return true
		},
	}
	l.List(c)
	return folders
}
