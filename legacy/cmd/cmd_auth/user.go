package cmd_auth

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdAuthUser struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdAuthUser) Name() string {
	return "user"
}

func (z *CmdAuthUser) Desc() string {
	return "cmd.auth.user.desc"
}

func (z *CmdAuthUser) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdAuthUser) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdAuthUser) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}
	svc := sv_profile.NewProfile(ctx)
	profile, err := svc.Current()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.ExecContext.Msg("cmd.auth.user.success.authorised").Tell()
	z.report.Init(z.ExecContext)
	z.report.Report(profile)
	z.report.Close()
}
