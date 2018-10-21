package team

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAM_INFO = "team/info"
)

type WorkerTeamInfo struct {
	workflow.SimpleWorkerImpl
	ApiManagement *dbx_api.Context
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
	req := dbx_rpc.RpcRequest{
		Endpoint: "team/get_info",
	}
	res, ea, _ := req.Call(w.ApiManagement)
	if ea.IsFailure() {
		w.Pipeline.HandleGeneralFailure(ea)
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
