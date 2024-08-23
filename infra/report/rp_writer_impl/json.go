package rp_writer_impl

import (
	"bytes"
	"encoding/json"
	"github.com/itchyny/gojq"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_writer"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

func NewJsonWriter(name string, ctl app_control.Control, toStdout bool) rp_writer.Writer {
	return &jsonWriter{
		name:     name,
		toStdout: toStdout,
		ctl:      ctl,
	}
}

type jsonWriter struct {
	name     string
	index    int
	path     string
	toStdout bool
	file     *os.File
	w        io.Writer
	mutex    sync.Mutex
	ctl      app_control.Control
	warnZero sync.Once
}

func (z *jsonWriter) Name() string {
	return z.name
}

func (z *jsonWriter) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) (err error) {
	z.ctl = ctl
	if z.toStdout {
		z.w = es_stdout.NewDirectOut()
		return nil
	}
	l := ctl.Log()
	ro := &rp_model.ReportOpts{}
	for _, o := range opts {
		o(ro)
	}

	z.path = filepath.Join(ctl.Workspace().Report(), z.Name()+ro.ReportSuffix+".json")
	l = l.With(esl.String("path", z.path))
	l.Debug("Create new json report")
	z.file, err = os.Create(z.path)
	if err != nil {
		l.Error("Unable to create file", esl.Error(err))
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

	l := z.ctl.Log().With(esl.String("path", z.path))
	z.index++
	if r == nil {
		z.warnZero.Do(func() {
			l.Error("Empty row found")
		})
		return
	}

	filter, enabled := z.ctl.Feature().UIReportFilter()
	var filterQuery *gojq.Query
	var err error
	if enabled {
		l.Debug("Filter enabled", esl.String("filter", filter))
		filterQuery, err = gojq.Parse(filter)
		if err != nil {
			l.Debug("Unable to parse filter query", esl.Error(err))
			filterQuery = nil // ignore filter
		}
	}

	raw := z.findRaw(r)
	if raw != nil {
		if filterQuery != nil {
			v0, err := es_json.ParseAny(raw)
			if err != nil {
				l.Debug("Unable to unmarshal", esl.Error(err))
				return
			}
			v1, err := es_json.QuerySingle(v0, filterQuery)
			if err != nil {
				l.Debug("Unable to filter", esl.Error(err))
				return
			}
			switch v2 := v1.(type) {
			case string:
				z.w.Write([]byte(v2))
			case []byte:
				z.w.Write(v2)
			default:
				v4, _ := json.Marshal(v2)
				z.w.Write(v4)
			}

			z.w.Write([]byte("\n"))
			return
		} else {
			z.w.Write(raw)
			z.w.Write([]byte("\n"))
			return
		}
	}

	if filterQuery != nil {
		v, err := es_json.QuerySingle(r, filterQuery)
		if err != nil {
			l.Debug("Unable to filter", esl.Error(err))
			return
		}
		switch v1 := v.(type) {
		case string:
			z.w.Write([]byte(v1))
		case []byte:
			z.w.Write(v1)
		default:
			v2, _ := json.Marshal(v)
			z.w.Write(v2)
		}
		z.w.Write([]byte("\n"))
		return
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "")
	err = enc.Encode(r)
	if err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		return
	}
	z.w.Write(buf.Bytes())
}

func (z *jsonWriter) Close() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.ctl.Log().With(esl.String("path", z.path))

	if z.file != nil {
		err := z.file.Close()
		l.Debug("File closed", esl.Error(err))

		if z.index < 1 && z.ctl.Feature().IsProduction() && !z.ctl.Feature().IsTest() {
			l.Debug("Try removing empty report file")
			err := os.Remove(z.path)
			l.Debug("Removed or had an error (ignore)", esl.Error(err))
		}
		z.file = nil
	}
}
