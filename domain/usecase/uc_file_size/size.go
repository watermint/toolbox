package uc_file_size

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_file_size"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
)

type Scale interface {
	Size(path mo_path.Path, depth int) (sizes map[mo_path.Path]mo_file_size.Size, err error)
}

func New(ctx api_context.Context) Scale {
	return &scaleImpl{
		ctx: ctx,
	}
}

type scaleImpl struct {
	ctx api_context.Context
}

func (z *scaleImpl) Size(path mo_path.Path, depth int) (sizes map[mo_path.Path]mo_file_size.Size, err error) {
	sizes = make(map[mo_path.Path]mo_file_size.Size)
	svc := sv_file.NewFiles(z.ctx)

	var traverse func(path mo_path.Path, d int) (current mo_file_size.Size, err error)
	traverse = func(path mo_path.Path, d int) (current mo_file_size.Size, err error) {
		current = mo_file_size.Size{
			Path: path.Path(),
		}
		entries, err := svc.List(path)
		if err != nil {
			return
		}
		for _, entry := range entries {
			current.CountDescendant++
			if f, e := entry.File(); e {
				current.CountFile++
				current.Size += f.Size
			}
			if f, e := entry.Folder(); e {
				current.CountFolder++
				ds, err := traverse(path.ChildPath(f.Name()), d+1)
				if err != nil {
					return current, err
				}
				current = current.Plus(ds)
			}
		}
		if d < depth {
			sizes[mo_path.NewPath(current.Path)] = current
		}
		return current, nil
	}

	_, err = traverse(path, 0)
	if err != nil {
		return nil, err
	}
	return sizes, nil
}
