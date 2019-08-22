package cmd_linkedapp

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_linkedapp"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_linkedapp"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdMemberLinkedAppList struct {
	*cmd2.SimpleCommandlet

	OptWithMemberEmail bool
	report             app_report.Factory
}

func (CmdMemberLinkedAppList) Name() string {
	return "list"
}

func (CmdMemberLinkedAppList) Desc() string {
	return "cmd.team.linkedapp.list.desc"
}

func (z *CmdMemberLinkedAppList) Usage() func(cmd2.CommandUsage) {
	return func(u cmd2.CommandUsage) {
		z.ExecContext.Msg("cmd.team.linkedapp.list.desc").Tell()
	}
}

func (z *CmdMemberLinkedAppList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdMemberLinkedAppList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	svm := sv_member.New(ctx)
	memberList, err := svm.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	members := mo_member.MapByTeamMemberId(memberList)

	sva := sv_linkedapp.New(ctx)
	apps, err := sva.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	for _, app := range apps {
		m := &mo_member.Member{}
		m.TeamMemberId = app.TeamMemberId

		if m0, e := members[m.TeamMemberId]; e {
			m = m0
		}

		ma := mo_linkedapp.NewMemberLinkedApp(m, app)
		z.report.Report(ma)
	}
}
