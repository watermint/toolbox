package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
)

type CmdTeamFolderList struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (CmdTeamFolderList) Name() string {
	return "list"
}

func (CmdTeamFolderList) Desc() string {
	return "cmd.teamfolder.list.desc"
}

func (CmdTeamFolderList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamFolderList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_teamfolder.New(ctx)
	folders, err := svc.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	for _, f := range folders {
		z.report.Report(f)
	}
}
