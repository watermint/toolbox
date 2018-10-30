package cmd_namespace_file

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_file"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"go.uber.org/zap"
)

type CmdTeamNamespaceFileList struct {
	*cmdlet.SimpleCommandlet
	optIncludeMediaInfo bool
	optIncludeDeleted   bool

	apiContext *dbx_api.Context
	report     cmdlet.Report
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
	f.BoolVar(&c.optIncludeDeleted, "include-deleted", false, descIncludeDeleted)

	descIncludeMediaInfo := "Include media info (metadata of photo and video)"
	f.BoolVar(&c.optIncludeMediaInfo, "include-media-info", false, descIncludeMediaInfo)
}

type NamespaceFile struct {
	Namespace *dbx_team.Namespace `json:"namespace"`
	File      *dbx_file.File      `json:"file"`
}

type NamespaceFolder struct {
	Namespace *dbx_team.Namespace `json:"namespace"`
	Folder    *dbx_file.Folder    `json:"folder"`
}

type NamespaceDeleted struct {
	Namespace *dbx_team.Namespace `json:"namespace"`
	Deleted   *dbx_file.Deleted
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

	onNamespace := func(namespace *dbx_team.Namespace) bool {
		log := c.Log().With(zap.String("ns", namespace.NamespaceId))
		log.Debug("list_folder")
		c.Log().Info("Scanning folder",
			zap.String("namespace_type", namespace.NamespaceType),
			zap.String("namespace_id", namespace.NamespaceId),
			zap.String("name", namespace.Name),
		)
		lf := dbx_file.ListFolder{
			AsAdminId:                       admin.TeamMemberId,
			IncludeMediaInfo:                c.optIncludeMediaInfo,
			IncludeDeleted:                  c.optIncludeDeleted,
			IncludeHasExplicitSharedMembers: true,
			IncludeMountedFolders:           false,
			OnError:                         c.DefaultErrorHandler,
		}
		lf.OnFile = func(file *dbx_file.File) bool {
			nf := NamespaceFile{
				Namespace: namespace,
				File:      file,
			}
			c.report.Report(nf)
			return true
		}
		lf.OnFolder = func(folder *dbx_file.Folder) bool {
			nf := NamespaceFolder{
				Namespace: namespace,
				Folder:    folder,
			}
			c.report.Report(nf)
			lf.List(apiFile, folder.FolderId)
			return true
		}
		lf.OnDelete = func(deleted *dbx_file.Deleted) bool {
			nd := NamespaceDeleted{
				Namespace: namespace,
				Deleted:   deleted,
			}
			c.report.Report(nd)
			return true
		}
		return lf.List(apiFile, "ns:"+namespace.NamespaceId)
	}

	l := dbx_team.NamespaceList{
		OnError: c.DefaultErrorHandler,
		OnEntry: onNamespace,
	}
	l.List(apiFile)
}
