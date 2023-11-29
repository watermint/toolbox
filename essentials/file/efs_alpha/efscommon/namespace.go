package efscommon

import (
	efs_deprecated2 "github.com/watermint/toolbox/essentials/file/efs_alpha"
)

const (
	nsIdentitySeparator = ":"
)

func NewNamespace(fs efs_deprecated2.FileSystem, ns string) efs_deprecated2.Namespace {
	return &namespaceImpl{
		ns: ns,
		fs: fs,
	}
}

type namespaceImpl struct {
	ns string
	fs efs_deprecated2.FileSystem
}

func (z namespaceImpl) Identity() string {
	fsId := z.fs.Identity()
	if fsId == "" {
		return z.ns
	}
	return fsId + nsIdentitySeparator + z.ns
}

func (z namespaceImpl) FileSystem() efs_deprecated2.FileSystem {
	return z.fs
}

func (z namespaceImpl) String() string {
	return z.ns
}

func (z namespaceImpl) Equals(other efs_deprecated2.Namespace) bool {
	return other.FileSystem().Equals(z.fs) && z.Identity() == other.Identity()
}

func (z namespaceImpl) IsDefault() bool {
	return z.ns == ""
}
