package nw_capture

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"regexp"
	"time"
)

type Capture interface {
	Rpc(req api_request.Request, res api_response.Response, resErr error, latency int64)
}

func currentImpl(cap *zap.Logger) Capture {
	return &captureImpl{
		capture: cap,
	}
}

func Current() Capture {
	cap := app_root.Capture()
	return currentImpl(cap)
}

type Record struct {
	Timestamp      time.Time         `json:"timestamp"`
	RequestMethod  string            `json:"req_method"`
	RequestUrl     string            `json:"req_url"`
	RequestParam   string            `json:"req_param,omitempty"`
	RequestHeaders map[string]string `json:"req_headers"`
	ResponseCode   int               `json:"res_code"`
	ResponseBody   string            `json:"res_body,omitempty"`
	ResponseError  string            `json:"res_error,omitempty"`
	Latency        int64             `json:"latency"`
}

type mockImpl struct {
}

func (mockImpl) Rpc(req api_request.Request, res api_response.Response, resErr error, latency int64) {
	// ignore
}

var (
	tokenMatcher = regexp.MustCompile(`\w`)
)

func NewCapture(cap *zap.Logger) Capture {
	return &captureImpl{
		capture: cap,
	}
}

type captureImpl struct {
	capture *zap.Logger
}

func (z *captureImpl) Rpc(req api_request.Request, res api_response.Response, resErr error, latency int64) {
	type Req struct {
		RequestMethod  string            `json:"method"`
		RequestUrl     string            `json:"url"`
		RequestParam   string            `json:"param,omitempty"`
		RequestHeaders map[string]string `json:"headers"`
	}
	type Res struct {
		ResponseCode    int               `json:"code"`
		ResponseBody    string            `json:"body,omitempty"`
		ResponseHeaders map[string]string `json:"headers"`
		ResponseJson    json.RawMessage   `json:"json,omitempty"`
		ResponseError   string            `json:"error,omitempty"`
	}

	// request
	rq := Req{}
	rq.RequestMethod = req.Method()
	rq.RequestUrl = req.Url()
	rq.RequestParam = req.ParamString()
	rq.RequestHeaders = make(map[string]string)
	for k, v := range req.Headers() {
		// Anonymize token
		if k == api_request.ReqHeaderAuthorization {
			rq.RequestHeaders[k] = "Bearer <secret>"
		} else {
			rq.RequestHeaders[k] = v
		}
	}

	// response
	rs := Res{}
	rs.ResponseCode = res.StatusCode()
	resBody, _ := res.Result()
	if len(resBody) == 0 {
		rs.ResponseBody = ""
	} else if resBody[0] == '[' || resBody[0] == '{' {
		rs.ResponseJson = []byte(resBody)
	} else {
		rs.ResponseBody = resBody
	}
	if resErr != nil {
		rs.ResponseError = resErr.Error()
	}
	rs.ResponseHeaders = res.Headers()

	z.capture.Debug("",
		zap.Any("req", rq),
		zap.Any("res", rs),
		zap.Int64("latency", latency),
	)
}
