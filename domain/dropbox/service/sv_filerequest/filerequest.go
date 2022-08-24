package sv_filerequest

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type FileRequest interface {
	List() (requests []*mo_filerequest.FileRequest, err error)
	Create(title string, destination mo_path.DropboxPath, opts ...CreateOpt) (req *mo_filerequest.FileRequest, err error)
	Delete(ids ...string) (requests []*mo_filerequest.FileRequest, err error)
	DeleteAllClosed() (requests []*mo_filerequest.FileRequest, err error)
	Update(fr *mo_filerequest.FileRequest) (req *mo_filerequest.FileRequest, err error)
}

type CreateOpt func(opts *CreateOpts) *CreateOpts

type CreateOpts struct {
	Open                     bool
	Deadline                 string
	DeadlineAllowLateUploads string
}
type Deadline struct {
	Deadline         string `json:"deadline,omitempty"`
	AllowLateUploads string `json:"allow_late_uploads,omitempty"`
}

func OptDeadline(deadline string) CreateOpt {
	return func(opts *CreateOpts) *CreateOpts {
		opts.Deadline = deadline
		return opts
	}
}
func OptAllowLateUploads(tag string) CreateOpt {
	return func(opts *CreateOpts) *CreateOpts {
		opts.DeadlineAllowLateUploads = tag
		return opts
	}
}

func New(ctx dbx_client.Client) FileRequest {
	return &fileRequestImpl{
		ctx: ctx,
	}
}

type fileRequestImpl struct {
	ctx dbx_client.Client
}

func (z *fileRequestImpl) Update(fr *mo_filerequest.FileRequest) (req *mo_filerequest.FileRequest, err error) {
	var deadline *Deadline
	if fr.Deadline != "" {
		deadline = &Deadline{
			Deadline:         fr.Deadline,
			AllowLateUploads: fr.DeadlineAllowLateUploads,
		}
	}

	co := struct {
		Id          string    `json:"id"`
		Title       string    `json:"title"`
		Destination string    `json:"destination"`
		Deadline    *Deadline `json:"deadline,omitempty"`
		Open        bool      `json:"open"`
	}{
		Id:          fr.Id,
		Title:       fr.Title,
		Destination: fr.Destination,
		Deadline:    deadline,
		Open:        fr.IsOpen,
	}
	res := z.ctx.Post("file_requests/update", api_request.Param(co))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	req = &mo_filerequest.FileRequest{}
	err = res.Success().Json().Model(req)
	return
}

func (z *fileRequestImpl) Delete(ids ...string) (requests []*mo_filerequest.FileRequest, err error) {
	p := struct {
		Ids []string `json:"ids"`
	}{
		Ids: ids,
	}

	res := z.ctx.List("file_requests/delete", api_request.Param(p)).Call(
		dbx_list.Continue("file_requests/list/continue"),
		dbx_list.ResultTag("file_requests"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			fr := &mo_filerequest.FileRequest{}
			if err := entry.Model(fr); err != nil {
				return err
			}
			requests = append(requests, fr)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return requests, nil
}

func (z *fileRequestImpl) DeleteAllClosed() (requests []*mo_filerequest.FileRequest, err error) {
	requests = make([]*mo_filerequest.FileRequest, 0)
	res := z.ctx.List("file_requests/delete_all_closed").Call(
		dbx_list.Continue("file_requests/list/continue"),
		dbx_list.ResultTag("file_requests"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			fr := &mo_filerequest.FileRequest{}
			if err := entry.Model(fr); err != nil {
				return err
			}
			requests = append(requests, fr)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return requests, nil
}

func (z *fileRequestImpl) Create(title string, destination mo_path.DropboxPath, opts ...CreateOpt) (req *mo_filerequest.FileRequest, err error) {
	uo := &CreateOpts{}
	for _, opt := range opts {
		opt(uo)
	}
	var deadline *Deadline
	if uo.Deadline != "" {
		deadline = &Deadline{
			Deadline:         uo.Deadline,
			AllowLateUploads: uo.DeadlineAllowLateUploads,
		}
	}
	co := struct {
		Title       string    `json:"title"`
		Destination string    `json:"destination"`
		Deadline    *Deadline `json:"deadline,omitempty"`
	}{
		Title:       title,
		Destination: destination.Path(),
		Deadline:    deadline,
	}
	res := z.ctx.Post("file_requests/create", api_request.Param(co))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	req = &mo_filerequest.FileRequest{}
	err = res.Success().Json().Model(req)
	return
}

func (z *fileRequestImpl) List() (requests []*mo_filerequest.FileRequest, err error) {
	requests = make([]*mo_filerequest.FileRequest, 0)

	res := z.ctx.List("file_requests/list_v2").Call(
		dbx_list.Continue("file_requests/list/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("file_requests"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			fr := &mo_filerequest.FileRequest{}
			if err := entry.Model(fr); err != nil {
				return err
			}
			requests = append(requests, fr)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return requests, nil
}
