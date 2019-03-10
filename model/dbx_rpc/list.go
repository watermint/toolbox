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
func (r *RpcList) List(c *dbx_api.Context, arg interface{}) bool {
	req := RpcRequest{
		Endpoint:   r.EndpointList,
		PathRoot:   r.PathRoot,
		AsAdminId:  r.AsAdminId,
		AsMemberId: r.AsMemberId,
		Param:      arg,
	}
	res, err := req.Call(c)

	if !r.handleResponse(c, res, err) {
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
	res, err := req.Call(c)

	if r.handleResponse(c, res, err) {
		return false
	}
	return true
}

func (r *RpcList) handleResponse(c *dbx_api.Context, res *RpcResponse, err error) bool {
	log := c.Log().With(zap.String("endpoint", r.EndpointList))

	if err != nil {
		if !r.OnError(err) || res == nil {
			return false
		}
		log.Debug("continue list (error handler returns continue)")
	}

	if res == nil {
		log.Debug("Response is null")
		return false
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
