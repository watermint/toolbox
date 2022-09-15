package sv_profile

import (
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/domain/github/model/mo_user"
)

type Current interface {
	User() (user *mo_user.User, err error)
}

func New(ctx gh_client.Client) Current {
	return &currentImpl{
		ctx: ctx,
	}
}

type currentImpl struct {
	ctx gh_client.Client
}

func (z *currentImpl) User() (user *mo_user.User, err error) {
	res := z.ctx.Get("user")
	if err, fa := res.Failure(); fa {
		return nil, err
	}
	user = &mo_user.User{}
	err = res.Success().Json().Model(user)
	return
}
