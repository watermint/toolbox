package app_report

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/app86/app_control"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"reflect"
)

func NewJsonForQuiet(name string, ctl app_control.Control) (r Report, err error) {
	r = &Json{
		Writer: os.Stdout,
		Ctl:    ctl,
	}
	return r, nil
}

func NewJson(name string, ctl app_control.Control) (r Report, err error) {
	p, err := ctl.Workspace().Descendant(reportPath)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(filepath.Join(p, name+".json"))
	if err != nil {
		return nil, err
	}
	r = &Json{
		File:   f,
		Writer: f,
		Ctl:    ctl,
	}
	return r, nil
}

type Json struct {
	File   *os.File
	Writer io.Writer
	Ctl    app_control.Control
}

func (z *Json) findRaw(row interface{}, orig interface{}) json.RawMessage {
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

func (z *Json) Row(row interface{}) {
	raw := z.findRaw(row, row)
	if raw != nil {
		z.Writer.Write(raw)
		z.Writer.Write([]byte("\n"))
		return
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "")
	err := enc.Encode(row)
	if err != nil {
		z.Ctl.Log().Debug("Unable to unmarshal", zap.Error(err))
		return
	}
	z.Writer.Write(buf.Bytes())
}

func (z *Json) Transaction(state State, row interface{}, result interface{}) {
	z.Row(Transaction{State: state(), Input: row, Result: result})
}

func (z *Json) Flush() {
}

func (z *Json) Close() {
	if z.File != nil {
		z.File.Close()
	}
}
