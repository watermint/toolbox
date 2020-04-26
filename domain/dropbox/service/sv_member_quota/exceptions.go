package sv_member_quota

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Exceptions interface {
	Add(teamMemberId string) (err error)
	Remove(teamMemberId string) (err error)
	List() (members []*mo_profile.Profile, err error)
}

func NewExceptions(ctx dbx_context.Context) Exceptions {
	return &exceptionsImpl{
		ctx: ctx,
	}
}

type exceptionsImpl struct {
	ctx dbx_context.Context
}

func (z *exceptionsImpl) Add(teamMemberId string) (err error) {
	type U struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	p := struct {
		Users []*U `json:"users"`
	}{
		Users: []*U{
			{
				Tag:          "team_member_id",
				TeamMemberId: teamMemberId,
			},
		},
	}
	res := z.ctx.Post("team/member_space_limits/excluded_users/add", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}

func (z *exceptionsImpl) Remove(teamMemberId string) (err error) {
	type U struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	p := struct {
		Users []*U `json:"users"`
	}{
		Users: []*U{
			{
				Tag:          "team_member_id",
				TeamMemberId: teamMemberId,
			},
		},
	}
	res := z.ctx.Post("team/member_space_limits/excluded_users/remove", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}

func (z *exceptionsImpl) List() (members []*mo_profile.Profile, err error) {
	members = make([]*mo_profile.Profile, 0)
	res := z.ctx.List("team/member_space_limits/excluded_users/list").Call(
		dbx_list.Continue("team/member_space_limits/excluded_users/list/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("users"),
		dbx_list.OnEntry(func(entry tjson.Json) error {
			p := &mo_profile.Profile{}
			if err := entry.Model(p); err != nil {
				return err
			}
			members = append(members, p)
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	return members, nil
}
