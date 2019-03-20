package dbx_sharing

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

const (
	AccessLevelOwner           = "owner"
	AccessLevelEditor          = "editor"
	AccessLevelViewer          = "viewer"
	AccessLevelViewerNoComment = "viewer_no_comment"
)

type SharedFolderMembers struct {
	AsMemberId string
	AsAdminId  string
	OnError    func(err error) bool
	OnUser     func(user *MembershipUser) bool
	OnGroup    func(group *MembershipGroup) bool
	OnInvitee  func(invitee *MembershipInvitee) bool
}

func (z *SharedFolderMembers) List(c *dbx_api.DbxContext, sharedFolderId string) bool {
	list := dbx_rpc.RpcList{
		AsAdminId:            z.AsAdminId,
		AsMemberId:           z.AsMemberId,
		EndpointList:         "sharing/list_folder_members",
		EndpointListContinue: "sharing/list_folder_members/continue",
		UseHasMore:           false,
		OnError:              z.OnError,
		OnResponseBody: func(body string) bool {
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
	type ListParam struct {
		SharedFolderId string `json:"shared_folder_id"`
	}
	lp := &ListParam{
		SharedFolderId: sharedFolderId,
	}

	return list.List(c, lp)
}

type AddMembers struct {
	Context       *dbx_api.DbxContext
	Quiet         bool
	CustomMessage string
	AsMemberId    string
	AsAdminId     string
}

// If you want to add members in team folder, please specify `team_folder_id`
// That is equals to `shared_folder_id`.
func (z *AddMembers) Add(sharedFolderId string, dropboxIds []string, accessLevel string) error {
	switch accessLevel {
	case AccessLevelOwner, AccessLevelEditor, AccessLevelViewer, AccessLevelViewerNoComment:
		z.Context.Log().Debug("adding member", zap.String("sharedFolderId", sharedFolderId), zap.Strings("dropboxIds", dropboxIds), zap.String("access", accessLevel))

	default:
		z.Context.Log().Error("invalid access level", zap.String("access", accessLevel))
		return errors.New("invalid access level")
	}

	type M struct {
		Tag       string `json:".tag"`
		DropboxId string `json:"dropbox_id"`
	}
	type A struct {
		Member      M      `json:"member"`
		AccessLevel string `json:"access_level"`
	}
	members := make([]*A, 0)
	for _, gid := range dropboxIds {
		members = append(members,
			&A{
				AccessLevel: accessLevel,
				Member: M{
					Tag:       "dropbox_id",
					DropboxId: gid,
				},
			},
		)
	}

	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		Members        []*A   `json:"members"`
		Quiet          bool   `json:"quiet,omitempty"`
		CustomMessage  string `json:"custom_message,omitempty"`
	}{
		SharedFolderId: sharedFolderId,
		Members:        members,
		Quiet:          z.Quiet,
		CustomMessage:  z.CustomMessage,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint:   "sharing/add_folder_member",
		Param:      p,
		AsMemberId: z.AsMemberId,
		AsAdminId:  z.AsAdminId,
	}
	_, err := req.Call(z.Context)
	if err != nil {
		return err
	}
	return nil
}

type RemoveMembers struct {
	Context    *dbx_api.DbxContext
	LeaveACopy bool
	AsMemberId string
	AsAdminId  string
}

func (z *RemoveMembers) Remove(sharedFolderId string, dropboxId string) error {
	type M struct {
		Tag       string `json:".tag"`
		DropboxId string `json:"dropbox_id"`
	}
	p := struct {
		SharedFolderId string `json:"shared_folder_id"`
		Member         M      `json:"member"`
		LeaveACopy     bool   `json:"leave_a_copy"`
	}{
		SharedFolderId: sharedFolderId,
		Member: M{
			Tag:       "dropbox_id",
			DropboxId: dropboxId,
		},
		LeaveACopy: z.LeaveACopy,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint:   "sharing/remove_folder_member",
		Param:      p,
		AsMemberId: z.AsMemberId,
		AsAdminId:  z.AsAdminId,
	}
	_, err := req.Call(z.Context)
	if err != nil {
		return err
	}
	return nil
}
