package cmd_group_member

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdGroupMemberList struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdGroupMemberList) Name() string {
	return "list"
}

func (z *CmdGroupMemberList) Desc() string {
	return "cmd.group.member.list.desc"
}

func (z *CmdGroupMemberList) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdGroupMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdGroupMemberList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessInfo())
	if err != nil {
		return
	}
	gsv := sv_group.New(ctx)
	groups, err := gsv.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	for _, group := range groups {
		msv := sv_group_member.New(ctx, group)
		members, err := msv.List()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		for _, m := range members {
			row := mo_group_member.NewGroupMember(group, m)
			z.report.Report(row)
		}
	}
}
