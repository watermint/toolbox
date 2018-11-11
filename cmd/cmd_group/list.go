package cmd_group

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/report"
)

type CmdGrouplist struct {
	*cmd.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            report.Factory
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

func (c *CmdGrouplist) Exec(args []string) {
	apiInfo, err := c.ExecContext.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.Init(c.Log())
	defer c.report.Close()

	gl := dbx_group.GroupList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(group *dbx_group.Group) bool {
			c.report.Report(group)
			return true
		},
	}
	gl.List(apiInfo)
}
