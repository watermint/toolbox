package api_capture

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/app/app_report/app_report_json"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

const (
	valuePathCapture = "api_capture.Capture"
)

type Capture interface {
	Rpc(req api_rpc.Request, res api_rpc.Response, resErr error)
}

func Current() Capture {
	ec := app.Root()
	if !ec.IsDebug() {
		return mockImpl{}
	}

	if c, e := ec.GetValue(valuePathCapture); e {
		switch ca := c.(type) {
		case Capture:
			return ca
		}
	}

	storage := app_report_json.JsonReport{
		ReportPath:    filepath.Join(ec.JobsPath(), "capture"),
		DefaultWriter: os.Stdout,
	}
	storage.Init(ec)
	ca := &captureImpl{
		storage: &storage,
	}
	ec.SetValue(valuePathCapture, ca)
	return ca
}

type Record struct {
	Timestamp      time.Time         `json:"timestamp"`
	RequestMethod  string            `json:"request_method"`
	RequestUrl     string            `json:"request_url"`
	RequestParam   string            `json:"request_param,omitempty"`
	RequestHeaders map[string]string `json:"request_headers"`
	ResponseCode   int               `json:"response_code"`
	ResponseBody   string            `json:"response_body"`
	ResponseError  string            `json:"response_error,omitempty"`
}

type mockImpl struct {
}

func (mockImpl) Rpc(req api_rpc.Request, res api_rpc.Response, resErr error) {
	// ignore
}

var (
	tokenMatcher = regexp.MustCompile(`\w`)
)

type captureImpl struct {
	storage app_report.Report
}

func (z *captureImpl) Rpc(req api_rpc.Request, res api_rpc.Response, resErr error) {
	rec := Record{
		Timestamp: time.Now(),
	}

	// request
	rec.RequestMethod = req.Method()
	rec.RequestUrl = req.Url()
	rec.RequestParam = req.Param()
	headers := make(map[string]string)
	for k, v := range req.Headers() {
		// Anonymize token
		if k == api_rpc.ReqHeaderAuthorization {
			headers[k] = "Bearer <secret>"
		} else {
			headers[k] = v
		}
	}
	rec.RequestHeaders = headers

	// response
	rec.ResponseCode = res.StatusCode()
	resBody, _ := res.Body()
	rec.ResponseBody = resBody
	if resErr != nil {
		rec.ResponseError = resErr.Error()
	}

	z.storage.Report(rec)
}
