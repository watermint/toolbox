package cmd_auth_business

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_profile"
)

type CmdAuthBusinessInfo struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdAuthBusinessInfo) Name() string {
	return "info"
}

func (z *CmdAuthBusinessInfo) Desc() string {
	return "cmd.auth.business.info.desc"
}

func (z *CmdAuthBusinessInfo) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdAuthBusinessInfo) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdAuthBusinessInfo) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessInfo())
	if err != nil {
		return
	}

	svc := sv_profile.NewTeam(ctx)
	profile, err := svc.Admin()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.ExecContext.Msg("cmd.auth.business.info.success.authorised").Tell()
	z.report.Init(z.ExecContext)
	z.report.Report(profile)
	z.report.Close()
}
