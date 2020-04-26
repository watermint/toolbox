package sv_namespace

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Namespace interface {
	List() (namespaces []*mo_namespace.Namespace, err error)
}

func New(ctx dbx_context.Context) Namespace {
	return &namespaceImpl{
		ctx: ctx,
	}
}

func newTest(ctx dbx_context.Context, limit int) Namespace {
	return &namespaceImpl{
		ctx:   ctx,
		limit: limit,
	}
}

type namespaceImpl struct {
	ctx   dbx_context.Context
	limit int
}

func (z *namespaceImpl) List() (namespaces []*mo_namespace.Namespace, err error) {
	namespaces = make([]*mo_namespace.Namespace, 0)
	p := struct {
		Limit int `json:"limit,omitempty"`
	}{
		Limit: z.limit,
	}

	req := z.ctx.List("team/namespaces/list", api_request.Param(p)).Call(
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("namespaces"),
		dbx_list.OnEntry(func(entry tjson.Json) error {
			n := &mo_namespace.Namespace{}
			if err := entry.Model(n); err != nil {
				return err
			}
			namespaces = append(namespaces, n)
			return nil
		}),
	)
	if err, fail := req.Failure(); fail {
		return nil, err
	}
	return namespaces, nil
}
