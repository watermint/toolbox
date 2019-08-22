package cmd_group

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdGroupList struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdGroupList) Name() string {
	return "list"
}

func (z *CmdGroupList) Desc() string {
	return "cmd.group.list.desc"
}

func (z *CmdGroupList) Usage() func(cmd2.CommandUsage) {
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
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.report.Init(z.ExecContext)
	for _, f := range groups {
		z.report.Report(f)
	}
	z.report.Close()
}
