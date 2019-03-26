package sv_group_member

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_list"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
)

type GroupMember interface {
	List() (members []*mo_group_member.Member, err error)
	Add(teamMemberIds []string) (group *mo_group.Group, err error)
	Remove(teamMemberIds []string) (group *mo_group.Group, err error)
}

func New(ctx api_context.Context, group *mo_group.Group) GroupMember {
	return &groupMemberImpl{
		ctx:     ctx,
		groupId: group.GroupId,
	}
}

func NewByGroupId(ctx api_context.Context, groupId string) GroupMember {
	return &groupMemberImpl{
		ctx:     ctx,
		groupId: groupId,
	}
}

type groupMemberImpl struct {
	ctx     api_context.Context
	groupId string
}

func (z *groupMemberImpl) List() (members []*mo_group_member.Member, err error) {
	type GS struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}
	p := struct {
		Group GS  `json:"group"`
		Limit int `json:"limit,omitempty"`
	}{
		Group: GS{
			Tag:     "group_id",
			GroupId: z.groupId,
		},
	}

	members = make([]*mo_group_member.Member, 0)
	err = z.ctx.List("team/groups/members/list").
		Continue("team/groups/members/list/continue").
		Param(p).
		ResultTag("members").
		UseHasMore(true).
		OnEntry(func(entry api_list.ListEntry) error {
			gm := &mo_group_member.Member{}
			if err := entry.Model(gm); err != nil {
				return err
			}
			members = append(members, gm)
			return nil
		}).Call()
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (z *groupMemberImpl) Add(teamMemberIds []string) (group *mo_group.Group, err error) {
	type GS struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}
	type U struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	type M struct {
		User       U      `json:"user"`
		AccessType string `json:"access_type"`
	}

	members := make([]*M, 0)
	for _, m := range teamMemberIds {
		members = append(members, &M{
			AccessType: "member",
			User: U{
				Tag:          "team_member_id",
				TeamMemberId: m,
			},
		})
	}
	p := struct {
		Group         GS   `json:"group"`
		Members       []*M `json:"members"`
		ReturnMembers bool `json:"return_members"`
	}{
		Group: GS{
			Tag:     "group_id",
			GroupId: z.groupId,
		},
		Members:       members,
		ReturnMembers: false,
	}

	group = &mo_group.Group{}
	a := z.ctx.Async("team/groups/members/add").
		Status("team/groups/job_status/get").
		Param(p)
	res, err := a.Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(group); err != nil {
		return nil, err
	}
	return group, nil
}

func (z *groupMemberImpl) Remove(teamMemberIds []string) (group *mo_group.Group, err error) {
	type GS struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}
	type U struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	users := make([]*U, 0)
	for _, m := range teamMemberIds {
		users = append(users, &U{
			Tag:          "team_member_id",
			TeamMemberId: m,
		})
	}
	p := struct {
		Group         GS   `json:"group"`
		Users         []*U `json:"users"`
		ReturnMembers bool `json:"return_members,omitempty"`
	}{
		Group: GS{
			Tag:     "group_id",
			GroupId: z.groupId,
		},
		Users: users,
	}

	group = &mo_group.Group{}
	a := z.ctx.Async("team/groups/members/remove").
		Status("team/groups/job_status/get").
		Param(p)
	res, err := a.Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(group); err != nil {
		return nil, err
	}
	return group, nil
}
