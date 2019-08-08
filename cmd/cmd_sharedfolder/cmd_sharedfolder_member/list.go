package cmd_sharedfolder_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
)

type CmdSharedFolderMemberList struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdSharedFolderMemberList) Name() string {
	return "list"
}

func (z *CmdSharedFolderMemberList) Desc() string {
	return "cmd.sharedfolder.member.list.desc"
}

func (z *CmdSharedFolderMemberList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdSharedFolderMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdSharedFolderMemberList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svNamespace := sv_sharedfolder.New(ctx)
	folders, err := svNamespace.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	for _, folder := range folders {
		svMember := sv_sharedfolder_member.New(ctx, folder)
		members, err := svMember.List()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}

		for _, member := range members {
			nm := mo_sharedfolder_member.NewSharedFolderMember(folder, member)
			z.report.Report(nm)
		}
	}
}
