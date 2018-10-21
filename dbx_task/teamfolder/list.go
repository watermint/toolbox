package teamfolder

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAMFOLDER_LIST = "teamfolder/list"
)

type WorkerTeamFolderList struct {
	workflow.SimpleWorkerImpl
	Api      *dbx_api.Context
	NextTask string
}

type ContextTeamFolderListResult struct {
	TeamFolderId string          `json:"team_folder_id"`
	Name         string          `json:"name"`
	TeamFolder   json.RawMessage `json:"team_folder"`
}

func (w *WorkerTeamFolderList) Prefix() string {
	return WORKER_TEAMFOLDER_LIST
}

func (w *WorkerTeamFolderList) Exec(task *workflow.Task) {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/team_folder/list",
		EndpointListContinue: "team/team_folder/list/continue",
		UseHasMore:           true,
		ResultTag:            "team_folders",
		HandlerError:         w.Pipeline.HandleGeneralFailure,
		HandlerEntry: func(folder gjson.Result) bool {
			teamFolderId := folder.Get("team_folder_id").String()

			c := ContextTeamFolderListResult{
				TeamFolderId: teamFolderId,
				Name:         folder.Get("name").String(),
				TeamFolder:   json.RawMessage(folder.Raw),
			}

			w.Pipeline.Enqueue(
				workflow.MarshalTask(
					w.NextTask,
					teamFolderId,
					c,
				),
			)
			return true
		},
	}

	list.List(w.Api, nil)
}
