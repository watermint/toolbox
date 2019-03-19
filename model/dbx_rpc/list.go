package dbx_rpc

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"go.uber.org/zap"
)

type RpcList struct {
	EndpointList         string
	EndpointListContinue string
	PathRoot             interface{}
	UseHasMore           bool
	AsMemberId           string
	AsAdminId            string
	ResultTag            string
	OnError              func(err error) bool
	OnResponse           func(body string) bool
	OnEntry              func(result gjson.Result) bool
}

// List and call handlers. Returns true when all operation succeed, otherwise false.
func (z *RpcList) List(c *dbx_api.DbxContext, arg interface{}) bool {
	req := RpcRequest{
		Endpoint:   z.EndpointList,
		PathRoot:   z.PathRoot,
		AsAdminId:  z.AsAdminId,
		AsMemberId: z.AsMemberId,
		Param:      arg,
	}
	res, err := req.Call(c)

	return z.handleResponse(c, res, err)
}

func (z *RpcList) listContinue(c *dbx_api.DbxContext, cursor string) bool {
	type ContinueParam struct {
		Cursor string `json:"cursor"`
	}
	req := RpcRequest{
		Endpoint: z.EndpointListContinue,
		Param: ContinueParam{
			Cursor: cursor,
		},
		AsAdminId:  z.AsAdminId,
		AsMemberId: z.AsMemberId,
	}
	res, err := req.Call(c)

	return z.handleResponse(c, res, err)
}

func (z *RpcList) handleResponse(c *dbx_api.DbxContext, res *RpcResponse, err error) bool {
	log := c.Log().With(zap.String("endpoint", z.EndpointList))

	if err != nil {
		if !z.OnError(err) || res == nil {
			return false
		}
		log.Debug("continue list (error handler returns continue)")
	}

	if res == nil {
		log.Debug("Response is null")
		return false
	}

	if z.OnResponse != nil && !z.OnResponse(res.Body) {
		log.Debug("endpoint handler body returned abort")
		return false
	}

	if !z.handleEntry(c, res) {
		return false
	}

	if cont, cursor := z.isContinue(c, res); cont {
		if !z.listContinue(c, cursor) {
			return false
		}
	}
	return true
}

func (z *RpcList) handleEntry(c *dbx_api.DbxContext, res *RpcResponse) bool {
	if z.OnEntry == nil {
		return true
	}

	log := c.Log().With(
		zap.String("endpoint", z.EndpointList),
		zap.String("result_tag", z.ResultTag),
	)

	results := gjson.Get(res.Body, z.ResultTag)
	if !results.Exists() {
		log.Debug("No result found")
		return true
	}

	if !results.IsArray() {
		log.Debug("result was not an array")
		return false
	}

	for _, e := range results.Array() {
		if !z.OnEntry(e) {
			log.Debug("handler returned abort")
			return false
		}
	}
	return true
}

func (z *RpcList) isContinue(c *dbx_api.DbxContext, res *RpcResponse) (cont bool, cursor string) {
	log := c.Log().With(
		zap.String("endpoint", z.EndpointList),
	)

	if z.UseHasMore {
		if gjson.Get(res.Body, "has_more").Bool() {
			cursor = gjson.Get(res.Body, "cursor").String()
			if cursor != "" {
				return true, cursor
			}
			log.Debug("has_more returned true, but no cursor found in the body")
			return false, ""
		} else {
			return false, ""
		}
	}

	cursor = gjson.Get(res.Body, "cursor").String()
	if cursor != "" {
		return true, cursor
	}
	return false, ""
}
