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
	OnError   func(annotation dbx_api.ErrorAnnotation) bool
	OnSuccess func(member *dbx_profile.Member) bool
	OnFailure func(email string, reason string) bool
}

func (m *MembersInvite) Invite(c *dbx_api.Context, members []*InviteMember) bool {
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

		if !m.handleInvite(c, batch) {
			return false
		}
	}
	return true
}

func (m *MembersInvite) handleInvite(c *dbx_api.Context, members []*InviteMember) bool {
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
	res, ea, _ := req.Call(c)
	if ea.IsFailure() {
		if m.OnError != nil {
			return m.OnError(ea)
		}
		return false
	}

	as := dbx_rpc.AsyncStatus{
		Endpoint:   "team/members/add/job_status/get",
		OnError:    m.OnError,
		OnComplete: m.handleComplete,
	}
	return as.Poll(c, res)
}

func (m *MembersInvite) handleComplete(complete gjson.Result) bool {
	if !complete.IsArray() {
		err := errors.New("unexpected data format: `complete` is not an array")
		annotation := dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		}
		if m.OnError != nil {
			return m.OnError(annotation)
		}
		return false
	}
	for _, c := range complete.Array() {
		tag := c.Get(dbx_api.ResJsonDotTag)
		if tag.String() == "success" {
			if !m.handleSuccess(c) {
				return false
			}
		} else {
			if !m.handleFailure(tag.String(), c) {
				return false
			}
		}
	}
	return true
}

func (m *MembersInvite) handleSuccess(complete gjson.Result) bool {
	member, ea, _ := dbx_profile.ParseMember(complete)
	if ea.IsFailure() {
		return m.OnError(ea)
	}

	if m.OnSuccess != nil {
		return m.OnSuccess(member)
	}
	return true
}

func (m *MembersInvite) handleFailure(tag string, complete gjson.Result) bool {
	email := complete.Get(tag).String()

	if m.OnFailure != nil {
		return m.OnFailure(email, tag)
	}
	return false
}
