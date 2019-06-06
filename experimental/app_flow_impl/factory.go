package app_flow_impl

import (
	"errors"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_msg"
	"os"
	"strings"
)

type Factory struct {
	FilePath string
}

func (z *Factory) EachRow(ctl app_control.Control, exec func(cols []string, rowIndex int) error) error {
	if z.FilePath == "" {
		ctl.UI().Error("flow.error.no_file_path")
		return errors.New("please specify data file")
	}
	st, err := os.Stat(z.FilePath)
	switch {
	case err != nil && os.IsNotExist(err):
		ctl.UI().Error("flow.error.file_not_found", app_msg.P("Path", z.FilePath))
		return err
	case err != nil:
		ctl.UI().Error("flow.error.unable_to_read",
			app_msg.P("Path", z.FilePath),
			app_msg.P("Error", err),
		)
		return err
	case st.IsDir():
		ctl.UI().Error("flow.error.not_a_file", app_msg.P("Path", z.FilePath))
		return errors.New("not a file")
	}

	switch {
	case strings.HasSuffix(strings.ToLower(z.FilePath), ".csv"):
		ctl.Log().Debug("Process with CSV")
		c, err := NewCsv(z.FilePath, ctl)
		if err != nil {
			return err
		}
		return c.EachRow(ctl, exec)
	}

	return errors.New("unsupported extension")
}
