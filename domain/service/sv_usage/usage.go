package sv_usage

import (
	"github.com/watermint/toolbox/domain/model/mo_usage"
	"github.com/watermint/toolbox/infra/api/api_context"
)

type Usage interface {
	Resolve() (usage *mo_usage.Usage, err error)
}

func New(ctx api_context.Context) Usage {
	return &usageImpl{
		ctx: ctx,
	}
}

type usageImpl struct {
	ctx api_context.Context
}

func (z *usageImpl) Resolve() (usage *mo_usage.Usage, err error) {
	res, err := z.ctx.Rpc("users/get_space_usage").Call()
	if err != nil {
		return nil, err
	}
	usage = &mo_usage.Usage{}
	if err = res.Model(usage); err != nil {
		return nil, err
	}
	return usage, nil
}
