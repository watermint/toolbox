package cmdlet

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/workflow"
	"github.com/watermint/toolbox/workflow/report"
)

type Report struct {
	ReportPath   string
	ReportFormat string

	DataHeaders []string
}

func (c *Report) FlagConfig(f *flag.FlagSet) {
	descReportPath := "Output file path of the report (default: STDOUT)"
	f.StringVar(&c.ReportPath, "report-path", "", descReportPath)

	descReportFormat := "Output file format (csv|jsonl) (default: jsonl)"
	f.StringVar(&c.ReportFormat, "report-format", "jsonl", descReportFormat)
}

func (c *Report) ReportStages() (reportTask string, stages []workflow.Worker, err error) {
	reportTask = report.WORKER_REPORT_JSONL
	switch c.ReportFormat {
	case "jsonl":
		reportTask = report.WORKER_REPORT_JSONL

	case "csv":
		reportTask = report.WORKER_REPORT_CSV

	default:
		seelog.Warnf("Unsupported report format [%s]", c.ReportFormat)
		return "", nil, err
	}

	return reportTask,
		[]workflow.Worker{
			&report.WorkerReportJsonl{
				ReportPath: c.ReportPath,
			},
			&report.WorkerReportCsv{
				ReportPath:  c.ReportPath,
				DataHeaders: c.DataHeaders,
			},
		}, nil

}
