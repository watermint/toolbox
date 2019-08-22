package cmd_auth_business

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdAuthBusinessAudit struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdAuthBusinessAudit) Name() string {
	return "audit"
}

func (z *CmdAuthBusinessAudit) Desc() string {
	return "cmd.auth.business.audit.desc"
}

func (z *CmdAuthBusinessAudit) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdAuthBusinessAudit) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdAuthBusinessAudit) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessAudit())
	if err != nil {
		return
	}

	svc := sv_profile.NewTeam(ctx)
	profile, err := svc.Admin()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.ExecContext.Msg("cmd.auth.business.audit.success.authorised").Tell()
	z.report.Init(z.ExecContext)
	z.report.Report(profile)
	z.report.Close()
}
