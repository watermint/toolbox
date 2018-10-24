package dbx_team

import (
	"context"
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type TeamInfo struct {
	TeamId string          `json:"team_id"`
	Info   json.RawMessage `json:"info"`
}

type TeamInfoList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(info *TeamInfo) bool
}

func (t *TeamInfoList) List(c *context.Context) bool {
	req := dbx_rpc.RpcRequest{
		Endpoint: "team/get_info",
	}
	res, ea, _ := req.Call(c)
	if ea.IsFailure() {
		if t.OnError != nil {
			return t.OnError(ea)
		}
		return false
	}

	teamId := gjson.Get(res.Body, "team_id").String()
	team := &TeamInfo{
		TeamId: teamId,
		Info:   json.RawMessage(res.Body),
	}

	if t.OnEntry != nil {
		return t.OnEntry(team)
	}
	return true
}
