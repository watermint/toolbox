package dbx_task

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type Profile struct {
	Email        string          `json:"email"`
	AccountId    string          `json:"account_id"`
	TeamMemberId string          `json:"team_member_id"`
	Profile      json.RawMessage `json:"profile"`
}

func ParseProfile(p gjson.Result) (profile *Profile, annotation dbx_api.ErrorAnnotation, err error) {
	email := p.Get("email")
	if !email.Exists() {
		err = errors.New("required field `email` not found")
		annotation = dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		}
		return
	}

	profile = &Profile{
		Email:        email.String(),
		AccountId:    p.Get("account_id").String(),
		TeamMemberId: p.Get("team_member_id").String(),
		Profile:      json.RawMessage(p.Raw),
	}
	return profile, dbx_api.ErrorAnnotation{ErrorType: dbx_api.ErrorSuccess}, nil
}

func AuthenticatedAdmin(c *dbx_api.Context) (admin *Profile, annotation dbx_api.ErrorAnnotation, err error) {
	req := dbx_rpc.RpcRequest{
		Endpoint: "team/token/get_authenticated_admin",
	}
	res, annotation, err := req.Call(c)
	if annotation.IsSuccess() {
		return ParseProfile(gjson.Get(res.Body, "admin_profile"))
	} else {
		return
	}
}

func CurrentAccount(c *dbx_api.Context) (account *Profile, annotation dbx_api.ErrorAnnotation, err error) {
	req := dbx_rpc.RpcRequest{
		Endpoint: "users/get_current_account",
	}
	res, annotation, err := req.Call(c)
	if annotation.IsSuccess() {
		return ParseProfile(gjson.Get(res.Body, "i"))
	} else {
		return
	}
}
