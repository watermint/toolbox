package cmd_device

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_device"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_device"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdTeamDeviceList struct {
	*cmd.SimpleCommandlet
	report        app_report.Factory
	optDeviceType string
}

func (CmdTeamDeviceList) Name() string {
	return "list"
}

func (CmdTeamDeviceList) Desc() string {
	return "cmd.team.device.list.desc"
}

func (CmdTeamDeviceList) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamDeviceList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descDeviceType := z.ExecContext.Msg("cmd.team.device.list.flag.device").T()
	f.StringVar(&z.optDeviceType, "device", "", descDeviceType)
}

func (z *CmdTeamDeviceList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	svm := sv_member.New(ctx)
	memberList, err := svm.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}
	members := mo_member.MapByTeamMemberId(memberList)

	svd := sv_device.New(ctx)
	devices, err := svd.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
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
