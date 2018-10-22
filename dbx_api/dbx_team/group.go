package dbx_team

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type Group struct {
	GroupId             string `json:"group_id,omitempty"`
	GroupName           string `json:"group_name,omitempty"`
	GroupManagementType string `json:"group_management_type,omitempty"`
	GroupExternalId     string `json:"group_external_id,omitempty"`
	MemberCount         int64  `json:"member_count,omitempty"`
}

func ParseGroup(g gjson.Result) (group *Group, annotation dbx_api.ErrorAnnotation, err error) {
	groupIdJson := g.Get("group_id")

	if !groupIdJson.Exists() {
		err = errors.New("required field `group_id` not found")
		annotation = dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		}
		return
	}

	group = &Group{
		GroupId:             groupIdJson.String(),
		GroupName:           g.Get("group_name").String(),
		GroupManagementType: g.Get("group_management_type.\\.tag").String(),
		GroupExternalId:     g.Get("group_external_id").String(),
		MemberCount:         g.Get("member_count").Int(),
	}
	return group, dbx_api.ErrorAnnotation{ErrorType: dbx_api.ErrorSuccess}, nil
}

type GroupList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(group *Group) bool
}

func (a *GroupList) List(c *dbx_api.Context) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/groups/list",
		EndpointListContinue: "team/groups/list/continue",
		UseHasMore:           true,
		ResultTag:            "groups",
		OnError:              a.OnError,
		OnEntry: func(r gjson.Result) bool {
			p, ea, _ := ParseGroup(r)
			if ea.IsSuccess() && a.OnEntry != nil {
				return a.OnEntry(p)
			} else {
				if a.OnError != nil {
					a.OnError(ea)
				}
				return false
			}
		},
	}

	return list.List(c, nil)
}

type GroupMember struct {
	GroupId      string            `json:"group_id"`
	GroupName    string            `json:"group_name"`
	TeamMemberId string            `json:"team_member_id"`
	AccessType   string            `json:"access_type"`
	Profile      *dbx_task.Profile `json:"profile"`
}

type GroupMemberList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(r *GroupMember) bool
}

func (a *GroupMemberList) List(c *dbx_api.Context, group *Group) bool {
	type GroupSelector struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}
	type ListParam struct {
		Group GroupSelector `json:"group"`
	}

	lp := ListParam{
		Group: GroupSelector{
			Tag:     "group_id",
			GroupId: group.GroupId,
		},
	}

	list := dbx_rpc.RpcList{
		EndpointList:         "team/groups/members/list",
		EndpointListContinue: "team/groups/members/list/continue",
		UseHasMore:           true,
		ResultTag:            "members",
		OnError:              a.OnError,
		OnEntry: func(r gjson.Result) bool {
			accessType := r.Get("access_type\\.tag").String()
			p, ea, _ := dbx_task.ParseProfile(r.Get("profile"))
			if ea.IsSuccess() && a.OnEntry != nil {
				gm := &GroupMember{
					GroupId:      group.GroupId,
					GroupName:    group.GroupName,
					TeamMemberId: p.TeamMemberId,
					AccessType:   accessType,
					Profile:      p,
				}
				return a.OnEntry(gm)
			} else {
				if a.OnError != nil {
					a.OnError(ea)
				}
				return false
			}
		},
	}

	return list.List(c, lp)
}
