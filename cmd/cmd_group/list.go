package cmd_group

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/report"
)

type CmdGrouplist struct {
	*cmd.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            report.Factory
}

func (z *CmdGrouplist) Name() string {
	return "list"
}

func (z *CmdGrouplist) Desc() string {
	return "List groups"
}

func (z *CmdGrouplist) Usage() string {
	return ""
}

func (z *CmdGrouplist) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)

	descCsv := "Include removed members"
	f.BoolVar(&z.optIncludeRemoved, "include-removed", false, descCsv)
}

func (z *CmdGrouplist) Exec(args []string) {
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
			z.report.Report(group)
			return true
		},
	}
	gl.List(apiInfo)
}
