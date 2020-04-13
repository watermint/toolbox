package sv_file_content

import (
	"errors"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"go.uber.org/zap"
	"os"
)

type Download interface {
	Download(path mo_path.DropboxPath) (entry mo_file.Entry, localPath mo_path2.FileSystemPath, err error)
}

func NewDownload(ctx dbx_context.Context) Download {
	return &downloadImpl{ctx: ctx}
}

type downloadImpl struct {
	ctx dbx_context.Context
}

func (z *downloadImpl) Download(path mo_path.DropboxPath) (entry mo_file.Entry, localPath mo_path2.FileSystemPath, err error) {
	l := z.ctx.Log()
	p := struct {
		Path string `json:"path"`
	}{
		Path: path.Path(),
	}

	res, err := z.ctx.Download("files/download").Param(p).Call()
	if err != nil {
		return nil, nil, err
	}
	if !res.IsContentDownloaded() {
		return nil, nil, errors.New("content was not downloaded")
	}
	entry = &mo_file.Metadata{}
	if err := res.Model(entry); err != nil {
		// Try remove downloaded file
		if removeErr := os.Remove(res.ContentFilePath().Path()); removeErr != nil {
			l.Debug("Unable to remove downloaded file", zap.Error(err), zap.String("path", res.ContentFilePath().Path()))
			// fall through
		}

		return nil, nil, err
	}
	return entry, res.ContentFilePath(), nil
}
