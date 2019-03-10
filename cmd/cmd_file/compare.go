package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_file/compare"
	"github.com/watermint/toolbox/report"
)

type CmdFileCompare struct {
	*cmd.SimpleCommandlet
	report          report.Factory
	optLeftAccount  string
	optLeftPath     string
	optRightAccount string
	optRightPath    string
}

func (CmdFileCompare) Name() string {
	return "compare"
}

func (CmdFileCompare) Desc() string {
	return "cmd.file.compare.desc"
}

func (CmdFileCompare) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdFileCompare) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descLeftAccount := z.ExecContext.Msg("cmd.file.compare.flag.left_account").Text()
	f.StringVar(&z.optLeftAccount, "left-account", "compare-left", descLeftAccount)

	descRightAccount := z.ExecContext.Msg("cmd.file.compare.flag.right_account").Text()
	f.StringVar(&z.optRightAccount, "right-account", "compare-right", descRightAccount)

	descLeftPath := z.ExecContext.Msg("cmd.file.compare.flag.left_path").Text()
	f.StringVar(&z.optLeftPath, "left-path", "", descLeftPath)

	descRightPath := z.ExecContext.Msg("cmd.file.compare.flag.right_path").Text()
	f.StringVar(&z.optRightPath, "right-path", "", descRightPath)
}

func (z *CmdFileCompare) Exec(args []string) {
	if z.optLeftAccount == "" ||
		z.optRightAccount == "" ||
		z.optLeftPath == "" ||
		z.optRightPath == "" {

		z.ExecContext.Msg("cmd.file.compare.err.not_enough_params").TellError()
		return
	}

	// Ask for LEFT account authentication
	z.ExecContext.Msg("cmd.file.compare.prompt.ask_left_account_auth").Tell()
	auLeft := dbx_auth.NewAuth(z.ExecContext, z.optLeftAccount)
	acLeft, err := auLeft.Auth(dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	// Ask for RIGHT account authentication
	z.ExecContext.Msg("cmd.file.compare.prompt.ask_right_account_auth").Tell()
	auRight := dbx_auth.NewAuth(z.ExecContext, z.optRightAccount)
	acRight, err := auRight.Auth(dbx_auth.DropboxTokenFull)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	ba := compare.BetweenAccounts{
		ExecContext:       z.ExecContext,
		LeftAccountAlias:  z.optLeftAccount,
		LeftPath:          z.optLeftPath,
		LeftApi:           acLeft,
		RightAccountAlias: z.optRightAccount,
		RightPath:         z.optRightPath,
		RightApi:          acRight,
		OnDiff: func(diff compare.Diff) {
			z.report.Report(diff)
		},
	}
	ba.Compare()
}
