package cmdlet

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"io"
	"os"
)

type Report struct {
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

func (c *Report) Open() error {
	if c.ReportPath == "" {
		c.reportWriter = os.Stdout
	} else {
		if f, err := os.Open(c.ReportPath); err != nil {
			seelog.Errorf("Unable to open report file [%s]. Fallback to STDOUT", c.ReportPath)
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
		seelog.Errorf("Report is not opened. Fallback to stdout")
		c.reportWriter = os.Stdout
	}
	r, err := json.Marshal(row)
	if err != nil {
		seelog.Debugf("Marshal error[%s] Data[%v]", err, row)
		seelog.Warnf("Unable to marshal report due to error[%s].", err)
		return
	}
	fmt.Fprintln(c.reportWriter, string(r))
}
