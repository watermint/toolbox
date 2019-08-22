package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdTeamFolderList struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (CmdTeamFolderList) Name() string {
	return "list"
}

func (CmdTeamFolderList) Desc() string {
	return "cmd.teamfolder.list.desc"
}

func (CmdTeamFolderList) Usage() func(cmd2.CommandUsage) {
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
