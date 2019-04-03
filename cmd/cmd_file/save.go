package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_url"
)

type CmdFileSave struct {
	*cmd.SimpleCommandlet
	report  app_report.Factory
	optPath string
	optUrl  string
}

func (z *CmdFileSave) Name() string {
	return "save"
}

func (z *CmdFileSave) Desc() string {
	return "cmd.file.save.desc"
}

func (z *CmdFileSave) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdFileSave) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descPath := z.ExecContext.Msg("cmd.file.save.flag.path").T()
	f.StringVar(&z.optPath, "path", "", descPath)

	descUrl := z.ExecContext.Msg("cmd.file.save.flag.url").T()
	f.StringVar(&z.optUrl, "url", "", descUrl)
}

func (z *CmdFileSave) Exec(args []string) {
	if z.optUrl == "" || z.optPath == "" {
		z.ExecContext.Msg("cmd.file.save.err.not_enough_params").TellError()
		return
	}

	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}

	svc := sv_file_url.New(ctx)
	entry, err := svc.Save(mo_path.NewPath(z.optPath), z.optUrl)
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	if err := z.report.Init(z.ExecContext); err != nil {
		return
	}
	z.report.Report(entry)
	z.report.Close()
}
