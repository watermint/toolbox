package nw_replay

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/essentials/http/es_response_impl"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_client"
	"github.com/watermint/toolbox/essentials/network/nw_request"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Code    int               `json:"code"`
	Proto   string            `json:"proto"`
	Body    string            `json:"body"`
	Error   string            `json:"error"`
	Headers map[string]string `json:"headers"`
}

func (z Response) Http() *http.Response {
	r := http.Response{
		Status:     http.StatusText(z.Code),
		StatusCode: z.Code,
		Proto:      z.Proto,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(z.Body))),
	}
	r.Header = http.Header{}
	for k, v := range z.Headers {
		r.Header.Add(k, v)
	}
	return &r
}

func NewSequentialReplay(r []Response) nw_client.Http {
	return &seqReplay{
		records: r,
	}
}

type seqReplay struct {
	index   int
	records []Response
}

func (z *seqReplay) Call(clientHash string, endpoint string, req *http.Request) (res *http.Response, latency time.Duration, err error) {
	if z.index < len(z.records) {
		fakeLatency := time.Duration(rand.Int63n(1000000) + 1000000)
		rec := z.records[z.index]
		z.index++

		if rec.Code < 0 {
			return nil, fakeLatency, errors.New(rec.Error)
		}
		res = rec.Http()
		return res, fakeLatency, nil
	}
	return nil, 0, qt_errors.ErrorMock
}

// hash -> Response mapping
func NewHashReplay(responses kv_storage.Storage) nw_client.Rest {
	return &hashReplay{
		responses: responses,
	}
}

var (
	ErrorNoReplayFound = errors.New("no replay response data found")
)

type hashReplay struct {
	responses kv_storage.Storage
}

func (z hashReplay) Call(ctx api_context.Context, builder nw_client.RequestBuilder) (res es_response.Response) {
	l := ctx.Log()

	hr, err := builder.Build()
	if err != nil {
		l.Debug("Unable to build the http request", esl.Error(err))
		return es_response_impl.NewNoResponse(err)
	}

	recReq := nw_request.Req{}
	recReq.Apply(ctx, builder, hr)

	l = l.With(esl.String("endpoint", hr.URL.String()))

	_ = z.responses.View(func(kvs kv_kvs.Kvs) error {
		capData, err := kvs.GetJson(recReq.RequestHash)
		if err != nil {
			l.Debug("No replay found for the hash", esl.String("hash", recReq.RequestHash))
			res = es_response_impl.NewNoResponse(ErrorNoReplayFound)
			return err
		}
		capResponses := make([]*Response, 0)
		if err := json.Unmarshal(capData, &capResponses); err != nil {
			l.Debug("Unable to unmarshal", esl.Error(err))
			res = es_response_impl.NewNoResponse(ErrorNoReplayFound)
			return err
		}
		var capRes *Response
		if x := len(capResponses); x < 1 {
			capRes = capResponses[0]
		} else {
			capRes = capResponses[rand.Intn(len(capResponses))]
		}
		res = es_response_impl.New(ctx, capRes.Http())
		return nil
	})
	return
}
