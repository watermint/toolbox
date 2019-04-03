package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
)

type CmdFileMetadata struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdFileMetadata) Name() string {
	return "metadata"
}

func (z *CmdFileMetadata) Desc() string {
	return "cmd.file.metadata.desc"
}

func (z *CmdFileMetadata) Usage() func(usage cmd.CommandUsage) {
	return func(usage cmd.CommandUsage) {
		z.ExecContext.Msg("cmd.file.metadata.usage").WithData(usage).Tell()
	}
}

func (z *CmdFileMetadata) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdFileMetadata) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_file.NewFiles(ctx)

	for _, p := range args {
		md, err := svc.Resolve(mo_path.NewPath(p))
		if err != nil {
			ctx.ErrorMsg(err).TellError()
			continue
		}
		if file, e := md.File(); e {
			z.report.Report(file)
		}
		if folder, e := md.Folder(); e {
			z.report.Report(folder)
		}
		if deleted, e := md.Deleted(); e {
			z.report.Report(deleted)
		}
	}
}
