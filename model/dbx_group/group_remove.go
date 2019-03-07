package dbx_group

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type Remove struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
}

func (z *Remove) Remove(c *dbx_api.Context, groupId string) bool {
	p := struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}{
		Tag:     "group_id",
		GroupId: groupId,
	}

	req := dbx_rpc.RpcRequest{
		Endpoint: "team/groups/delete",
		Param:    p,
	}
	res, ea, _ := req.Call(c)
	if ea.IsFailure() {
		z.OnError(ea)
		return false
	}
	as := dbx_rpc.AsyncStatus{
		Endpoint: "team/groups/job_status/get",
		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			c.Log().Error("error", zap.Error(annotation.Error))
			return true
		},
		OnComplete: func(complete gjson.Result) bool {
			c.Log().Debug("complete", zap.Any("complete", complete))
			return true
		},
	}
	if as.Poll(c, res) {
		return true
	} else {
		return false
	}
}
