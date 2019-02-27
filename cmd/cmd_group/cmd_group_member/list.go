package cmd_group_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/report"
)

type CmdGroupMemberList struct {
	*cmd.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            report.Factory
}

func (z *CmdGroupMemberList) Name() string {
	return "list"
}

func (z *CmdGroupMemberList) Desc() string {
	return "List group members"
}

func (z *CmdGroupMemberList) Usage() string {
	return ""
}

func (z *CmdGroupMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)
}

func (z *CmdGroupMemberList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiInfo, err := au.Auth(dbx_auth.DropboxTokenBusinessInfo)
	if err != nil {
		return
	}

	z.report.Init(z.Log())
	defer z.report.Close()

	gl := dbx_group.GroupList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(group *dbx_group.Group) bool {

			gml := dbx_group.GroupMemberList{
				OnError: z.DefaultErrorHandler,
				OnEntry: func(gm *dbx_group.GroupMember) bool {
					z.report.Report(gm)
					return true
				},
			}
			gml.List(apiInfo, group)

			return true
		},
	}
	gl.List(apiInfo)
}
