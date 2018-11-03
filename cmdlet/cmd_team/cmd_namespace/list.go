package cmd_namespace

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_namespace"
	"github.com/watermint/toolbox/report"
)

type CmdTeamNamespaceList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     report.Factory
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

	c.report.Init(c.Log())
	defer c.report.Close()

	l := dbx_namespace.NamespaceList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(namespace *dbx_namespace.Namespace) bool {
			c.report.Report(namespace)
			return true
		},
	}
	l.List(apiInfo)
}
