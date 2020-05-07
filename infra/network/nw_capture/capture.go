package nw_capture

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/log/stats/es_http"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/network/nw_client"
	"net/http"
	"strings"
)

func New(client nw_client.Http) nw_client.Rest {
	return &Client{httpClient: client}
}

type Client struct {
	httpClient nw_client.Http
}

func (z *Client) Call(ctx api_context.Context, req nw_client.RequestBuilder) (res es_response.Response) {
	l := ctx.Log()
	hReq, err := req.Build()
	if err != nil {
		l.Debug("Unable to make http request", es_log.Error(err))
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

	// Monitor stats
	es_http.Log(hReq, hRes)

	// Capture
	cp := NewCapture(ctx.Capture())
	cp.WithResponse(req, hReq, res, err, latency.Nanoseconds())

	return res
}

type Capture interface {
	WithResponse(rb nw_client.RequestBuilder, req *http.Request, res es_response.Response, resErr error, latency int64)
}

func NewCapture(cap es_log.Logger) Capture {
	return &captureImpl{
		capture: cap,
	}
}

type captureImpl struct {
	capture es_log.Logger
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

func (z *Req) Apply(rb nw_client.RequestBuilder, req *http.Request) {
	z.RequestMethod = req.Method
	z.RequestUrl = req.URL.String()
	z.RequestParam = rb.Param()
	z.RequestHeaders = make(map[string]string)
	z.ContentLength = req.ContentLength
	for k, v := range req.Header {
		v0 := v[0]
		// Anonymize token
		if k == api_request.ReqHeaderAuthorization {
			vv := strings.Split(v0, " ")
			z.RequestHeaders[k] = vv[0] + " <secret>"
		} else {
			z.RequestHeaders[k] = v0
		}
	}
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
	z.ContentLength = res.Success().ContentLength()
	if res.Success().IsFile() {
		z.ResponseBody = ""
	} else {
		z.ResponseBody = res.Success().BodyString()
	}
	if resErr != nil {
		z.ResponseError = resErr.Error()
	}
	z.ResponseHeaders = res.Headers()
}

func (z *captureImpl) WithResponse(rb nw_client.RequestBuilder, req *http.Request, res es_response.Response, resErr error, latency int64) {
	// request
	rq := Req{}
	rq.Apply(rb, req)

	// response
	rs := Res{}
	rs.Apply(res, resErr)

	z.capture.Debug("",
		es_log.Any("req", rq),
		es_log.Any("res", rs),
		es_log.Int64("latency", latency),
	)
}

func (z *captureImpl) NoResponse(rb nw_client.RequestBuilder, req *http.Request, resErr error, latency int64) {
	// request
	rq := Req{}
	rq.Apply(rb, req)

	// response
	rs := Res{}
	if resErr != nil {
		rs.ResponseError = resErr.Error()
	}

	z.capture.Debug("",
		es_log.Any("req", rq),
		es_log.Any("res", rs),
		es_log.Int64("latency", latency),
	)
}
