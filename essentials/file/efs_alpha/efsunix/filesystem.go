package efsunix

import (
	efs_deprecated2 "github.com/watermint/toolbox/essentials/file/efs_alpha"
	"github.com/watermint/toolbox/essentials/file/efs_alpha/efscommon"
)

const (
	identity = "efsunix"
)

var (
	unixFileName = efscommon.NewName(efscommon.DefineNameInvalidChars('/'))
)

type fsImpl struct {
}

func (z fsImpl) Identity() string {
	return identity
}

func (z fsImpl) Path(path string) (efs_deprecated2.Path, efs_deprecated2.PathOutcome) {
	//ap, err := filepath.Abs(path)
	//
	//// the error of filepath.Abs may be caused by retrieving current path
	//if err != nil {
	//	return nil, efscommon.newpath
	//}
	panic("implement me")
}

func (z fsImpl) Equals(other efs_deprecated2.FileSystem) bool {
	return other.Identity() == identity
}

func (z fsImpl) CurrentPath() (efs_deprecated2.Path, efs_deprecated2.CurrentPathOutcome) {
	panic("implement me")
}

func (z fsImpl) NameRule() efs_deprecated2.Name {
	return unixFileName
}
