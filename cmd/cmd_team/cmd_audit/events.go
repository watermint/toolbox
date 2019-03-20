package cmd_audit

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdTeamAuditEvents struct {
	*cmd.SimpleCommandlet
	report       app_report.Factory
	optAccountId string
}

func (CmdTeamAuditEvents) Name() string {
	return "events"
}

func (CmdTeamAuditEvents) Desc() string {
	return "cmd.team.audit.events.desc"
}

func (CmdTeamAuditEvents) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamAuditEvents) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descAccountId := z.ExecContext.Msg("cmd.team.audit.events.flag.account_id").T()
	f.StringVar(&z.optAccountId, "account-id", "", descAccountId)
}

func (z *CmdTeamAuditEvents) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessAudit)
	if err != nil {
		return
	}
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_activity.New(ctx)
	err = svc.All(func(event *mo_activity.Event) error {
		z.report.Report(event)
		return nil
	})
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}
}
