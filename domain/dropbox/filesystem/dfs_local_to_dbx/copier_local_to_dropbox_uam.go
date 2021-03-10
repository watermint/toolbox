package dfs_local_to_dbx

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_relocation"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
)

func NewLocalToDropboxUpAndMove(ctx dbx_context.Context, workPath mo_path.DropboxPath, opts ...sv_file_content.UploadOpt) es_filesystem.Connector {
	return &connLocalToDropboxUpAndMoveStrategy{
		ctx:        ctx,
		uploadOpts: opts,
		workPath:   workPath,
	}
}

type connLocalToDropboxUpAndMoveStrategy struct {
	ctx        dbx_context.Context
	uploadOpts []sv_file_content.UploadOpt
	workPath   mo_path.DropboxPath
}

func (z connLocalToDropboxUpAndMoveStrategy) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	return nil
}

func (z connLocalToDropboxUpAndMoveStrategy) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload to work)")
	cp := es_filesystem.NewCopyPair(source, target)

	targetDbxPath, err := filesystem.ToDropboxPath(target)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		onFailure(cp, err)
		return
	}

	svc := sv_file_content.NewUpload(z.ctx, z.uploadOpts...)
	dbxEntry, dbxErr := svc.Overwrite(z.workPath, source.Path().Path())
	if dbxErr != nil {
		l.Debug("Unable to upload file", esl.Error(dbxErr))
		onFailure(cp, filesystem.NewError(dbxErr))
		return
	}

	movedDbxEntry, dbxErr := sv_file_relocation.New(z.ctx).Move(dbxEntry.Path(), targetDbxPath)
	if dbxErr != nil {
		l.Debug("Unable to upload file", esl.Error(dbxErr))
		onFailure(cp, filesystem.NewError(dbxErr))
		return
	}

	l.Debug("successfully uploaded", esl.Any("entry", movedDbxEntry.Concrete()))
	onSuccess(cp, filesystem.NewEntry(dbxEntry))
}

func (z connLocalToDropboxUpAndMoveStrategy) Shutdown() (err es_filesystem.FileSystemError) {
	z.ctx.Log().Debug("Shutdown")
	return nil
}
