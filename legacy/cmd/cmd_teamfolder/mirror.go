package cmd_teamfolder

import (
	"flag"
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report_legacy"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdTeamFolderMirror struct {
	*cmd2.SimpleCommandlet

	optSrcTeamAlias   string
	optDstTeamAlias   string
	optAllTeamFolders bool

	report app_report_legacy.Factory
}

func (CmdTeamFolderMirror) Name() string {
	return "mirror"
}

func (CmdTeamFolderMirror) Desc() string {
	return "cmd.teamfolder.mirror.desc"
}

func (CmdTeamFolderMirror) Usage() func(usage cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderMirror) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descFromAccount := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.src_account").T()
	f.StringVar(&z.optSrcTeamAlias, "alias-src", "mirror-src", descFromAccount)

	descToAccount := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.dst_account").T()
	f.StringVar(&z.optDstTeamAlias, "alias-dest", "mirror-dst", descToAccount)

	descAll := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.all").T()
	f.BoolVar(&z.optAllTeamFolders, "all", false, descAll)
}

func (z *CmdTeamFolderMirror) Exec(args []string) {
	if z.optSrcTeamAlias == "" ||
		z.optDstTeamAlias == "" {

		z.ExecContext.Msg("cmd.teamfolder.mirror.err.not_enough_params").TellError()
		return
	}
	if z.optSrcTeamAlias == z.optDstTeamAlias {
		z.ExecContext.Msg("cmd.teamfolder.mirror.err.same_team").TellError()
		return
	}
	if len(args) < 1 && !z.optAllTeamFolders {
		z.ExecContext.Msg("cmd.teamfolder.mirror.err.not_enough_arguments").TellError()
		return
	}
	var err error

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

	// Ask for SRC account authentication
	z.ExecContext.Msg("cmd.teamfolder.mirror.prompt.ask_src_mgmt_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optSrcTeamAlias,
	}).Tell()
	ctxMgtSrc, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.PeerName(z.optSrcTeamAlias), api_auth_impl.BusinessManagement())
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

	// Ask for DST account authentication
	z.ExecContext.Msg("cmd.teamfolder.mirror.prompt.ask_dst_mgmt_account_auth").WithData(struct {
		Alias string
	}{
		Alias: z.optDstTeamAlias,
	}).Tell()
	ctxMgtDst, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.PeerName(z.optDstTeamAlias), api_auth_impl.BusinessManagement())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	ucm := uc_teamfolder_mirror.New(ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst, &z.report)

	if z.optAllTeamFolders {
		uc, err := ucm.AllFolderScope()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		if err = ucm.Mirror(uc); err != nil {
			api_util.UIMsgFromError(err).TellError()
		}
	} else {
		uc, err := ucm.PartialScope(args)
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		if err = ucm.Mirror(uc); err != nil {
			api_util.UIMsgFromError(err).TellError()
		}
	}
}
