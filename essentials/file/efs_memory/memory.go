// Package efs_memory is the in-memory file system (for mostly used for temporary operation or testing)
package efs_memory

import (
	"github.com/watermint/toolbox/essentials/file/efs_base"
)

type InMemory interface {
	efs_base.FileSystemBase
}
