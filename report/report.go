package report

import (
	"errors"
	"flag"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/report/report_csv"
	"github.com/watermint/toolbox/report/report_json"
	"go.uber.org/zap"
	"io"
	"os"
)

type Report interface {
	Init(ec *app.ExecContext) error
	Close()
	Report(row interface{}) error
}

type Factory struct {
	ExecContext   *app.ExecContext
	report        Report
	DefaultWriter io.Writer
	ReportHeader  bool
	ReportUseBom  bool
	ReportPath    string
	ReportFormat  string
}

func (z *Factory) FlagConfig(f *flag.FlagSet) {
	descReportPath := z.ExecContext.Msg("report.common.flag.report_path").T()
	f.StringVar(&z.ReportPath, "report-path", "", descReportPath)

	descReportFormat := z.ExecContext.Msg("report.common.flag.report_format").T()
	f.StringVar(&z.ReportFormat, "report-format", "json", descReportFormat)

	descUseBom := z.ExecContext.Msg("report.common.flag.use_bom").T()
	f.BoolVar(&z.ReportUseBom, "report-usebom", false, descUseBom)

	descReportHeader := z.ExecContext.Msg("report.common.flag.with_header").T()
	f.BoolVar(&z.ReportHeader, "report-header", true, descReportHeader)
}

func (z *Factory) Init(ec *app.ExecContext) error {
	if z.DefaultWriter == nil {
		z.DefaultWriter = os.Stdout
	}

	switch z.ReportFormat {
	case "csv":
		z.report = &report_csv.CsvReport{
			DefaultWriter: z.DefaultWriter,
			ReportPath:    z.ReportPath,
			ReportHeader:  z.ReportHeader,
			ReportUseBom:  z.ReportUseBom,
		}
		return z.report.Init(ec)

	case "json":
		z.report = &report_json.JsonReport{
			DefaultWriter: z.DefaultWriter,
			ReportPath:    z.ReportPath,
		}
		return z.report.Init(ec)

	default:
		z.ExecContext.Log().Error(
			"unsupported report format",
			zap.String("specified_format", z.ReportFormat),
		)
		return errors.New("unsupported format")
	}
}

func (z *Factory) Report(row interface{}) error {
	if z.report == nil {
		z.ExecContext.Log().Fatal("open report before write")
		return errors.New("report was not opened")
	}

	return z.report.Report(row)
}

func (z *Factory) Close() {
	if z.report == nil {
		z.ExecContext.Log().Debug("Report already closed")
		return
	}
	z.report.Close()
}
