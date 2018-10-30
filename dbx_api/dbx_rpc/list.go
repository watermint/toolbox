package dbx_rpc

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"go.uber.org/zap"
)

type RpcList struct {
	EndpointList         string
	EndpointListContinue string
	UseHasMore           bool
	AsMemberId           string
	AsAdminId            string
	ResultTag            string
	OnError              func(annotation dbx_api.ErrorAnnotation) bool
	OnResponse           func(body string) bool
	OnEntry              func(result gjson.Result) bool
}

// List and call handlers. Returns true when all operation succeed, otherwise false.
func (r *RpcList) List(c *dbx_api.Context, arg interface{}) bool {
	req := RpcRequest{
		Endpoint:   r.EndpointList,
		AsAdminId:  r.AsAdminId,
		AsMemberId: r.AsMemberId,
		Param:      arg,
	}
	res, et, _ := req.Call(c)

	if !r.handleResponse(c, res, et) {
		return false
	}
	return true
}

func (r *RpcList) listContinue(c *dbx_api.Context, cursor string) bool {
	type ContinueParam struct {
		Cursor string `json:"cursor"`
	}
	req := RpcRequest{
		Endpoint: r.EndpointListContinue,
		Param: ContinueParam{
			Cursor: cursor,
		},
		AsAdminId:  r.AsAdminId,
		AsMemberId: r.AsMemberId,
	}
	res, et, _ := req.Call(c)

	if r.handleResponse(c, res, et) {
		return false
	}
	return true
}

func (r *RpcList) handleResponse(c *dbx_api.Context, res *RpcResponse, et dbx_api.ErrorAnnotation) bool {
	log := c.Log().With(zap.String("endpoint", r.EndpointList))

	if et.IsFailure() {
		if !r.OnError(et) {
			return false
		}
		log.Debug("continue list (error handler returns continue)")
	}

	if r.OnResponse != nil && !r.OnResponse(res.Body) {
		log.Debug("endpoint handler body returned abort")
		return false
	}

	if !r.handleEntry(c, res) {
		return false
	}

	if cont, cursor := r.isContinue(c, res); cont {
		if !r.listContinue(c, cursor) {
			return false
		}
	}
	return true
}

func (r *RpcList) handleEntry(c *dbx_api.Context, res *RpcResponse) bool {
	if r.OnEntry == nil {
		return true
	}

	log := c.Log().With(
		zap.String("endpoint", r.EndpointList),
		zap.String("result_tag", r.ResultTag),
	)

	results := gjson.Get(res.Body, r.ResultTag)
	if !results.Exists() {
		log.Debug("No result found")
		return true
	}

	if !results.IsArray() {
		log.Debug("result was not an array")
		return false
	}

	for _, e := range results.Array() {
		if !r.OnEntry(e) {
			log.Debug("handler returned abort")
			return false
		}
	}
	return true
}

func (r *RpcList) isContinue(c *dbx_api.Context, res *RpcResponse) (cont bool, cursor string) {
	log := c.Log().With(
		zap.String("endpoint", r.EndpointList),
	)

	if r.UseHasMore {
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
