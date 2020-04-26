package sv_member_quota

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member_quota"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Quota interface {
	Resolve(teamMemberId string) (quota *mo_member_quota.Quota, err error)
	Update(quota *mo_member_quota.Quota) (updated *mo_member_quota.Quota, err error)
	Remove(teamMemberId string) (err error)
}

func NewQuota(ctx dbx_context.Context) Quota {
	return &quotaImpl{
		ctx: ctx,
	}
}

type quotaImpl struct {
	ctx dbx_context.Context
}

func (z *quotaImpl) Remove(teamMemberId string) (err error) {
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

	res := z.ctx.Post("team/member_space_limits/remove_custom_quota", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}

func (z *quotaImpl) Resolve(teamMemberId string) (quota *mo_member_quota.Quota, err error) {
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

	res := z.ctx.Post("team/member_space_limits/get_custom_quota", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	quota = &mo_member_quota.Quota{}
	err = res.Success().Json().FindModel(tjson.PathArrayFirst, quota)
	return
}

func (z *quotaImpl) Update(quota *mo_member_quota.Quota) (updated *mo_member_quota.Quota, err error) {
	type U struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	type Q struct {
		User    U   `json:"user"`
		QuotaGb int `json:"quota_gb"`
	}
	p := struct {
		UsersAndQuotas []Q `json:"users_and_quotas"`
	}{
		UsersAndQuotas: []Q{
			{
				User: U{
					Tag:          "team_member_id",
					TeamMemberId: quota.TeamMemberId,
				},
				QuotaGb: quota.Quota,
			},
		},
	}

	res := z.ctx.Post("team/member_space_limits/set_custom_quota", api_request.Param(p))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	quota = &mo_member_quota.Quota{}
	err = res.Success().Json().FindModel(tjson.PathArrayFirst, quota)
	return
}
