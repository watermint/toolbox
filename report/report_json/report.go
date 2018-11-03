package report_json

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
)

type JsonReport struct {
	logger       *zap.Logger
	ReportPath   string
	reportFile   *os.File
	reportWriter io.Writer
}

func (c *JsonReport) Init(logger *zap.Logger) error {
	c.logger = logger

	if c.ReportPath == "" {
		c.reportWriter = os.Stdout
	} else {
		if f, err := os.Create(c.ReportPath); err != nil {
			c.logger.Error(
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

func (c *JsonReport) Close() {
	if c.reportFile != nil {
		c.reportFile.Close()
		c.reportFile = nil
	}
}

func (c *JsonReport) Report(row interface{}) error {
	if c.reportWriter == nil {
		c.logger.Error("Report is not opened. Fallback to stdout")
		c.reportWriter = os.Stdout
	}

	m, err := json.Marshal(row)
	if err != nil {
		c.logger.Debug("marshal error", zap.Error(err), zap.Any("data", row))
		c.logger.Error("Unable to marshal report due to error", zap.Error(err))
		return err
	}
	_, err = fmt.Fprintln(c.reportWriter, string(m))
	if err != nil {
		c.logger.Debug("Unable to write data to the file. Fallback to log")
		c.logger.Warn(
			"Can't write data into the file",
			zap.String("report_path", c.ReportPath),
			zap.String("report", string(m)),
		)
		return errors.New("could not write data into the file")
	}
	return nil
}
