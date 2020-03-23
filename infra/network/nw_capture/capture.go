package nw_capture

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

type Capture interface {
	WithResponse(req api_request.Request, res api_response.Response, resErr error, latency int64)
	NoResponse(req api_request.Request, resErr error, latency int64)
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

func NewCapture(cap *zap.Logger) Capture {
	return &captureImpl{
		capture: cap,
	}
}

type captureImpl struct {
	capture *zap.Logger
}

type Record struct {
	Time    string `json:"time"`
	Req     *Req   `json:"req"`
	Res     *Res   `json:"res"`
	Latency int64  `json:"latency"`
}

type Req struct {
	RequestMethod  string            `json:"method"`
	RequestUrl     string            `json:"url"`
	RequestParam   string            `json:"param,omitempty"`
	RequestHeaders map[string]string `json:"headers"`
	ContentLength  int64             `json:"content_length"`
}

func (z *Req) Apply(req api_request.Request) {
	z.RequestMethod = req.Method()
	z.RequestUrl = req.Url()
	z.RequestParam = req.ParamString()
	z.RequestHeaders = make(map[string]string)
	z.ContentLength = req.ContentLength()
	for k, v := range req.Headers() {
		// Anonymize token
		if k == api_request.ReqHeaderAuthorization {
			z.RequestHeaders[k] = "Bearer <secret>"
		} else {
			z.RequestHeaders[k] = v
		}
	}
}

type Res struct {
	ResponseCode    int               `json:"code"`
	ResponseBody    string            `json:"body,omitempty"`
	ResponseHeaders map[string]string `json:"headers"`
	ResponseJson    json.RawMessage   `json:"json,omitempty"`
	ResponseError   string            `json:"error,omitempty"`
	ContentLength   int64             `json:"content_length"`
}

func (z *Res) Apply(res api_response.Response, resErr error) {
	z.ResponseCode = res.StatusCode()
	z.ContentLength = res.ContentLength()
	resBody, _ := res.Result()
	if len(resBody) == 0 {
		z.ResponseBody = ""
	} else if resBody[0] == '[' || resBody[0] == '{' {
		z.ResponseJson = []byte(resBody)
	} else {
		z.ResponseBody = resBody
	}
	if resErr != nil {
		z.ResponseError = resErr.Error()
	}
	z.ResponseHeaders = res.Headers()
}

func (z *captureImpl) WithResponse(req api_request.Request, res api_response.Response, resErr error, latency int64) {
	// request
	rq := Req{}
	rq.Apply(req)

	// response
	rs := Res{}
	rs.Apply(res, resErr)

	z.capture.Debug("",
		zap.Any("req", rq),
		zap.Any("res", rs),
		zap.Int64("latency", latency),
	)
}

func (z *captureImpl) NoResponse(req api_request.Request, resErr error, latency int64) {
	// request
	rq := Req{}
	rq.Apply(req)

	// response
	rs := Res{}
	if resErr != nil {
		rs.ResponseError = resErr.Error()
	}

	z.capture.Debug("",
		zap.Any("req", rq),
		zap.Any("res", rs),
		zap.Int64("latency", latency),
	)
}
