package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_url"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdFileSave struct {
	*cmd2.SimpleCommandlet
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

func (z *CmdFileSave) Usage() func(cmd2.CommandUsage) {
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
		api_util.UIMsgFromError(err).TellError()
		return
	}

	if err := z.report.Init(z.ExecContext); err != nil {
		return
	}
	z.report.Report(entry)
	z.report.Close()
}
