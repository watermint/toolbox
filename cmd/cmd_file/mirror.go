package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_file/compare"
	"github.com/watermint/toolbox/model/dbx_file/copy_ref"
	"github.com/watermint/toolbox/report"
)

type CmdFileMirror struct {
	*cmd.SimpleCommandlet
	optSrcAccount string
	optDstAccount string
	optSrcPath    string
	optDstPath    string
	optVerify     bool
	report        report.Factory
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
	auSrc := dbx_auth.NewAuth(z.ExecContext, z.optSrcAccount)
	acSrc, err := auSrc.Auth(dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	// Ask for TO account authentication
	z.ExecContext.Msg("cmd.file.mirror.prompt.ask_dst_account_auth").Tell()
	auDst := dbx_auth.NewAuth(z.ExecContext, z.optDstAccount)
	acDst, err := auDst.Auth(dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	m := copy_ref.Mirror{
		SrcApi:          acSrc,
		SrcPath:         z.optSrcPath,
		SrcAccountAlias: z.optSrcAccount,
		DstApi:          acDst,
		DstPath:         z.optDstPath,
		DstAccountAlias: z.optDstAccount,
		ExecContext:     z.ExecContext,
	}
	m.Mirror()

	if z.optVerify {
		z.report.Init(z.ExecContext)
		defer z.report.Close()

		ba := compare.BetweenAccounts{
			ExecContext:       z.ExecContext,
			LeftAccountAlias:  z.optSrcAccount,
			LeftPath:          z.optSrcPath,
			LeftApi:           acSrc,
			RightAccountAlias: z.optDstAccount,
			RightPath:         z.optDstPath,
			RightApi:          acDst,
			OnDiff: func(diff compare.Diff) {
				z.report.Report(diff)
			},
		}
		ba.Compare()
	}
}
