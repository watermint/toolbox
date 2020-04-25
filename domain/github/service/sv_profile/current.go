package sv_profile

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_user"
)

type Current interface {
	User() (user *mo_user.User, err error)
}

func New(ctx gh_context.Context) Current {
	return &currentImpl{
		ctx: ctx,
	}
}

type currentImpl struct {
	ctx gh_context.Context
}

func (z *currentImpl) User() (user *mo_user.User, err error) {
	res, err := z.ctx.Get("user").Call()
	if err != nil {
		return nil, err
	}
	user = &mo_user.User{}
	if _, err := res.Body().Json().Model(user); err != nil {
		return nil, err
	}
	return user, nil
}
