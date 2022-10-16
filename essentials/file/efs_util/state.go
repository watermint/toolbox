package efs_util

import "github.com/watermint/toolbox/essentials/file/efs_base"

type StateUtilOps interface {
	IsNotExist(path efs_base.Path) (bool, efs_base.FsError)
}
