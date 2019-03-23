package cmd_auth_business

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_profile"
)

type CmdAuthBusinessManagement struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdAuthBusinessManagement) Name() string {
	return "management"
}

func (z *CmdAuthBusinessManagement) Desc() string {
	return "cmd.auth.business.management.desc"
}

func (z *CmdAuthBusinessManagement) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdAuthBusinessManagement) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdAuthBusinessManagement) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessManagement())
	if err != nil {
		return
	}

	svc := sv_profile.NewTeam(ctx)
	profile, err := svc.Admin()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	z.ExecContext.Msg("cmd.auth.business.file.success.authorised").Tell()
	z.report.Init(z.ExecContext)
	z.report.Report(profile)
	z.report.Close()
}
