package dbx_async_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/essentials/format/tjson"
	"github.com/watermint/toolbox/essentials/http/response"
	"go.uber.org/zap"
	"strings"
	"time"
)

var (
	ErrorUnexpectedResponseDataFormat = errors.New("unexpected response data format for async job")
	ErrorAsyncJobIdNotFound           = errors.New("async job id not found in the response")
	ErrorAsyncJobFailed               = errors.New("async job failed")
)

func New(ctx dbx_context.Context, endpoint string, asMemberId, asAdminId string, base dbx_context.PathRoot) dbx_async.Async {
	return &asyncImpl{
		ctx:             ctx,
		requestEndpoint: endpoint,
		asMemberId:      asMemberId,
		asAdminId:       asAdminId,
		base:            base,
		pollInterval:    time.Duration(3) * 1000 * time.Millisecond,
	}
}

type asyncImpl struct {
	ctx             dbx_context.Context
	asMemberId      string
	asAdminId       string
	base            dbx_context.PathRoot
	param           interface{}
	pollInterval    time.Duration
	requestEndpoint string
	statusEndpoint  string
}

func (z asyncImpl) Param(p interface{}) dbx_async.Async {
	z.param = p
	return &z
}

func (z asyncImpl) Status(endpoint string) dbx_async.Async {
	z.statusEndpoint = endpoint
	return &z
}

func (z asyncImpl) PollInterval(second int) dbx_async.Async {
	z.pollInterval = time.Duration(second) * 1000 * time.Millisecond
	return &z
}

func (z asyncImpl) poll(res response.Response) (asyncRes dbx_async.Response, resErr error) {
	return z.handlePoll(res, "")
}

func (z asyncImpl) handleNoDotTag(res response.Response, resJson tjson.Json) (asyncRes dbx_async.Response, resErr error) {
	l := z.ctx.Log()

	if asyncJobIdTag, found := resJson.Find("async_job_id"); found {
		if asyncJobId, found := asyncJobIdTag.String(); !found {
			return dbx_async.NewIncomplete(res), nil
		} else {
			asyncJobIdTrimSpace := strings.TrimSpace(asyncJobId)
			if asyncJobIdTrimSpace == "" {
				return dbx_async.NewIncomplete(res), nil

			} else {
				l.Debug("Wait for async", zap.Duration("wait", z.pollInterval))
				time.Sleep(z.pollInterval)
				return z.handleAsyncJobId(res, resJson, asyncJobIdTrimSpace)
			}
		}
	}

	return nil, ErrorUnexpectedResponseDataFormat
}

func (z asyncImpl) handleTag(res response.Response, resJson tjson.Json, tag, asyncJobId string) (asyncRes dbx_async.Response, resErr error) {
	l := z.ctx.Log().With(zap.String("tag", tag), zap.String("asyncJobId", asyncJobId))

	switch tag {
	case "async_job_id":
		l.Debug("Waiting for complete", zap.Duration("wait", z.pollInterval))
		return z.handleAsyncJobId(res, resJson, "")

	case "complete":
		l.Debug("Complete")
		if cmp, found := resJson.Find("complete"); found {
			asyncRes = dbx_async.NewCompleted(res, cmp)
		} else {
			asyncRes = dbx_async.NewIncomplete(res)
		}
		return asyncRes, nil

	case "in_progress":
		l.Debug("In Progress", zap.Duration("wait", z.pollInterval))
		time.Sleep(z.pollInterval)
		return z.handleAsyncJobId(res, resJson, asyncJobId)

	case "failed":
		l.Debug("Failed", zap.ByteString("body", resJson.Raw()))

		if reason, found := resJson.Find("failed"); found {
			l.Debug("Reason of failure", zap.ByteString("reason", reason.Raw()))
		}
		return nil, ErrorAsyncJobFailed

	default:
		if errTag, found := resJson.Find("error.\\.tag"); found {
			l.Debug("Endpoint specific error", zap.ByteString("error_tag", errTag.Raw()))
			return nil, dbx_error.ParseApiError(res.Body().BodyString())
		}
		l.Debug("Unknown data format")
		return nil, ErrorUnexpectedResponseDataFormat
	}
}

func (z asyncImpl) handlePoll(res response.Response, asyncJobId string) (asyncRes dbx_async.Response, resErr error) {
	resJson, err := res.Body().AsJson()
	if err != nil {
		return nil, err
	}

	l := z.ctx.Log().With(zap.String("async_job_id", asyncJobId))
	l.Debug("Handle poll", zap.ByteString("body", resJson.Raw()))
	if tagJson, found := resJson.Find("\\.tag"); !found {
		return z.handleNoDotTag(res, resJson)
	} else if tag, found := tagJson.String(); found {
		return z.handleTag(res, resJson, tag, asyncJobId)
	}
	return nil, ErrorUnexpectedResponseDataFormat
}

func (z asyncImpl) findAsyncJobId(resJson tjson.Json, asyncJobId string) (newAsyncJobId string, err error) {
	l := z.ctx.Log().With(zap.String("asyncJobId", asyncJobId))
	if asyncJobId != "" {
		return asyncJobId, nil
	}
	if asyncJobIdTag, found := resJson.Find("async_job_id"); !found {
		l.Debug("async job id tag not found")
		return "", ErrorAsyncJobIdNotFound
	} else {
		if id, found := asyncJobIdTag.String(); found {
			l.Debug("updating async job id value", zap.String("id", id))
			return id, nil
		} else {
			l.Debug("async job id tag is not a string")
			return "", ErrorAsyncJobIdNotFound
		}
	}
}

func (z *asyncImpl) handleAsyncJobId(res response.Response, resJson tjson.Json, asyncJobId string) (asyncRes dbx_async.Response, resErr error) {
	l := z.ctx.Log()

	if aji, err := z.findAsyncJobId(resJson, asyncJobId); err != nil {
		return nil, err
	} else {
		p := struct {
			AsyncJobId string `json:"async_job_id"`
		}{
			AsyncJobId: aji,
		}
		ll := l.With(zap.String("asyncJobId", aji))
		ll.Debug("make status call")
		res, err := z.ctx.Post(z.statusEndpoint).Param(p).Call()
		if err != nil {
			ll.Debug("error on status call", zap.Error(err))
			return nil, err
		}
		return z.handlePoll(res, asyncJobId)
	}
}

func (z *asyncImpl) Call() (res dbx_async.Response, resErr error) {
	rpcRes, err := z.ctx.Post(z.requestEndpoint).Param(z.param).Call()
	if err != nil {
		return nil, err
	}
	return z.poll(rpcRes)
}
