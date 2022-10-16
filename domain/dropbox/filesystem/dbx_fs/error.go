package dbx_fs

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
	dbxErr dbx_error.DropboxError
}

func (z dbxError) IsMockError() bool {
	return z.err == qt_errors.ErrorMock
}

func (z dbxError) Error() string {
	if z.err != nil {
		return z.err.Error()
	} else if z.dbxErr != nil {
		return z.dbxErr.Summary()
	} else {
		return "dbx_error: undefined error"
	}
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

const (
	cacheErrorNotFound cacheErrorType = iota
)

type cacheErrorType int

type cacheError struct {
	errorType cacheErrorType
}

func (z cacheError) Error() string {
	switch z.errorType {
	case cacheErrorNotFound:
		return "not found"
	default:
		return "other error"
	}
}

func (z cacheError) IsPathNotFound() bool {
	return z.errorType == cacheErrorNotFound
}

func (z cacheError) IsConflict() bool {
	return false
}

func (z cacheError) IsNoPermission() bool {
	return false
}

func (z cacheError) IsInsufficientSpace() bool {
	return false
}

func (z cacheError) IsDisallowedName() bool {
	return false
}

func (z cacheError) IsInvalidEntryDataFormat() bool {
	return false
}

func (z cacheError) IsMockError() bool {
	return false
}

func NotFoundError() es_filesystem.FileSystemError {
	return &cacheError{
		errorType: cacheErrorNotFound,
	}
}
