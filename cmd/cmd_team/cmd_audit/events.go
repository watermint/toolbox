package cmd_audit

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_activity"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/report"
)

type CmdTeamAuditEvents struct {
	*cmd.SimpleCommandlet
	report       report.Factory
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

	descAccountId := z.ExecContext.Msg("cmd.team.audit.events.flag.account_id").Text()
	f.StringVar(&z.optAccountId, "account-id", "", descAccountId)
}

func (z *CmdTeamAuditEvents) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiAudit, err := au.Auth(dbx_auth.DropboxTokenBusinessAudit)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	list := dbx_activity.ActivityLog{
		AccountId: z.optAccountId,
		OnError:   z.DefaultErrorHandler,
		OnEvent: func(event dbx_activity.Event) bool {
			z.report.Report(event)
			return true
		},
	}
	list.Events(apiAudit)
}
