package sv_group_member

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
)

type GroupMember interface {
	Add(teamMemberIds []string) error
	Delete(teamMemberIds []string) error
}

type groupMemberImpl struct {
	dc      api_context.Context
	groupId string
}

func (z *groupMemberImpl) Add(teamMemberIds []string) error {
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

	a := z.dc.Async("team/groups/members/add").
		Status("team/groups/job_status/get").
		Param(p)

	if _, err := a.Call(); err != nil {
		return err
	}
	return nil
}

func (z *groupMemberImpl) Delete(teamMemberId []string) error {
	panic("implement me")
}
