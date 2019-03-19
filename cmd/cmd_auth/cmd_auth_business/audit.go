package cmd_auth_business

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdAuthBusinessAudit struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdAuthBusinessAudit) Name() string {
	return "audit"
}

func (z *CmdAuthBusinessAudit) Desc() string {
	return "cmd.auth.business.audit.desc"
}

func (z *CmdAuthBusinessAudit) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdAuthBusinessAudit) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdAuthBusinessAudit) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessAudit)
	if err != nil {
		return
	}

	svc := sv_profile.NewTeam(ctx)
	profile, err := svc.Admin()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	z.ExecContext.Msg("cmd.auth.business.audit.success.authorised").Tell()
	z.report.Init(z.ExecContext)
	z.report.Report(profile)
	z.report.Close()
}
