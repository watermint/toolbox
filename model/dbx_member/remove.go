package dbx_member

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type MemberRemove struct {
	OnError   func(annotation dbx_api.ErrorAnnotation) bool
	OnSuccess func(email string) bool
	OnFailure func(email string, reason dbx_api.ApiError) bool
}

// Detach user. Convert user account from Dropbox Business to
// Dropbox Basic. Call /members/remove with `keep_account=true`
func (z *MemberRemove) Detach(c *dbx_api.Context, email string) {
	z.Remove(c, email, false, true)
}

func (z *MemberRemove) Remove(c *dbx_api.Context,
	email string,
	wipeData bool,
	keepAccount bool,
) bool {
	type Selector struct {
		Tag   string `json:".tag"`
		Email string `json:"email"`
	}
	type Param struct {
		User        Selector `json:"user"`
		WipeData    bool     `json:"wipe_data"`
		KeepAccount bool     `json:"keep_account"`
	}

	arg := Param{
		User: Selector{
			Tag:   "email",
			Email: email,
		},
		WipeData:    wipeData,
		KeepAccount: keepAccount,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/members/remove",
		Param:    arg,
	}
	res, ea, _ := req.Call(c)
	if ea.IsFailure() {
		if z.OnError != nil {
			return z.OnError(ea)
		}
		return false
	}
	as := dbx_rpc.AsyncStatus{
		Endpoint: "team/members/remove/job_status/get",
		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			if annotation.ErrorType == dbx_api.ErrorEndpointSpecific {
				switch e := annotation.Error.(type) {
				case dbx_api.ApiError:
					return z.OnFailure(email, e)

				default:
					return z.OnError(annotation)
				}
			}
			return z.OnError(annotation)
		},
		OnComplete: func(complete gjson.Result) bool {
			return z.OnSuccess(email)
		},
	}
	return as.Poll(c, res)
}
