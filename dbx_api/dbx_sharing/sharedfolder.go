package dbx_sharing

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type Member struct {
	AccessType  string          `json:"access_type"`
	Permissions json.RawMessage `json:"permissions,omitempty"`
	IsInherited bool            `json:"is_inherited,omitempty"`
}

type User struct {
	Member
	UserAccountId string `json:"user_account_id,omitempty"`
	Email         string `json:"email"`
	DisplayName   string `json:"display_name,omitempty"`
	SameTeam      bool   `json:"same_team,omitempty"`
	TeamMemberId  string `json:"team_member_id,omitempty"`
}

type Group struct {
	Member
	GroupName           string `json:"group_name,omitempty"`
	GroupId             string `json:"group_id"`
	GroupManagementType string `json:"group_management_type,omitempty"`
	GroupType           string `json:"group_type,omitempty"`
	IsMember            bool   `json:"is_member,omitempty"`
	IsOwner             bool   `json:"is_owner,omitempty"`
	SameTeam            bool   `json:"same_team,omitempty"`
	MemberCount         int    `json:"member_count,omitempty"`
}

type Invitee struct {
	Member
	Email string `json:"email"`
}

type SharedFolderMembers struct {
	AsMemberId string
	AsAdminId  string
	OnError    func(annotation dbx_api.ErrorAnnotation) bool
	OnUser     func(user *User) bool
	OnGroup    func(group *Group) bool
	OnInvitee  func(invitee *Invitee) bool
}

func (s *SharedFolderMembers) List(c *dbx_api.Context, sharedFolderId string) bool {
	list := dbx_rpc.RpcList{
		AsAdminId:            s.AsAdminId,
		AsMemberId:           s.AsMemberId,
		EndpointList:         "sharing/list_folder_members",
		EndpointListContinue: "sharing/list_folder_members/continue",
		UseHasMore:           false,
		OnError:              s.OnError,
		OnResponse: func(body string) bool {
			//users := gjson.Get(body, "users")
			//if s.OnUser != nil && users.Exists() && users.IsArray() {
			//	for _, u := range users.Array() {
			//		if !s.OnUser(u) {
			//			return false
			//		}
			//	}
			//}
			seelog.Info(body)
			return true
		},
	}
	type ListParam struct {
		SharedFolderId string `json:"shared_folder_id"`
	}
	lp := &ListParam{
		SharedFolderId: sharedFolderId,
	}

	return list.List(c, lp)
}
