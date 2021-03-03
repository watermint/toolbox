package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
)

func NewLocalToDropbox(ctx dbx_context.Context, opts ...sv_file_content.UploadOpt) es_filesystem.Connector {
	return &copierLocalToDropbox{
		ctx:        ctx,
		uploadOpts: opts,
	}
}

type copierLocalToDropbox struct {
	ctx        dbx_context.Context
	uploadOpts []sv_file_content.UploadOpt
}

func (z copierLocalToDropbox) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	return nil
}

func (z copierLocalToDropbox) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload)")
	cp := es_filesystem.NewCopyPair(source, target)

	targetDbxPath, err := ToDropboxPath(target.Ancestor())
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		onFailure(cp, err)
		return
	}

	svc := sv_file_content.NewUpload(z.ctx, z.uploadOpts...)
	dbxEntry, dbxErr := svc.Overwrite(targetDbxPath, source.Path().Path())
	if dbxErr != nil {
		l.Debug("Unable to upload file", esl.Error(dbxErr))
		onFailure(cp, NewError(dbxErr))
		return
	}

	l.Debug("successfully uploaded", esl.Any("entry", dbxEntry.Concrete()))
	onSuccess(cp, NewEntry(dbxEntry))
}

func (z copierLocalToDropbox) Shutdown() (err es_filesystem.FileSystemError) {
	z.ctx.Log().Debug("Shutdown")
	return nil
}
