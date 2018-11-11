package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_member"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/report"
)

type CmdMemberList struct {
	*cmd.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            report.Factory
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

func (c *CmdMemberList) Exec(args []string) {
	apiInfo, err := c.ExecContext.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.Init(c.Log())
	defer c.report.Close()

	l := dbx_member.MembersList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(member *dbx_profile.Member) bool {
			c.report.Report(member)
			return true
		},
	}
	l.List(apiInfo, c.optIncludeRemoved)
}
