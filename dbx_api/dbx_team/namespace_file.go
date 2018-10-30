package dbx_team

import (
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_file"
	"go.uber.org/zap"
)

type NamespaceFile struct {
	Namespace *Namespace     `json:"namespace"`
	File      *dbx_file.File `json:"file"`
}

type NamespaceFolder struct {
	Namespace *Namespace       `json:"namespace"`
	Folder    *dbx_file.Folder `json:"folder"`
}

type NamespaceDeleted struct {
	Namespace *Namespace `json:"namespace"`
	Deleted   *dbx_file.Deleted
}

// Recursively list files under namespaces
type ListNamespaceFile struct {
	AsAdminId              string
	OptIncludeMediaInfo    bool
	OptIncludeDeleted      bool
	OptIncludeAppFolder    bool
	OptIncludeMemberFolder bool
	OptIncludeTeamFolder   bool
	OptIncludeSharedFolder bool
	OnError                func(annotation dbx_api.ErrorAnnotation) bool
	OnFolder               func(folder *NamespaceFolder) bool
	OnFile                 func(file *NamespaceFile) bool
	OnDelete               func(deleted *NamespaceDeleted) bool
	OnNamespace            func(namespace *Namespace) bool
}

func (l *ListNamespaceFile) List(c *dbx_api.Context) bool {
	onNamespace := func(namespace *Namespace) bool {
		log := c.Log().With(
			zap.String("ns", namespace.NamespaceId),
			zap.String("type", namespace.NamespaceType),
		)
		log.Debug("scanning namespace")

		if namespace.NamespaceType == "app_folder" && !l.OptIncludeAppFolder {
			log.Debug("Skip")
			return true
		}
		if namespace.NamespaceType == "shared_folder" && !l.OptIncludeSharedFolder {
			log.Debug("Skip")
			return true
		}
		if namespace.NamespaceType == "team_folder" && !l.OptIncludeTeamFolder {
			log.Debug("Skip")
			return true
		}
		if namespace.NamespaceType == "team_member_folder" && !l.OptIncludeMemberFolder {
			log.Debug("Skip")
			return true
		}

		if l.OnNamespace != nil {
			if !l.OnNamespace(namespace) {
				log.Debug("abort process namespace due to `OnNamespace` returned false")
				return false
			}
		}
		lf := dbx_file.ListFolder{
			AsAdminId:                       l.AsAdminId,
			IncludeMediaInfo:                l.OptIncludeMediaInfo,
			IncludeDeleted:                  l.OptIncludeDeleted,
			IncludeHasExplicitSharedMembers: true,
			IncludeMountedFolders:           false,
			OnError:                         l.OnError,
		}
		lf.OnFile = func(file *dbx_file.File) bool {
			if l.OnFile == nil {
				return true
			}
			nf := &NamespaceFile{
				Namespace: namespace,
				File:      file,
			}
			return l.OnFile(nf)
		}
		lf.OnFolder = func(folder *dbx_file.Folder) bool {
			// recursive call
			lf.List(c, folder.FolderId)
			if l.OnFolder == nil {
				return true
			}

			nf := &NamespaceFolder{
				Namespace: namespace,
				Folder:    folder,
			}
			return l.OnFolder(nf)
		}
		lf.OnDelete = func(deleted *dbx_file.Deleted) bool {
			if l.OnDelete == nil {
				return true
			}

			nd := &NamespaceDeleted{
				Namespace: namespace,
				Deleted:   deleted,
			}
			return l.OnDelete(nd)
		}
		return lf.List(c, "ns:"+namespace.NamespaceId)
	}

	nsl := NamespaceList{
		OnError: l.OnError,
		OnEntry: onNamespace,
	}
	return nsl.List(c)
}
