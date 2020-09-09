package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewLocalToDropbox(ctx dbx_context.Context, opts ...sv_file_content.UploadOpt) es_filesystem.Connector {
	return &connLocalToDropbox{
		ctx:        ctx,
		uploadOpts: opts,
	}
}

type connLocalToDropbox struct {
	ctx        dbx_context.Context
	uploadOpts []sv_file_content.UploadOpt
}

func (z connLocalToDropbox) Copy(source es_filesystem.Entry, target es_filesystem.Path) (err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload)")

	targetDbxPath, err := ToDropboxPath(target)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return err
	}

	svc := sv_file_content.NewUpload(z.ctx, z.uploadOpts...)
	dbxEntry, dbxErr := svc.Overwrite(targetDbxPath, source.Path().Path())
	if dbxErr != nil {
		l.Debug("Unable to upload file", esl.Error(dbxErr))
		return NewError(dbxErr)
	}

	l.Debug("successfully uploaded", esl.Any("entry", dbxEntry.Concrete()))
	return nil
}
