package sv_file_content

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/time/ut_format"
	"github.com/watermint/toolbox/infra/api/api_request"
	"os"
	"time"
)

var (
	ErrorDownloadUrlNotFound = errors.New("temporary link not found")
)

type Download interface {
	Download(path mo_path.DropboxPath) (entry mo_file.Entry, localPath mo_path2.FileSystemPath, err error)
	DownloadUrl(path mo_path.DropboxPath) (url string, err error)
}

type Opts struct {
	SharedLinkUrl string `json:"shared_link_url"`
	Path          string `json:"path"`
}

func NewDownload(ctx dbx_context.Context) Download {
	return &downloadImpl{ctx: ctx}
}

type downloadImpl struct {
	ctx dbx_context.Context
}

func (z *downloadImpl) DownloadUrl(path mo_path.DropboxPath) (url string, err error) {
	l := z.ctx.Log()
	p := struct {
		Path string `json:"path"`
	}{
		Path: path.Path(),
	}

	res := z.ctx.Post("files/get_temporary_link", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		l.Debug("Unable to retrieve temporary download link", esl.Error(err))
		return "", err
	}

	resJson := res.Success().Json()
	if url, found := resJson.FindString("link"); found {
		return url, nil
	}
	return "", ErrorDownloadUrlNotFound
}

func (z *downloadImpl) Download(path mo_path.DropboxPath) (entry mo_file.Entry, localPath mo_path2.FileSystemPath, err error) {
	l := z.ctx.Log()
	p := struct {
		Path string `json:"path"`
	}{
		Path: path.Path(),
	}

	q, err := dbx_request.DropboxApiArg(p)
	if err != nil {
		l.Debug("unable to marshal parameter", esl.Error(err))
		return nil, nil, err
	}

	res := z.ctx.Download("files/download", q)
	if err, f := res.Failure(); f {
		return nil, nil, err
	}
	contentFilePath, err := res.Success().AsFile()
	if err != nil {
		return nil, nil, err
	}
	resData := dbx_context.ContentResponseData(res)

	entry = &mo_file.Metadata{}
	if err := resData.Model(entry); err != nil {
		// Try remove downloaded file
		if removeErr := os.Remove(contentFilePath); removeErr != nil {
			l.Debug("Unable to remove downloaded file",
				esl.Error(err),
				esl.String("path", contentFilePath))
			// fall through
		}

		return nil, nil, err
	}

	// update file timestamp
	clientModified := entry.Concrete().ClientModified
	ftm, ok := ut_format.ParseTimestamp(clientModified)
	if !ok {
		l.Debug("Unable to parse client modified", esl.String("client_modified", clientModified))
	} else if err := os.Chtimes(contentFilePath, time.Now(), ftm); err != nil {
		l.Debug("Unable to change time", esl.Error(err))
	}
	return entry, mo_path2.NewFileSystemPath(contentFilePath), nil
}
