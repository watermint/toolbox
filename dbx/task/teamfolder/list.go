package teamfolder

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAMFOLDER_LIST = "teamfolder/list"
)

type WorkerTeamFolderList struct {
	workflow.SimpleWorkerImpl
	Api      *api.ApiContext
	NextTask string
}

type ContextTeamFolderList struct {
	Cursor string `json:"cursor"`
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
	tc := &ContextTeamFolderList{}
	workflow.UnmarshalContext(task, tc)

	if tc.Cursor == "" {
		w.list(task)
	} else {
		w.listContinue(tc.Cursor, task)
	}
}

func (w *WorkerTeamFolderList) list(task *workflow.Task) {
	seelog.Info("Loading team folder list")
	cont, res, _ := w.Pipeline.TaskRpc(task, w.Api, "team/team_folder/list", nil)
	if !cont {
		return
	}

	w.processResult(res, task)
}

func (w *WorkerTeamFolderList) listContinue(cursor string, task *workflow.Task) {
	type ListContinueParam struct {
		Cursor string `json:"cursor"`
	}
	lp := ListContinueParam{
		Cursor: cursor,
	}

	seelog.Debugf("team_folder/list/continue (cursor: %s)", cursor)
	cont, res, _ := w.Pipeline.TaskRpc(task, w.Api, "team/team_folder/list/continue", lp)
	if !cont {
		return
	}

	w.processResult(res, task)
}

func (w *WorkerTeamFolderList) processResult(res *api.ApiRpcResponse, task *workflow.Task) {
	folders := gjson.Get(res.Body, "team_folders")
	if !folders.Exists() {
		seelog.Debugf("`team_folders` data not found")
		return
	}

	for _, folder := range folders.Array() {
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
	}

	hasMoreJson := gjson.Get(res.Body, "has_more")
	if hasMoreJson.Exists() && hasMoreJson.Bool() {
		cursorJson := gjson.Get(res.Body, "cursor")
		if !cursorJson.Exists() {
			seelog.Debug("Cursor not found in the response (has_more appear and true)")
			return
		}
		c := ContextTeamFolderList{
			Cursor: cursorJson.String(),
		}

		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.Prefix(),
				cursorJson.String(),
				c,
			),
		)
	}
}
