package efs_posix

import (
	"github.com/watermint/toolbox/essentials/file/efs_base"
	"github.com/watermint/toolbox/essentials/file/efs_util"
	"os"
	"time"
)

type FileOpenOpts struct {
	// O_RDONLY. open the file read-only.
	ReadOnly bool

	// O_WRONLY. open the file write-only.
	WriteOnly bool

	// O_RDWR. open the file read-write.
	ReadWrite bool

	// O_APPEND. append data to the file when writing.
	Append bool

	// O_CREATE. create a new file if none exists.
	Create bool

	// O_EXCL. used with Create, file must not exist.
	Exclusive bool

	// O_SYNC. open for synchronous I/O.
	Sync bool

	// O_TRUNC. truncate regular writable file when opened.
	Truncate bool
}

type FileOpenOpt func(o FileOpenOpts) FileOpenOpts

type FileOps interface {
	FilePosixOpen(path efs_base.Path, opts ...FileOpenOpt) (f *os.File, fsError efs_base.FsError)
}

type PermMode uint32

type PermOps interface {
	PosixPermSet(path efs_base.Path, mode PermMode) efs_base.FsError
	PosixPermGet(path efs_base.Path) (mode PermMode, fsError efs_base.FsError)
	PosixSetOwner(path efs_base.Path, uid, gid int) efs_base.FsError
	// PosixGetOwner
	// impl note. https://pirivan.gitlab.io/post/how-to-retrieve-file-ownership-information-in-golang/
	PosixGetOwner(path efs_base.Path) (uid, gid int, fsError efs_base.FsError)
}

type SymlinkOps interface {
	PosixSymlinkCreate(path, target efs_base.Path) efs_base.FsError
	PosixSymlinkInfo(path efs_base.Path) (target efs_base.Path, fsError efs_base.FsError)
}

type HardLinkOps interface {
	PosixHardLinkCreate(path, target efs_base.Path) efs_base.FsError
	PosixHardLinkInfo(path efs_base.Path) (target efs_base.Path, fsError efs_base.FsError)
}

type TypeOps interface {
	IsSymlink(path efs_base.Path) (bool, efs_base.FsError)
	IsExecutable(path efs_base.Path) (bool, efs_base.FsError)
	IsHidden(path efs_base.Path) (bool, efs_base.FsError)
	IsReadable(path efs_base.Path) (bool, efs_base.FsError)
	IsWritable(path efs_base.Path) (bool, efs_base.FsError)
}

type TimeOps interface {
	TimeModifiedSet(path efs_base.Path, t time.Time) efs_base.FsError
	TimeAccessGet() (time.Time, efs_base.FsError)
	TimeAccessSet(path efs_base.Path, t time.Time) efs_base.FsError
	TimeCreatedGet() (time.Time, efs_base.FsError)
	TimeCreatedSet(path efs_base.Path, t time.Time) efs_base.FsError
}

type PosixFileSystemBase interface {
	efs_base.FileSystemBase
	PermOps
	SymlinkOps
	HardLinkOps
	TypeOps
	TimeOps
}

type PosixFileSystem interface {
	PosixFileSystemBase
	efs_util.FileUtilOps
	efs_util.FolderUtilOps
	efs_util.StateUtilOps
}
