package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_compare"
	"github.com/watermint/toolbox/domain/usecase/uc_file_mirror"
)

type CmdFileMirror struct {
	*cmd.SimpleCommandlet
	optSrcAccount string
	optDstAccount string
	optSrcPath    string
	optDstPath    string
	optVerify     bool
	report        app_report.Factory
}

func (CmdFileMirror) Name() string {
	return "mirror"
}

func (CmdFileMirror) Desc() string {
	return "cmd.file.mirror.desc"
}

func (CmdFileMirror) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdFileMirror) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descFromAccount := z.ExecContext.Msg("cmd.file.mirror.flag.src_account").T()
	f.StringVar(&z.optSrcAccount, "src-account", "mirror-src", descFromAccount)

	descToAccount := z.ExecContext.Msg("cmd.file.mirror.flag.dst_account").T()
	f.StringVar(&z.optDstAccount, "dest-account", "mirror-dest", descToAccount)

	descSrcPath := z.ExecContext.Msg("cmd.file.mirror.flag.src_path").T()
	f.StringVar(&z.optSrcPath, "src-path", "", descSrcPath)

	descDstPath := z.ExecContext.Msg("cmd.file.mirror.flag.dst_path").T()
	f.StringVar(&z.optDstPath, "dest-path", "", descDstPath)

	descVerify := z.ExecContext.Msg("cmd.file.mirror.flag.verify").T()
	f.BoolVar(&z.optVerify, "verify", false, descVerify)
}

func (z *CmdFileMirror) Exec(args []string) {
	if z.optSrcAccount == "" ||
		z.optDstAccount == "" ||
		z.optSrcPath == "" ||
		z.optDstPath == "" {

		z.ExecContext.Msg("cmd.file.mirror.err.not_enough_params").TellError()
		return
	}

	// Ask for FROM account authentication
	z.ExecContext.Msg("cmd.file.mirror.prompt.ask_src_account_auth").Tell()
	ctxSrc, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.PeerName(z.optSrcAccount), api_auth_impl.Full())
	if err != nil {
		return
	}

	// Ask for TO account authentication
	z.ExecContext.Msg("cmd.file.mirror.prompt.ask_dst_account_auth").Tell()
	ctxDst, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.PeerName(z.optDstAccount), api_auth_impl.Full())
	if err != nil {
		return
	}

	srcPath := mo_path.NewPath(z.optSrcPath)
	dstPath := mo_path.NewPath(z.optDstPath)

	ucm := uc_file_mirror.NewFiles(ctxSrc, ctxDst)
	err = ucm.Mirror(srcPath, dstPath)

	if err != nil {
		ctxSrc.ErrorMsg(err).TellError()
		return
	}

	if z.optVerify {
		z.report.Init(z.ExecContext)
		defer z.report.Close()

		rep := func(diff mo_file_diff.Diff) error {
			z.report.Report(diff)
			return nil
		}
		ucc := uc_file_compare.New(ctxSrc, ctxDst)
		_, err := ucc.Diff(rep, uc_file_compare.LeftPath(srcPath), uc_file_compare.RightPath(dstPath))
		if err != nil {
			ctxSrc.ErrorMsg(err).TellError()
			return
		}
	}
}
