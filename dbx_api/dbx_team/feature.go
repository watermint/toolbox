package dbx_team

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type Feature struct {
	Feature string          `json:"feature"`
	Value   json.RawMessage `json:"value"`
}

type FeatureList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(feature *Feature) bool
}

func (w *FeatureList) List(c *dbx_api.Context) bool {
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
	res, ea, _ := req.Call(c)
	if ea.IsFailure() {
		if w.OnError != nil {
			w.OnError(ea)
		}
		return false
	}

	values := gjson.Get(res.Body, "values")
	for _, v := range values.Array() {
		feature := v.Get("\\.tag").String()

		f := &Feature{
			Feature: feature,
			Value:   json.RawMessage(v.Raw),
		}
		if w.OnEntry != nil {
			return w.OnEntry(f)
		}
		return false
	}
	return true
}
