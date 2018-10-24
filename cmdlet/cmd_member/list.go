package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
)

type CmdMemberList struct {
	*cmdlet.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            cmdlet.Report
}

func (c *CmdMemberList) Name() string {
	return "list"
}

func (c *CmdMemberList) Desc() string {
	return "List members"
}

func (CmdMemberList) Usage() string {
	return ""
}

func (c *CmdMemberList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)

	descCsv := "Include removed members"
	f.BoolVar(&c.optIncludeRemoved, "include-removed", false, descCsv)
}

func (c *CmdMemberList) Exec(ec *infra.ExecContext, args []string) {
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

	l := dbx_team.MembersList{
		OnError: cmdlet.DefaultErrorHandler,
		OnEntry: func(profile *dbx_profile.Profile) bool {
			c.report.Report(profile)
			return true
		},
	}
	l.List(apiInfo, c.optIncludeRemoved)
}
