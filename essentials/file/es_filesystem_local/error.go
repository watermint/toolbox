package es_filesystem_local

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"os"
)

func NewError(err error) es_filesystem.FileSystemError {
	if err == nil {
		return nil
	} else {
		return &errImpl{
			err: err,
		}
	}
}

var (
	ErrorInvalidEntryDataFormat = errors.New("invalid entry data format")
)

type errImpl struct {
	err error
}

func (z errImpl) IsMockError() bool {
	return false
}

func (z errImpl) Error() string {
	return z.err.Error()
}

func (z errImpl) IsPathNotFound() bool {
	return os.IsNotExist(z.err)
}

func (z errImpl) IsConflict() bool {
	return os.IsExist(z.err)
}

func (z errImpl) IsNoPermission() bool {
	return os.IsPermission(z.err)
}

func (z errImpl) IsInsufficientSpace() bool {
	return false
}

func (z errImpl) IsDisallowedName() bool {
	return false
}

func (z errImpl) IsInvalidEntryDataFormat() bool {
	return z.err == ErrorInvalidEntryDataFormat
}
