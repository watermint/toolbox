package cmd_sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdSharedLinkRemove struct {
	*cmd.SimpleCommandlet
	report  app_report.Factory
	optPath string
}

func (z *CmdSharedLinkRemove) Name() string {
	return "remove"
}

func (z *CmdSharedLinkRemove) Desc() string {
	return "cmd.sharedlink.remove.desc"
}

func (z *CmdSharedLinkRemove) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdSharedLinkRemove) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descPath := z.ExecContext.Msg("cmd.sharedlink.remove.flag.path").T()
	f.StringVar(&z.optPath, "path", "", descPath)
}

func (z *CmdSharedLinkRemove) Exec(args []string) {
	if z.optPath == "" {
		z.ExecContext.Msg("cmd.sharedlink.remove.err.not_enough_arguments").TellError()
		return
	}

	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_sharedlink.New(ctx)
	links, err := svc.ListByPath(mo_path.NewPath(z.optPath))
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}
	if len(links) < 1 {
		z.ExecContext.Msg("cmd.sharedlink.remove.err.no_link_found").TellError()
		return
	}
	for _, link := range links {
		err := svc.Delete(link)
		if err != nil {
			ctx.ErrorMsg(err).TellError()
			continue
		}
		z.report.Report(link)
	}
}
