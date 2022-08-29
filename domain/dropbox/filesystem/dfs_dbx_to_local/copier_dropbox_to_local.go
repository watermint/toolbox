package dfs_dbx_to_local

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
)

func NewDropboxToLocal(ctx dbx_client.Client) es_filesystem.Connector {
	return &copierDropboxToLocal{
		ctx:    ctx,
		target: es_filesystem_local.NewFileSystem(),
	}
}

type copierDropboxToLocal struct {
	ctx    dbx_client.Client
	target es_filesystem.FileSystem
}

func (z copierDropboxToLocal) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	return nil
}

func (z copierDropboxToLocal) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (download)")
	cp := es_filesystem.NewCopyPair(source, target)

	sourcePath, err := filesystem.ToDropboxPath(source.Path())
	if err != nil {
		l.Debug("Unable to convert path format", esl.Error(err))
		onFailure(cp, err)
		return
	}

	downloadUrl, dbxErr := sv_file_content.NewDownload(z.ctx).DownloadUrl(sourcePath)
	if dbxErr != nil {
		l.Debug("Unable to download", esl.Error(dbxErr))
		onFailure(cp, filesystem.NewError(dbxErr))
		return
	}
	l.Debug("Download url", esl.String("url", downloadUrl))

	dlErr := es_download.Download(l, downloadUrl, target.Path())
	if dlErr != nil {
		l.Debug("Download failure", esl.Error(dlErr))
		onFailure(cp, filesystem.NewError(dlErr))
		return
	}

	if entry, fsErr := z.target.Info(target); fsErr != nil {
		onFailure(cp, fsErr)
	} else {
		onSuccess(cp, entry)
	}
}

func (z copierDropboxToLocal) Shutdown() (err es_filesystem.FileSystemError) {
	z.ctx.Log().Debug("Shutdown")
	return nil
}
