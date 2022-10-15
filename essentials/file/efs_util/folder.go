package efs_util

import "github.com/watermint/toolbox/essentials/file/efs_base"

type FolderUtilOps interface {
	ListFolderEntries(path efs_base.Path) (entries []efs_base.FolderEntry, fsError efs_base.FsError)
}
