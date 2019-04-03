package sv_member_quota

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_member_quota"
)

type Quota interface {
	Resolve(teamMemberId string) (quota *mo_member_quota.Quota, err error)
	Update(quota *mo_member_quota.Quota) (updated *mo_member_quota.Quota, err error)
	Remove(teamMemberId string) (err error)
}

func NewQuota(ctx api_context.Context) Quota {
	return &quotaImpl{
		ctx: ctx,
	}
}

type quotaImpl struct {
	ctx api_context.Context
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

	_, err = z.ctx.Request("team/member_space_limits/remove_custom_quota").Param(p).Call()
	return err
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

	quota = &mo_member_quota.Quota{}
	res, err := z.ctx.Request("team/member_space_limits/get_custom_quota").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.ModelArrayFirst(quota); err != nil {
		return nil, err
	}
	return quota, nil
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

	quota = &mo_member_quota.Quota{}
	res, err := z.ctx.Request("team/member_space_limits/set_custom_quota").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err = res.ModelArrayFirst(quota); err != nil {
		return nil, err
	}
	return quota, nil
}
