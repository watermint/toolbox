package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/report"
)

type CmdTeamInfo struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     report.Factory
}

func (CmdTeamInfo) Name() string {
	return "info"
}

func (CmdTeamInfo) Desc() string {
	return "Team info"
}

func (CmdTeamInfo) Usage() string {
	return ""
}

func (c *CmdTeamInfo) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdTeamInfo) Exec(args []string) {
	apiInfo, err := c.ExecContext.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.Open(c.Log())
	defer c.report.Close()

	l := dbx_team.TeamInfoList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(info *dbx_team.TeamInfo) bool {
			c.report.Report(info)
			return true
		},
	}
	l.List(apiInfo)
}
