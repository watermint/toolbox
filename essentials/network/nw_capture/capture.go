package nw_capture

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
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
	cp := NewCapture(ctx.Capture())
	cp.WithResponse(req, hReq, res, err, latency.Nanoseconds())

	return res
}

type Capture interface {
	WithResponse(rb nw_client.RequestBuilder, req *http.Request, res es_response.Response, resErr error, latency int64)
}

func NewCapture(cap esl.Logger) Capture {
	return &captureImpl{
		capture: cap,
	}
}

type captureImpl struct {
	capture esl.Logger
}

type Record struct {
	Time    string `json:"time"`
	Req     *Req   `json:"req"`
	Res     *Res   `json:"res"`
	Latency int64  `json:"latency"`
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

type Req struct {
	RequestMethod  string            `json:"method"`
	RequestUrl     string            `json:"url"`
	RequestParam   string            `json:"param,omitempty"`
	RequestHeaders map[string]string `json:"headers"`
	ContentLength  int64             `json:"content_length"`
	RequestHash    string            `json:"hash"`
}

type HashSeed struct {
	Url    string            `json:"u"`
	Method string            `json:"m"`
	Param  string            `json:"p"`
	Length int64             `json:"l"`
	Header map[string]string `json:"h"`
}

func (z HashSeed) Hash() string {
	seed := "u" + z.Url +
		"m" + z.Method +
		"p" + z.Param +
		"l" + fmt.Sprintf("%x", z.Length)
	for k, v := range z.Header {
		seed += "h" + k + ":" + v
	}
	h := sha256.Sum256([]byte(seed))
	return base64.RawStdEncoding.EncodeToString(h[:])
}

func (z *Req) Apply(rb nw_client.RequestBuilder, req *http.Request) {
	url := req.URL.String()
	param := rb.Param()
	z.RequestHash = HashSeed{
		Url:    url,
		Method: req.Method,
		Param:  param,
		Length: req.ContentLength,
		Header: z.RequestHeaders,
	}.Hash()

	if ruf, ok := rb.(nw_client.RequestUrlFilter); ok {
		url = ruf.FilterUrl(url)
	}
	z.RequestMethod = req.Method
	z.RequestUrl = url
	z.RequestParam = param
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

	if res.IsSuccess() {
		z.ContentLength = res.Success().ContentLength()
		if res.Success().IsFile() {
			z.ResponseBody = ""
		} else {
			z.ResponseBody = res.Success().BodyString()
		}
	} else {
		if resErr != nil {
			z.ResponseError = resErr.Error()
		}
		z.ContentLength = res.Alt().ContentLength()
		if res.Alt().IsFile() {
			z.ResponseBody = ""
		} else {
			z.ResponseBody = res.Alt().BodyString()
		}
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
		esl.Any("req", rq),
		esl.Any("res", rs),
		esl.Int64("latency", latency),
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
		esl.Any("req", rq),
		esl.Any("res", rs),
		esl.Int64("latency", latency),
	)
}
