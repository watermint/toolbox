package api_capture

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_rpc"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"regexp"
	"time"
)

const (
	valuePathCapture = "api_capture.Capture"
)

type Capture interface {
	Rpc(req api_rpc.Request, res api_rpc.Response, resErr error, latency int64)
}

func currentKitchen(cap *zap.Logger) Capture {
	return &kitchenImpl{
		capture: cap,
	}
}

func Current() Capture {
	cap := app_root.Capture()
	return currentKitchen(cap)
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

func (mockImpl) Rpc(req api_rpc.Request, res api_rpc.Response, resErr error, latency int64) {
	// ignore
}

var (
	tokenMatcher = regexp.MustCompile(`\w`)
)

func NewCapture(cap *zap.Logger) Capture {
	return &kitchenImpl{
		capture: cap,
	}
}

type kitchenImpl struct {
	capture *zap.Logger
}

func (z *kitchenImpl) Rpc(req api_rpc.Request, res api_rpc.Response, resErr error, latency int64) {
	type Req struct {
		RequestMethod  string            `json:"method"`
		RequestUrl     string            `json:"url"`
		RequestParam   string            `json:"param,omitempty"`
		RequestHeaders map[string]string `json:"headers"`
	}
	type Res struct {
		ResponseCode  int             `json:"code"`
		ResponseBody  string          `json:"body,omitempty"`
		ResponseJson  json.RawMessage `json:"json,omitempty"`
		ResponseError string          `json:"error,omitempty"`
	}

	// request
	rq := Req{}
	rq.RequestMethod = req.Method()
	rq.RequestUrl = req.Url()
	rq.RequestParam = req.Param()
	headers := make(map[string]string)
	for k, v := range req.Headers() {
		// Anonymize token
		if k == api_rpc.ReqHeaderAuthorization {
			headers[k] = "Bearer <secret>"
		} else {
			headers[k] = v
		}
	}
	rq.RequestHeaders = headers

	// response
	rs := Res{}
	rs.ResponseCode = res.StatusCode()
	resBody, _ := res.Body()
	if resBody[0] == '[' || resBody[0] == '{' {
		rs.ResponseJson = []byte(resBody)
	} else {
		rs.ResponseBody = resBody
	}
	if resErr != nil {
		rs.ResponseError = resErr.Error()
	}

	z.capture.Debug("", zap.Any("req", rq), zap.Any("res", rs), zap.Int64("latency", latency))
}
