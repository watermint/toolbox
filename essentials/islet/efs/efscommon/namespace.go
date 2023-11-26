package efscommon

import "github.com/watermint/toolbox/essentials/islet/efs"

const (
	nsIdentitySeparator = ":"
)

func NewNamespace(fs efs.FileSystem, ns string) efs.Namespace {
	return &namespaceImpl{
		ns: ns,
		fs: fs,
	}
}

type namespaceImpl struct {
	ns string
	fs efs.FileSystem
}

func (z namespaceImpl) Identity() string {
	fsId := z.fs.Identity()
	if fsId == "" {
		return z.ns
	}
	return fsId + nsIdentitySeparator + z.ns
}

func (z namespaceImpl) FileSystem() efs.FileSystem {
	return z.fs
}

func (z namespaceImpl) String() string {
	return z.ns
}

func (z namespaceImpl) Equals(other efs.Namespace) bool {
	return other.FileSystem().Equals(z.fs) && z.Identity() == other.Identity()
}

func (z namespaceImpl) IsDefault() bool {
	return z.ns == ""
}
