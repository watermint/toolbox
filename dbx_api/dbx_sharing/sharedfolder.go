package dbx_sharing

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type SharedFolderMembers struct {
	AsMemberId string
	AsAdminId  string
	OnError    func(annotation dbx_api.ErrorAnnotation) bool
	OnUser     func(user *MembershipUser) bool
	OnGroup    func(group *MembershipGroup) bool
	OnInvitee  func(invitee *MembershipInvitee) bool
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
			users := gjson.Get(body, "users")
			if s.OnUser != nil && users.Exists() && users.IsArray() {
				for _, u := range users.Array() {
					user := ParseMembershipUser(u, c.Log())
					if user == nil {
						continue
					}
					if !s.OnUser(user) {
						return false
					}
				}
			}
			groups := gjson.Get(body, "groups")
			if s.OnGroup != nil && groups.Exists() && groups.IsArray() {
				for _, g := range groups.Array() {
					group := ParseMembershipGroup(g, c.Log())
					if group == nil {
						continue
					}
					if !s.OnGroup(group) {
						return false
					}
				}
			}
			invitees := gjson.Get(body, "invitees")
			if s.OnInvitee != nil && invitees.Exists() && invitees.IsArray() {
				for _, v := range invitees.Array() {
					invitee := ParseMembershipInvitee(v, c.Log())
					if invitee == nil {
						continue
					}
					if !s.OnInvitee(invitee) {
						return false
					}
				}
			}
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
