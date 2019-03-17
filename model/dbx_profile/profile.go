package dbx_profile

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type Profile struct {
	Email        string          `json:"email"`
	AccountId    string          `json:"account_id"`
	TeamMemberId string          `json:"team_member_id"`
	Profile      json.RawMessage `json:"profile"`
}

type Member struct {
	Profile *Profile `json:"profile"`
	Role    string   `json:"role"`
}

func ParseMember(p gjson.Result) (member *Member, err error) {
	profile, err := ParseProfile(p.Get("profile"))
	if err != nil {
		return nil, err
	}
	role := p.Get("role." + dbx_api.ResJsonDotTag).String()

	member = &Member{
		Profile: profile,
		Role:    role,
	}
	return member, nil
}

func ParseProfile(p gjson.Result) (profile *Profile, err error) {
	email := p.Get("email")
	if !email.Exists() {
		err = errors.New("required field `email` not found")
		return
	}

	profile = &Profile{
		Email:        email.String(),
		AccountId:    p.Get("account_id").String(),
		TeamMemberId: p.Get("team_member_id").String(),
		Profile:      json.RawMessage(p.Raw),
	}
	return profile, nil
}

func AuthenticatedAdmin(c *dbx_api.DbxContext) (admin *Profile, err error) {
	req := dbx_rpc.RpcRequest{
		Endpoint: "team/token/get_authenticated_admin",
	}
	res, err := req.Call(c)
	if err != nil {
		return nil, err
	}
	return ParseProfile(gjson.Get(res.Body, "admin_profile"))
}

func CurrentAccount(c *dbx_api.DbxContext) (account *Profile, err error) {
	req := dbx_rpc.RpcRequest{
		Endpoint: "users/get_current_account",
	}
	res, err := req.Call(c)
	if err != nil {
		return nil, err
	}
	return ParseProfile(gjson.Parse(res.Body))
}
