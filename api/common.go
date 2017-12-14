package api

import (
	"errors"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"io"
	"path/filepath"
	"strings"
	"time"
)

const (
	THINSDK_RETRY_REASON_NORETRY = iota
	THINSDK_RETRY_REASON_THROTTLING
	THINSDK_RETRY_REASON_TEMPORARY_NETWORK_ERROR
	THINSDK_RETRY_REASON_HEURISTIC
)

const (
	THINSDK_API_CALL_RETRY_INTERVAL = 60
	THINSDK_API_CALL_RETRY_MAX      = 100
)

type RetryReason int

type ApiContext struct {
	Config dropbox.Config
}

type DropboxPath struct {
	Path string
}

func NewDropboxPath(path string) *DropboxPath {
	return &DropboxPath{
		Path: path,
	}
}

func (d *DropboxPath) CleanPath() string {
	p := filepath.ToSlash(filepath.Clean(d.Path))
	if p == "/" {
		return ""
	} else {
		return p
	}
}

func RebaseTimeForAPI(t time.Time) time.Time {
	return t.UTC().Round(time.Second)
}

func IsRetriableError(err error) RetryReason {
	if err == nil {
		return THINSDK_RETRY_REASON_NORETRY
	}
	if strings.HasPrefix(err.Error(), "too_many_requests") {
		return THINSDK_RETRY_REASON_THROTTLING
	}
	if strings.HasPrefix(err.Error(), "too_many_write_operations") {
		return THINSDK_RETRY_REASON_THROTTLING
	}
	if strings.Contains(err.Error(), "An existing connection was forcibly closed by the remote host") {
		return THINSDK_RETRY_REASON_TEMPORARY_NETWORK_ERROR
	}
	if strings.HasPrefix(err.Error(), "unexpected end of JSON input") {
		return THINSDK_RETRY_REASON_HEURISTIC
	}
	return THINSDK_RETRY_REASON_NORETRY
}

func (a *ApiContext) call0(api func() error) (err error) {
	sameErrorCount := 0
	lastError := ""

	for {
		err = api()
		if err == nil {
			return
		}
		if lastError == err.Error() {
			sameErrorCount++
			if THINSDK_API_CALL_RETRY_MAX <= sameErrorCount {
				seelog.Warnf("Exceed retry maximum number[%d] for same error[%s]", THINSDK_API_CALL_RETRY_MAX, err.Error())
				return
			}
		}
		lastError = err.Error()

		if r := IsRetriableError(err); r != THINSDK_RETRY_REASON_NORETRY {
			seelog.Debugf("Retriable error[%s] reason[%d]: Wait for [%d] seconds, then retry", err, r, THINSDK_API_CALL_RETRY_INTERVAL)
			time.Sleep(THINSDK_API_CALL_RETRY_INTERVAL * time.Second)
			continue
		} else {
			return
		}
	}
}

func (a *ApiContext) call1(api func() (interface{}, error)) (res interface{}, err error) {
	sameErrorCount := 0
	lastError := ""

	for {
		res, err = api()
		if err == nil {
			return
		}
		if lastError == err.Error() {
			sameErrorCount++
			if THINSDK_API_CALL_RETRY_MAX <= sameErrorCount {
				seelog.Warnf("Exceed retry maximum number[%d] for same error[%s]", THINSDK_API_CALL_RETRY_MAX, err.Error())
				return
			}
		}
		lastError = err.Error()

		if r := IsRetriableError(err); r != THINSDK_RETRY_REASON_NORETRY {
			seelog.Debugf("Retriable error[%s] reason[%d]: Wait for [%d] seconds, then retry", err, r, THINSDK_API_CALL_RETRY_INTERVAL)
			time.Sleep(THINSDK_API_CALL_RETRY_INTERVAL * time.Second)
			continue
		} else {
			return
		}
	}
}

// wrapper for files.ListFolderResult
func (a *ApiContext) filesListFolderResult(f func() (*files.ListFolderResult, error)) (r *files.ListFolderResult, err error) {
	r0, err := a.call1(func() (interface{}, error) {
		return f()
	})
	if err != nil {
		return
	}
	switch r1 := r0.(type) {
	case *files.ListFolderResult:
		return r1, err
	default:
		seelog.Warn("unexpected result type")
		return nil, errors.New("unexpected result type")
	}
}

// wrapper for files.RelocationResult
func (a *ApiContext) filesRelocationResult(f func() (*files.RelocationResult, error)) (r *files.RelocationResult, err error) {
	r0, err := a.call1(func() (interface{}, error) {
		return f()
	})
	if err != nil {
		return
	}
	switch r1 := r0.(type) {
	case *files.RelocationResult:
		return r1, err
	default:
		seelog.Warn("unexpected result type")
		return nil, errors.New("unexpected result type")
	}
}

// wrapper for files.DeleteResult
func (a *ApiContext) filesDeleteResult(f func() (*files.DeleteResult, error)) (r *files.DeleteResult, err error) {
	r0, err := a.call1(func() (interface{}, error) {
		return f()
	})
	if err != nil {
		return
	}
	switch r1 := r0.(type) {
	case *files.DeleteResult:
		return r1, err
	default:
		seelog.Warn("unexpected result type")
		return nil, errors.New("unexpected result type")
	}
}

// wrapper for files.MetadataResult
func (a *ApiContext) filesMetadataResult(f func() (files.IsMetadata, error)) (r files.IsMetadata, err error) {
	r0, err := a.call1(func() (interface{}, error) {
		return f()
	})
	if err != nil {
		return
	}
	switch r1 := r0.(type) {
	case files.IsMetadata:
		return r1, err
	default:
		seelog.Warn("unexpected result type")
		return nil, errors.New("unexpected result type")
	}
}

// wrapper for files.FileMetadataResult
func (a *ApiContext) filesFileMetadataResult(f func() (*files.FileMetadata, error)) (r *files.FileMetadata, err error) {
	r0, err := a.call1(func() (interface{}, error) {
		return f()
	})
	if err != nil {
		return
	}
	switch r1 := r0.(type) {
	case *files.FileMetadata:
		return r1, err
	default:
		seelog.Warn("unexpected result type")
		return nil, errors.New("unexpected result type")
	}
}

// wrapper for UploadSessionStartResult
func (a *ApiContext) filesUploadSessionStartResult(f func() (*files.UploadSessionStartResult, error)) (r *files.UploadSessionStartResult, err error) {
	r0, err := a.call1(func() (interface{}, error) {
		return f()
	})
	if err != nil {
		return
	}
	switch r1 := r0.(type) {
	case *files.UploadSessionStartResult:
		return r1, err
	default:
		seelog.Warn("unexpected result type")
		return nil, errors.New("unexpected result type")
	}
}

// API /files/list_folder
func (a *ApiContext) FilesListFolder(arg *files.ListFolderArg) (r *files.ListFolderResult, err error) {
	return a.filesListFolderResult(func() (*files.ListFolderResult, error) {
		return files.New(a.Config).ListFolder(arg)
	})
}

// API /files/list_folder_continue
func (a *ApiContext) FilesListFolderContinue(arg *files.ListFolderContinueArg) (r *files.ListFolderResult, err error) {
	return a.filesListFolderResult(func() (*files.ListFolderResult, error) {
		return files.New(a.Config).ListFolderContinue(arg)
	})
}

// API /files/copy_v2
func (a *ApiContext) FilesCopyV2(arg *files.RelocationArg) (r *files.RelocationResult, err error) {
	return a.filesRelocationResult(func() (*files.RelocationResult, error) {
		return files.New(a.Config).CopyV2(arg)
	})
}

// API /files/move
func (a *ApiContext) FilesMoveV2(arg *files.RelocationArg) (r *files.RelocationResult, err error) {
	return a.filesRelocationResult(func() (*files.RelocationResult, error) {
		return files.New(a.Config).MoveV2(arg)
	})
}

// API /files/delete
func (a *ApiContext) FilesDeleteV2(arg *files.DeleteArg) (r *files.DeleteResult, err error) {
	return a.filesDeleteResult(func() (*files.DeleteResult, error) {
		return files.New(a.Config).DeleteV2(arg)
	})
}

// API /files/get_metadata
func (a *ApiContext) FilesGetMetadata(arg *files.GetMetadataArg) (r files.IsMetadata, err error) {
	return a.filesMetadataResult(func() (files.IsMetadata, error) {
		return files.New(a.Config).GetMetadata(arg)
	})
}

// API /files/upload
func (a *ApiContext) FilesUpload(ci *files.CommitInfo, content io.Reader) (r *files.FileMetadata, err error) {
	return a.filesFileMetadataResult(func() (metadata *files.FileMetadata, err error) {
		return files.New(a.Config).Upload(ci, content)
	})
}

// API /files/upload_session/start
func (a *ApiContext) FilesUploadSessionStart(arg *files.UploadSessionStartArg, content io.Reader) (r *files.UploadSessionStartResult, err error) {
	return a.filesUploadSessionStartResult(func() (*files.UploadSessionStartResult, error) {
		return files.New(a.Config).UploadSessionStart(arg, content)
	})
}

// API /files/upload_session/append_v2
func (a *ApiContext) FilesUploadSessionAppendV2(arg *files.UploadSessionAppendArg, content io.Reader) (err error) {
	return a.call0(func() error {
		return files.New(a.Config).UploadSessionAppendV2(arg, content)
	})
}

// API /files/upload_session/finish
func (a *ApiContext) FilesUploadSessionFinish(arg *files.UploadSessionFinishArg, content io.Reader) (r *files.FileMetadata, err error) {
	return a.filesFileMetadataResult(func() (metadata *files.FileMetadata, err error) {
		return files.New(a.Config).UploadSessionFinish(arg, content)
	})
}
