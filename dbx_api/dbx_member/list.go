package dbx_member

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type MembersList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(member *dbx_profile.Member) bool
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
			m, ea, _ := dbx_profile.ParseMember(member)
			if ea.IsSuccess() {
				return a.OnEntry(m)
			} else {
				if a.OnError != nil {
					return a.OnError(ea)
				}
				return false
			}
		},
	}

	return list.List(c, lp)
}

func (a *MembersList) ListAsMap(c *dbx_api.Context, includeRemoved bool) map[string]*dbx_profile.Member {
	members := make(map[string]*dbx_profile.Member)
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
			m, ea, _ := dbx_profile.ParseMember(member)
			if ea.IsSuccess() {
				members[m.Profile.Email] = m
				return true
			} else {
				if a.OnError != nil {
					return a.OnError(ea)
				}
				return false
			}
		},
	}

	if !list.List(c, lp) {
		return nil
	}
	return members
}
