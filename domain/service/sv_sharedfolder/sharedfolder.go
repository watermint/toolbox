package sv_sharedfolder

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
)

type SharedFolder interface {
	List() (sf []*mo_sharedfolder.SharedFolder, err error)
}

func New(ctx api_context.Context) SharedFolder {
	return &sharedFolderImpl{
		ctx: ctx,
	}
}

type sharedFolderImpl struct {
	ctx   api_context.Context
	limit int
}

func (z *sharedFolderImpl) List() (sf []*mo_sharedfolder.SharedFolder, err error) {
	p := struct {
		Limit int `json:"limit,omitempty"`
	}{
		Limit: z.limit,
	}
	sf = make([]*mo_sharedfolder.SharedFolder, 0)
	req := z.ctx.List("sharing/list_folders").
		Continue("sharing/list_folders/continue").
		Param(p).
		UseHasMore(false).
		ResultTag("entries").
		OnEntry(func(entry api_list.ListEntry) error {
			f := &mo_sharedfolder.SharedFolder{}
			if err := entry.Model(f); err != nil {
				return err
			}
			sf = append(sf, f)
			return nil
		})
	if err := req.Call(); err != nil {
		return nil, err
	}
	return sf, err
}
