package report_csv

import (
	"go.uber.org/zap"
	"io"
	"os"
)

type CsvReport struct {
	logger       *zap.Logger
	ReportPath   string
	OmitHeader   bool
	reportFile   *os.File
	reportWriter io.Writer
}

func (r *CsvReport) Open(logger *zap.Logger) error {
	r.logger = logger

	if r.ReportPath == "" {
		r.reportWriter = os.Stdout
	} else {
		if f, err := os.Create(r.ReportPath); err != nil {
			r.logger.Error(
				"unable to open report file. Fallback to STDOUT",
				zap.String("file", r.ReportPath),
			)
			r.reportWriter = os.Stdout
			return err
		} else {
			r.reportFile = f
			r.reportWriter = f
		}
	}
	return nil
}

func (r *CsvReport) Report(row interface{}) error {
	return nil
}

func (r *CsvReport) Close() {
	if r.reportFile != nil {
		r.reportFile.Close()
		r.reportFile = nil
	}
}
