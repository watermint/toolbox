package filesystem

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

var (
	ErrorInvalidEntryDataFormat = errors.New("invalid entry data format")
	ErrorInvalidEntryType       = errors.New("invalid entry type")
)

func NewError(err error) es_filesystem.FileSystemError {
	return &dbxError{
		err:    err,
		dbxErr: dbx_error.NewErrors(err),
	}
}

type dbxError struct {
	err    error
	dbxErr dbx_error.Errors
}

func (z dbxError) IsMockError() bool {
	return z.err == qt_errors.ErrorMock
}

func (z dbxError) Error() string {
	return z.err.Error()
}

func (z dbxError) IsPathNotFound() bool {
	return z.dbxErr.Path().IsNotFound()
}

func (z dbxError) IsConflict() bool {
	return z.dbxErr.Path().IsConflict()
}

func (z dbxError) IsNoPermission() bool {
	panic("implement me")
}

func (z dbxError) IsInsufficientSpace() bool {
	panic("implement me")
}

func (z dbxError) IsDisallowedName() bool {
	panic("implement me")
}

func (z dbxError) IsInvalidEntryDataFormat() bool {
	return z.err == ErrorInvalidEntryDataFormat
}
