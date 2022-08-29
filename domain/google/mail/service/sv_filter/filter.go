package sv_filter

import (
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_filter"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Filter interface {
	List() (filters []*mo_filter.Filter, err error)
	Resolve(id string) (filter *mo_filter.Filter, err error)
	Delete(id string) error
	Add(opts ...Opt) (filter *mo_filter.Filter, err error)
}

type Criteria struct {
	From           string `json:"from,omitempty"`
	To             string `json:"to,omitempty"`
	Subject        string `json:"subject,omitempty"`
	Query          string `json:"query,omitempty"`
	NegatedQuery   string `json:"negatedQuery,omitempty"`
	HasAttachment  *bool  `json:"hasAttachment,omitempty"`
	ExcludeChats   *bool  `json:"excludeChats,omitempty"`
	Size           *int   `json:"size,omitempty"`
	SizeComparison string `json:"sizeComparison,omitempty"`
}

type Action struct {
	AddLabelIds    []string `json:"addLabelIds,omitempty"`
	RemoveLabelIds []string `json:"removeLabelIds,omitempty"`
	Forward        string   `json:"forward,omitempty"`
}

type Opts struct {
	Criteria Criteria `json:"criteria,omitempty"`
	Action   Action   `json:"action,omitempty"`
}

func (z Opts) Apply(opts ...Opt) Opts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

type Opt func(o Opts) Opts

func From(v string) Opt {
	return func(o Opts) Opts {
		o.Criteria.From = v
		return o
	}
}
func To(v string) Opt {
	return func(o Opts) Opts {
		o.Criteria.To = v
		return o
	}
}
func Subject(v string) Opt {
	return func(o Opts) Opts {
		o.Criteria.Subject = v
		return o
	}
}
func Query(v string) Opt {
	return func(o Opts) Opts {
		o.Criteria.Query = v
		return o
	}
}
func NegatedQuery(v string) Opt {
	return func(o Opts) Opts {
		o.Criteria.NegatedQuery = v
		return o
	}
}
func HasAttachment(v bool) Opt {
	return func(o Opts) Opts {
		o.Criteria.HasAttachment = &v
		return o
	}
}
func ExcludeChats(v bool) Opt {
	return func(o Opts) Opts {
		o.Criteria.ExcludeChats = &v
		return o
	}
}
func Size(v int) Opt {
	return func(o Opts) Opts {
		o.Criteria.Size = &v
		return o
	}
}
func SizeComparison(v string) Opt {
	return func(o Opts) Opts {
		o.Criteria.SizeComparison = v
		return o
	}
}
func AddLabelIds(v []string) Opt {
	return func(o Opts) Opts {
		o.Action.AddLabelIds = v
		return o
	}
}
func RemoveLabelIds(v []string) Opt {
	return func(o Opts) Opts {
		o.Action.RemoveLabelIds = v
		return o
	}
}
func Forward(v string) Opt {
	return func(o Opts) Opts {
		o.Action.Forward = v
		return o
	}
}

func New(ctx goog_client.Client, userId string) Filter {
	return &filterImpl{
		ctx:    ctx,
		userId: userId,
	}
}

type filterImpl struct {
	ctx    goog_client.Client
	userId string
}

func (z filterImpl) Delete(id string) error {
	res := z.ctx.Delete("gmail/v1/users/" + z.userId + "/settings/filters/" + id)
	if err, f := res.Failure(); f {
		return err
	}
	return nil
}

func (z filterImpl) Add(opts ...Opt) (filter *mo_filter.Filter, err error) {
	ao := Opts{}.Apply(opts...)
	res := z.ctx.Post("gmail/v1/users/"+z.userId+"/settings/filters", api_request.Param(&ao))
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
	if _, found := j.Find("filter"); !found {
		return filters, nil
	}
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
