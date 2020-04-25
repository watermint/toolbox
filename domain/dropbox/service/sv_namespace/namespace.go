package sv_namespace

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/essentials/format/tjson"
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

	req := z.ctx.List("team/namespaces/list").
		Continue("team/namespaces/list/continue").
		Param(p).
		UseHasMore(true).
		ResultTag("namespaces").
		OnEntry(func(entry tjson.Json) error {
			n := &mo_namespace.Namespace{}
			if _, err := entry.Model(n); err != nil {
				return err
			}
			namespaces = append(namespaces, n)
			return nil
		})
	if err := req.Call(); err != nil {
		return nil, err
	}
	return namespaces, nil
}
