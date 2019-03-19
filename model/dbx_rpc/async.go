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

	OnError    func(err error) bool
	OnComplete func(complete gjson.Result) bool
}

func (z *AsyncStatus) Poll(c *dbx_api.DbxContext, res *RpcResponse) bool {
	return z.handlePoll(c, res, "")
}

func (z *AsyncStatus) handlePoll(c *dbx_api.DbxContext, res *RpcResponse, asyncJobId string) bool {
	resJson := gjson.Parse(res.Body)

	log := c.Log().With(zap.String("async_job_id", asyncJobId))
	log.Debug("Handle poll", zap.String("body", res.Body))
	tag := resJson.Get(dbx_api.ResJsonDotTag)

	if !tag.Exists() {
		asyncJobId := resJson.Get("async_job_id")
		if asyncJobId.Exists() {
			time.Sleep(time.Duration(3) * time.Second)
			return z.handleAsyncJobId(c, res, asyncJobId.String())

		} else {
			err := errors.New("unexpected data format: `.tag` not found")
			return z.OnError(err)
		}
	}

	switch tag.String() {
	case "async_job_id":
		log.Debug("Waiting for complete")

		time.Sleep(time.Duration(3) * time.Second)
		return z.handleAsyncJobId(c, res, "")

	case "complete":
		log.Debug("Complete")
		if z.OnComplete != nil {
			cmp := resJson.Get("complete")
			if cmp.Exists() {
				return z.OnComplete(cmp)
			} else {
				return z.OnComplete(resJson)
			}
		}
		return true

	case "in_progress":
		log.Debug("In Progress")
		time.Sleep(time.Duration(3) * time.Second)
		return z.handleAsyncJobId(c, res, asyncJobId)

	case "failed":
		log.Debug("Failed")
		// TODO Log entire message
		if z.OnError == nil {
			return false
		}

		reasonTag := gjson.Get(res.Body, "failed")
		reason := reasonTag.String()
		if !reasonTag.Exists() {
			reason = "operation failed with unknown reason"
		}
		err := errors.New(reason)
		return z.OnError(err)
	}

	tag = gjson.Get(res.Body, "error."+dbx_api.ResJsonDotTag)
	if tag.Exists() {
		log.Debug("Endpoint specific error", zap.String("error_tag", tag.String()))
		return z.OnError(dbx_api.ParseApiError(res.Body))
	}

	c.Log().Debug("Unknown error", zap.Int("res_code", res.StatusCode), zap.String("res_body", res.Body))

	return false
}

func (z *AsyncStatus) handleAsyncJobId(c *dbx_api.DbxContext, res *RpcResponse, asyncJobId string) bool {
	if asyncJobId == "" {
		asyncJobIdTag := gjson.Get(res.Body, "async_job_id")

		if !asyncJobIdTag.Exists() {
			err := errors.New("unexpected data format: `async_job_id` not found")
			return z.OnError(err)
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
		Endpoint:   z.Endpoint,
		Param:      p,
		AsMemberId: z.AsMemberId,
		AsAdminId:  z.AsAdminId,
	}
	res, err := req.Call(c)

	if err != nil {
		if z.OnError != nil {
			return z.OnError(err)
		}
		return false
	}

	return z.handlePoll(c, res, asyncJobId)
}
