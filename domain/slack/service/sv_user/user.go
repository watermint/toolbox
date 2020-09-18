package sv_user

import (
	"github.com/watermint/toolbox/domain/slack/api/work_context"
	"github.com/watermint/toolbox/domain/slack/api/work_pagination"
	"github.com/watermint/toolbox/domain/slack/model/mo_user"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type User interface {
	Resolve(id string) (u *mo_user.User, err error)
}

func New(ctx work_context.Context) User {
	return &userImpl{
		ctx: ctx,
	}
}

func NewCached(user User) User {
	return &cachedImpl{
		user:       user,
		cache:      make(map[string]*mo_user.User),
		cacheError: make(map[string]error),
	}
}

type cachedImpl struct {
	user       User
	cache      map[string]*mo_user.User
	cacheError map[string]error
}

func (z *cachedImpl) Resolve(id string) (u *mo_user.User, err error) {
	if u, ok := z.cache[id]; ok {
		return u, nil
	}

	if e, ok := z.cacheError[id]; ok {
		return nil, e
	}

	u, err = z.user.Resolve(id)
	if err != nil {
		z.cacheError[id] = err
		return nil, err
	}
	z.cache[id] = u
	return u, nil
}

type userImpl struct {
	ctx work_context.Context
}

func (z userImpl) Resolve(id string) (u *mo_user.User, err error) {
	pg := work_pagination.New(z.ctx).WithEndpoint("user.info")
	err = pg.OnData("user", func(entry es_json.Json) error {
		u = &mo_user.User{}
		if err := entry.Model(u); err != nil {
			return err
		}
		return nil
	})
	return
}
