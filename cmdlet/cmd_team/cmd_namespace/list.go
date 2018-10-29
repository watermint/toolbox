package cmd_namespace

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
)

type CmdTeamNamespaceList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     cmdlet.Report
}

func (CmdTeamNamespaceList) Name() string {
	return "list"
}

func (CmdTeamNamespaceList) Desc() string {
	return "List all namespaces of the team"
}

func (CmdTeamNamespaceList) Usage() string {
	return ""
}

func (c *CmdTeamNamespaceList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdTeamNamespaceList) Exec(args []string) {
	apiInfo, err := c.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	c.report.Open(c)
	defer c.report.Close()

	l := dbx_team.NamespaceList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(namespace *dbx_team.Namespace) bool {
			c.report.Report(namespace)
			return true
		},
	}
	l.List(apiInfo)
}
