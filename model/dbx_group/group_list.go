package dbx_group

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type GroupList struct {
	OnError func(err error) bool
	OnEntry func(group *Group) bool
}

func (a *GroupList) List(c *dbx_api.DbxContext) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/groups/list",
		EndpointListContinue: "team/groups/list/continue",
		UseHasMore:           true,
		ResultTag:            "groups",
		OnError:              a.OnError,
		OnEntry: func(r gjson.Result) bool {
			g := &Group{}
			err := c.ParseModel(g, r)
			if err != nil {
				return a.OnError(err)
			}
			return a.OnEntry(g)
		},
	}

	return list.List(c, nil)
}
