package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
)

type CmdMemberList struct {
	*cmd.SimpleCommandlet

	optIncludeRemoved bool
	apiContext        *dbx_api.Context
	report            app_report.Factory
}

func (z *CmdMemberList) Name() string {
	return "list"
}

func (z *CmdMemberList) Desc() string {
	return "cmd.member.list.desc"
}

func (CmdMemberList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
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

	z.report.Init(z.ExecContext)
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
