package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
)

func NewDropboxToLocal(ctx dbx_context.Context) es_filesystem.Connector {
	return &connDropboxToLocal{
		ctx:    ctx,
		target: es_filesystem_local.NewFileSystem(),
	}
}

type connDropboxToLocal struct {
	ctx    dbx_context.Context
	target es_filesystem.FileSystem
}

func (z connDropboxToLocal) Copy(source es_filesystem.Entry, target es_filesystem.Path) (copied es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (download)")

	sourcePath, err := ToDropboxPath(source.Path())
	if err != nil {
		l.Debug("Unable to convert path format", esl.Error(err))
		return nil, err
	}

	dbxEntry, localPath, dbxErr := sv_file_content.NewDownload(z.ctx).Download(sourcePath)
	if dbxErr != nil {
		l.Debug("Unable to download", esl.Error(dbxErr))
		return nil, NewError(dbxErr)
	}
	l.Debug("Downloaded", esl.Any("dbxEntry", dbxEntry.Concrete()))

	osErr := os.Rename(localPath.Path(), target.Path())
	if osErr != nil {
		l.Debug("Unable to move tmp to target location", esl.Error(osErr))
		return nil, NewError(osErr)
	}

	return z.target.Info(target)
}
