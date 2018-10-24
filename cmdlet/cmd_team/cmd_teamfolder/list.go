package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
)

type CmdTeamTeamFolderList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     cmdlet.Report
}

func (CmdTeamTeamFolderList) Name() string {
	return "list"
}

func (CmdTeamTeamFolderList) Desc() string {
	return "List all team folder of the team"
}

func (CmdTeamTeamFolderList) Usage() string {
	return ""
}

func (c *CmdTeamTeamFolderList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdTeamTeamFolderList) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	apiInfo, err := ec.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.Open()
	defer c.report.Close()

	l := dbx_team.TeamFolderList{
		OnError: cmdlet.DefaultErrorHandler,
		OnEntry: func(teamFolder *dbx_team.TeamFolder) bool {
			c.report.Report(teamFolder)
			return true
		},
	}
	l.List(apiInfo)
}
