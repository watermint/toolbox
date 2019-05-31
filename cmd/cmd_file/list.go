package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
)

type CmdFileList struct {
	*cmd.SimpleCommandlet
	report       app_report.Factory
	optRecursive bool
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

	descRecursive := z.ExecContext.Msg("cmd.file.list.flag.recursive").T()
	f.BoolVar(&z.optRecursive, "recursive", false, descRecursive)
}

func (z *CmdFileList) Exec(args []string) {
	if len(args) < 1 {
		z.ExecContext.Msg("cmd.file.list.err.no_argument").TellError()
		return
	}
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_file.NewFiles(ctx)

	opts := make([]sv_file.ListOpt, 0)
	if z.optRecursive {
		opts = append(opts, sv_file.Recursive())
	}

	var listPath func(path string) error
	listPath = func(path string) error {
		p := mo_path.NewPath(path)
		err := svc.ListChunked(p, func(entry mo_file.Entry) {
			if file, e := entry.File(); e {
				z.report.Report(file)
			}
			if folder, e := entry.Folder(); e {
				z.report.Report(folder)
			}
			if deleted, e := entry.Deleted(); e {
				z.report.Report(deleted)
			}
		}, opts...)
		if err != nil {
			z.ExecContext.Msg("cmd.file.list.err.failure").WithData(struct {
				Path   string
				Reason string
			}{
				Path:   path,
				Reason: api_util.UIMsgFromError(err).T(),
			}).TellError()
		}
		return err
	}

	for _, path := range args {
		listPath(path)
	}
}
