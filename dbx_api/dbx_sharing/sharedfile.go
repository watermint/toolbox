package dbx_sharing

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type SharedFileMembers struct {
	AsMemberId string
	AsAdminId  string
	OnError    func(annotation dbx_api.ErrorAnnotation) bool
	OnUser     func(user *MembershipUser) bool
	OnGroup    func(group *MembershipGroup) bool
	OnInvitee  func(invitee *MembershipInvitee) bool
}

func (z *SharedFileMembers) List(c *dbx_api.Context, file string) bool {
	type Arg struct {
		File string `json:"file"`
	}
	list := dbx_rpc.RpcList{
		AsAdminId:            z.AsAdminId,
		AsMemberId:           z.AsMemberId,
		EndpointList:         "sharing/list_file_members",
		EndpointListContinue: "sharing/list_file_members/continue",
		UseHasMore:           false,
		OnError:              z.OnError,
		OnResponse: func(body string) bool {
			users := gjson.Get(body, "users")
			if z.OnUser != nil && users.Exists() && users.IsArray() {
				for _, u := range users.Array() {
					user := ParseMembershipUser(u, c.Log())
					if user == nil {
						continue
					}
					if !z.OnUser(user) {
						return false
					}
				}
			}
			groups := gjson.Get(body, "groups")
			if z.OnGroup != nil && groups.Exists() && groups.IsArray() {
				for _, g := range groups.Array() {
					group := ParseMembershipGroup(g, c.Log())
					if group == nil {
						continue
					}
					if !z.OnGroup(group) {
						return false
					}
				}
			}
			invitees := gjson.Get(body, "invitees")
			if z.OnInvitee != nil && invitees.Exists() && invitees.IsArray() {
				for _, v := range invitees.Array() {
					invitee := ParseMembershipInvitee(v, c.Log())
					if invitee == nil {
						continue
					}
					if !z.OnInvitee(invitee) {
						return false
					}
				}
			}

			return true
		},
	}

	arg := &Arg{
		File: file,
	}
	return list.List(c, arg)
}
