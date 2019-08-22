package app_io

import (
	"errors"
	"github.com/watermint/toolbox/infra/util/ut_encoding"
	app2 "github.com/watermint/toolbox/legacy/app"
	"io"
	"os"
)

type CsvLoader interface {
	OnRow(rowLoader func(cols []string) error) CsvLoader
	Load() error
}

func NewCsvLoader(ec *app2.ExecContext, filepath string) CsvLoader {
	return &csvLoaderImpl{
		ec:       ec,
		filePath: filepath,
	}
}

type csvLoaderImpl struct {
	ec        *app2.ExecContext
	filePath  string
	rowLoader func(cols []string) error
}

func (z *csvLoaderImpl) OnRow(rowLoader func(row []string) error) CsvLoader {
	z.rowLoader = rowLoader
	return z
}

func (z *csvLoaderImpl) Load() error {
	if z.rowLoader == nil {
		return errors.New("no rowLoader")
	}
	if z.filePath == "" {
		z.ec.Msg("app.common.io.csv_loader.err.no_filepath").TellError()
		return errors.New("please specify csv file path")
	}

	f, err := os.Open(z.filePath)
	if err != nil {
		z.ec.Msg("app.common.io.csv_loader.err.cant_read").WithData(struct {
			File string
		}{
			File: z.filePath,
		}).TellError()
		return err
	}
	defer f.Close()
	csv := ut_encoding.NewBomAwareCsvReader(f)

	for {
		cols, err := csv.Read()
		if err == io.EOF {
			return nil
		}
		if len(cols) < 1 {
			continue
		}

		if err := z.rowLoader(cols); err != nil {
			return err
		}
	}
}
