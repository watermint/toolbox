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

func (w *ListTeamFolder) List(c *dbx_api.Context) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/team_folder/list",
		EndpointListContinue: "team/team_folder/list/continue",
		UseHasMore:           true,
		ResultTag:            "team_folders",
		OnError:              w.OnError,
		OnEntry: func(folder gjson.Result) bool {
			tf := &TeamFolder{}
			err := c.ParseModel(tf, folder)
			if err != nil {
				return w.OnError(err)
			}
			return w.OnEntry(tf)
		},
	}

	return list.List(c, nil)
}
