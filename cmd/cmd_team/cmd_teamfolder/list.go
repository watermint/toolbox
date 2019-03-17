package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_teamfolder"
)

type CmdTeamTeamFolderList struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.DbxContext
	report     app_report.Factory
}

func (CmdTeamTeamFolderList) Name() string {
	return "list"
}

func (CmdTeamTeamFolderList) Desc() string {
	return "cmd.team.teamfolder.list.desc"
}

func (CmdTeamTeamFolderList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamTeamFolderList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamTeamFolderList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	l := dbx_teamfolder.ListTeamFolder{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(teamFolder *dbx_teamfolder.TeamFolder) bool {
			z.report.Report(teamFolder)
			return true
		},
	}
	l.List(apiFile)
}
