package sv_namespace

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Namespace interface {
	List() (namespaces []*mo_namespace.Namespace, err error)
	ListEach(f func(entry *mo_namespace.Namespace) bool) error
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

func (z *namespaceImpl) ListEach(f func(entry *mo_namespace.Namespace) bool) error {
	p := struct {
		Limit int `json:"limit,omitempty"`
	}{
		Limit: z.limit,
	}

	ErrorBreak := errors.New("break continue")

	req := z.ctx.List("team/namespaces/list", api_request.Param(p)).Call(
		dbx_list.Continue("team/namespaces/list/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("namespaces"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			n := &mo_namespace.Namespace{}
			if err := entry.Model(n); err != nil {
				return err
			}
			if !f(n) {
				return ErrorBreak
			}
			return nil
		}),
	)
	if err, fail := req.Failure(); fail {
		if err == ErrorBreak {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (z *namespaceImpl) List() (namespaces []*mo_namespace.Namespace, err error) {
	namespaces = make([]*mo_namespace.Namespace, 0)
	err = z.ListEach(func(entry *mo_namespace.Namespace) bool {
		namespaces = append(namespaces, entry)
		return true
	})
	return
}
