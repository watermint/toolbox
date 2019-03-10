package cmd_namespace_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_size"
	"github.com/watermint/toolbox/report"
)

type CmdTeamNamespaceFileSize struct {
	*cmd.SimpleCommandlet
	report report.Factory
	nsz    dbx_size.NamespaceSizes
}

func (CmdTeamNamespaceFileSize) Name() string {
	return "size"
}

func (CmdTeamNamespaceFileSize) Desc() string {
	return "cmd.team.namespace.file.size.desc"
}

func (CmdTeamNamespaceFileSize) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamNamespaceFileSize) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descIncludeTeamFolder := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.include_team_folder").T()
	f.BoolVar(&z.nsz.OptIncludeTeamFolder, "include-team-folder", true, descIncludeTeamFolder)

	descIncludeSharedFolder := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.include_shared_folder").T()
	f.BoolVar(&z.nsz.OptIncludeSharedFolder, "include-shared-folder", true, descIncludeSharedFolder)

	descIncludeAppFolder := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.include_app_folder").T()
	f.BoolVar(&z.nsz.OptIncludeAppFolder, "include-app-folder", false, descIncludeAppFolder)

	descIncludeMemberFolder := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.include_member_folder").T()
	f.BoolVar(&z.nsz.OptIncludeMemberFolder, "include-member-folder", false, descIncludeMemberFolder)

	descUseCached := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.cache").T()
	f.StringVar(&z.nsz.OptCachePath, "cache", "", descUseCached)

	descOptDepth := z.ExecContext.Msg("cmd.team.namespace.file.size.flag.depth").T()
	f.IntVar(&z.nsz.OptDepth, "depth", 2, descOptDepth)
}

func (z *CmdTeamNamespaceFileSize) Exec(args []string) {
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	z.nsz.Init(z.ExecContext)
	z.nsz.Load(apiFile)

	z.Log().Info("Reporting result")
	for _, sz := range z.nsz.Sizes {
		z.report.Report(sz)
	}
}
