package cmd_group

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/report"
)

type CmdGroupList struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.Context
	report     report.Factory
}

func (z *CmdGroupList) Name() string {
	return "list"
}

func (z *CmdGroupList) Desc() string {
	return "cmd.group.list.desc"
}

func (z *CmdGroupList) Usage() string {
	return ""
}

func (z *CmdGroupList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdGroupList) Exec(args []string) {
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
			z.report.Report(group)
			return true
		},
	}
	gl.List(apiInfo)
}
