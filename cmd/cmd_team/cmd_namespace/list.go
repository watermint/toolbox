package cmd_namespace

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
)

type CmdTeamNamespaceList struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (CmdTeamNamespaceList) Name() string {
	return "list"
}

func (CmdTeamNamespaceList) Desc() string {
	return "cmd.team.namespace.list.desc"
}

func (CmdTeamNamespaceList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamNamespaceList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamNamespaceList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_namespace.New(ctx)
	namespaces, err := svc.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	for _, n := range namespaces {
		z.report.Report(n)
	}
}
