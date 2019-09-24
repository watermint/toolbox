package sv_filerequest

import (
	"github.com/watermint/toolbox/domain/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_list"
)

type FileRequest interface {
	List() (requests []*mo_filerequest.FileRequest, err error)
	Create(title string, destination mo_path.Path, opts ...UpdateOpt) (req *mo_filerequest.FileRequest, err error)
}

type UpdateOpt func(opts *UpdateOpts) *UpdateOpts

type UpdateOpts struct {
	Open                     bool
	Deadline                 string
	DeadlineAllowLateUploads string
}
type Deadline struct {
	Deadline         string `json:"deadline,omitempty"`
	AllowLateUploads string `json:"allow_late_uploads,omitempty"`
}

func New(ctx api_context.Context) FileRequest {
	return &fileRequestImpl{
		ctx: ctx,
	}
}

type fileRequestImpl struct {
	ctx api_context.Context
}

func (z *fileRequestImpl) Create(title string, destination mo_path.Path, opts ...UpdateOpt) (req *mo_filerequest.FileRequest, err error) {
	uo := &UpdateOpts{}
	for _, opt := range opts {
		opt(uo)
	}
	co := struct {
		Title       string   `json:"title"`
		Destination string   `json:"destination"`
		Deadline    Deadline `json:"deadline,omitempty"`
	}{
		Title:       title,
		Destination: destination.Path(),
		Deadline: Deadline{
			Deadline:         uo.Deadline,
			AllowLateUploads: uo.DeadlineAllowLateUploads,
		},
	}
	fr := &mo_filerequest.FileRequest{}
	res, err := z.ctx.Request("file_requests/create").Param(co).Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(fr); err != nil {
		return nil, err
	}
	return fr, nil
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
