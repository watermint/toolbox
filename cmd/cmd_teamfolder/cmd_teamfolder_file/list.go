package cmd_teamfolder_file

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_namespace"
	"github.com/watermint/toolbox/model/dbx_profile"
	"go.uber.org/zap"
)

type CmdTeamFolderFileList struct {
	*cmd.SimpleCommandlet
	optIncludeMediaInfo bool
	optIncludeDeleted   bool
	report              app_report.Factory
	namespaceFile       dbx_namespace.ListNamespaceFile
}

func (CmdTeamFolderFileList) Name() string {
	return "list"
}

func (CmdTeamFolderFileList) Desc() string {
	return "cmd.teamfolder.file.list.desc"
}

func (CmdTeamFolderFileList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamFolderFileList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descIncludeDeleted := z.ExecContext.Msg("cmd.teamfolder.file.list.flag.include_deleted").T()
	f.BoolVar(&z.namespaceFile.OptIncludeDeleted, "include-deleted", false, descIncludeDeleted)

	descIncludeMediaInfo := z.ExecContext.Msg("cmd.teamfolder.file.list.flag.include_media_info").T()
	f.BoolVar(&z.namespaceFile.OptIncludeMediaInfo, "include-media-info", false, descIncludeMediaInfo)
}

func (z *CmdTeamFolderFileList) Exec(args []string) {
	z.namespaceFile.OptIncludeMemberFolder = false
	z.namespaceFile.OptIncludeAppFolder = false
	z.namespaceFile.OptIncludeTeamFolder = true
	z.namespaceFile.OptIncludeSharedFolder = false

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
		z.ExecContext.Msg("cmd.teamfolder.file.list.progress.scan").WithData(struct {
			Id   string
			Name string
		}{
			Id:   namespace.NamespaceId,
			Name: namespace.Name,
		})
		z.Log().Info("Scanning team folder",
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
