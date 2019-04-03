package cmd_group

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_group"
)

type CmdGroupList struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdGroupList) Name() string {
	return "list"
}

func (z *CmdGroupList) Desc() string {
	return "cmd.group.list.desc"
}

func (z *CmdGroupList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdGroupList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdGroupList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessInfo())
	if err != nil {
		return
	}

	svc := sv_group.New(ctx)
	groups, err := svc.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	z.report.Init(z.ExecContext)
	for _, f := range groups {
		z.report.Report(f)
	}
	z.report.Close()
}
