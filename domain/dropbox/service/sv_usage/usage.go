package sv_usage

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_usage"
)

type Usage interface {
	Resolve() (usage *mo_usage.Usage, err error)
}

func New(ctx dbx_client.Client) Usage {
	return &usageImpl{
		ctx: ctx,
	}
}

type usageImpl struct {
	ctx dbx_client.Client
}

func (z *usageImpl) Resolve() (usage *mo_usage.Usage, err error) {
	res := z.ctx.Post("users/get_space_usage")
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	usage = &mo_usage.Usage{}
	err = res.Success().Json().Model(usage)
	return
}
