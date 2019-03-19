package app_report_json

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_report/app_report_column"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"reflect"
)

type JsonReport struct {
	ec            *app.ExecContext
	ReportPath    string
	DefaultWriter io.Writer
	files         map[string]*os.File
	writers       map[string]io.Writer
}

func (z *JsonReport) findRaw(row interface{}, orig interface{}) json.RawMessage {
	var rv reflect.Value
	switch r := row.(type) {
	case reflect.Value:
		rv = r
	default:
		rv = reflect.ValueOf(row)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
	}
	rt := rv.Type()
	_, e := rt.FieldByName("Raw")
	if !e {
		return nil
	}
	rvf := rv.FieldByName("Raw")
	if rvf.Type().Kind() != reflect.TypeOf(json.RawMessage{}).Kind() {
		return nil
	}
	return rvf.Bytes()
}

func (z *JsonReport) prepare(row interface{}) (f *os.File, w io.Writer, err error) {
	name := app_report_column.RowName(row)
	if zw, ok := z.writers[name]; ok {
		w = zw
	}
	if zf, ok := z.files[name]; ok {
		f = zf
	}
	if w != nil {
		return
	}

	open := func(name string) (f *os.File, w io.Writer, err2 error) {
		if z.ReportPath == "" {
			return nil, z.DefaultWriter, nil
		}
		if st, err := os.Stat(z.ReportPath); os.IsNotExist(err) {
			err = os.MkdirAll(z.ReportPath, 0701)
			if err != nil {
				z.ec.Log().Error(
					"Unable to create report path",
					zap.Error(err),
					zap.String("path", z.ReportPath),
				)
				return nil, z.DefaultWriter, err
			}
		} else if err != nil {
			z.ec.Log().Error(
				"Unable to acquire information about the path",
				zap.Error(err),
				zap.String("path", z.ReportPath),
			)
			return nil, z.DefaultWriter, err
		} else if !st.IsDir() {
			z.ec.Log().Error(
				"Report path is not a directory",
				zap.Error(err),
				zap.String("path", z.ReportPath),
			)
			return nil, z.DefaultWriter, nil
		}

		filePath := filepath.Join(z.ReportPath, name+".json")
		z.ec.Log().Debug("Opening report file", zap.String("path", filePath))
		if zf, err := os.Create(filePath); err != nil {
			z.ec.Log().Error(
				"unable to create report file, fallback to default writer",
				zap.String("path", filePath),
				zap.Error(err),
			)
			return nil, z.DefaultWriter, nil
		} else {
			return zf, zf, nil
		}
	}

	if f != nil {
		f.Close()
		z.ec.Log().Fatal("File opened but no writer and/or parser available")
	}
	f, w, err = open(name)
	if err != nil {
		return nil, nil, err
	}

	z.files[name] = f
	z.writers[name] = w

	return
}

func (z *JsonReport) Init(ec *app.ExecContext) error {
	z.ec = ec
	if z.files == nil {
		z.files = make(map[string]*os.File)
	}
	if z.writers == nil {
		z.writers = make(map[string]io.Writer)
	}
	return nil
}

func (z *JsonReport) Close() {
	for _, f := range z.files {
		f.Close()
	}
}

func (z *JsonReport) Report(row interface{}) error {
	f, w, err := z.prepare(row)
	if err != nil {
		return err
	}
	raw := z.findRaw(row, row)
	if raw != nil {
		w.Write(raw)
		w.Write([]byte("\n"))

		return nil
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "")
	err = enc.Encode(row)
	if err != nil {
		fn := ""
		if f != nil {
			fn = f.Name()
		}
		z.ec.Log().Error("Couldn't write report", zap.Error(err), zap.String("file", fn))
	}
	w.Write(buf.Bytes())
	w.Write([]byte("\n"))

	return nil
}
