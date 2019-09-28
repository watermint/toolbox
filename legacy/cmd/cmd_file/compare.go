package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_compare"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/legacy/app/app_report_legacy"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdFileCompare struct {
	*cmd2.SimpleCommandlet
	report          app_report_legacy.Factory
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

func (CmdFileCompare) Usage() func(usage cmd2.CommandUsage) {
	return nil
}

func (z *CmdFileCompare) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descLeftAccount := z.ExecContext.Msg("cmd.file.compare.flag.left_account").T()
	f.StringVar(&z.optLeftAccount, "alias-left", "compare-left", descLeftAccount)

	descRightAccount := z.ExecContext.Msg("cmd.file.compare.flag.right_account").T()
	f.StringVar(&z.optRightAccount, "alias-right", "compare-right", descRightAccount)

	descLeftPath := z.ExecContext.Msg("cmd.file.compare.flag.left_path").T()
	f.StringVar(&z.optLeftPath, "path-left", "", descLeftPath)

	descRightPath := z.ExecContext.Msg("cmd.file.compare.flag.right_path").T()
	f.StringVar(&z.optRightPath, "path-right", "", descRightPath)
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
	ctxLeft, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.PeerName(z.optLeftAccount), api_auth_impl.Full())
	if err != nil {
		return
	}

	// Ask for RIGHT account authentication
	z.ExecContext.Msg("cmd.file.compare.prompt.ask_right_account_auth").Tell()
	ctxRight, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.PeerName(z.optRightAccount), api_auth_impl.Full())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	diffReport := func(diff mo_file_diff.Diff) error {
		z.report.Report(diff)
		return nil
	}

	ucc := uc_file_compare.New(ctxLeft, ctxRight)
	_, _ = ucc.Diff(diffReport,
		uc_file_compare.LeftPath(mo_path.NewPath(z.optLeftPath)),
		uc_file_compare.RightPath(mo_path.NewPath(z.optRightPath)),
	)
}
