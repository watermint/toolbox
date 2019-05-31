package cmd_auth_business

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_profile"
)

type CmdAuthBusinessFile struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdAuthBusinessFile) Name() string {
	return "file"
}

func (z *CmdAuthBusinessFile) Desc() string {
	return "cmd.auth.business.file.desc"
}

func (z *CmdAuthBusinessFile) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdAuthBusinessFile) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdAuthBusinessFile) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}
	svc := sv_profile.NewTeam(ctx)
	profile, err := svc.Admin()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.ExecContext.Msg("cmd.auth.business.file.success.authorised").Tell()
	z.report.Init(z.ExecContext)
	z.report.Report(profile)
	z.report.Close()
}
