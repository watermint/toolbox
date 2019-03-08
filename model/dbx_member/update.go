package dbx_member

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type UpdateMember struct {
	NewExternalid            string `json:"new_externalid,omitempty"`
	NewGivenName             string `json:"new_given_name,omitempty"`
	NewSurname               string `json:"new_surname,omitempty"`
	NewPersistentId          string `json:"new_persistent_id,omitempty"`
	NewIsDirectoryRestricted bool   `json:"new_is_directory_restricted,omitempty"`
	NewEmail                 string `json:"new_email"`
}

type MemberUpdate struct {
	OnError   func(err error) bool
	OnSuccess func(m *dbx_profile.Member) bool
}

func (z *MemberUpdate) Update(c *dbx_api.Context, email string, m *UpdateMember) bool {
	type Selector struct {
		Tag   string `json:".tag"`
		Email string `json:"email"`
	}
	type Arg struct {
		User                     Selector `json:"user"`
		NewExternalid            string   `json:"new_externalid,omitempty"`
		NewGivenName             string   `json:"new_given_name,omitempty"`
		NewSurname               string   `json:"new_surname,omitempty"`
		NewPersistentId          string   `json:"new_persistent_id,omitempty"`
		NewIsDirectoryRestricted bool     `json:"new_is_directory_restricted,omitempty"`
	}

	a := Arg{
		User: Selector{
			Tag:   "email",
			Email: email,
		},
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
	res, err := req.Call(c)
	if err != nil {
		return z.OnError(err)
	}

	um, err := dbx_profile.ParseMember(gjson.Parse(res.Body))
	if err != nil {
		return z.OnError(err)
	}
	if z.OnSuccess != nil {
		return z.OnSuccess(um)
	}
	return true
}
