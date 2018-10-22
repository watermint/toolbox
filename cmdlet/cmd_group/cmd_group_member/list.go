package cmd_group_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
)

type CmdGroupMemberList struct {
	*cmdlet.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            cmdlet.Report
}

func (c *CmdGroupMemberList) Name() string {
	return "list"
}

func (c *CmdGroupMemberList) Desc() string {
	return "List group members"
}

func (c *CmdGroupMemberList) Usage() string {
	return ""
}

func (c *CmdGroupMemberList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdGroupMemberList) Exec(ec *infra.ExecContext, args []string) {
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

			gml := dbx_team.GroupMemberList{
				OnError: cmdlet.DefaultErrorHandler,
				OnEntry: func(gm *dbx_team.GroupMember) bool {
					c.report.Report(gm)
					return true
				},
			}
			gml.List(apiInfo, group)

			return true
		},
	}
	gl.List(apiInfo)
}
