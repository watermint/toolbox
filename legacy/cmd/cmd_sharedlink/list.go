package cmd_sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdSharedLinkList struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdSharedLinkList) Name() string {
	return "list"
}

func (z *CmdSharedLinkList) Desc() string {
	return "cmd.sharedlink.list.desc"
}

func (z *CmdSharedLinkList) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdSharedLinkList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdSharedLinkList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_sharedlink.New(ctx)
	links, err := svc.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}

	for _, link := range links {
		z.report.Report(link)
	}
}
