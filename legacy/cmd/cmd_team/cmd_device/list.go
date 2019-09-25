package cmd_device

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_device"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_device"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report_legacy"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdTeamDeviceList struct {
	*cmd2.SimpleCommandlet
	report        app_report_legacy.Factory
	optDeviceType string
}

func (CmdTeamDeviceList) Name() string {
	return "list"
}

func (CmdTeamDeviceList) Desc() string {
	return "cmd.team.device.list.desc"
}

func (CmdTeamDeviceList) Usage() func(usage cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamDeviceList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descDeviceType := z.ExecContext.Msg("cmd.team.device.list.flag.device").T()
	f.StringVar(&z.optDeviceType, "device", "", descDeviceType)
}

func (z *CmdTeamDeviceList) Exec(args []string) {
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

	svd := sv_device.New(ctx)
	devices, err := svd.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	z.report.Init(z.ExecContext)
	defer z.report.Close()
	for _, dev := range devices {
		if member, e := members[dev.EntryTeamMemberId()]; e {
			md := mo_device.NewMemberSession(member, dev)
			z.report.Report(md)
		}
	}
}
