package team

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAM_INFO = "team/info"
)

type WorkerTeamInfo struct {
	workflow.SimpleWorkerImpl
	ApiManagement *api.ApiContext
	NextTask      string
}

type ContextTeamInfo struct {
	TeamId string          `json:"team_id"`
	Info   json.RawMessage `json:"info"`
}

func (w *WorkerTeamInfo) Prefix() string {
	return WORKER_TEAM_INFO
}

func (w *WorkerTeamInfo) Exec(task *workflow.Task) {
	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/get_info", nil)
	if !cont {
		return
	}

	teamId := gjson.Get(res.Body, "team_id").String()

	w.Pipeline.Enqueue(
		workflow.MarshalTask(
			w.NextTask,
			teamId,
			ContextTeamInfo{
				TeamId: teamId,
				Info:   json.RawMessage(res.Body),
			},
		),
	)
}
