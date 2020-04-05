package uc_file_relocation

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_relocation"
	"github.com/watermint/toolbox/infra/api/api_context"
	"go.uber.org/zap"
)

type Relocation interface {
	// options: allow_shared_folder, allow_ownership_transfer, auto_rename
	Copy(from, to mo_path.DropboxPath) (err error)

	// options: allow_shared_folder, allow_ownership_transfer, auto_rename
	Move(from, to mo_path.DropboxPath) (err error)
}

func New(ctx api_context.DropboxApiContext) Relocation {
	return &relocationImpl{
		ctx: ctx,
	}
}

type relocationImpl struct {
	ctx api_context.DropboxApiContext
}

func (z *relocationImpl) Copy(from, to mo_path.DropboxPath) (err error) {
	return z.relocation(from, to, func(from, to mo_path.DropboxPath) (err error) {
		svc := sv_file_relocation.New(z.ctx)
		_, err = svc.Copy(from, to)
		return err
	})
}

func (z *relocationImpl) Move(from, to mo_path.DropboxPath) (err error) {
	return z.relocation(from, to, func(from, to mo_path.DropboxPath) (err error) {
		svc := sv_file_relocation.New(z.ctx)
		_, err = svc.Move(from, to)
		return err
	})
}

func (z *relocationImpl) relocation(from, to mo_path.DropboxPath,
	reloc func(from, to mo_path.DropboxPath) (err error)) (err error) {
	l := z.ctx.Log().With(zap.String("from", from.Path()), zap.String("to", to.Path()))

	svc := sv_file.NewFiles(z.ctx)

	fromEntry, err := svc.Resolve(from)
	if err != nil {
		l.Debug("Cannot resolve from", zap.Error(err))
		return err
	}
	var fromToTag string
	if to.LogicalPath() == "/" {
		fromToTag = fromEntry.Tag() + "-folder"
		l = l.With(zap.String("fromTag", fromEntry.Tag()), zap.String("toTag", "root"))
	} else {
		toEntry, err := svc.Resolve(to)
		if err != nil {
			es := dbx_util.ErrorSummary(err)
			switch es {
			case "path/not_found":
				l.Debug("To not found. Do relocate", zap.Error(err))
				return reloc(from, to)
			}
			l.Debug("Invalid path to relocate, or restricted", zap.Error(err), zap.String("summary", es))
			return err
		}
		fromToTag = fromEntry.Tag() + "-" + toEntry.Tag()
		l = l.With(zap.String("fromTag", fromEntry.Tag()), zap.String("toTag", toEntry.Tag()))
	}

	switch fromToTag {
	case "file-file":
		l.Debug("Do relocate")
		return reloc(from, to)

	case "file-folder", "folder-folder":
		l.Debug("Do relocate into folder")
		toPath := to.ChildPath(fromEntry.Name())
		return reloc(from, toPath)

	case "folder-file":
		l.Debug("Not a folder")
		return errors.New("not a folder")

	default:
		return errors.New("unsupported file/folder type")
	}
}
