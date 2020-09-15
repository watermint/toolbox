package filesystem

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_relocation"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewDropboxToDropbox(ctx dbx_context.Context) es_filesystem.Connector {
	return &copierDropboxToDropboxSingle{
		ctx: ctx,
	}
}

type copierDropboxToDropboxSingle struct {
	ctx dbx_context.Context
}

func (z copierDropboxToDropboxSingle) Copy(source es_filesystem.Entry, target es_filesystem.Path) (copied es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy")

	sourceDbxPath, err := ToDropboxPath(source.Path())
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return nil, err
	}

	targetDbxPath, err := ToDropboxPath(target)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return nil, err
	}

	dbxEntry, dbxErr := sv_file_relocation.New(z.ctx).Copy(sourceDbxPath, targetDbxPath)
	if dbxErr != nil {
		l.Debug("Unable to copy", esl.Error(dbxErr))
		return nil, err
	}

	l.Debug("successfully copied", esl.Any("entry", dbxEntry.Concrete()))
	return NewEntry(dbxEntry), nil
}
