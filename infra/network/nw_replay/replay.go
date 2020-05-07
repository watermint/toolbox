package nw_replay

import (
	"bytes"
	"errors"
	"github.com/watermint/toolbox/infra/network/nw_client"
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

func NewReplay(r []Response) nw_client.Http {
	return &Replay{
		records: r,
	}
}

type Replay struct {
	index   int
	records []Response
}

func (z *Replay) Call(clientHash string, endpoint string, req *http.Request) (res *http.Response, latency time.Duration, err error) {
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
