package cmd_group_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
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

func (c *CmdGroupMemberList) Exec(args []string) {
	apiInfo, err := c.ExecContext.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.Open(c)
	defer c.report.Close()

	gl := dbx_team.GroupList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(group *dbx_team.Group) bool {

			gml := dbx_team.GroupMemberList{
				OnError: c.DefaultErrorHandler,
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
