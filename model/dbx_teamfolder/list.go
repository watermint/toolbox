package dbx_teamfolder

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type ListTeamFolder struct {
	OnError func(err error) bool
	OnEntry func(teamFolder *TeamFolder) bool
}

func (z *ListTeamFolder) List(c *dbx_api.DbxContext) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/team_folder/list",
		EndpointListContinue: "team/team_folder/list/continue",
		UseHasMore:           true,
		ResultTag:            "team_folders",
		OnError:              z.OnError,
		OnEntry: func(folder gjson.Result) bool {
			tf := &TeamFolder{}
			err := c.ParseModel(tf, folder)
			if err != nil {
				return z.OnError(err)
			}
			return z.OnEntry(tf)
		},
	}

	return list.List(c, nil)
}
