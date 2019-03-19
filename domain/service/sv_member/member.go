package sv_member

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/model/mo_member"
)

type Member interface {
	Update(member *mo_member.Member) (updated *mo_member.Member, err error)
	List() (members []*mo_member.Member, err error)
	Resolve(teamMemberId string) (member *mo_member.Member, err error)
}

func New(ctx api_context.Context) Member {
	return &memberImpl{
		ctx: ctx,
	}
}

func newTest(ctx api_context.Context) Member {
	return &memberImpl{
		ctx:   ctx,
		limit: 3,
	}
}

type memberImpl struct {
	ctx            api_context.Context
	includeDeleted bool
	limit          int
}

func (z *memberImpl) Update(member *mo_member.Member) (updated *mo_member.Member, err error) {
	type US struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	p := struct {
		User            US     `json:"user"`
		NewEmail        string `json:"new_email,omitempty"`
		NewExternalId   string `json:"new_external_id,omitempty"`
		NewGivenName    string `json:"new_given_name,omitempty"`
		NewSurname      string `json:"new_surname,omitempty"`
		NewPersistentId string `json:"new_persistent_id,omitempty"`
	}{
		User: US{
			Tag:          "team_member_id",
			TeamMemberId: member.TeamMemberId,
		},
		NewEmail:        member.Email,
		NewExternalId:   member.ExternalId,
		NewGivenName:    member.GivenName,
		NewSurname:      member.Surname,
		NewPersistentId: member.PersistentId,
	}
	req := z.ctx.Request("team/members/set_profile").Param(p)
	res, err := req.Call()
	if err != nil {
		return nil, err
	}
	updated = &mo_member.Member{}
	if err = res.Model(updated); err != nil {
		return nil, err
	}
	return updated, nil
}

func (z *memberImpl) Resolve(teamMemberId string) (member *mo_member.Member, err error) {
	type US struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	p := struct {
		Members []US `json:"members"`
	}{
		Members: []US{
			{
				Tag:          "team_member_id",
				TeamMemberId: teamMemberId,
			},
		},
	}
	member = &mo_member.Member{}
	res, err := z.ctx.Request("team/members/get_info").Param(p).Call()
	if err != nil {
		return nil, err
	}
	if err := res.ModelArrayFirst(member); err != nil {
		return nil, err
	}
	return member, nil
}

func (z *memberImpl) List() (members []*mo_member.Member, err error) {
	members = make([]*mo_member.Member, 0)
	p := struct {
		IncludeRemoved bool `json:"include_removed,omitempty"`
		Limit          int  `json:"limit,omitempty"`
	}{
		IncludeRemoved: z.includeDeleted,
		Limit:          z.limit,
	}

	req := z.ctx.List("team/members/list").
		Continue("team/members/list/continue").
		Param(p).
		UseHasMore(true).
		ResultTag("members").
		OnEntry(func(entry api_list.ListEntry) error {
			m := &mo_member.Member{}
			if err := entry.Model(m); err != nil {
				return err
			}
			members = append(members, m)
			return nil
		})
	if err := req.Call(); err != nil {
		return nil, err
	}
	return members, nil
}
