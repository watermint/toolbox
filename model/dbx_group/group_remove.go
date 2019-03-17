package dbx_group

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type Remove struct {
	OnError   func(err error) bool
	OnSuccess func()
}

func (z *Remove) Remove(c *dbx_api.DbxContext, groupId string) bool {
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
	res, err := req.Call(c)
	if err != nil {
		z.OnError(err)
		return false
	}
	as := dbx_rpc.AsyncStatus{
		Endpoint: "team/groups/job_status/get",
		OnError: func(err error) bool {
			c.Log().Error("error", zap.Error(err))
			return true
		},
		OnComplete: func(complete gjson.Result) bool {
			c.Log().Debug("complete", zap.Any("complete", complete))
			return true
		},
	}
	if as.Poll(c, res) {
		z.OnSuccess()
		return true
	} else {
		return false
	}
}
