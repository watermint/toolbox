package cmdlet

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/watermint/toolbox/dbx_api"
	"go.uber.org/zap"
	"io"
	"os"
)

type Report struct {
	cmd          Commandlet
	ReportPath   string
	ReportFormat string
	reportFile   *os.File
	reportWriter io.Writer
}

func (c *Report) FlagConfig(f *flag.FlagSet) {
	descReportPath := "Output file path of the report (default: STDOUT)"
	f.StringVar(&c.ReportPath, "report-path", "", descReportPath)

	//descReportFormat := "Output file format (jsonl) (default: jsonl)"
	//f.StringVar(&c.ReportFormat, "report-format", "jsonl", descReportFormat)

	// Make supported format only for JSONL for while
	c.ReportFormat = "jsonl"
}

func (c *Report) Open(cmd Commandlet) error {
	c.cmd = cmd
	if c.ReportPath == "" {
		c.reportWriter = os.Stdout
	} else {
		if f, err := os.Create(c.ReportPath); err != nil {
			cmd.Log().Error(
				"unable to open report file. Fallback to STDOUT",
				zap.String("file", c.ReportPath),
			)
			c.reportWriter = os.Stdout
			return err
		} else {
			c.reportFile = f
			c.reportWriter = f
		}
	}
	return nil
}

func (c *Report) Close() {
	if c.reportFile != nil {
		c.reportFile.Close()
		c.reportFile = nil
	}
}

func (c *Report) Report(row interface{}) {
	if c.reportWriter == nil {
		c.cmd.Log().Error("Report is not opened. Fallback to stdout")
		c.reportWriter = os.Stdout
	}
	r, err := json.Marshal(row)
	if err != nil {
		c.cmd.Log().Debug("marshal error", zap.Error(err), zap.Any("data", row))
		c.cmd.Log().Error("Unable to marshal report due to error", zap.Error(err))
		return
	}
	_, err = fmt.Fprintln(c.reportWriter, string(r))
	if err != nil {
		c.cmd.Log().Debug("Unable to write data to the file. Fallback to log")
		c.cmd.Log().Warn(
			"Can't write data into the file",
			zap.String("report_path", c.ReportPath),
			zap.String("report", string(r)),
		)
		an := dbx_api.ErrorAnnotation{
			ErrorType: dbx_api.ErrorOperationFailed,
			Error:     err,
		}

		c.cmd.Log().Debug("Report error to default error handler")
		c.cmd.DefaultErrorHandler(an)
	}
}
