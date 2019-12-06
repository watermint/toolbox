package nw_monitor

import (
	"fmt"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Monitor interface {
	// All implementation should aware req/res might be nil
	Log(req *http.Request, res *http.Response)
}

type Aggregator interface {
	Monitor

	Summary() (callCount, reqContentLen, resContentLen int64)
}

type Averager interface {
	Aggregator

	Traffic() (callPerMin, reqBps, resBps int64)
}

const (
	monitorIntervalMins = 5
	reportInterval      = monitorIntervalMins * time.Minute
)

var (
	mon   = newTimeSeries(monitorIntervalMins)
	total = &counterImpl{}
)

func Log(req *http.Request, res *http.Response) {
	mon.Log(req, res)
	total.Log(req, res)
}

func LaunchReporting(ui app_ui.UI, l *zap.Logger) {
	go func() {
		for {
			time.Sleep(reportInterval)
			cpm, qps, sps := mon.Traffic()
			ui.Info("run.network.progress.stats", app_msg.P{
				"CallPerMin": cpm,
				"ReqKps":     fmt.Sprintf("%.2f", float64(qps)/1024.0),
				"ResKps":     fmt.Sprintf("%.2f", float64(sps)/1024.0),
			})

			tcc, tql, tsl := total.Summary()
			cc, ql, sl := mon.Summary()
			l.Debug("Network stats",
				zap.Int64("CallPerMin", cpm),
				zap.Int64("ReqBytesPerSec", qps),
				zap.Int64("ResBytesPerSec", sps),
				zap.Int64("IntervalCallCount", cc),
				zap.Int64("IntervalReqContentLen", ql),
				zap.Int64("IntervalResContentLen", sl),
				zap.Int64("TotalCallCount", tcc),
				zap.Int64("TotalReqContentLen", tql),
				zap.Int64("TotalResContentLen", tsl),
			)
		}
	}()
}

func DumpStats(l *zap.Logger) {
	cc, ql, sl := mon.Summary()
	l.Debug("Network summary",
		zap.Int64("CallCount", cc),
		zap.Int64("ReqContentLength", ql),
		zap.Int64("ResContentLength", sl),
	)
}

type counterImpl struct {
	callCount        int64
	reqContentLength int64
	resContentLength int64
}

func (z *counterImpl) Summary() (callCount, reqContentLen, resContentLen int64) {
	return z.callCount, z.reqContentLength, z.resContentLength
}

func (z *counterImpl) Log(req *http.Request, res *http.Response) {
	atomic.AddInt64(&z.callCount, 1)
	if req != nil && req.ContentLength > 0 {
		atomic.AddInt64(&z.reqContentLength, req.ContentLength)
	}
	if res != nil && res.ContentLength > 0 {
		atomic.AddInt64(&z.resContentLength, res.ContentLength)
	}
}

func newTimeSeries(numMinutes int) Averager {
	return &timeSeriesImpl{
		numUnit:   numMinutes,
		precision: time.Minute,
		history:   make(map[time.Time]Aggregator),
		latest:    &counterImpl{},
	}
}

type timeSeriesImpl struct {
	numUnit     int
	precision   time.Duration
	history     map[time.Time]Aggregator
	latest      Aggregator
	latestTime  time.Time
	latestMutex sync.Mutex
}

func (z *timeSeriesImpl) Traffic() (callPerMin, reqBps, resBps int64) {
	z.latestMutex.Lock()
	defer z.latestMutex.Unlock()

	var tcc, tql, tsl int64
	var dur time.Duration

	t := time.Now().Truncate(z.precision)
	threshold := t.Add(time.Duration(-z.numUnit) * z.precision)
	for k, ag := range z.history {
		if threshold.Before(k) {
			cc, ql, sl := ag.Summary()
			tcc += cc
			tql += ql
			tsl += sl
			dur += z.precision
		}
	}
	if z.latestTime.Equal(t) {
		cc, ql, sl := z.latest.Summary()
		tcc += cc
		tql += ql
		tsl += sl
		dur += z.precision
	}

	if dur == 0 {
		return 0, 0, 0
	}

	return tcc * int64(time.Minute) / int64(dur),
		tql * int64(time.Second) / int64(dur),
		tsl * int64(time.Second) / int64(dur)
}

func (z *timeSeriesImpl) Log(req *http.Request, res *http.Response) {
	z.latestMutex.Lock()
	defer z.latestMutex.Unlock()

	t := time.Now().Truncate(z.precision)

	if z.latestTime.Equal(t) {
		z.latest.Log(req, res)
	} else {
		z.history[z.latestTime] = z.latest
		z.latest = &counterImpl{}
		z.latest.Log(req, res)
		z.latestTime = t

		// remove old history
		threshold := t.Add(time.Duration(-z.numUnit) * z.precision)
		for k := range z.history {
			if threshold.After(k) {
				delete(z.history, k)
			}
		}
	}
}

func (z *timeSeriesImpl) Summary() (callCount, reqContentLen, resContentLen int64) {
	z.latestMutex.Lock()
	defer z.latestMutex.Unlock()

	t := time.Now().Truncate(z.precision)
	threshold := t.Add(time.Duration(-z.numUnit) * z.precision)
	for k, ag := range z.history {
		if threshold.Before(k) {
			cc, ql, sl := ag.Summary()
			callCount += cc
			reqContentLen += ql
			resContentLen += sl
		}
	}
	if z.latestTime.Equal(t) {
		cc, ql, sl := z.latest.Summary()
		callCount += cc
		reqContentLen += ql
		resContentLen += sl
	}
	return
}

type monitorImpl struct {
}
