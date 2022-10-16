package dbx_fs_copier_dbx_to_dbx

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_relocation"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
)

func NewDropboxToDropbox(ctx dbx_client.Client) es_filesystem.Connector {
	return &copierDropboxToDropboxSingle{
		ctx: ctx,
	}
}

type copierDropboxToDropboxSingle struct {
	ctx dbx_client.Client
}

func (z copierDropboxToDropboxSingle) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	return nil
}

func (z copierDropboxToDropboxSingle) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy")
	cp := es_filesystem.NewCopyPair(source, target)

	sourceDbxPath, err := dbx_fs.ToDropboxPath(source.Path())
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		onFailure(cp, err)
		return
	}

	targetDbxPath, err := dbx_fs.ToDropboxPath(target)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		onFailure(cp, err)
		return
	}

	dbxEntry, dbxErr := sv_file_relocation.New(z.ctx).Copy(sourceDbxPath, targetDbxPath)
	if dbxErr != nil {
		l.Debug("Unable to copy", esl.Error(dbxErr))
		onFailure(cp, err)
		return
	}

	l.Debug("successfully copied", esl.Any("entry", dbxEntry.Concrete()))
	onSuccess(cp, dbx_fs.NewEntry(dbxEntry))
}

func (z copierDropboxToDropboxSingle) Shutdown() (err es_filesystem.FileSystemError) {
	z.ctx.Log().Debug("Shutdown")
	return nil
}
