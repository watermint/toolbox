package dbx_fs

import (
	"errors"
	"testing"

	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

func TestNewError(t *testing.T) {
	testErr := errors.New("test error")
	fsErr := NewError(testErr)

	if fsErr == nil {
		t.Fatal("Expected non-nil error")
	}

	dbxErr, ok := fsErr.(*dbxError)
	if !ok {
		t.Fatal("Expected dbxError type")
	}

	if dbxErr.err != testErr {
		t.Error("Error not set correctly")
	}
}

func TestDbxError_IsMockError(t *testing.T) {
	// Test with mock error
	mockErr := &dbxError{
		err: qt_errors.ErrorMock,
	}
	if !mockErr.IsMockError() {
		t.Error("Expected IsMockError to return true for ErrorMock")
	}

	// Test with regular error
	regularErr := &dbxError{
		err: errors.New("regular error"),
	}
	if regularErr.IsMockError() {
		t.Error("Expected IsMockError to return false for regular error")
	}
}

func TestDbxError_Error(t *testing.T) {
	testCases := []struct {
		name     string
		dbxErr   *dbxError
		expected string
	}{
		{
			name: "with err",
			dbxErr: &dbxError{
				err: errors.New("test error"),
			},
			expected: "test error",
		},
		{
			name: "with nil err and nil dbxErr",
			dbxErr: &dbxError{
				err:    nil,
				dbxErr: nil,
			},
			expected: "dbx_error: undefined error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.dbxErr.Error()
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

func TestDbxError_IsInvalidEntryDataFormat(t *testing.T) {
	// Test with ErrorInvalidEntryDataFormat
	invalidErr := &dbxError{
		err: ErrorInvalidEntryDataFormat,
	}
	if !invalidErr.IsInvalidEntryDataFormat() {
		t.Error("Expected IsInvalidEntryDataFormat to return true")
	}

	// Test with different error
	otherErr := &dbxError{
		err: errors.New("other error"),
	}
	if otherErr.IsInvalidEntryDataFormat() {
		t.Error("Expected IsInvalidEntryDataFormat to return false")
	}
}

func TestDbxError_UnimplementedMethods(t *testing.T) {
	dbxErr := &dbxError{
		err: errors.New("test"),
	}

	// Test IsNoPermission - should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected IsNoPermission to panic")
		}
	}()
	dbxErr.IsNoPermission()
}

func TestDbxError_IsInsufficientSpace(t *testing.T) {
	dbxErr := &dbxError{
		err: errors.New("test"),
	}

	// Test IsInsufficientSpace - should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected IsInsufficientSpace to panic")
		}
	}()
	dbxErr.IsInsufficientSpace()
}

func TestDbxError_IsDisallowedName(t *testing.T) {
	dbxErr := &dbxError{
		err: errors.New("test"),
	}

	// Test IsDisallowedName - should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected IsDisallowedName to panic")
		}
	}()
	dbxErr.IsDisallowedName()
}

func TestCacheError_Error(t *testing.T) {
	testCases := []struct {
		name      string
		errorType cacheErrorType
		expected  string
	}{
		{
			name:      "not found",
			errorType: cacheErrorNotFound,
			expected:  "not found",
		},
		{
			name:      "other error",
			errorType: cacheErrorType(999),
			expected:  "other error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := &cacheError{errorType: tc.errorType}
			if err.Error() != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, err.Error())
			}
		})
	}
}

func TestCacheError_IsPathNotFound(t *testing.T) {
	// Test with not found error
	notFoundErr := &cacheError{errorType: cacheErrorNotFound}
	if !notFoundErr.IsPathNotFound() {
		t.Error("Expected IsPathNotFound to return true")
	}

	// Test with other error
	otherErr := &cacheError{errorType: cacheErrorType(999)}
	if otherErr.IsPathNotFound() {
		t.Error("Expected IsPathNotFound to return false")
	}
}

func TestCacheError_OtherMethods(t *testing.T) {
	err := &cacheError{errorType: cacheErrorNotFound}

	// All these should return false
	if err.IsConflict() {
		t.Error("Expected IsConflict to return false")
	}
	if err.IsNoPermission() {
		t.Error("Expected IsNoPermission to return false")
	}
	if err.IsInsufficientSpace() {
		t.Error("Expected IsInsufficientSpace to return false")
	}
	if err.IsDisallowedName() {
		t.Error("Expected IsDisallowedName to return false")
	}
	if err.IsInvalidEntryDataFormat() {
		t.Error("Expected IsInvalidEntryDataFormat to return false")
	}
	if err.IsMockError() {
		t.Error("Expected IsMockError to return false")
	}
}

func TestNotFoundError(t *testing.T) {
	err := NotFoundError()
	if err == nil {
		t.Fatal("Expected non-nil error")
	}

	cacheErr, ok := err.(*cacheError)
	if !ok {
		t.Fatal("Expected cacheError type")
	}

	if cacheErr.errorType != cacheErrorNotFound {
		t.Error("Expected cacheErrorNotFound type")
	}

	if !err.IsPathNotFound() {
		t.Error("Expected IsPathNotFound to return true")
	}
}

func TestErrorConstants(t *testing.T) {
	// Test error constants
	if ErrorInvalidEntryDataFormat.Error() != "invalid entry data format" {
		t.Error("ErrorInvalidEntryDataFormat has unexpected message")
	}
	if ErrorInvalidEntryType.Error() != "invalid entry type" {
		t.Error("ErrorInvalidEntryType has unexpected message")
	}
}

// Test that all methods in FileSystemError interface are implemented
func TestFileSystemErrorInterface(t *testing.T) {
	var _ es_filesystem.FileSystemError = &dbxError{}
	var _ es_filesystem.FileSystemError = &cacheError{}
}
