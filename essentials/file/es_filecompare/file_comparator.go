package es_filecompare

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
)

type FileComparator interface {
	Compare(source, target es_filesystem.Entry) (same bool, err es_filesystem.FileSystemError)
}
