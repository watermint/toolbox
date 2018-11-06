package dbx_member

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type UpdateMember struct {
	NewEmail                 string `json:"new_email,omitempty"`
	NewExternalid            string `json:"new_externalid,omitempty"`
	NewGivenName             string `json:"new_given_name,omitempty"`
	NewSurname               string `json:"new_surname,omitempty"`
	NewPersistentId          string `json:"new_persistent_id,omitempty"`
	NewIsDirectoryRestricted bool   `json:"new_is_directory_restricted,omitempty"`
}

type MemberUpdate struct {
	OnError   func(annotation dbx_api.ErrorAnnotation) bool
	OnSuccess func(m *dbx_profile.Member) bool
}

func (z *MemberUpdate) Update(c *dbx_api.Context, teamMemberId string, m *UpdateMember) bool {
	type Selector struct {
		Tag          string `json:".tag"`
		TeamMemberId string `json:"team_member_id"`
	}
	type Arg struct {
		User                     Selector `json:"user"`
		NewEmail                 string   `json:"new_email,omitempty"`
		NewExternalid            string   `json:"new_externalid,omitempty"`
		NewGivenName             string   `json:"new_given_name,omitempty"`
		NewSurname               string   `json:"new_surname,omitempty"`
		NewPersistentId          string   `json:"new_persistent_id,omitempty"`
		NewIsDirectoryRestricted bool     `json:"new_is_directory_restricted,omitempty"`
	}

	a := Arg{
		User: Selector{
			Tag:          "team_member_id",
			TeamMemberId: teamMemberId,
		},
		NewEmail:                 m.NewEmail,
		NewExternalid:            m.NewExternalid,
		NewGivenName:             m.NewGivenName,
		NewSurname:               m.NewSurname,
		NewPersistentId:          m.NewPersistentId,
		NewIsDirectoryRestricted: m.NewIsDirectoryRestricted,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/members/set_profile",
		Param:    a,
	}
	res, ea, _ := req.Call(c)
	if ea.IsFailure() {
		if z.OnError != nil {
			return z.OnError(ea)
		}
		return false
	}

	um, ea, _ := dbx_profile.ParseMember(gjson.Parse(res.Body))
	if ea.IsFailure() {
		if z.OnError != nil {
			return z.OnError(ea)
		}
		return false
	}
	if z.OnSuccess != nil {
		return z.OnSuccess(um)
	}
	return true
}
