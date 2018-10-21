package dbx_rpc

import (
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
)

type RpcList struct {
	EndpointList         string
	EndpointListContinue string
	UseHasMore           bool
	AsMemberId           string
	AsAdminId            string
	ResultTag            string
	HandlerError         func(annotation dbx_api.ErrorAnnotation) bool
	HandlerBody          func(body string) bool
	HandlerEntry         func(result gjson.Result) bool
}

// List and call handlers. Returns true when all operation succeed, otherwise false.
func (r *RpcList) List(c *dbx_api.Context, arg interface{}) bool {
	seelog.Debugf("Endpoint[%s]", r.EndpointList)
	req := RpcRequest{
		Endpoint:   r.EndpointList,
		AsAdminId:  r.AsAdminId,
		AsMemberId: r.AsMemberId,
		Param:      arg,
	}
	res, et, _ := req.Call(c)

	if r.handleResponse(c, res, et) {
		return false
	}
	return true
}

func (r *RpcList) listContinue(c *dbx_api.Context, cursor string) bool {
	type ContinueParam struct {
		Cursor string `json:"cursor"`
	}
	seelog.Debugf("Endpoint[%s] Cursor[%s] Continue", r.EndpointListContinue, cursor)
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
	if et.IsFailure() {
		if !r.HandlerError(et) {
			return false
		}
		seelog.Debugf("Endpoint[%s] Continue list (error handler returns continue)", r.EndpointList)
	}

	if r.HandlerBody != nil && !r.HandlerBody(res.Body) {
		seelog.Debugf("Endpoint[%s] Handler Body returned abort", r.EndpointList)
		return false
	}

	if !r.handleEntry(res) {
		return false
	}

	if cont, cursor := r.isContinue(res); cont {
		if !r.listContinue(c, cursor) {
			return false
		}
	}
	return true
}

func (r *RpcList) handleEntry(res *RpcResponse) bool {
	if r.HandlerEntry == nil {
		seelog.Warnf("No entry handler found for Endpoint[%s]", r.EndpointList)
		return false
	}

	results := gjson.Get(res.Body, r.ResultTag)
	if !results.Exists() {
		seelog.Debugf("Endpoint[%s] no results found for tag [%s]", r.EndpointList, r.ResultTag)
		return true
	}

	if !results.IsArray() {
		seelog.Debugf("Endpoint[%s] result[%s] was not an array", r.EndpointList, r.ResultTag)
		return false
	}

	for _, e := range results.Array() {
		if !r.HandlerEntry(e) {
			seelog.Debugf("Endpoint[%s] Handler Entry returned abort. Entry[%s]", r.EndpointList, e.Raw)
			return false
		}
	}
	return true
}

func (r *RpcList) isContinue(res *RpcResponse) (cont bool, cursor string) {
	if r.UseHasMore {
		if gjson.Get(res.Body, "has_more").Bool() {
			cursor = gjson.Get(res.Body, "cursor").String()
			if cursor != "" {
				return true, cursor
			}
			seelog.Debugf("Endpoint[%s] has_more returned true, but no cursor represented", r.EndpointList)
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
