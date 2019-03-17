package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_teamfolder"
)

type CmdTeamFolderList struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.DbxContext
	report     app_report.Factory
}

func (CmdTeamFolderList) Name() string {
	return "list"
}

func (CmdTeamFolderList) Desc() string {
	return "cmd.teamfolder.list.desc"
}

func (CmdTeamFolderList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamFolderList) Exec(args []string) {
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
