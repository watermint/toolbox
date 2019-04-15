package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_member"
)

type CmdMemberList struct {
	*cmd.SimpleCommandlet

	optIncludeRemoved bool
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
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessInfo())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_member.New(ctx)
	members, err := svc.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	for _, m := range members {
		z.report.Report(m)
	}
}
