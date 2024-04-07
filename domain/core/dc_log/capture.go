package dc_log

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/montanaflynn/stats"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"gorm.io/gorm"
	"net/url"
	"strings"
	"time"
)

var (
	ErrorJobNotFound = errors.New("job not found")
)

const (
	DefaultTimeIntervalSeconds = 3600
)

type CaptureAggregator interface {
	AddByJob(job app_job.History) error
	AddById(jobId string) error
	AddByCliPath(cliPath string) error

	AggregatePopulation(handler func(r *Population)) error
	AggregateLatency(handler func(r *Latency)) error
	AggregateTimeSeries(handler func(r *TimeSeries)) error
}

type Population struct {
	Url        string  `json:"url"`
	Code       int     `json:"code"`
	Population int64   `json:"population"`
	Proportion float64 `json:"proportion"`
}

type Latency struct {
	Url        string  `json:"url"`
	Code       int     `json:"code"`
	Population int     `json:"population"`
	Mean       float64 `json:"mean"`
	Median     float64 `json:"median"`
	P50        float64 `json:"p_50"`
	P70        float64 `json:"p_70"`
	P90        float64 `json:"p_90"`
}

type TimeSeries struct {
	Time      string `json:"time"`
	Url       string `json:"url"`
	Code2xx   int64  `json:"code_2xx"`
	Code3xx   int64  `json:"code_3xx"`
	Code4xx   int64  `json:"code_4xx"`
	Code429   int64  `json:"code_429"`
	Code5xx   int64  `json:"code_5xx"`
	CodeOther int64  `json:"code_other"`
}

type UrlCode struct {
	ReqUrl  string
	ResCode int
}

type Record struct {
	Timestamp        time.Time `json:"timestamp"`
	ReqMethod        string    `json:"req_method"`
	ReqUrl           string    `json:"req_url"`
	ResCode          int       `json:"res_code"`
	ResContentLength int64     `json:"res_content_length,omitempty"`
	Latency          float64   `json:"latency,omitempty"`
}

func UrlFormat(reqUrl string, shorten bool) string {
	if !shorten {
		return reqUrl
	}

	parsed, err := url.Parse(reqUrl)
	if err != nil {
		return reqUrl
	}
	return parsed.Path
}

type CaptureAggregatorOpts struct {
	TimeIntervalSeconds int
	Shorten             bool
}

func (z CaptureAggregatorOpts) Apply(opts []CaptureAggregatorOpt) CaptureAggregatorOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return z.Apply(opts[1:])
	}
}

type CaptureAggregatorOpt func(o CaptureAggregatorOpts) CaptureAggregatorOpts

func OptTimeInterval(intervalSeconds int) CaptureAggregatorOpt {
	return func(o CaptureAggregatorOpts) CaptureAggregatorOpts {
		o.TimeIntervalSeconds = intervalSeconds
		return o
	}
}
func OptShorten(enabled bool) CaptureAggregatorOpt {
	return func(o CaptureAggregatorOpts) CaptureAggregatorOpts {
		o.Shorten = enabled
		return o
	}
}

func NewCaptureAggregator(db *gorm.DB, c app_control.Control, opts ...CaptureAggregatorOpt) CaptureAggregator {
	return &caImpl{
		db:  db,
		ctl: c,
		opts: CaptureAggregatorOpts{
			TimeIntervalSeconds: DefaultTimeIntervalSeconds,
		}.Apply(opts),
	}
}

type caImpl struct {
	db   *gorm.DB
	ctl  app_control.Control
	opts CaptureAggregatorOpts
}

func (z caImpl) AggregateTimeSeries(handler func(r *TimeSeries)) error {
	l := z.ctl.Log()
	urls := make([]string, 0)
	err := z.db.Model(&Record{}).
		Distinct("req_url").
		Find(&urls).Error
	if err != nil {
		l.Debug("Unable to list urls", esl.Error(err))
		return err
	}

	for _, u := range urls {
		timeBegin := time.Time{}
		err := z.db.Model(&Record{}).
			Select("timestamp").
			Where("req_url = ?", u).
			Order("timestamp").
			First(&timeBegin).Error
		if err != nil {
			l.Debug("Unable to find time begin", esl.Error(err))
			return err
		}

		timeEnd := time.Time{}
		err = z.db.Model(&Record{}).
			Select("timestamp").
			Where("req_url = ?", u).
			Order("timestamp DESC").
			First(&timeEnd).Error
		if err != nil {
			l.Debug("Unable to find time end", esl.Error(err))
			return err
		}

		timeBeginInterval := timeBegin.Truncate(time.Duration(z.opts.TimeIntervalSeconds) * time.Second)
		timeEndInterval := timeEnd.Truncate(time.Duration(z.opts.TimeIntervalSeconds) * time.Second).Add(time.Duration(z.opts.TimeIntervalSeconds) * time.Second)

		for t := timeBeginInterval; t.Before(timeEndInterval); t = t.Add(time.Duration(z.opts.TimeIntervalSeconds) * time.Second) {
			codes := make([]int, 0)

			err = z.db.Model(&Record{}).
				Select("res_code").
				Where("req_url = ? AND timestamp >= ? AND timestamp < ?", u, t, t.Add(time.Duration(z.opts.TimeIntervalSeconds)*time.Second)).
				Find(&codes).Error
			if err != nil {
				l.Debug("Unable to find time series", esl.Error(err))
				return err
			}

			code2xx := int64(0)
			code3xx := int64(0)
			code4xx := int64(0)
			code429 := int64(0)
			code5xx := int64(0)
			codeOther := int64(0)
			for _, c := range codes {
				if c == 429 {
					code429++
					continue
				}

				switch c / 100 {
				case 2:
					code2xx++
				case 3:
					code3xx++
				case 4:
					code4xx++
				case 5:
					code5xx++
				default:
					codeOther++
				}
			}

			handler(&TimeSeries{
				Time:      t.Format("2006-01-02 15:04:05"),
				Url:       UrlFormat(u, z.opts.Shorten),
				Code2xx:   code2xx,
				Code3xx:   code3xx,
				Code4xx:   code4xx,
				Code429:   code429,
				Code5xx:   code5xx,
				CodeOther: codeOther,
			})
		}
	}

	return nil
}

func (z caImpl) AggregatePopulation(handler func(r *Population)) error {
	l := z.ctl.Log()
	urls := make([]string, 0)
	err := z.db.Model(&Record{}).
		Distinct("req_url").
		Find(&urls).Error
	if err != nil {
		l.Debug("Unable to list urls", esl.Error(err))
		return err
	}

	urlPopulation := make(map[string]int64)
	for _, u := range urls {
		var population int64
		err := z.db.Model(&Record{}).
			Where("req_url = ?", u).
			Count(&population).Error
		if err != nil {
			l.Debug("Unable to count population", esl.Error(err))
			return err
		}
		urlPopulation[u] = population
	}

	urlCodes := make([]UrlCode, 0)
	err = z.db.Model(&Record{}).
		Distinct("req_url", "res_code").
		Find(&urlCodes).Error
	if err != nil {
		l.Debug("Unable to list url codes", esl.Error(err))
		return err
	}

	for _, mc := range urlCodes {
		population, ok := urlPopulation[mc.ReqUrl]
		if !ok {
			l.Debug("Unable to find population", esl.String("url", mc.ReqUrl))
			return err
		}
		var urlCodePopulation int64
		err = z.db.Model(&Record{}).
			Where("req_url = ? AND res_code = ?", mc.ReqUrl, mc.ResCode).
			Count(&urlCodePopulation).Error
		if err != nil {
			l.Debug("Unable to count url code population", esl.Error(err))
			return err
		}

		handler(&Population{
			Url:        UrlFormat(mc.ReqUrl, z.opts.Shorten),
			Code:       mc.ResCode,
			Population: urlCodePopulation,
			Proportion: float64(urlCodePopulation) / float64(population),
		})
	}
	return nil
}

func (z caImpl) AggregateLatency(handler func(r *Latency)) error {
	l := z.ctl.Log()
	urlCodes := make([]UrlCode, 0)
	err := z.db.Model(&Record{}).
		Distinct("req_url", "res_code").
		Find(&urlCodes).Error
	if err != nil {
		l.Debug("Unable to list url codes", esl.Error(err))
		return err
	}

	for _, mc := range urlCodes {
		latencies := make([]float64, 0)
		err := z.db.Model(&Record{}).
			Select("latency").
			Where("req_url = ? AND res_code = ?", mc.ReqUrl, mc.ResCode).
			Find(&latencies).Error
		if err != nil {
			l.Debug("Unable to list latencies", esl.Error(err))
			return err
		}

		// calculate percentiles
		mean, err := stats.Mean(latencies)
		if err != nil {
			l.Debug("Unable to calculate mean", esl.Error(err))
			return err
		}
		median, err := stats.Median(latencies)
		if err != nil {
			l.Debug("Unable to calculate median", esl.Error(err))
			return err
		}
		p50, err := stats.Percentile(latencies, 50)
		if err != nil {
			l.Debug("Unable to calculate p50", esl.Error(err))
			return err
		}
		p70, err := stats.Percentile(latencies, 70)
		if err != nil {
			l.Debug("Unable to calculate p70", esl.Error(err))
			return err
		}
		p90, err := stats.Percentile(latencies, 90)
		if err != nil {
			l.Debug("Unable to calculate p90", esl.Error(err))
			return err
		}

		handler(&Latency{
			Url:        UrlFormat(mc.ReqUrl, z.opts.Shorten),
			Code:       mc.ResCode,
			Population: len(latencies),
			Mean:       mean,
			Median:     median,
			P50:        p50,
			P70:        p70,
			P90:        p90,
		})
	}
	return nil
}

func (z caImpl) AddByJob(job app_job.History) error {
	l := z.ctl.Log().With(esl.String("jobPath", job.JobPath()))
	if err := z.db.AutoMigrate(&Record{}); err != nil {
		l.Debug("Unable to migrate table", esl.Error(err))
		return err
	}
	logs, err := job.Logs()
	if err != nil {
		l.Debug("Unable to retrieve logs", esl.Error(err))
		return err
	}

	for _, lf := range logs {
		if lf.Type() != app_job.LogFileTypeCapture {
			continue
		}
		l.Debug("Processing log", esl.String("logPath", lf.Path()))
		// read log file
		var capFile bytes.Buffer
		if err := lf.CopyTo(&capFile); err != nil {
			l.Debug("Unable to copy log file", esl.Error(err))
			return err
		}

		// parse log file
		lines := strings.Split(capFile.String(), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			// parse line
			capRecord := nw_capture.Record{}
			if err := json.Unmarshal([]byte(line), &capRecord); err != nil {
				l.Debug("Unable to parse line", esl.Error(err))
				return err
			}

			ts, err := time.Parse("2006-01-02T15:04:05.999Z0700", capRecord.Time)
			if err != nil {
				l.Debug("Unable to parse timestamp", esl.Error(err))
				return err
			}

			// save to db
			rec := Record{
				Timestamp: ts,
				// nanoseconds to seconds
				Latency: float64(capRecord.Latency) / 1e9,
			}

			if capRecord.Req != nil {
				rec.ReqMethod = capRecord.Req.RequestMethod
				rec.ReqUrl = capRecord.Req.RequestUrl
			}

			if capRecord.Res != nil {
				rec.ResCode = capRecord.Res.ResponseCode
				rec.ResContentLength = capRecord.Res.ContentLength
			}

			if err := z.db.Create(&rec).Error; err != nil {
				l.Debug("Unable to save record", esl.Error(err))
				return err
			}
		}
	}

	return nil
}

func (z caImpl) AddByCliPath(cliPath string) error {
	l := z.ctl.Log().With(esl.String("cliPath", cliPath))
	historian := app_job_impl.NewHistorian(z.ctl.Workspace())
	jobs, err := historian.Histories()
	if err != nil {
		l.Debug("Unable to list jobs", esl.Error(err))
		return err
	}
	found := false
	for _, h := range jobs {
		hcp := h.StartLog().Name
		if hcp == cliPath {
			if err := z.AddByJob(h); err != nil {
				l.Debug("Unable to add job by path", esl.Error(err))
				return err
			} else {
				found = true
			}
		}
	}
	if found {
		return nil
	} else {
		return ErrorJobNotFound
	}
}

func (z caImpl) AddById(jobId string) error {
	l := z.ctl.Log().With(esl.String("jobId", jobId))
	historian := app_job_impl.NewHistorian(z.ctl.Workspace())
	jobs, err := historian.Histories()
	if err != nil {
		l.Debug("Unable to list jobs", esl.Error(err))
		return err
	}
	for _, h := range jobs {
		if h.JobId() == jobId {
			return z.AddByJob(h)
		}
	}
	return ErrorJobNotFound
}
