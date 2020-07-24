package sv_label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/model/mo_label"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewCached(ctx goog_context.Context, userId string) Label {
	return &labelCacheImpl{
		ctx:   ctx,
		label: New(ctx, userId),
	}
}

type labelCacheImpl struct {
	ctx    goog_context.Context
	label  Label
	labels []*mo_label.Label
}

func (z *labelCacheImpl) ResolveByNames(names []string) (labels []*mo_label.Label, missing []string, err error) {
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

func (z *labelCacheImpl) Add(name string, opts ...Opt) (label *mo_label.Label, err error) {
	z.labels = nil
	return z.label.Add(name, opts...)
}

func (z *labelCacheImpl) Update(id string, opts ...Opt) (label *mo_label.Label, err error) {
	z.labels = nil
	return z.label.Update(id, opts...)
}

func (z *labelCacheImpl) Remove(id string) error {
	z.labels = nil
	return z.label.Remove(id)
}

func (z *labelCacheImpl) ResolveByName(name string) (label *mo_label.Label, err error) {
	l := z.ctx.Log()
	if z.labels == nil {
		l.Debug("Retrieve list")
		z.labels, err = z.label.List()
		if err != nil {
			l.Debug("Unable to update cache", esl.Error(err))
			return nil, err
		}
	}

	l.Debug("Resolve with cache", esl.String("name", name))
	for _, label := range z.labels {
		if label.Name == name {
			return label, nil
		}
	}
	return nil, ErrorLabelNotFoundForName
}

func (z *labelCacheImpl) List() (labels []*mo_label.Label, err error) {
	l := z.ctx.Log()
	if z.labels == nil {
		l.Debug("Retrieve list")
		z.labels, err = z.label.List()
		if err != nil {
			return nil, err
		}
	}
	return z.labels, nil
}
