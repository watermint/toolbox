package efs_memory

import (
	"github.com/watermint/toolbox/essentials/file/efs_base"
	"io"
	"time"
)

type memImpl struct {
	// path -> file content
	content map[string][]byte

	// path -> modified time
	pathModified map[string]time.Time
}

func (z memImpl) Path(path string, names ...string) (efs_base.Path, efs_base.PathError) {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FilePut(path efs_base.Path, data io.Reader) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FileGet(path efs_base.Path) (data io.ReadCloser, fsErr efs_base.FsError) {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FileSize(path efs_base.Path) (size int64, fsError efs_base.FsError) {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FileCopy(src, dst efs_base.Path) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FileMove(src, dst efs_base.Path) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FileDelete(path efs_base.Path) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FolderCreate(path efs_base.Path) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FolderList(path efs_base.Path, onEntry func(entry efs_base.FolderEntry) bool) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FolderCopy(src, dst efs_base.Path) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FolderMove(src, dst efs_base.Path) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) FolderDelete(path efs_base.Path) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) IsFolder(path efs_base.Path) (bool, efs_base.FsError) {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) IsFile(path efs_base.Path) (bool, efs_base.FsError) {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) IsExist(path efs_base.Path) (bool, efs_base.FsError) {
	//TODO implement me
	panic("implement me")
}

func (z memImpl) TimeModifiedGet(path efs_base.Path) (time.Time, efs_base.FsError) {
	//TODO implement me
	panic("implement me")
}
