package sv_filter

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/model/mo_filter"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Filter interface {
	List() (filters []*mo_filter.Filter, err error)
	Resolve(id string) (filter *mo_filter.Filter, err error)
}

func New(ctx goog_context.Context, userId string) Filter {
	return &filterImpl{
		ctx:    ctx,
		userId: userId,
	}
}

type filterImpl struct {
	ctx    goog_context.Context
	userId string
}

func (z filterImpl) List() (filters []*mo_filter.Filter, err error) {
	res := z.ctx.Get("gmail/v1/users/" + z.userId + "/settings/filters")
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	filters = make([]*mo_filter.Filter, 0)
	err = j.FindArrayEach("filter", func(e es_json.Json) error {
		m := &mo_filter.Filter{}
		if err := e.Model(m); err != nil {
			return err
		}
		filters = append(filters, m)
		return nil
	})
	return filters, err
}

func (z filterImpl) Resolve(id string) (filter *mo_filter.Filter, err error) {
	res := z.ctx.Get("gmail/v1/users/" + z.userId + "/settings/filters/" + id)
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	filter = &mo_filter.Filter{}
	if err := j.Model(filter); err != nil {
		return nil, err
	} else {
		return filter, nil
	}
}
