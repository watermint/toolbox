package uc_file_traverse

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
)

// Recursive file list
type Traverse interface {
	Traverse(path mo_path.Path, onEntry func(entry mo_file.Entry) error) error
}

func New(ctx api_context.Context) Traverse {
	return &traverseImpl{
		ctx: ctx,
	}
}

type traverseImpl struct {
	ctx api_context.Context
}

func (z *traverseImpl) traversePath(svc sv_file.Files, path mo_path.Path, onEntry func(entry mo_file.Entry) error) error {
	entries, err := svc.List(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if err := onEntry(entry); err != nil {
			return err
		}
		if f, e := entry.Folder(); e {
			err = z.traversePath(svc, path.ChildPath(f.Name()), onEntry)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (z *traverseImpl) Traverse(path mo_path.Path, onEntry func(entry mo_file.Entry) error) error {
	svc := sv_file.NewFiles(z.ctx)
	return z.traversePath(svc, path, onEntry)
}
