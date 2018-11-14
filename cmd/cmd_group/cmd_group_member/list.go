package cmd_group_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/report"
)

type CmdGroupMemberList struct {
	*cmd.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            report.Factory
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

	c.report.Init(c.Log())
	defer c.report.Close()

	gl := dbx_group.GroupList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(group *dbx_group.Group) bool {

			gml := dbx_group.GroupMemberList{
				OnError: c.DefaultErrorHandler,
				OnEntry: func(gm *dbx_group.GroupMember) bool {
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
