package cmd_member_mirror

import (
	"flag"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/usecase/uc_member_mirror"
	"go.uber.org/zap"
)

type CmdMemberMirrorFiles struct {
	*cmd.SimpleCommandlet
	optCsv          string
	optSrcTeamAlias string
	optDstTeamAlias string
	report          app_report.Factory
}

func (z *CmdMemberMirrorFiles) Name() string {
	return "files"
}

func (z *CmdMemberMirrorFiles) Desc() string {
	return "cmd.member.mirror.files.desc"
}

func (z *CmdMemberMirrorFiles) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdMemberMirrorFiles) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descCsv := "CSV file path"
	f.StringVar(&z.optCsv, "csv", "", descCsv)

	descFromAccount := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.src_account").T()
	f.StringVar(&z.optSrcTeamAlias, "alias-src", "migration-src", descFromAccount)

	descToAccount := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.dst_account").T()
	f.StringVar(&z.optDstTeamAlias, "alias-dest", "migration-dst", descToAccount)
}

func (z *CmdMemberMirrorFiles) Exec(args []string) {
	// Ask for SRC account authentication
	z.ExecContext.Msg("cmd.teamfolder.mirror.prompt.ask_src_file_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optSrcTeamAlias,
	}).Tell()
	ctxFileSrc, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.PeerName(z.optSrcTeamAlias), api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	// Ask for DST account authentication
	z.ExecContext.Msg("cmd.teamfolder.mirror.prompt.ask_dst_file_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optDstTeamAlias,
	}).Tell()
	ctxFileDst, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.PeerName(z.optDstTeamAlias), api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	type MirrorFilesReport struct {
		Result   string `json:"result"`
		SrcEmail string `json:"src_email"`
		DstEmail string `json:"dst_email"`
		Reason   string `json:"reason"`
	}

	if err = z.report.Init(z.ExecContext); err != nil {
		return
	}
	defer z.report.Close()

	ucm := uc_member_mirror.New(ctxFileSrc, ctxFileDst)
	err = app_io.NewCsvLoader(z.ExecContext, z.optCsv).
		OnRow(func(cols []string) error {
			if len(cols) < 2 {
				z.Log().Warn("The column have not enough data. Skip", zap.Strings("cols", cols))
				r := &MirrorFilesReport{
					Result: "Skip",
					Reason: "The column have not enough data",
				}
				z.report.Report(r)
				return nil
			}
			srcEmail := cols[0]
			dstEmail := cols[1]

			err := ucm.Mirror(srcEmail, dstEmail)
			r := &MirrorFilesReport{
				SrcEmail: srcEmail,
				DstEmail: dstEmail,
			}
			if err != nil {
				r.Result = "Failure"
				r.Reason = ctxFileSrc.ErrorMsg(err).T()
				z.report.Report(r)
				return nil
			}

			r.Result = "Success"
			r.Reason = ""
			z.report.Report(r)

			return nil
		}).Load()
}
