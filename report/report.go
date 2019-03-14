package report

import (
	"errors"
	"flag"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/report/report_csv"
	"github.com/watermint/toolbox/report/report_json"
	"github.com/watermint/toolbox/report/report_xlsx"
	"io"
	"os"
	"path/filepath"
)

type Report interface {
	Init(ec *app.ExecContext) error
	Close()
	Report(row interface{}) error
}

type Factory struct {
	ExecContext   *app.ExecContext
	reports       []Report
	DefaultWriter io.Writer
	ReportHeader  bool
	ReportUseBom  bool
	ReportPath    string
	ReportFormat  string
}

func (z *Factory) FlagConfig(f *flag.FlagSet) {
	descReportPath := z.ExecContext.Msg("report.common.flag.report_path").T()
	f.StringVar(&z.ReportPath, "report-path", filepath.Join(z.ExecContext.JobsPath(), "reports"), descReportPath)

	//descReportFormat := z.ExecContext.Msg("report.common.flag.report_format").T()
	//f.StringVar(&z.ReportFormat, "report-format", "json", descReportFormat)

	//descUseBom := z.ExecContext.Msg("report.common.flag.use_bom").T()
	//f.BoolVar(&z.ReportUseBom, "report-usebom", false, descUseBom)

	//descReportHeader := z.ExecContext.Msg("report.common.flag.with_header").T()
	//f.BoolVar(&z.ReportHeader, "report-header", true, descReportHeader)
}

func (z *Factory) Init(ec *app.ExecContext) error {
	if z.DefaultWriter == nil {
		z.DefaultWriter = os.Stdout
	}
	if z.reports == nil {
		z.reports = make([]Report, 0)
		z.reports = append(z.reports, &report_json.JsonReport{
			DefaultWriter: z.DefaultWriter,
			ReportPath:    "",
		})
		z.reports = append(z.reports, &report_json.JsonReport{
			DefaultWriter: z.DefaultWriter,
			ReportPath:    z.ReportPath,
		})
		z.reports = append(z.reports, &report_csv.CsvReport{
			DefaultWriter: z.DefaultWriter,
			ReportPath:    z.ReportPath,
			ReportHeader:  true,
			ReportUseBom:  false,
		})
		z.reports = append(z.reports, &report_xlsx.XlsxReport{
			ReportPath: z.ReportPath,
		})

		for _, r := range z.reports {
			if err := r.Init(ec); err != nil {
				return err
			}
		}
	}
	return nil
}

func (z *Factory) Report(row interface{}) error {
	if z.reports == nil {
		z.ExecContext.Log().Fatal("open report before write")
		return errors.New("report was not opened")
	}

	for _, r := range z.reports {
		if err := r.Report(row); err != nil {
			return err
		}
	}
	return nil
}

func (z *Factory) Close() {
	if z.reports == nil {
		z.ExecContext.Log().Debug("Report already closed")
		return
	}
	for _, r := range z.reports {
		r.Close()
	}

	z.ExecContext.Msg("report.common.done.tell_location").WithData(struct {
		Path string
	}{
		z.ReportPath,
	}).TellSuccess()
}
