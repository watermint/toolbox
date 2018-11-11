package cmd_teamfolder_file

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
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

func (c *CmdTeamTeamFolderFileList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)

	descIncludeDeleted := "Include deleted folders/files"
	f.BoolVar(&c.namespaceFile.OptIncludeDeleted, "include-deleted", false, descIncludeDeleted)

	descIncludeMediaInfo := "Include media info (metadata of photo and video)"
	f.BoolVar(&c.namespaceFile.OptIncludeMediaInfo, "include-media-info", false, descIncludeMediaInfo)
}

func (c *CmdTeamTeamFolderFileList) Exec(args []string) {
	c.namespaceFile.OptIncludeMemberFolder = false
	c.namespaceFile.OptIncludeAppFolder = false
	c.namespaceFile.OptIncludeTeamFolder = true
	c.namespaceFile.OptIncludeSharedFolder = false

	apiFile, err := c.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	admin, ea, _ := dbx_profile.AuthenticatedAdmin(apiFile)
	if ea.IsFailure() {
		c.DefaultErrorHandler(ea)
		return
	}
	c.report.Init(c.Log())
	defer c.report.Close()

	c.namespaceFile.AsAdminId = admin.TeamMemberId
	c.namespaceFile.OnError = c.DefaultErrorHandler
	c.namespaceFile.OnNamespace = func(namespace *dbx_namespace.Namespace) bool {
		c.Log().Info("Scanning team folder",
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
