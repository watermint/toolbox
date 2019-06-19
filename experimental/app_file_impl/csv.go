package app_file_impl

import (
	"encoding/csv"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_file"
	"github.com/watermint/toolbox/experimental/app_msg"
	"io"
	"os"
)

func NewCsv(filePath string, ctl app_control.Control) (app_file.RowDataFile, error) {
	f, err := os.Open(filePath)
	if err != nil {
		ctl.UI().Error("flow.error.unable_to_read",
			app_msg.P("Path", filePath),
			app_msg.P("Error", err),
		)
		return nil, err
	}
	r := csv.NewReader(f)
	rdf := &CsvDataFile{
		Ctl:      ctl,
		FilePath: filePath,
		File:     f,
		Reader:   r,
	}
	return rdf, nil
}

type CsvDataFile struct {
	Ctl      app_control.Control
	FilePath string
	File     *os.File
	Reader   *csv.Reader
}

func (z *CsvDataFile) EachRow(ctl app_control.Control, exec func(cols []string, rowIndex int) error) error {
	defer z.File.Close()
	for i := 0; ; i++ {
		cols, err := z.Reader.Read()
		switch {
		case err == io.EOF:
			return nil

		case err != nil:
			ctl.UI().Error("flow.error.unable_to_read",
				app_msg.P("Path", z.FilePath),
				app_msg.P("Error", err),
			)
			return err

		default:
			if err := exec(cols, i); err != nil {
				return err
			}
		}
	}
}
