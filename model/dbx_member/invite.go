package dbx_member

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

const (
	AdminTierTeamAdmin           = "team_admin"
	AdminTierUserManagementAdmin = "user_management_admin"
	AdminTierSupportAdmin        = "support_admin"
	AdminTierMemberOnly          = "member_only"
)

type FailureReport struct {
	Email  string `json:"email,omitempty"`
	Reason string `json:"reason,omitempty"`
}
type InviteReport struct {
	Result  string              `json:"result"`
	Success *dbx_profile.Member `json:"success,omitempty"`
	Failure *FailureReport      `json:"failure,omitempty"`
}

type InviteMember struct {
	MemberEmail           string `json:"member_email"`
	MemberGivenName       string `json:"member_given_name,omitempty"`
	MemberSurname         string `json:"member_surname,omitempty"`
	MemberExternalId      string `json:"member_external_id,omitempty"`
	MemberPersistentId    string `json:"member_persistent_id,omitempty"`
	SendWelcomeEmail      bool   `json:"send_welcome_email"`
	Role                  string `json:"role,omitempty"`
	IsDirectoryRestricted bool   `json:"is_directory_restricted,omitempty"`
}

type MembersInvite struct {
	OnError   func(err error) bool
	OnSuccess func(member *dbx_profile.Member) bool
	OnFailure func(email string, reason string) bool
}

func (z *MembersInvite) Invite(c *dbx_api.Context, members []*InviteMember) bool {
	chunkSize := 20
	var batch []*InviteMember
	for len(members) > 0 {
		if len(members) >= chunkSize {
			batch = members[:chunkSize]
			members = members[chunkSize:]
		} else {
			batch = members
			members = make([]*InviteMember, 0)
		}

		if !z.handleInvite(c, batch) {
			return false
		}
	}
	return true
}

func (z *MembersInvite) handleInvite(c *dbx_api.Context, members []*InviteMember) bool {
	type NewMembers struct {
		NewMembers []*InviteMember `json:"new_members"`
		ForceAsync bool            `json:"force_async"`
	}

	arg := NewMembers{
		NewMembers: members,
		ForceAsync: true,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/members/add",
		Param:    arg,
	}
	res, err := req.Call(c)
	if err != nil {
		return z.OnError(err)
	}

	as := dbx_rpc.AsyncStatus{
		Endpoint:   "team/members/add/job_status/get",
		OnError:    z.OnError,
		OnComplete: z.handleComplete,
	}
	return as.Poll(c, res)
}

func (z *MembersInvite) handleComplete(complete gjson.Result) bool {
	if !complete.IsArray() {
		err := errors.New("unexpected data format: `complete` is not an array")
		return z.OnError(err)
	}
	for _, c := range complete.Array() {
		tag := c.Get(dbx_api.ResJsonDotTag)
		if tag.String() == "success" {
			if !z.handleSuccess(c) {
				return false
			}
		} else {
			if !z.handleFailure(tag.String(), c) {
				return false
			}
		}
	}
	return true
}

func (z *MembersInvite) handleSuccess(complete gjson.Result) bool {
	member, err := dbx_profile.ParseMember(complete)
	if err != nil {
		return z.OnError(err)
	}

	if z.OnSuccess != nil {
		return z.OnSuccess(member)
	}
	return true
}

func (z *MembersInvite) handleFailure(tag string, complete gjson.Result) bool {
	email := complete.Get(tag).String()

	if z.OnFailure != nil {
		return z.OnFailure(email, tag)
	}
	return false
}
