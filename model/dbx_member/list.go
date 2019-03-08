package dbx_member

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type MembersList struct {
	OnError func(err error) bool
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
			m, err := dbx_profile.ParseMember(member)
			if err != nil {
				return a.OnError(err)
			}
			return a.OnEntry(m)
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
			m, err := dbx_profile.ParseMember(member)
			if err != nil {
				return a.OnError(err)
			}
			members[m.Profile.Email] = m
			return true
		},
	}

	if !list.List(c, lp) {
		return nil
	}
	return members
}
