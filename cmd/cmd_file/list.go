package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/model/dbx_auth"
)

type CmdFileList struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (CmdFileList) Name() string {
	return "list"
}

func (CmdFileList) Desc() string {
	return "cmd.file.list.desc"
}

func (CmdFileList) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdFileList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdFileList) Exec(args []string) {
	if len(args) < 1 {
		z.ExecContext.Msg("cmd.file.list.err.no_argument").TellError()
		return
	}
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_file.NewFiles(ctx)

	for _, path := range args {
		files, err := svc.List(mo_path.NewPath(path))
		if err != nil {
			z.ExecContext.Msg("cmd.file.list.err.failure").WithData(struct {
				Path   string
				Reason string
			}{
				Path:   path,
				Reason: ctx.ErrorMsg(err).T(),
			}).TellError()
			continue
		}

		for _, f := range files {
			z.report.Report(f)
		}
	}
}
