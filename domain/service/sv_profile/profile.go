package sv_profile

import (
	"github.com/watermint/toolbox/domain/model/mo_profile"
	"github.com/watermint/toolbox/infra/api/api_context"
)

type Profile interface {
	Current() (profile *mo_profile.Profile, err error)
}

func NewProfile(ctx api_context.Context) Profile {
	return &profileImpl{
		ctx: ctx,
	}
}

type Team interface {
	Admin() (profile *mo_profile.Profile, err error)
}

func NewTeam(ctx api_context.Context) Team {
	return &teamImpl{
		ctx: ctx,
	}
}

type profileImpl struct {
	ctx api_context.Context
}

func (z *profileImpl) Current() (profile *mo_profile.Profile, err error) {
	profile = &mo_profile.Profile{}
	res, err := z.ctx.Rpc("users/get_current_account").Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

type teamImpl struct {
	ctx api_context.Context
}

func (z *teamImpl) Admin() (profile *mo_profile.Profile, err error) {
	profile = &mo_profile.Profile{}
	res, err := z.ctx.Rpc("team/token/get_authenticated_admin").Call()
	if err != nil {
		return nil, err
	}
	if err = res.ModelWithPath(profile, "admin_profile"); err != nil {
		return nil, err
	}
	return profile, nil
}
