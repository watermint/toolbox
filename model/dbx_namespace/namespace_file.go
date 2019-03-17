package dbx_namespace

import (
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_file"
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
	OnError                func(err error) bool
	OnFolder               func(folder *NamespaceFolder) bool
	OnFile                 func(file *NamespaceFile) bool
	OnDelete               func(deleted *NamespaceDeleted) bool
	OnNamespace            func(namespace *Namespace) bool
}

func (z *ListNamespaceFile) List(c *dbx_api.DbxContext) bool {
	onNamespace := func(namespace *Namespace) bool {
		log := c.Log().With(
			zap.String("ns", namespace.NamespaceId),
			zap.String("type", namespace.NamespaceType),
		)
		log.Debug("scanning namespace")

		if namespace.NamespaceType == "app_folder" && !z.OptIncludeAppFolder {
			log.Debug("Skip")
			return true
		}
		if namespace.NamespaceType == "shared_folder" && !z.OptIncludeSharedFolder {
			log.Debug("Skip")
			return true
		}
		if namespace.NamespaceType == "team_folder" && !z.OptIncludeTeamFolder {
			log.Debug("Skip")
			return true
		}
		if namespace.NamespaceType == "team_member_folder" && !z.OptIncludeMemberFolder {
			log.Debug("Skip")
			return true
		}

		if z.OnNamespace != nil {
			if !z.OnNamespace(namespace) {
				log.Debug("abort process namespace due to `OnNamespace` returned false")
				return false
			}
		}
		lf := dbx_file.ListFolder{
			AsAdminId:                       z.AsAdminId,
			IncludeMediaInfo:                z.OptIncludeMediaInfo,
			IncludeDeleted:                  z.OptIncludeDeleted,
			IncludeHasExplicitSharedMembers: true,
			IncludeMountedFolders:           false,
			OnError:                         z.OnError,
		}
		lf.OnFile = func(file *dbx_file.File) bool {
			if z.OnFile == nil {
				return true
			}
			nf := &NamespaceFile{
				Namespace: namespace,
				File:      file,
			}
			return z.OnFile(nf)
		}
		lf.OnFolder = func(folder *dbx_file.Folder) bool {
			// recursive call
			lf.List(c, folder.FolderId)
			if z.OnFolder == nil {
				return true
			}

			nf := &NamespaceFolder{
				Namespace: namespace,
				Folder:    folder,
			}
			return z.OnFolder(nf)
		}
		lf.OnDelete = func(deleted *dbx_file.Deleted) bool {
			if z.OnDelete == nil {
				return true
			}

			nd := &NamespaceDeleted{
				Namespace: namespace,
				Deleted:   deleted,
			}
			return z.OnDelete(nd)
		}
		return lf.List(c, "ns:"+namespace.NamespaceId)
	}

	nsl := NamespaceList{
		OnError: z.OnError,
		OnEntry: onNamespace,
	}
	return nsl.List(c)
}
