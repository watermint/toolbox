package cmd_auth_business

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdAuthBusinessFile struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdAuthBusinessFile) Name() string {
	return "file"
}

func (z *CmdAuthBusinessFile) Desc() string {
	return "cmd.auth.business.file.desc"
}

func (z *CmdAuthBusinessFile) Usage() func(cmd2.CommandUsage) {
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
