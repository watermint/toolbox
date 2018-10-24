package cmd_group

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
)

type CmdGrouplist struct {
	*cmdlet.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            cmdlet.Report
}

func (c *CmdGrouplist) Name() string {
	return "list"
}

func (c *CmdGrouplist) Desc() string {
	return "List groups"
}

func (c *CmdGrouplist) Usage() string {
	return ""
}

func (c *CmdGrouplist) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)

	descCsv := "Include removed members"
	f.BoolVar(&c.optIncludeRemoved, "include-removed", false, descCsv)
}

func (c *CmdGrouplist) Exec(ec *infra.ExecContext, args []string) {
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

	gl := dbx_team.GroupList{
		OnError: cmdlet.DefaultErrorHandler,
		OnEntry: func(group *dbx_team.Group) bool {
			c.report.Report(group)
			return true
		},
	}
	gl.List(apiInfo)
}
