package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdTeamFolderList struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.DbxContext
	report     app_report.Factory
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
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_teamfolder.New(ctx)
	folders, err := svc.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	for _, f := range folders {
		z.report.Report(f)
	}
}
