package cmd_migrate

import (
	"flag"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/usecase/uc_team_migration"
)

type CmdTeamMigrateCheck struct {
	*cmd.SimpleCommandlet
	report                 app_report.Factory
	optSrcTeamAlias        string
	optDstTeamAlias        string
	optMembersAll          bool
	optMembersCsv          string
	optTeamFoldersAll      bool
	optTeamFoldersCsv      string
	optAll                 bool
	optGroupsOnlyRelated   bool
	optKeepDesktopSessions bool
}

func (z *CmdTeamMigrateCheck) Name() string {
	return "check"
}

func (z *CmdTeamMigrateCheck) Desc() string {
	return "cmd.team.migrate.check.desc"
}

func (z *CmdTeamMigrateCheck) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamMigrateCheck) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descFromAccount := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.src_account").T()
	f.StringVar(&z.optSrcTeamAlias, "alias-src", "migration-src", descFromAccount)

	descToAccount := z.ExecContext.Msg("cmd.teamfolder.mirror.flag.dst_account").T()
	f.StringVar(&z.optDstTeamAlias, "alias-dest", "migration-dst", descToAccount)

	descMembersAll := z.ExecContext.Msg("cmd.team.migrate.check.flag.members_all").T()
	f.BoolVar(&z.optMembersAll, "member-all", false, descMembersAll)

	descMembersCsv := z.ExecContext.Msg("cmd.team.migrate.check.flag.members_csv").T()
	f.StringVar(&z.optMembersCsv, "member-csv", "", descMembersCsv)

	descTeamFolderAll := z.ExecContext.Msg("cmd.team.migrate.check.flag.teamfolder_all").T()
	f.BoolVar(&z.optTeamFoldersAll, "teamfolder-all", false, descTeamFolderAll)

	descTeamFolderCsv := z.ExecContext.Msg("cmd.team.migrate.check.flag.teamfolder_csv").T()
	f.StringVar(&z.optTeamFoldersCsv, "teamfolder-csv", "", descTeamFolderCsv)

	descAll := z.ExecContext.Msg("cmd.team.migrate.check.flag.all").T()
	f.BoolVar(&z.optAll, "all", false, descAll)

	descGroupsOnlyRelated := z.ExecContext.Msg("cmd.team.migrate.check.flag.groups_only_related").T()
	f.BoolVar(&z.optGroupsOnlyRelated, "groups-only-related", false, descGroupsOnlyRelated)

	descKeepDesktopSessions := z.ExecContext.Msg("cmd.team.migrate.check.flag.keep_desktop_sessions").T()
	f.BoolVar(&z.optKeepDesktopSessions, "keep-desktop-sessions", false, descKeepDesktopSessions)
}

func (z *CmdTeamMigrateCheck) Exec(args []string) {
	var err error

	teamFolderNames := make([]string, 0)
	memberEmails := make([]string, 0)

	if z.optTeamFoldersCsv != "" {
		err = app_io.NewCsvLoader(z.ExecContext, z.optTeamFoldersCsv).
			OnRow(func(cols []string) error {
				if len(cols) < 1 {
					return nil
				}
				teamFolderNames = append(teamFolderNames, cols[0])
				return nil
			}).Load()
		if err != nil {
			return
		}
	}
	if z.optMembersCsv != "" {
		err = app_io.NewCsvLoader(z.ExecContext, z.optMembersCsv).
			OnRow(func(cols []string) error {
				if len(cols) < 1 {
					return nil
				}
				memberEmails = append(memberEmails, cols[0])
				return nil
			}).Load()
		if err != nil {
			return
		}
	}

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

	opts := make([]uc_team_migration.ScopeOpt, 0)
	if z.optMembersCsv != "" {
		opts = append(opts, uc_team_migration.MembersSpecifiedEmail(memberEmails))
	}
	if z.optMembersAll {
		opts = append(opts, uc_team_migration.MembersAllExceptAdmin())
	}
	if z.optTeamFoldersCsv != "" {
		opts = append(opts, uc_team_migration.TeamFoldersSpecifiedName(teamFolderNames))
	}
	if z.optTeamFoldersAll {
		opts = append(opts, uc_team_migration.TeamFoldersAll())
	}
	if z.optGroupsOnlyRelated {
		opts = append(opts, uc_team_migration.GroupsOnlyRelated())
	}
	if z.optAll {
		opts = append(opts, uc_team_migration.MembersAllExceptAdmin(), uc_team_migration.TeamFoldersAll())
	}

	ucm := uc_team_migration.New(z.ExecContext, ctxFileSrc, ctxMgtSrc, ctxFileDst, ctxMgtDst)
	mc, err := ucm.Scope(opts...)
	if err != nil {
		ctxFileSrc.ErrorMsg(err).TellError()
		return
	}
	if err = ucm.Preflight(mc); err != nil {
		ctxFileSrc.ErrorMsg(err).TellError()
	}
}
