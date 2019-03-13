package cmd_namespace_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_namespace"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
)

type CmdTeamNamespaceFileList struct {
	*cmd.SimpleCommandlet
	optIncludeMediaInfo bool
	optIncludeDeleted   bool
	report              report.Factory
	namespaceFile       dbx_namespace.ListNamespaceFile
}

func (CmdTeamNamespaceFileList) Name() string {
	return "list"
}

func (CmdTeamNamespaceFileList) Desc() string {
	return "cmd.team.namespace.file.list.desc"
}

func (CmdTeamNamespaceFileList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamNamespaceFileList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descIncludeDeleted := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_deleted").T()
	f.BoolVar(&z.namespaceFile.OptIncludeDeleted, "include-deleted", false, descIncludeDeleted)

	descIncludeMediaInfo := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_media_info").T()
	f.BoolVar(&z.namespaceFile.OptIncludeMediaInfo, "include-media-info", false, descIncludeMediaInfo)

	descIncludeTeamFolder := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_team_folder").T()
	f.BoolVar(&z.namespaceFile.OptIncludeTeamFolder, "include-team-folder", true, descIncludeTeamFolder)

	descIncludeSharedFolder := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_shared_folder").T()
	f.BoolVar(&z.namespaceFile.OptIncludeSharedFolder, "include-shared-folder", true, descIncludeSharedFolder)

	descIncludeAppFolder := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_app_folder").T()
	f.BoolVar(&z.namespaceFile.OptIncludeAppFolder, "include-app-folder", false, descIncludeAppFolder)

	descIncludeMemberFolder := z.ExecContext.Msg("cmd.team.namespace.file.list.flag.include_member_folder").T()
	f.BoolVar(&z.namespaceFile.OptIncludeMemberFolder, "include-member-folder", false, descIncludeMemberFolder)
}

func (z *CmdTeamNamespaceFileList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	admin, err := dbx_profile.AuthenticatedAdmin(apiFile)
	if err != nil {
		z.DefaultErrorHandler(err)
		return
	}
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	z.namespaceFile.AsAdminId = admin.TeamMemberId
	z.namespaceFile.OnError = z.DefaultErrorHandler
	z.namespaceFile.OnNamespace = func(namespace *dbx_namespace.Namespace) bool {
		z.ExecContext.Msg("cmd.team.namespace.file.list.progress.scan_folder").WithData(struct {
			Type string
			Id   string
			Name string
		}{
			Type: namespace.NamespaceType,
			Id:   namespace.NamespaceId,
			Name: namespace.Name,
		})
		z.Log().Info("Scanning folder",
			zap.String("namespace_type", namespace.NamespaceType),
			zap.String("namespace_id", namespace.NamespaceId),
			zap.String("name", namespace.Name),
		)
		return true
	}
	z.namespaceFile.OnFolder = func(folder *dbx_namespace.NamespaceFolder) bool {
		z.report.Report(folder)
		return true
	}
	z.namespaceFile.OnFile = func(file *dbx_namespace.NamespaceFile) bool {
		z.report.Report(file)
		return true
	}
	z.namespaceFile.OnDelete = func(deleted *dbx_namespace.NamespaceDeleted) bool {
		z.report.Report(deleted)
		return true
	}
	z.namespaceFile.List(apiFile)
}
