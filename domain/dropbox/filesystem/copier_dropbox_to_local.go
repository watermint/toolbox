package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewDropboxToLocal(ctx dbx_context.Context) es_filesystem.Connector {
	return &copierDropboxToLocal{
		ctx:    ctx,
		target: es_filesystem_local.NewFileSystem(),
	}
}

type copierDropboxToLocal struct {
	ctx    dbx_context.Context
	target es_filesystem.FileSystem
}

func (z copierDropboxToLocal) Copy(source es_filesystem.Entry, target es_filesystem.Path) (copied es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (download)")

	sourcePath, err := ToDropboxPath(source.Path())
	if err != nil {
		l.Debug("Unable to convert path format", esl.Error(err))
		return nil, err
	}

	downloadUrl, dbxErr := sv_file_content.NewDownload(z.ctx).DownloadUrl(sourcePath)
	if dbxErr != nil {
		l.Debug("Unable to download", esl.Error(dbxErr))
		return nil, NewError(dbxErr)
	}
	l.Debug("Download url", esl.String("url", downloadUrl))

	dlErr := es_download.Download(l, downloadUrl, target.Path())
	if dlErr != nil {
		l.Debug("Download failure", esl.Error(dlErr))
		return nil, NewError(dlErr)
	}

	return z.target.Info(target)
}
