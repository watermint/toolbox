package sv_label

import (
	"errors"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/model/mo_label"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
)

var (
	ErrorLabelNotFound = errors.New("label not found")
)

type Label interface {
	Add(name string, opts ...Opt) (label *mo_label.Label, err error)
	Update(id string, opts ...Opt) (label *mo_label.Label, err error)
	Remove(id string) error
	Resolve(id string) (label *mo_label.Label, err error)
	ResolveByName(name string) (label *mo_label.Label, err error)
	// Return labels when all labels found.
	ResolveByNames(names []string) (labels []*mo_label.Label, missing []string, err error)
	List() (labels []*mo_label.Label, err error)
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

func (z labelImpl) Resolve(id string) (label *mo_label.Label, err error) {
	labels, err := z.List()
	if err != nil {
		return nil, err
	}
	for _, label := range labels {
		if label.Id == id {
			return label, nil
		}
	}
	return nil, ErrorLabelNotFound
}

func (z labelImpl) ResolveByNames(names []string) (labels []*mo_label.Label, missing []string, err error) {
	l := z.ctx.Log()
	labels = make([]*mo_label.Label, 0)
	missing = make([]string, 0)
	var missingErr error
	for _, name := range names {
		l.Debug("Resolve", esl.String("name", name))
		label, err := z.ResolveByName(name)
		if err != nil {
			missingErr = err
			missing = append(missing, name)
		} else {
			labels = append(labels, label)
		}
	}
	if missingErr == nil {
		return labels, missing, nil
	} else {
		return nil, missing, missingErr
	}
}

func (z labelImpl) ResolveByName(name string) (label *mo_label.Label, err error) {
	labels, err := z.List()
	if err != nil {
		return nil, err
	}
	for _, x := range labels {
		if x.Name == name {
			return x, nil
		}
	}
	return nil, ErrorLabelNotFound
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
	if _, found := j.Find("labels"); !found {
		return labels, nil
	}
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
