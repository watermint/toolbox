package dbx_async_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/log/esl"
	"strings"
	"time"
)

const (
	DefaultPollInterval = 3
)

var (
	ErrorAsyncJobIdNotFound = errors.New("async job id not found in the response")
)

func New(ctx dbx_client.Client, endpoint string, reqData []api_request.RequestDatum) dbx_async.Async {
	return &asyncImpl{
		ctx:         ctx,
		reqData:     reqData,
		reqEndpoint: endpoint,
	}
}

type asyncImpl struct {
	ctx         dbx_client.Client
	reqData     []api_request.RequestDatum
	reqEndpoint string
}

func (z asyncImpl) pollDuration(ao dbx_async.AsyncOpts) time.Duration {
	if ao.PollInterval > 0 {
		return time.Duration(ao.PollInterval) * time.Second
	} else {
		return DefaultPollInterval * time.Second
	}
}

func (z asyncImpl) handleNoDotTag(ao dbx_async.AsyncOpts, res es_response.Response, resJson es_json.Json) (asyncRes dbx_async.Response) {
	l := z.ctx.Log()

	if asyncJobIdTag, found := resJson.Find("async_job_id"); found {
		if asyncJobId, found := asyncJobIdTag.String(); !found {
			return dbx_async.NewIncomplete(res)
		} else {
			asyncJobIdTrimSpace := strings.TrimSpace(asyncJobId)
			if asyncJobIdTrimSpace == "" {
				return dbx_async.NewIncomplete(res)

			} else {
				l.Debug("Wait for async", esl.Duration("wait", z.pollDuration(ao)))
				time.Sleep(z.pollDuration(ao))
				return z.handleAsyncJobId(ao, res, resJson, asyncJobIdTrimSpace)
			}
		}
	}

	return dbx_async.NewIncomplete(res)
}

func (z asyncImpl) handleTag(ao dbx_async.AsyncOpts, res es_response.Response, resJson es_json.Json, tag, asyncJobId string) (asyncRes dbx_async.Response) {
	l := z.ctx.Log().With(esl.String("tag", tag), esl.String("asyncJobId", asyncJobId))

	switch tag {
	case "async_job_id":
		l.Debug("Waiting for complete", esl.Duration("wait", z.pollDuration(ao)))
		if jid, ok := resJson.FindString("async_job_id"); ok {
			return z.handleAsyncJobId(ao, res, resJson, jid)
		} else {
			return z.handleAsyncJobId(ao, res, resJson, "")
		}

	case "complete":
		l.Debug("Complete")
		if cmp, found := resJson.Find("complete"); found {
			return dbx_async.NewCompleted(res, cmp)
		} else {
			return dbx_async.NewIncomplete(res)
		}

	case "in_progress":
		l.Debug("In Progress", esl.Duration("wait", z.pollDuration(ao)))
		time.Sleep(z.pollDuration(ao))
		return z.handleAsyncJobId(ao, res, resJson, asyncJobId)

	case "failed":
		l.Debug("Failed", esl.ByteString("body", resJson.Raw()))

		if reason, found := resJson.Find("failed"); found {
			l.Debug("Reason of failure", esl.ByteString("reason", reason.Raw()))
		}
		return dbx_async.NewIncomplete(res)

	default:
		l.Debug("Unknown data format")
		return dbx_async.NewIncomplete(res)
	}
}

func (z asyncImpl) handlePoll(ao dbx_async.AsyncOpts, res es_response.Response, asyncJobId string) (asyncRes dbx_async.Response) {
	resJson, err := res.Success().AsJson()
	if err != nil {
		return dbx_async.NewIncomplete(res)
	}

	l := z.ctx.Log().With(esl.String("async_job_id", asyncJobId))
	l.Debug("Handle poll", esl.ByteString("body", resJson.Raw()))
	if tagJson, found := resJson.Find("\\.tag"); !found {
		return z.handleNoDotTag(ao, res, resJson)
	} else if tag, found := tagJson.String(); found {
		return z.handleTag(ao, res, resJson, tag, asyncJobId)
	}
	return dbx_async.NewIncomplete(res)
}

func (z asyncImpl) findAsyncJobId(resJson es_json.Json, asyncJobId string) (newAsyncJobId string, err error) {
	l := z.ctx.Log().With(esl.String("asyncJobId", asyncJobId))
	if asyncJobId != "" {
		return asyncJobId, nil
	}
	if asyncJobIdTag, found := resJson.Find("async_job_id"); !found {
		l.Debug("async job id tag not found")
		return "", ErrorAsyncJobIdNotFound
	} else {
		if id, found := asyncJobIdTag.String(); found {
			l.Debug("updating async job id value", esl.String("id", id))
			return id, nil
		} else {
			l.Debug("async job id tag is not a string")
			return "", ErrorAsyncJobIdNotFound
		}
	}
}

func (z asyncImpl) handleAsyncJobId(ao dbx_async.AsyncOpts, res es_response.Response, resJson es_json.Json, asyncJobId string) (asyncRes dbx_async.Response) {
	l := z.ctx.Log()

	if aji, err := z.findAsyncJobId(resJson, asyncJobId); err != nil {
		return dbx_async.NewIncomplete(res)
	} else {
		p := struct {
			AsyncJobId string `json:"async_job_id"`
		}{
			AsyncJobId: aji,
		}
		ll := l.With(esl.String("asyncJobId", aji))
		ll.Debug("make status call")
		res := z.ctx.Post(ao.StatusEndpoint, api_request.Param(p))
		if !res.IsSuccess() {
			return dbx_async.NewIncomplete(res)
		}
		return z.handlePoll(ao, res, asyncJobId)
	}
}

func (z asyncImpl) Call(opts ...dbx_async.AsyncOpt) dbx_async.Response {
	ao := dbx_async.Combined(opts)
	rpcRes := z.ctx.Post(z.reqEndpoint, z.reqData...)
	if !rpcRes.IsSuccess() {
		return dbx_async.NewIncomplete(rpcRes)
	}
	return z.handlePoll(ao, rpcRes, "")
}
