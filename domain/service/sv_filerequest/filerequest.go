package sv_filerequest

import (
	"github.com/watermint/toolbox/domain/model/mo_filerequest"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
)

type FileRequest interface {
	List() (requests []*mo_filerequest.FileRequest, err error)
}

func New(ctx api_context.Context) FileRequest {
	return &fileRequestImpl{
		ctx: ctx,
	}
}

type fileRequestImpl struct {
	ctx api_context.Context
}

func (z *fileRequestImpl) List() (requests []*mo_filerequest.FileRequest, err error) {
	requests = make([]*mo_filerequest.FileRequest, 0)

	req := z.ctx.List("file_requests/list_v2").
		Continue("file_requests/list/continue").
		UseHasMore(true).
		ResultTag("file_requests").
		OnEntry(func(entry api_list.ListEntry) error {
			fr := &mo_filerequest.FileRequest{}
			if err := entry.Model(fr); err != nil {
				return err
			}
			requests = append(requests, fr)
			return nil
		})
	if err := req.Call(); err != nil {
		return nil, err
	}
	return requests, nil
}
