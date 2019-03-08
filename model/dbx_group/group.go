package dbx_group

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

const (
	ManagementTypeUser    = "user_managed"
	ManagementTypeCompany = "company_managed"
	ManagementTypeSystem  = "system_managed"
)

type Group struct {
	Raw                 json.RawMessage `json:"-"`
	GroupId             string          `path:"group_id" json:"group_id"`
	GroupName           string          `path:"group_name" json:"group_name"`
	GroupManagementType string          `path:"group_management_type.\\.tag" json:"group_management_type"`
	GroupExternalId     string          `path:"group_external_id" json:"group_external_id,omitempty"`
	MemberCount         int64           `path:"member_count" json:"member_count,omitempty"`
}

type GroupMember struct {
	GroupId      string               `json:"group_id"`
	GroupName    string               `json:"group_name"`
	TeamMemberId string               `json:"team_member_id"`
	AccessType   string               `json:"access_type"`
	Profile      *dbx_profile.Profile `json:"profile"`
}

type GroupMemberList struct {
	OnError func(err error) bool
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
			p, err := dbx_profile.ParseProfile(r.Get("profile"))
			if err != nil {
				a.OnError(err)
				return false
			}

			gm := &GroupMember{
				GroupId:      group.GroupId,
				GroupName:    group.GroupName,
				TeamMemberId: p.TeamMemberId,
				AccessType:   accessType,
				Profile:      p,
			}
			return a.OnEntry(gm)
		},
	}

	return list.List(c, lp)
}
