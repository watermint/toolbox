package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_relocation"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewLocalToDropbox(ctx dbx_context.Context, opts ...sv_file_content.UploadOpt) es_filesystem.Connector {
	return &connLocalToDropbox{
		ctx:        ctx,
		uploadOpts: opts,
	}
}

func NewLocalToDropboxUpAndMove(ctx dbx_context.Context, workPath mo_path.DropboxPath, opts ...sv_file_content.UploadOpt) es_filesystem.Connector {
	return &connLocalToDropboxUpAndMoveStrategy{
		ctx:        ctx,
		uploadOpts: opts,
		workPath:   workPath,
	}
}

type connLocalToDropbox struct {
	ctx        dbx_context.Context
	uploadOpts []sv_file_content.UploadOpt
}

func (z connLocalToDropbox) Copy(source es_filesystem.Entry, target es_filesystem.Path) (copied es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload)")

	targetDbxPath, err := ToDropboxPath(target.Ancestor())
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return nil, err
	}

	svc := sv_file_content.NewUpload(z.ctx, z.uploadOpts...)
	dbxEntry, dbxErr := svc.Overwrite(targetDbxPath, source.Path().Path())
	if dbxErr != nil {
		l.Debug("Unable to upload file", esl.Error(dbxErr))
		return nil, NewError(dbxErr)
	}

	l.Debug("successfully uploaded", esl.Any("entry", dbxEntry.Concrete()))
	return NewEntry(dbxEntry), nil
}

type connLocalToDropboxUpAndMoveStrategy struct {
	ctx        dbx_context.Context
	uploadOpts []sv_file_content.UploadOpt
	workPath   mo_path.DropboxPath
}

func (z connLocalToDropboxUpAndMoveStrategy) Copy(source es_filesystem.Entry, target es_filesystem.Path) (copied es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload to work)")

	targetDbxPath, err := ToDropboxPath(target)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return nil, err
	}

	svc := sv_file_content.NewUpload(z.ctx, z.uploadOpts...)
	dbxEntry, dbxErr := svc.Overwrite(z.workPath, source.Path().Path())
	if dbxErr != nil {
		l.Debug("Unable to upload file", esl.Error(dbxErr))
		return nil, NewError(dbxErr)
	}

	movedDbxEntry, dbxErr := sv_file_relocation.New(z.ctx).Move(dbxEntry.Path(), targetDbxPath)
	if dbxErr != nil {
		l.Debug("Unable to upload file", esl.Error(dbxErr))
		return nil, NewError(dbxErr)
	}

	l.Debug("successfully uploaded", esl.Any("entry", movedDbxEntry.Concrete()))
	return NewEntry(dbxEntry), nil

}
