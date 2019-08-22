package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdTeamInfo struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (CmdTeamInfo) Name() string {
	return "info"
}

func (CmdTeamInfo) Desc() string {
	return "cmd.team.info.desc"
}

func (CmdTeamInfo) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamInfo) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamInfo) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessInfo())
	if err != nil {
		return
	}
	svc := sv_team.New(ctx)
	info, err := svc.Info()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.report.Init(z.ExecContext)
	z.report.Report(info)
	z.report.Close()
}
