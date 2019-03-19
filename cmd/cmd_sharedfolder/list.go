package cmd_sharedfolder

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdSharedFolderList struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdSharedFolderList) Name() string {
	return "list"
}

func (z *CmdSharedFolderList) Desc() string {
	return "cmd.sharedfolder.list.desc"
}

func (z *CmdSharedFolderList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdSharedFolderList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdSharedFolderList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_sharedfolder.New(ctx)
	folders, err := svc.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	for _, f := range folders {
		z.report.Report(f)
	}
}
