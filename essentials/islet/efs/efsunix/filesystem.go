package efsunix

import (
	"github.com/watermint/toolbox/essentials/islet/efs"
	"github.com/watermint/toolbox/essentials/islet/efs/efscommon"
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

func (z fsImpl) Path(path string) (efs.Path, efs.PathOutcome) {
	//ap, err := filepath.Abs(path)
	//
	//// the error of filepath.Abs may be caused by retrieving current path
	//if err != nil {
	//	return nil, efscommon.newpath
	//}
	panic("implement me")
}

func (z fsImpl) Equals(other efs.FileSystem) bool {
	return other.Identity() == identity
}

func (z fsImpl) CurrentPath() (efs.Path, efs.CurrentPathOutcome) {
	panic("implement me")
}

func (z fsImpl) NameRule() efs.Name {
	return unixFileName
}
