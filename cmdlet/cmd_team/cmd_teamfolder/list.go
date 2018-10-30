package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_teamfolder"
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

func (c *CmdTeamTeamFolderList) Exec(args []string) {
	apiInfo, err := c.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	c.report.Open(c)
	defer c.report.Close()

	l := dbx_teamfolder.ListTeamFolder{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(teamFolder *dbx_teamfolder.TeamFolder) bool {
			c.report.Report(teamFolder)
			return true
		},
	}
	l.List(apiInfo)
}
