package report_csv

import (
	"encoding/csv"
	"github.com/watermint/toolbox/report/report_column"
	"go.uber.org/zap"
	"os"
)

type CsvReport struct {
	logger       *zap.Logger
	ReportPath   string
	ReportHeader bool
	reportFile   *os.File
	reportWriter *csv.Writer
	marshaller   *report_column.ColumnZ
}

func (r *CsvReport) Open(logger *zap.Logger) error {
	r.logger = logger
	r.marshaller = &report_column.ColumnZ{
		Log: logger,
	}

	if r.ReportPath == "" {
		r.reportWriter = csv.NewWriter(os.Stdout)
	} else {
		if f, err := os.Create(r.ReportPath); err != nil {
			r.logger.Error(
				"unable to open report file. Fallback to STDOUT",
				zap.String("file", r.ReportPath),
			)
			r.reportWriter = csv.NewWriter(os.Stdout)
			return err
		} else {
			r.reportFile = f
			r.reportWriter = csv.NewWriter(f)
		}
	}
	return nil
}

func (r *CsvReport) Report(row interface{}) error {
	//	outHeader := r.ReportHeader
	headers := r.marshaller.Header(row)
	r.reportWriter.Write(headers)

	cols := r.marshaller.Values(headers, row)
	r.reportWriter.Write(cols)
	r.reportWriter.Flush()

	return nil
}

func (r *CsvReport) Close() {
	if r.reportFile != nil {
		r.reportFile.Close()
		r.reportFile = nil
	}
}
