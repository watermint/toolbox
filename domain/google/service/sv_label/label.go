package sv_label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/model/mo_label"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Label interface {
	List() (labels []*mo_label.Label, err error)
}

func New(ctx goog_context.Context, userId string) Label {
	return &labelImpl{
		ctx:    ctx,
		userId: userId,
	}
}

type labelImpl struct {
	ctx    goog_context.Context
	userId string
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
