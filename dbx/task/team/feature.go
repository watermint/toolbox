package team

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAM_FEATURES = "team/features"
)

type WorkerTeamFeatures struct {
	workflow.SimpleWorkerImpl
	ApiManagement *api.ApiContext
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

	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/features/get_values", param)
	if !cont {
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
