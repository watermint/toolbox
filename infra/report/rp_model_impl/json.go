package rp_model_impl

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

type jsonWriter struct {
	name     string
	index    int
	path     string
	toStdout bool
	file     *os.File
	w        io.Writer
	mutex    sync.Mutex
	ctl      app_control.Control
}

func (z *jsonWriter) Name() string {
	return z.name
}

func (z *jsonWriter) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) (err error) {
	z.ctl = ctl
	if z.toStdout {
		z.w = os.Stdout
		return nil
	}
	l := ctl.Log()

	z.path = filepath.Join(ctl.Workspace().Report(), z.Name()+".json")
	l = l.With(zap.String("path", z.path))
	l.Debug("Create new json report")
	z.file, err = os.Create(z.path)
	if err != nil {
		l.Error("Unable to create file", zap.Error(err))
		return err
	}
	z.w = z.file
	return nil
}

func (z *jsonWriter) findRaw(row interface{}) json.RawMessage {
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

func (z *jsonWriter) Row(r interface{}) {
	z.mutex.Lock()
	defer z.mutex.Unlock()
	z.index++
	l := z.ctl.Log().With(zap.String("path", z.path))

	raw := z.findRaw(r)
	if raw != nil {
		z.w.Write(raw)
		z.w.Write([]byte("\n"))
		return
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "")
	err := enc.Encode(r)
	if err != nil {
		l.Debug("Unable to unmarshal", zap.Error(err))
		return
	}
	z.w.Write(buf.Bytes())
}

func (z *jsonWriter) Close() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.ctl.Log().With(zap.String("path", z.path))

	if z.file != nil {
		err := z.file.Close()
		l.Debug("File closed", zap.Error(err))

		if z.index < 1 && z.ctl.IsProduction() {
			l.Debug("Try removing empty report file")
			err := os.Remove(z.path)
			l.Debug("Removed or had an error (ignore)", zap.Error(err))
		}
		z.file = nil
	}
}
