package api_async_impl

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_async"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_error"
	"github.com/watermint/toolbox/infra/api/api_response"
	"go.uber.org/zap"
	"strings"
	"time"
)

func New(ctx api_context.Context, endpoint string, asMemberId, asAdminId string, base api_context.PathRoot) api_async.Async {
	return &asyncImpl{
		ctx:             ctx,
		requestEndpoint: endpoint,
		asMemberId:      asMemberId,
		asAdminId:       asAdminId,
		base:            base,
		pollInterval:    time.Duration(3) * time.Second,
	}
}

type asyncImpl struct {
	ctx             api_context.Context
	asMemberId      string
	asAdminId       string
	base            api_context.PathRoot
	param           interface{}
	pollInterval    time.Duration
	requestEndpoint string
	statusEndpoint  string
	success         func(res api_async.Response) error
	failure         func(err error) error
}

func (z *asyncImpl) Param(p interface{}) api_async.Async {
	z.param = p
	return z
}

func (z *asyncImpl) Status(endpoint string) api_async.Async {
	z.statusEndpoint = endpoint
	return z
}

func (z *asyncImpl) PollInterval(second int) api_async.Async {
	z.pollInterval = time.Duration(second) * time.Second
	return z
}

func (z *asyncImpl) OnSuccess(success func(res api_async.Response) error) api_async.Async {
	z.success = success
	return z
}

func (z *asyncImpl) OnFailure(failure func(err error) error) api_async.Async {
	z.failure = failure
	return z
}

func (z *asyncImpl) poll(res api_response.Response) (asyncRes api_async.Response, resErr error) {
	return z.handlePoll(res, "")
}

func (z *asyncImpl) handlePoll(res api_response.Response, asyncJobId string) (asyncRes api_async.Response, resErr error) {
	resJson, err := res.Json()
	if err != nil {
		return nil, err
	}

	log := z.ctx.Log().With(zap.String("async_job_id", asyncJobId))
	log.Debug("Handle poll", zap.String("body", resJson.Raw))
	tag := resJson.Get("\\.tag")

	if !tag.Exists() {
		asyncJobIdTag := resJson.Get("async_job_id")
		if asyncJobIdTag.Exists() {
			asyncJobId := strings.TrimSpace(asyncJobIdTag.String())
			if asyncJobId == "" {
				asyncRes = &responseImpl{
					res:            res,
					complete:       resJson,
					completeExists: false,
				}
				return asyncRes, nil

			} else {
				log.Debug("Wait for async", zap.Duration("wait", z.pollInterval))
				time.Sleep(z.pollInterval)
				return z.handleAsyncJobId(res, asyncJobIdTag.String())
			}
		} else {
			err := errors.New("unexpected data format: `.tag` not found")
			if z.failure != nil {
				if err := z.failure(err); err != nil {
					return nil, err
				}
				// fall through
			} else {
				return nil, err
			}
		}
	}

	switch tag.String() {
	case "async_job_id":
		log.Debug("Waiting for complete", zap.Duration("wait", z.pollInterval))
		return z.handleAsyncJobId(res, "")

	case "complete":
		log.Debug("Complete")
		cmp := resJson.Get("complete")
		if cmp.Exists() {
			asyncRes = &responseImpl{
				res:            res,
				complete:       cmp,
				completeExists: true,
			}
		} else {
			asyncRes = &responseImpl{
				res:            res,
				complete:       resJson,
				completeExists: false,
			}
		}
		if z.success != nil {
			if err = z.success(asyncRes); err != nil {
				return nil, err
			}
		}
		return asyncRes, nil

	case "in_progress":
		log.Debug("In Progress", zap.Duration("wait", z.pollInterval))
		time.Sleep(z.pollInterval)
		return z.handleAsyncJobId(res, asyncJobId)

	case "failed":
		log.Debug("Failed", zap.String("body", resJson.Raw))
		if z.failure == nil {
			return nil, errors.New("failed") // TODO create error type for "failed"
		}

		reasonTag := resJson.Get("failed")
		reason := reasonTag.String()
		if !reasonTag.Exists() {
			reason = "operation failed with unknown reason"
		}
		err := errors.New(reason)
		if err2 := z.failure(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}

	if errorTag := resJson.Get("error.\\.tag"); errorTag.Exists() {
		log.Debug("Endpoint specific error", zap.String("error_tag", tag.String()))
		body, _ := res.Result()
		err = api_error.ParseApiError(body)
		if err2 := z.failure(err); err2 != nil {
			return nil, err2
		}
		return nil, err
	}

	z.ctx.Log().Debug("Unknown status tag", zap.String("tag", tag.String()), zap.Int("res_code", res.StatusCode()), zap.String("res_body", resJson.Raw))
	return nil, errors.New("unknown status tag")
}

func (z *asyncImpl) handleAsyncJobId(res api_response.Response, asyncJobId string) (asyncRes api_async.Response, resErr error) {
	resJson, err := res.Json()
	if err != nil {
		return nil, err
	}

	if asyncJobId == "" {
		asyncJobIdTag := resJson.Get("async_job_id")

		if !asyncJobIdTag.Exists() {
			err := errors.New("unexpected data format: `async_job_id` not found")
			if z.failure != nil {
				if err2 := z.failure(err); err2 != nil {
					return nil, err2
				}
			}
			return nil, err
		}
		asyncJobId = asyncJobIdTag.String()
	}

	p := struct {
		AsyncJobId string `json:"async_job_id"`
	}{
		AsyncJobId: asyncJobId,
	}

	res, err = z.ctx.Rpc(z.statusEndpoint).Param(p).Call()
	if err != nil {
		if z.failure != nil {
			if err2 := z.failure(err); err2 != nil {
				return nil, err2
			}
		}
		return nil, err
	}

	return z.handlePoll(res, asyncJobId)
}

func (z *asyncImpl) Call() (res api_async.Response, resErr error) {
	rpcRes, err := z.ctx.Rpc(z.requestEndpoint).Param(z.param).Call()
	if err != nil {
		return nil, err
	}
	return z.poll(rpcRes)
}
