package cmd_group_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_group"
)

type CmdGroupMemberList struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.DbxContext
	report     app_report.Factory
}

func (z *CmdGroupMemberList) Name() string {
	return "list"
}

func (z *CmdGroupMemberList) Desc() string {
	return "cmd.group.member.list.desc"
}

func (z *CmdGroupMemberList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdGroupMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdGroupMemberList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiInfo, err := au.Auth(dbx_auth.DropboxTokenBusinessInfo)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
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
