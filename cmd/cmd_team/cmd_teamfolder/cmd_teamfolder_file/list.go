package cmd_teamfolder_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_namespace"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
)

type CmdTeamTeamFolderFileList struct {
	*cmd.SimpleCommandlet
	optIncludeMediaInfo bool
	optIncludeDeleted   bool
	report              report.Factory
	namespaceFile       dbx_namespace.ListNamespaceFile
}

func (CmdTeamTeamFolderFileList) Name() string {
	return "list"
}

func (CmdTeamTeamFolderFileList) Desc() string {
	return "List files/folders in all team folders of the team"
}

func (CmdTeamTeamFolderFileList) Usage() string {
	return ""
}

func (z *CmdTeamTeamFolderFileList) FlagConfig(f *flag.FlagSet) {
	z.report.FlagConfig(f)

	descIncludeDeleted := "Include deleted folders/files"
	f.BoolVar(&z.namespaceFile.OptIncludeDeleted, "include-deleted", false, descIncludeDeleted)

	descIncludeMediaInfo := "Include media info (metadata of photo and video)"
	f.BoolVar(&z.namespaceFile.OptIncludeMediaInfo, "include-media-info", false, descIncludeMediaInfo)
}

func (z *CmdTeamTeamFolderFileList) Exec(args []string) {
	z.namespaceFile.OptIncludeMemberFolder = false
	z.namespaceFile.OptIncludeAppFolder = false
	z.namespaceFile.OptIncludeTeamFolder = true
	z.namespaceFile.OptIncludeSharedFolder = false

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	admin, ea, _ := dbx_profile.AuthenticatedAdmin(apiFile)
	if ea.IsFailure() {
		z.DefaultErrorHandler(ea)
		return
	}
	z.report.Init(z.Log())
	defer z.report.Close()

	z.namespaceFile.AsAdminId = admin.TeamMemberId
	z.namespaceFile.OnError = z.DefaultErrorHandler
	z.namespaceFile.OnNamespace = func(namespace *dbx_namespace.Namespace) bool {
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
