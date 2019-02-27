package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/report"
)

type CmdMemberList struct {
	*cmd.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            report.Factory
}

func (z *CmdMemberList) Name() string {
	return "list"
}

func (z *CmdMemberList) Desc() string {
	return "List members"
}

func (CmdMemberList) Usage() string {
	return ""
}

func (z *CmdMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)

	descCsv := "Include removed members"
	f.BoolVar(&z.optIncludeRemoved, "include-removed", false, descCsv)
}

func (z *CmdMemberList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiInfo, err := au.Auth(dbx_auth.DropboxTokenBusinessInfo)
	if err != nil {
		return
	}

	z.report.Init(z.Log())
	defer z.report.Close()

	l := dbx_member.MembersList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(member *dbx_profile.Member) bool {
			z.report.Report(member)
			return true
		},
	}
	l.List(apiInfo, z.optIncludeRemoved)
}
