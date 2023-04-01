package sv_user

import (
	"github.com/watermint/toolbox/domain/figma/api/fg_client"
	"github.com/watermint/toolbox/domain/figma/model/mo_user"
)

type User interface {
	Current() (user *mo_user.User, err error)
}

func New(client fg_client.Client) User {
	return &userImpl{
		client: client,
	}
}

type userImpl struct {
	client fg_client.Client
}

func (z userImpl) Current() (user *mo_user.User, err error) {
	res := z.client.Get("me")
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	user = &mo_user.User{}
	if err = res.Success().Json().Model(user); err != nil {
		return nil, err
	}
	return user, nil
}
