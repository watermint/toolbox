package cmd_namespace_file

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api/dbx_namespace"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"go.uber.org/zap"
)

type CmdTeamNamespaceFileList struct {
	*cmdlet.SimpleCommandlet
	optIncludeMediaInfo bool
	optIncludeDeleted   bool
	report              cmdlet.Report
	namespaceFile       dbx_namespace.ListNamespaceFile
}

func (CmdTeamNamespaceFileList) Name() string {
	return "list"
}

func (CmdTeamNamespaceFileList) Desc() string {
	return "List files/folders in all namespaces of the team"
}

func (CmdTeamNamespaceFileList) Usage() string {
	return ""
}

func (c *CmdTeamNamespaceFileList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)

	descIncludeDeleted := "Include deleted folders/files"
	f.BoolVar(&c.namespaceFile.OptIncludeDeleted, "include-deleted", false, descIncludeDeleted)

	descIncludeMediaInfo := "Include media info (metadata of photo and video)"
	f.BoolVar(&c.namespaceFile.OptIncludeMediaInfo, "include-media-info", false, descIncludeMediaInfo)

	descIncludeTeamFolder := "Include team folders"
	f.BoolVar(&c.namespaceFile.OptIncludeTeamFolder, "include-team-folder", true, descIncludeTeamFolder)

	descIncludeSharedFolder := "Include shared folders"
	f.BoolVar(&c.namespaceFile.OptIncludeSharedFolder, "include-shared-folder", true, descIncludeSharedFolder)

	descIncludeAppFolder := "Include app folders"
	f.BoolVar(&c.namespaceFile.OptIncludeAppFolder, "include-app-folder", false, descIncludeAppFolder)

	descIncludeMemberFolder := "Include team member folders"
	f.BoolVar(&c.namespaceFile.OptIncludeMemberFolder, "include-member-folder", false, descIncludeMemberFolder)
}

func (c *CmdTeamNamespaceFileList) Exec(args []string) {
	apiFile, err := c.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	admin, ea, _ := dbx_profile.AuthenticatedAdmin(apiFile)
	if ea.IsFailure() {
		c.DefaultErrorHandler(ea)
		return
	}
	c.report.Open(c)
	defer c.report.Close()

	c.namespaceFile.AsAdminId = admin.TeamMemberId
	c.namespaceFile.OnError = c.DefaultErrorHandler
	c.namespaceFile.OnNamespace = func(namespace *dbx_namespace.Namespace) bool {
		c.Log().Info("Scanning folder",
			zap.String("namespace_type", namespace.NamespaceType),
			zap.String("namespace_id", namespace.NamespaceId),
			zap.String("name", namespace.Name),
		)
		return true
	}
	c.namespaceFile.OnFolder = func(folder *dbx_namespace.NamespaceFolder) bool {
		c.report.Report(folder)
		return true
	}
	c.namespaceFile.OnFile = func(file *dbx_namespace.NamespaceFile) bool {
		c.report.Report(file)
		return true
	}
	c.namespaceFile.OnDelete = func(deleted *dbx_namespace.NamespaceDeleted) bool {
		c.report.Report(deleted)
		return true
	}
	c.namespaceFile.List(apiFile)
}
