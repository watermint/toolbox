package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/usecase/uc_team_migration"
)

type CmdTeamMigrate struct {
	*cmd.SimpleCommandlet
	report          app_report.Factory
	optSrcTeamAlias string
	optDstTeamAlias string
}

func (z *CmdTeamMigrate) Name() string {
	return "migrate"
}

func (z *CmdTeamMigrate) Desc() string {
	return "cmd.team.migrate.desc"
}

func (z *CmdTeamMigrate) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamMigrate) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descFromAccount := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.src_account").T()
	f.StringVar(&z.optSrcTeamAlias, "alias-src", "mirror-src", descFromAccount)

	descToAccount := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.dst_account").T()
	f.StringVar(&z.optDstTeamAlias, "alias-dest", "mirror-dst", descToAccount)
}

func (z *CmdTeamMigrate) Exec(args []string) {
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

	ucm := uc_team_migration.New(z.ExecContext, ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst)
	mc, err := ucm.Scope()
	if err != nil {
		ctxFileSrc.ErrorMsg(err).TellError()
		return
	}
	if err = ucm.Migrate(mc); err != nil {
		ctxFileSrc.ErrorMsg(err).TellError()
	}
}