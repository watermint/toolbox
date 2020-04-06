package sv_member_quota

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/infra/api/api_list"
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
	_, err = z.ctx.Rpc("team/member_space_limits/excluded_users/add").Param(p).Call()
	return err
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
	_, err = z.ctx.Rpc("team/member_space_limits/excluded_users/remove").Param(p).Call()
	return err
}

func (z *exceptionsImpl) List() (members []*mo_profile.Profile, err error) {
	members = make([]*mo_profile.Profile, 0)
	err = z.ctx.List("team/member_space_limits/excluded_users/list").
		Continue("team/member_space_limits/excluded_users/list/continue").
		UseHasMore(true).
		ResultTag("users").
		OnEntry(func(entry api_list.ListEntry) error {
			p := &mo_profile.Profile{}
			if err := entry.Model(p); err != nil {
				return err
			}
			members = append(members, p)
			return nil
		}).
		Call()
	if err != nil {
		return nil, err
	}
	return members, nil
}
