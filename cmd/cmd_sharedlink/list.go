package cmd_sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdSharedLinkList struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdSharedLinkList) Name() string {
	return "list"
}

func (z *CmdSharedLinkList) Desc() string {
	return "cmd.sharedlink.list.desc"
}

func (z *CmdSharedLinkList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdSharedLinkList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdSharedLinkList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_sharedlink.New(ctx)
	links, err := svc.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	for _, link := range links {
		z.report.Report(link)
	}
}
