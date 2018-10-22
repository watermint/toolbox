package dbx_team

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type MembersList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(profile *dbx_task.Profile) bool
}

func (a *MembersList) List(c *dbx_api.Context, includeRemoved bool) bool {
	type ListParam struct {
		IncludeRemoved bool `json:"include_removed"`
	}
	lp := ListParam{
		IncludeRemoved: includeRemoved,
	}

	list := dbx_rpc.RpcList{
		EndpointList:         "team/members/list",
		EndpointListContinue: "team/members/list/continue",
		UseHasMore:           true,
		ResultTag:            "members",
		OnError:              a.OnError,
		OnEntry: func(member gjson.Result) bool {
			p, ea, _ := dbx_task.ParseProfile(member)
			if ea.IsSuccess() {
				return a.OnEntry(p)
			} else {
				if a.OnError != nil {
					a.OnError(ea)
				}
				return false
			}
		},
	}

	return list.List(c, lp)
}
