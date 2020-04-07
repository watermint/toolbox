package sv_usage

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_usage"
)

type Usage interface {
	Resolve() (usage *mo_usage.Usage, err error)
}

func New(ctx dbx_context.Context) Usage {
	return &usageImpl{
		ctx: ctx,
	}
}

type usageImpl struct {
	ctx dbx_context.Context
}

func (z *usageImpl) Resolve() (usage *mo_usage.Usage, err error) {
	res, err := z.ctx.Post("users/get_space_usage").Call()
	if err != nil {
		return nil, err
	}
	usage = &mo_usage.Usage{}
	if err = res.Model(usage); err != nil {
		return nil, err
	}
	return usage, nil
}
