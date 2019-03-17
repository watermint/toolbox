package uc_relocation

import (
	"errors"
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_relocation"
	"go.uber.org/zap"
)

type Relocation interface {
	// options: allow_shared_folder, allow_ownership_transfer, auto_rename
	Copy(from, to mo_path.Path) (msg app_ui.UIMessage, err error)

	// options: allow_shared_folder, allow_ownership_transfer, auto_rename
	Move(from, to mo_path.Path) (msg app_ui.UIMessage, err error)
}

func New(ctx api_context.Context) Relocation {
	return &relocationImpl{
		ctx: ctx,
	}
}

type relocationImpl struct {
	ctx api_context.Context
}

func (z *relocationImpl) Copy(from, to mo_path.Path) (msg app_ui.UIMessage, err error) {
	return z.relocation(from, to, func(from, to mo_path.Path) (msg app_ui.UIMessage, err error) {
		svc := sv_relocation.New(z.ctx)
		_, err = svc.Copy(from, to)
		return z.ctx.ErrorMsg(err), err
	})
}

func (z *relocationImpl) Move(from, to mo_path.Path) (msg app_ui.UIMessage, err error) {
	return z.relocation(from, to, func(from, to mo_path.Path) (msg app_ui.UIMessage, err error) {
		svc := sv_relocation.New(z.ctx)
		_, err = svc.Move(from, to)
		return z.ctx.ErrorMsg(err), err
	})
}

func (z *relocationImpl) relocation(from, to mo_path.Path,
	reloc func(from, to mo_path.Path) (msg app_ui.UIMessage, err error)) (msg app_ui.UIMessage, err error) {
	log := z.ctx.Log().With(zap.String("from", from.Path()), zap.String("to", to.Path()))

	svc := sv_file.NewFiles(z.ctx)

	fromEntry, err := svc.Resolve(from)
	if err != nil {
		log.Debug("Cannot resolve from", zap.Error(err))
		msg := z.ctx.Msg("usecase.relocation.err.from_not_found").WithData(struct {
			FromPath string
		}{
			FromPath: from.Path(),
		})
		return msg, err
	}
	var fromToTag string
	if to.LogicalPath() == "/" {
		fromToTag = fromEntry.Tag() + "-folder"
		log = log.With(zap.String("fromTag", fromEntry.Tag()), zap.String("toTag", "root"))
	} else {
		toEntry, err := svc.Resolve(to)
		if err != nil {
			switch api_util.ErrorSummary(err) {
			case "path/not_found":
				log.Debug("To not found. Do relocate", zap.Error(err))
				return reloc(from, to)
			}
			log.Debug("Invalid path to relocate, or restricted", zap.Error(err))
		}
		fromToTag = fromEntry.Tag() + "-" + toEntry.Tag()
		log = log.With(zap.String("fromTag", fromEntry.Tag()), zap.String("toTag", toEntry.Tag()))
	}

	switch fromToTag {
	case "file-file":
		log.Debug("Do relocate")
		return reloc(from, to)

	case "file-folder", "folder-folder":
		log.Debug("Do relocate into folder")
		toPath := to.ChildPath(fromEntry.Name())
		return reloc(from, toPath)

	case "folder-file":
		log.Debug("Not a folder")
		msg = z.ctx.Msg("usecase.relocation.err.not_a_folder").WithData(struct {
			ToPath string
		}{
			ToPath: to.Path(),
		})
		return msg, errors.New("not a folder")

	default:
		return z.ctx.Msg("usecase.relocation.err.unsupported_type"), errors.New("unsupported file/folder type")
	}
}
