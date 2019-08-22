package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdMemberList struct {
	*cmd2.SimpleCommandlet

	optIncludeRemoved bool
	report            app_report.Factory
}

func (z *CmdMemberList) Name() string {
	return "list"
}

func (z *CmdMemberList) Desc() string {
	return "cmd.member.list.desc"
}

func (CmdMemberList) Usage() func(cmd2.CommandUsage) {
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
		api_util.UIMsgFromError(err).TellError()
		return
	}

	for _, m := range members {
		z.report.Report(m)
	}
}
