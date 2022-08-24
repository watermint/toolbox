package nw_capture

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_request"
	"github.com/watermint/toolbox/infra/api/api_client"
	"net/http"
)

func New(client nw_client.Http) nw_client.Rest {
	return &Client{httpClient: client}
}

type Client struct {
	httpClient nw_client.Http
}

func (z *Client) Call(ctx api_client.Client, req nw_client.RequestBuilder) (res es_response.Response) {
	l := ctx.Log()
	hReq, err := req.Build()
	if err != nil {
		l.Debug("Unable to make http request", esl.Error(err))
		return es_response_impl.NewNoResponse(err)
	}

	// Call
	hRes, latency, err := z.httpClient.Call(ctx.ClientHash(), req.Endpoint(), hReq)

	// Make response
	if err != nil {
		res = es_response_impl.NewTransportErrorHttpResponse(err, hRes)
	} else {
		res = es_response_impl.New(ctx, hRes)
	}

	// Capture
	cp := NewCapture(ctx)
	cp.WithResponse(req, hReq, res, err, latency.Nanoseconds())

	return res
}

type Capture interface {
	WithResponse(rb nw_client.RequestBuilder, req *http.Request, res es_response.Response, resErr error, latency int64)
}

func NewCapture(ctx api_client.Client) Capture {
	return &captureImpl{
		ctx: ctx,
	}
}

type captureImpl struct {
	ctx api_client.Client
}

type Record struct {
	Time    string          `json:"time"`
	Req     *nw_request.Req `json:"req"`
	Res     *Res            `json:"res"`
	Latency int64           `json:"latency"`
}

func (z Record) IsSuccess() bool {
	if z.Res == nil {
		return false
	}
	switch z.Res.ResponseCode / 100 {
	case 4, 5:
		return false
	}
	return true
}

type Res struct {
	ResponseCode    int               `json:"code"`
	ResponseProto   string            `json:"proto,omitempty"`
	ResponseBody    string            `json:"body,omitempty"`
	ResponseHeaders map[string]string `json:"headers"`
	ResponseJson    json.RawMessage   `json:"json,omitempty"`
	ResponseError   string            `json:"error,omitempty"`
	ContentLength   int64             `json:"content_length"`
}

func (z *Res) Apply(res es_response.Response, resErr error) {
	z.ResponseCode = res.Code()
	z.ResponseProto = res.Proto()

	if res.IsSuccess() {
		z.ContentLength = res.Success().ContentLength()
		if !res.IsTextContentType() {
			z.ResponseBody = ""
		} else if res.Success().IsFile() {
			z.ResponseBody = ""
		} else {
			z.ResponseBody = res.Success().BodyString()
		}
	} else {
		if resErr != nil {
			z.ResponseError = resErr.Error()
		}
		z.ContentLength = res.Alt().ContentLength()
		if !res.IsTextContentType() {
			z.ResponseBody = ""
		} else if res.Alt().IsFile() {
			z.ResponseBody = ""
		} else {
			z.ResponseBody = res.Alt().BodyString()
		}
	}

	z.ResponseHeaders = res.Headers()
}

func (z *captureImpl) WithResponse(rb nw_client.RequestBuilder, req *http.Request, res es_response.Response, resErr error, latency int64) {
	// request
	rq := nw_request.Req{}
	rq.Apply(z.ctx, rb, req)

	// response
	rs := Res{}
	rs.Apply(res, resErr)

	z.ctx.Capture().Debug("",
		esl.Any("req", rq),
		esl.Any("res", rs),
		esl.Int64("latency", latency),
	)
}

func (z *captureImpl) NoResponse(rb nw_client.RequestBuilder, req *http.Request, resErr error, latency int64) {
	// request
	rq := nw_request.Req{}
	rq.Apply(z.ctx, rb, req)

	// response
	rs := Res{}
	if resErr != nil {
		rs.ResponseError = resErr.Error()
	}

	z.ctx.Capture().Debug("",
		esl.Any("req", rq),
		esl.Any("res", rs),
		esl.Int64("latency", latency),
	)
}
