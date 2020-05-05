package es_http

import (
	"net/http"
	"testing"
	"time"
)

func TestTimeSeriesImpl_Log(t *testing.T) {
	p := 10 * time.Millisecond
	ts := &timeSeriesImpl{
		numUnit:   3,
		precision: p,
		history:   make(map[time.Time]Aggregator),
		latest:    &counterImpl{},
	}

	log := func(reqLen, resLen int64, a Aggregator) {
		req := &http.Request{ContentLength: reqLen}
		res := &http.Response{ContentLength: resLen}
		a.Log(req, res)
	}

	now := time.Now().Truncate(10 * time.Millisecond)
	time.Sleep(p - time.Now().Sub(now))
	{
		log(123, 456, ts)
		cc, ql, sl := ts.Summary()
		if cc != 1 {
			t.Error(cc)
		}
		if ql != 123 {
			t.Error(ql)
		}
		if sl != 456 {
			t.Error(sl)
		}
	}
	{
		log(123, 456, ts)
		cc, ql, sl := ts.Summary()
		if cc != 2 {
			t.Error(cc)
		}
		if ql != 123*2 {
			t.Error(ql)
		}
		if sl != 456*2 {
			t.Error(sl)
		}
	}
	time.Sleep(p)
	{
		log(123, 456, ts)
		cc, ql, sl := ts.Summary()
		if cc != 3 {
			t.Error(cc)
		}
		if ql != 123*3 {
			t.Error(ql)
		}
		if sl != 456*3 {
			t.Error(sl)
		}
	}
	time.Sleep(4 * p)
	{
		cc, ql, sl := ts.Summary()
		if cc != 0 {
			t.Error(cc)
		}
		if ql != 0 {
			t.Error(ql)
		}
		if sl != 0 {
			t.Error(sl)
		}

		// should truncate old history
		log(123, 456, ts)
		if len(ts.history) != 0 {
			t.Error("invalid map length", len(ts.history))
		}
	}
	{
		cpm, qps, sps := ts.Traffic()
		if cpm < 0 {
			t.Error(cpm)
		}
		if qps < 1 {
			t.Error(qps)
		}
		if sps < 1 {
			t.Error(sps)
		}
	}
}
