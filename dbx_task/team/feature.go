package team

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAM_FEATURES = "team/features"
)

type WorkerTeamFeatures struct {
	workflow.SimpleWorkerImpl
	ApiManagement *dbx_api.Context
	NextTask      string
}

type ContextTeamFeature struct {
	Feature string          `json:"feature"`
	Value   json.RawMessage `json:"value"`
}

func (*WorkerTeamFeatures) Prefix() string {
	return WORKER_TEAM_FEATURES
}

func (w *WorkerTeamFeatures) Exec(task *workflow.Task) {
	type FeatureTag struct {
		Tag string `json:".tag"`
	}
	type FeatureParam struct {
		Values []FeatureTag `json:"features"`
	}

	param := FeatureParam{
		Values: []FeatureTag{
			{Tag: "upload_api_rate_limit"},
			{Tag: "has_team_shared_dropbox"},
			{Tag: "has_team_file_events"},
			{Tag: "has_team_selective_sync"},
		},
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/features/get_values",
		Param:    param,
	}
	res, ea, _ := req.Call(w.ApiManagement)
	if ea.IsFailure() {
		w.Pipeline.HandleGeneralFailure(ea)
		return
	}

	values := gjson.Get(res.Body, "values")
	for _, v := range values.Array() {
		feature := v.Get("\\.tag").String()

		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.NextTask,
				feature,
				ContextTeamFeature{
					Feature: feature,
					Value:   json.RawMessage(v.Raw),
				},
			),
		)
	}
}
