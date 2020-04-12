package mo_path

import (
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"path/filepath"
)

type FileSystemPath interface {
	Path

	Drive() string
}

type ExistingFileSystemPath interface {
	FileSystemPath
	ShouldExist() bool
}

func NewFileSystemPath(path string) FileSystemPath {
	p, err := ut_filepath.FormatPathWithPredefinedVariables(path)
	if err != nil {
		p = path
	}
	return &fileSystemPathImpl{path: p}
}

func NewExistingFileSystemPath(path string) FileSystemPath {
	p, err := ut_filepath.FormatPathWithPredefinedVariables(path)
	if err != nil {
		p = path
	}
	return &fileSystemPathImpl{path: p, shouldExist: true}
}

type fileSystemPathImpl struct {
	path        string
	shouldExist bool
}

func (z *fileSystemPathImpl) ShouldExist() bool {
	return z.shouldExist
}

func (z *fileSystemPathImpl) Drive() string {
	return filepath.VolumeName(z.path)
}

func (z *fileSystemPathImpl) Path() string {
	return z.path
}
