package dbx_rpc

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"go.uber.org/zap"
	"time"
)

type AsyncStatus struct {
	Endpoint   string
	AsMemberId string
	AsAdminId  string
	OnError    func(annotation dbx_api.ErrorAnnotation) bool
	OnComplete func(complete gjson.Result) bool
}

func (a *AsyncStatus) Poll(c *dbx_api.Context, res *RpcResponse) bool {
	return a.handlePoll(c, res, "")
}

func (a *AsyncStatus) handlePoll(c *dbx_api.Context, res *RpcResponse, asyncJobId string) bool {
	log := c.Log().With(zap.String("async_job_id", asyncJobId))
	log.Debug(
		"handlePoll",
		zap.String("body", res.Body),
	)
	tag := gjson.Get(res.Body, dbx_api.ResJsonDotTag)

	if !tag.Exists() {
		err := errors.New("unexpected data format: `.tag` not found")
		annotation := dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorUnexpectedDataType,
			Error:     err,
		}
		if a.OnError != nil {
			return a.OnError(annotation)
		}
		return false
	}

	switch tag.String() {
	case "async_job_id":
		log.Debug("Waiting for complete")

		time.Sleep(time.Duration(3) * time.Second)
		return a.handleAsyncJobId(c, res, "")

	case "complete":
		log.Debug("complete")
		if a.OnComplete != nil {
			return a.OnComplete(gjson.Get(res.Body, "complete"))
		}
		return true

	case "in_progress":
		log.Debug("in_progress")
		time.Sleep(time.Duration(3) * time.Second)
		return a.handleAsyncJobId(c, res, asyncJobId)

	case "failed":
		log.Debug("failed")
		// TODO Log entire message
		if a.OnError == nil {
			return false
		}

		reasonTag := gjson.Get(res.Body, "failed")
		reason := reasonTag.String()
		if !reasonTag.Exists() {
			reason = "operation failed with unknown reason"
		}
		err := errors.New(reason)
		annotation := dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorEndpointSpecific,
			Error:     err,
		}

		return a.OnError(annotation)
	}

	tag = gjson.Get(res.Body, "error."+dbx_api.ResJsonDotTag)
	if tag.Exists() {
		log.Debug(
			"endpoint specific error",
			zap.String("error_tag", tag.String()),
		)
		annotation := dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorEndpointSpecific,
			Error:     dbx_api.ParseApiError(res.Body),
		}
		return a.OnError(annotation)
	}

	c.Log().Debug(
		"Unknown error",
		zap.Int("res_code", res.StatusCode),
		zap.String("res_body", res.Body),
	)

	return false
}

func (a *AsyncStatus) handleAsyncJobId(c *dbx_api.Context, res *RpcResponse, asyncJobId string) bool {
	if asyncJobId == "" {
		asyncJobIdTag := gjson.Get(res.Body, "async_job_id")

		if !asyncJobIdTag.Exists() {
			err := errors.New("unexpected data format: `async_job_id` not found")
			annotation := dbx_api.ErrorAnnotation{
				ErrorType: dbx_api.ErrorUnexpectedDataType,
				Error:     err,
			}
			if a.OnError != nil {
				return a.OnError(annotation)
			}
			return false
		}
		asyncJobId = asyncJobIdTag.String()
	}

	type AsyncParam struct {
		AsyncJobId string `json:"async_job_id"`
	}
	p := AsyncParam{
		AsyncJobId: asyncJobId,
	}

	req := RpcRequest{
		Endpoint:   a.Endpoint,
		Param:      p,
		AsMemberId: a.AsMemberId,
		AsAdminId:  a.AsAdminId,
	}
	res, ea, _ := req.Call(c)

	if ea.IsFailure() {
		if a.OnError != nil {
			return a.OnError(ea)
		}
		return false
	}

	return a.handlePoll(c, res, asyncJobId)
}
