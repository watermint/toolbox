package sv_label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/model/mo_label"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Label interface {
	Add(name string, opts ...Opt) (label *mo_label.Label, err error)
	Update(id string, opts ...Opt) (label *mo_label.Label, err error)
	Remove(id string) error
	List() (labels []*mo_label.Label, err error)
}

type Opts struct {
	LabelListVisibility   string
	MessageListVisibility string
	ColorBackground       string
	ColorText             string
	Name                  string
}
type Opt func(o Opts) Opts

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

func (z Opts) Param() LabelParam {
	p := LabelParam{
		Name:                  z.Name,
		LabelListVisibility:   z.LabelListVisibility,
		MessageListVisibility: z.MessageListVisibility,
	}
	if z.ColorText != "" || z.ColorBackground != "" {
		p.Color = &LabelColorParam{
			BackgroundColor: z.ColorBackground,
			TextColor:       z.ColorText,
		}
	}
	return p
}

func Name(v string) Opt {
	return func(o Opts) Opts {
		o.Name = v
		return o
	}
}

func LabelListVisibility(v string) Opt {
	return func(o Opts) Opts {
		o.LabelListVisibility = v
		return o
	}
}
func MessageListVisibility(v string) Opt {
	return func(o Opts) Opts {
		o.MessageListVisibility = v
		return o
	}
}
func ColorBackground(c string) Opt {
	return func(o Opts) Opts {
		o.ColorBackground = c
		return o
	}
}
func ColorText(c string) Opt {
	return func(o Opts) Opts {
		o.ColorText = c
		return o
	}
}

func New(ctx goog_context.Context, userId string) Label {
	return &labelImpl{
		ctx:    ctx,
		userId: userId,
	}
}

type LabelParam struct {
	LabelListVisibility   string           `json:"labelListVisibility,omitempty"`
	MessageListVisibility string           `json:"messageListVisibility,omitempty"`
	Color                 *LabelColorParam `json:"color,omitempty"`
	Name                  string           `json:"name"`
}

type labelImpl struct {
	ctx    goog_context.Context
	userId string
}

func (z labelImpl) Add(name string, opts ...Opt) (label *mo_label.Label, err error) {
	p := Opts{}.Apply(opts...).Apply(Name(name)).Param()
	res := z.ctx.Post("gmail/v1/users/"+z.userId+"/labels", api_request.Param(&p))
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	label = &mo_label.Label{}
	err = j.Model(label)
	return
}

func (z labelImpl) Update(id string, opts ...Opt) (label *mo_label.Label, err error) {
	p := Opts{}.Apply(opts...).Param()
	res := z.ctx.Put("gmail/v1/users/"+z.userId+"/labels/"+id, api_request.Param(&p))
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	label = &mo_label.Label{}
	err = j.Model(label)
	return
}

func (z labelImpl) Remove(id string) error {
	res := z.ctx.Delete("gmail/v1/users/" + z.userId + "/labels/" + id)
	if err, f := res.Failure(); f {
		return err
	}
	_, err := res.Success().AsJson()
	return err
}

func (z labelImpl) List() (labels []*mo_label.Label, err error) {
	res := z.ctx.Get("gmail/v1/users/" + z.userId + "/labels")
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	labels = make([]*mo_label.Label, 0)
	err = j.FindArrayEach("labels", func(e es_json.Json) error {
		m := &mo_label.Label{}
		if err := e.Model(m); err != nil {
			return err
		}
		labels = append(labels, m)
		return nil
	})
	return labels, err
}
