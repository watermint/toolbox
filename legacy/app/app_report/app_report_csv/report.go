package app_report_csv

import (
	"encoding/csv"
	"github.com/watermint/toolbox/infra/util/ut_encoding"
	app2 "github.com/watermint/toolbox/legacy/app"
	app_report_column2 "github.com/watermint/toolbox/legacy/app/app_report/app_report_column"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
)

type CsvReport struct {
	ec            *app2.ExecContext
	ReportUseBom  bool
	ReportPath    string
	ReportHeader  bool
	DefaultWriter io.Writer
	files         map[string]*os.File
	writers       map[string]*csv.Writer
	parsers       map[string]app_report_column2.Row
}

func (z *CsvReport) prepare(row interface{}) (f *os.File, w *csv.Writer, p app_report_column2.Row, err error) {
	name := app_report_column2.RowName(row)
	if zp, ok := z.parsers[name]; ok {
		p = zp
	}
	if zw, ok := z.writers[name]; ok {
		w = zw
	}
	if zf, ok := z.files[name]; ok {
		f = zf
	}

	if p != nil && w != nil {
		return
	}

	// TODO: generalise func, and deduplicate with report_json's func
	open := func(name string) (f *os.File, w *csv.Writer, err2 error) {
		if z.ReportPath == "" {
			return nil, csv.NewWriter(z.DefaultWriter), nil
		}
		if st, err := os.Stat(z.ReportPath); os.IsNotExist(err) {
			err = os.MkdirAll(z.ReportPath, 0701)
			if err != nil {
				z.ec.Log().Error("Unable to create report path", zap.Error(err), zap.String("path", z.ReportPath))
				return nil, csv.NewWriter(z.DefaultWriter), err
			}
		} else if err != nil {
			z.ec.Log().Error("Unable to acquire information about the path", zap.Error(err), zap.String("path", z.ReportPath))
			return nil, csv.NewWriter(z.DefaultWriter), err
		} else if !st.IsDir() {
			z.ec.Log().Error("Report path is not a directory", zap.Error(err), zap.String("path", z.ReportPath))
			return nil, csv.NewWriter(z.DefaultWriter), nil
		}
		filePath := filepath.Join(z.ReportPath, name+".csv")
		z.ec.Log().Debug("Opening report file", zap.String("path", filePath))
		if zf, err := os.Create(filePath); err != nil {
			z.ec.Log().Error("unable to create report file, fallback to stdout", zap.String("path", filePath), zap.Error(err))
			return nil, csv.NewWriter(z.DefaultWriter), nil
		} else if z.ReportUseBom {
			return zf, ut_encoding.NewBomAawareCsvWriter(zf), nil
		} else {
			return zf, csv.NewWriter(zf), nil
		}
	}

	if f != nil {
		f.Close()
		z.ec.Log().Fatal("File opened but no writer and/or parser available")
	}
	f, w, err = open(name)
	if err != nil {
		return nil, nil, nil, err
	}
	p = app_report_column2.NewRow(row, z.ec)

	z.files[name] = f
	z.writers[name] = w
	z.parsers[name] = p

	if z.ReportHeader {
		w.Write(p.Header())
	}
	return
}

func (z *CsvReport) Init(ec *app2.ExecContext) error {
	z.ec = ec
	if z.files == nil {
		z.files = make(map[string]*os.File)
	}
	if z.writers == nil {
		z.writers = make(map[string]*csv.Writer)
	}
	if z.parsers == nil {
		z.parsers = make(map[string]app_report_column2.Row)
	}
	return nil
}

func (z *CsvReport) Report(row interface{}) error {
	_, w, p, err := z.prepare(row)
	if err != nil {
		return err
	}
	w.Write(p.ValuesAsString(row))

	return nil
}

func (z *CsvReport) Close() {
	for _, w := range z.writers {
		w.Flush()
	}
	for _, f := range z.files {
		f.Close()
	}
}
