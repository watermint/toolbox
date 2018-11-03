package cmd_linkedapp

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/report"
)

type CmdMemberLinkedAppList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     report.Factory
}

func (CmdMemberLinkedAppList) Name() string {
	return "list"
}

func (CmdMemberLinkedAppList) Desc() string {
	return "List all applications linked to the team members' accounts"
}

func (CmdMemberLinkedAppList) Usage() string {
	return ""
}

func (c *CmdMemberLinkedAppList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdMemberLinkedAppList) Exec(args []string) {
	apiFile, err := c.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	c.report.Init(c.Log())
	defer c.report.Close()

	l := dbx_team.LinkedAppList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(app *dbx_team.LinkedApp) bool {
			c.report.Report(app)
			return true
		},
	}
	l.List(apiFile)
}
