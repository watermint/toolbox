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

	optSrcTeamAlias   string
	optDstTeamAlias   string
	optVerify         bool
	optAllTeamFolders bool

	report report.Factory

	srcTeamFolders map[string]*dbx_teamfolder.TeamFolder
	srcTeamAdminId string
	srcTempGroup   *dbx_group.Group
	srcFileApi     *dbx_api.Context
	srcMgmtApi     *dbx_api.Context

	dstTeamFolders map[string]*dbx_teamfolder.TeamFolder
	dstTeamAdminId string
	dstTempGroup   *dbx_group.Group
	dstFileApi     *dbx_api.Context
	dstMgmtApi     *dbx_api.Context
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

	descFromAccount := z.ExecContext.Msg("cmd.team.teamfolder.mirror.flag.src_account").T()
	f.StringVar(&z.optSrcTeamAlias, "src-account", "mirror-src", descFromAccount)

	descToAccount := z.ExecContext.Msg("cmd.team.teamfolder.mirror.flag.dst_account").T()
	f.StringVar(&z.optDstTeamAlias, "dest-account", "mirror-dst", descToAccount)

	descVerify := z.ExecContext.Msg("cmd.team.teamfolder.mirror.flag.verify").T()
	f.BoolVar(&z.optVerify, "verify", false, descVerify)

	descAll := z.ExecContext.Msg("cmd.team.teamfolder.mirror.flag.all").T()
	f.BoolVar(&z.optAllTeamFolders, "all", false, descAll)
}

func (z *CmdTeamTeamFolderMirror) Exec(args []string) {
	if z.optSrcTeamAlias == "" ||
		z.optDstTeamAlias == "" {

		z.ExecContext.Msg("cmd.team.teamfolder.mirror.err.not_enough_params").TellError()
		return
	}
	if z.optSrcTeamAlias == z.optDstTeamAlias {
		z.ExecContext.Msg("cmd.team.teamfolder.mirror.err.same_team").TellError()
		return
	}
	if len(args) < 1 && !z.optAllTeamFolders {
		z.ExecContext.Msg("cmd.team.teamfolder.mirror.err.not_enough_arguments").TellError()
		return
	}
	var err error

	// Ask for SRC account authentication
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.prompt.ask_src_file_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optSrcTeamAlias,
	}).Tell()
	auFrom := dbx_auth.NewAuth(z.ExecContext, z.optSrcTeamAlias)
	z.srcFileApi, err = auFrom.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	// Ask for SRC account authentication
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.prompt.ask_src_mgmt_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optSrcTeamAlias,
	}).Tell()
	z.srcMgmtApi, err = auFrom.Auth(dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}

	// Ask for DST account authentication
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.prompt.ask_dst_file_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optDstTeamAlias,
	}).Tell()
	auTo := dbx_auth.NewAuth(z.ExecContext, z.optDstTeamAlias)
	z.dstFileApi, err = auTo.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	// Ask for DST account authentication
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.prompt.ask_dst_mgmt_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optDstTeamAlias,
	}).Tell()
	z.dstMgmtApi, err = auTo.Auth(dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}

	// Identify SRC team admin
	var srcAdminEmail, dstAdminEmail string
	z.srcTeamAdminId, srcAdminEmail, err = z.identifyAdmin(z.srcFileApi)
	if err != nil {
		return
	}
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.identified_from_team_admin").WithData(struct {
		Alias        string
		TeamMemberId string
		Email        string
	}{
		Alias:        z.optSrcTeamAlias,
		TeamMemberId: z.srcTeamAdminId,
		Email:        srcAdminEmail,
	}).Tell()
	z.ExecContext.Log().Debug("from team admin", zap.String("teamMemberId", z.srcTeamAdminId), zap.String("email", srcAdminEmail))

	// Identify DST team admin
	z.dstTeamAdminId, dstAdminEmail, err = z.identifyAdmin(z.dstFileApi)
	if err != nil {
		return
	}
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.identified_to_team_admin").WithData(struct {
		Alias        string
		TeamMemberId string
		Email        string
	}{
		Alias:        z.optSrcTeamAlias,
		TeamMemberId: z.srcTeamAdminId,
		Email:        srcAdminEmail,
	}).Tell()
	z.ExecContext.Log().Debug("to team admin", zap.String("teamMemberId", z.dstTeamAdminId), zap.String("email", dstAdminEmail))

	// Create temporary group for mirroring (SRC)
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.create_tmp_group").WithData(struct {
		Alias string
	}{
		Alias: z.optSrcTeamAlias,
	}).Tell()
	z.ExecContext.Log().Debug("create temporary group for mirroring")
	z.srcTempGroup, err = z.createTempGroup(z.srcMgmtApi, z.optSrcTeamAlias)
	if err != nil || z.srcTempGroup == nil {
		z.Log().Debug("failed create temp group (src)", zap.Any("group", z.srcTempGroup), zap.Error(err))
		return
	}
	defer z.removeTempGroup(z.srcMgmtApi, z.srcTempGroup.GroupId)

	// Adding admin user into temporary group (SRC)
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.add_admin_to_tmp_group").WithData(struct {
		Email string
	}{
		Email: srcAdminEmail,
	}).Tell()
	z.ExecContext.Log().Debug("adding admin user into temporary group")
	err = z.addAdminIntoTempGroup(z.srcMgmtApi, z.srcTempGroup.GroupId, z.srcTeamAdminId, z.optSrcTeamAlias)
	if err != nil {
		return
	}

	// Create temporary group for mirroring (DST)
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.create_tmp_group").WithData(struct {
		Alias string
	}{
		Alias: z.optDstTeamAlias,
	}).Tell()
	z.ExecContext.Log().Debug("create temporary group for mirroring")
	z.dstTempGroup, err = z.createTempGroup(z.dstMgmtApi, z.optDstTeamAlias)
	if err != nil || z.dstTempGroup == nil {
		z.Log().Debug("failed create temp group (dst)", zap.Any("group", z.dstTempGroup), zap.Error(err))
		return
	}
	defer z.removeTempGroup(z.dstMgmtApi, z.dstTempGroup.GroupId)

	// Adding admin user into temporary group (DST)
	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.add_admin_to_tmp_group").WithData(struct {
		Email string
	}{
		Email: dstAdminEmail,
	}).Tell()
	z.ExecContext.Log().Debug("adding admin user into temporary group")
	err = z.addAdminIntoTempGroup(z.dstMgmtApi, z.dstTempGroup.GroupId, z.dstTeamAdminId, z.optDstTeamAlias)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	z.report.Close()

	z.srcTeamFolders = z.listTeamFolders(z.srcFileApi)
	z.dstTeamFolders = z.listTeamFolders(z.dstFileApi)

	if z.optAllTeamFolders {
		for _, t := range z.srcTeamFolders {
			z.mirrorTeamFolder(t.Name)
		}
	} else {
		for _, n := range args {
			z.mirrorTeamFolder(n)
		}
	}
}

func (z *CmdTeamTeamFolderMirror) removeTempGroup(api *dbx_api.Context, groupId string) bool {
	remove := dbx_group.Remove{
		OnError: func(err error) bool {
			z.Log().Error("unable to clean up temporary group", zap.String("group_id", groupId), zap.Error(err))
			return true
		},
		OnSuccess: func() {
			z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.tmp_group_removed")
		},
	}
	return remove.Remove(api, groupId)
}

func (z *CmdTeamTeamFolderMirror) createTempGroup(api *dbx_api.Context, alias string) (createdGroup *dbx_group.Group, err error) {
	groupName := fmt.Sprintf("%s-teamfolder-mirror-%x", app.AppName, time.Now().Unix())
	z.Log().Debug("temporary group name", zap.String("groupName", groupName), zap.String("alias", alias))

	c := dbx_group.Create{
		OnError: func(err error) bool {
			z.Log().Warn("unable to create temporary group", zap.Error(err))
			return true
		},
		OnSuccess: func(group dbx_group.Group) {
			z.Log().Debug("group created", zap.String("group_id", group.GroupId))
			createdGroup = &group
			z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.tmp_group_created").WithData(struct {
				Name  string
				Alias string
			}{
				Name:  group.GroupName,
				Alias: alias,
			}).Tell()
		},
	}
	err = c.Create(api, groupName, dbx_group.ManagementTypeCompany)
	return
}

func (z *CmdTeamTeamFolderMirror) addAdminIntoTempGroup(api *dbx_api.Context, groupId, adminId, alias string) error {
	log := z.Log().With(zap.String("group_id", groupId), zap.String("admin_id", adminId), zap.String("alias", alias))
	log.Debug("adding admin")
	add := group_members.Add{
		OnError: func(err error) bool {
			log.Warn("unable to add admin into temporary group", zap.Error(err))
			return true
		},
		OnSuccess: func(group dbx_group.Group) {
			log.Debug("group is ready")
		},
	}
	return add.AddMembers(api, groupId, []string{adminId})
}

func (z *CmdTeamTeamFolderMirror) mirrorTeamFolder(name string) {
	var err error
	srcTeamFolder, e := z.srcTeamFolders[strings.ToLower(name)]
	if !e {
		z.ExecContext.Msg("cmd.team.teamfolder.mirror.err.team_folder_not_found").WithData(struct {
			Name string
		}{
			Name: name,
		}).TellError()
		return
	}

	if srcTeamFolder.Status != dbx_teamfolder.StatusActive {
		z.ExecContext.Msg("cmd.team.teamfolder.mirror.prompt.skipped_archived_folder").WithData(struct {
			Name string
		}{
			Name: srcTeamFolder.Name,
		}).Tell()
		return
	}

	z.ExecContext.Msg("cmd.team.teamfolder.mirror.progress.mirroring_team_folder").WithData(struct {
		Name string
	}{
		Name: name,
	}).Tell()

	dstTeamFolder, e := z.dstTeamFolders[strings.ToLower(name)]
	if !e {
		dstTeamFolder, err = z.createTeamFolder(srcTeamFolder.Name, z.dstFileApi)
		if err != nil {
			z.Log().Error("unable to create team folder (DST)", zap.Error(err))
			return
		}
	}

	if dstTeamFolder.Status != dbx_teamfolder.StatusActive {
		z.Log().Error("destination team folder is not active.")
		return
	}

	addSrc := dbx_sharing.AddMembers{
		AsAdminId: z.srcTeamAdminId,
		Context:   z.srcFileApi,
		Quiet:     true,
	}
	err = addSrc.Add(srcTeamFolder.TeamFolderId, []string{z.srcTempGroup.GroupId}, dbx_sharing.AccessLevelEditor)
	if err != nil {
		z.Log().Error("unable to add group to team folder (SRC)", zap.String("groupName", z.srcTempGroup.GroupName), zap.Error(err))
		// TODO: log
		return
	}
	removeSrc := dbx_sharing.RemoveMembers{
		AsAdminId:  z.srcTeamAdminId,
		Context:    z.srcFileApi,
		LeaveACopy: false,
	}
	// TODO progress
	defer removeSrc.Remove(srcTeamFolder.TeamFolderId, z.srcTempGroup.GroupId)

	addDst := dbx_sharing.AddMembers{
		AsAdminId: z.dstTeamAdminId,
		Context:   z.dstFileApi,
		Quiet:     true,
	}
	err = addDst.Add(dstTeamFolder.TeamFolderId, []string{z.dstTempGroup.GroupId}, dbx_sharing.AccessLevelEditor)
	if err != nil {
		z.Log().Error("unable to add group to team folder (DST)", zap.String("groupName", z.dstTempGroup.GroupName), zap.Error(err))
		// TODO: log
		return
	}
	// TODO: progress
	removeDst := dbx_sharing.RemoveMembers{
		AsAdminId:  z.dstTeamAdminId,
		Context:    z.dstFileApi,
		LeaveACopy: false,
	}
	defer removeDst.Remove(dstTeamFolder.TeamFolderId, z.dstTempGroup.GroupId)

	m := copy_ref.Mirror{
		SrcAsMemberId:   z.srcTeamAdminId,
		SrcApi:          z.srcFileApi,
		SrcPath:         "/",
		SrcAccountAlias: z.optSrcTeamAlias,
		SrcNamespaceId:  srcTeamFolder.TeamFolderId,
		DstAsMemberId:   z.dstTeamAdminId,
		DstApi:          z.dstFileApi,
		DstPath:         "/",
		DstNamespaceId:  dstTeamFolder.TeamFolderId,
		DstAccountAlias: z.optDstTeamAlias,
		ExecContext:     z.ExecContext,
	}
	m.MirrorAncestors()

	if z.optVerify {
		ba := compare.BetweenAccounts{
			ExecContext:      z.ExecContext,
			LeftAsMemberId:   z.srcTeamAdminId,
			LeftAccountAlias: z.optSrcTeamAlias,
			LeftPath:         "/" + srcTeamFolder.Name,
			//LeftPathRoot:      dbx_api.NewPathRootNamespace(srcTeamFolder.TeamFolderId),
			LeftApi:           z.srcFileApi,
			RightAsMemberId:   z.dstTeamAdminId,
			RightAccountAlias: z.optDstTeamAlias,
			RightPath:         "/" + dstTeamFolder.Name,
			//RightPathRoot:     dbx_api.NewPathRootNamespace(dstTeamFolder.TeamFolderId),
			RightApi: z.dstFileApi,
			OnDiff: func(diff compare.Diff) {
				z.report.Report(diff)
			},
		}
		ba.Compare()
	}
}

func (z *CmdTeamTeamFolderMirror) identifyAdmin(c *dbx_api.Context) (teamMemberId string, email string, err error) {
	admin, err := dbx_profile.AuthenticatedAdmin(c)
	if err != nil {
		return "", "", err
	}
	return admin.TeamMemberId, admin.Email, nil
}

func (z *CmdTeamTeamFolderMirror) createTeamFolder(name string, dstApi *dbx_api.Context) (tf *dbx_teamfolder.TeamFolder, err error) {
	cr := dbx_teamfolder.Create{
		OnError: z.DefaultErrorHandler,
		OnSuccess: func(teamFolder dbx_teamfolder.TeamFolder) {
			z.Log().Debug("created", zap.Any("tf", teamFolder))
			tf = &teamFolder
		},
	}
	err = cr.Create(dstApi, name)
	if err != nil {
		z.Log().Warn("failed create team folder", zap.Error(err))
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
				z.Log().Debug("folder_name_already_used") //TODO: detailed log
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
		Alias: z.optDstTeamAlias,
	}).Tell()

	z.Log().Debug("team folder created", zap.Any("tf", tf))
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
