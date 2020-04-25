package nw_capture

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"go.uber.org/zap"
	"strings"
)

func New(client nw_client.Http) nw_client.Rest {
	return &Client{httpClient: client}
}

type Client struct {
	httpClient nw_client.Http
}

func (z *Client) Call(ctx api_context.Context, req api_request.Request) (res response.Response, err error) {
	l := ctx.Log()
	hReq, err := req.Make()
	if err != nil {
		l.Debug("Unable to make http request", zap.Error(err))
		return nil, err
	}

	// Call
	hRes, latency, err := z.httpClient.Call(ctx.ClientHash(), req.Endpoint(), hReq)

	// Make response
	cp := NewCapture(ctx.Capture())

	res = response.New(ctx, hRes)
	if bodyErr := res.Body().Error(); bodyErr != nil {
		l.Debug("Unable to make http response", zap.Error(bodyErr))
		cp.NoResponse(req, bodyErr, latency.Nanoseconds())
		return nil, bodyErr
	}
	cp.WithResponse(req, res, err, latency.Nanoseconds())

	return res, nil
}

type Capture interface {
	WithResponse(req api_request.Request, res response.Response, resErr error, latency int64)
	NoResponse(req api_request.Request, resErr error, latency int64)
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
			vv := strings.Split(v, " ")
			z.RequestHeaders[k] = vv[0] + " <secret>"
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

func (z *Res) Apply(res response.Response, resErr error) {
	z.ResponseCode = res.Code()
	z.ContentLength = res.Body().ContentLength()
	if res.Body().IsFile() {
		z.ResponseBody = ""
	} else {
		z.ResponseBody = res.Body().BodyString()
	}
	if resErr != nil {
		z.ResponseError = resErr.Error()
	}
	z.ResponseHeaders = res.Headers()
}

func (z *captureImpl) WithResponse(req api_request.Request, res response.Response, resErr error, latency int64) {
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
