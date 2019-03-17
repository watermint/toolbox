package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_team"
)

type CmdTeamInfo struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.DbxContext
	report     app_report.Factory
}

func (CmdTeamInfo) Name() string {
	return "info"
}

func (CmdTeamInfo) Desc() string {
	return "cmd.team.info.desc"
}

func (CmdTeamInfo) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamInfo) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamInfo) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiInfo, err := au.Auth(dbx_auth.DropboxTokenBusinessInfo)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	l := dbx_team.TeamInfoList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(info *dbx_team.TeamInfo) bool {
			z.report.Report(info)
			return true
		},
	}
	l.List(apiInfo)
}
