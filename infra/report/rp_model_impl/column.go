package rp_model_impl

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"go.uber.org/zap"
)

type Column interface {
	Header() []string
	Values(r interface{}) []interface{}
	ValueStrings(r interface{}) []string
}

func NewColumn(model interface{}, opts ...rp_model.ReportOpt) Column {
	ro := &rp_model.ReportOpts{}
	for _, opt := range opts {
		opt(ro)
	}
	ri := &columnImpl{
		opts: ro,
	}
	_ = ri.Parse(model)

	return ri
}

type columnImpl struct {
	header []string
	opts   *rp_model.ReportOpts
}

func (z *columnImpl) Parse(r interface{}) error {
	l := app_root.Log()
	b, err := json.Marshal(r)
	if err != nil {
		l.Debug("Unable to marshal", zap.Error(err))
		return err
	}
	if !gjson.ValidBytes(b) {
		l.Debug("Invalid JSON sequence")
		return errors.New("invalid row data format found")
	}
	s := gjson.ParseBytes(b)
	z.header = make([]string, 0)

	var parse func(t gjson.Result, prefix string)
	parse = func(t gjson.Result, prefix string) {
		t.ForEach(func(key, value gjson.Result) bool {
			switch {
			case value.IsArray():
				return true
			case key.String() == "Raw":
				return true
			case value.IsObject():
				parse(value, prefix+key.String()+".")
				return true
			default:
				colName := prefix + key.String()
				if !z.opts.IsHiddenColumn(colName) {
					z.header = append(z.header, colName)
				}
				return true
			}
		})
	}

	parse(s, "")

	return nil
}

func (z *columnImpl) Header() []string {
	return z.header
}

func (z *columnImpl) Values(r interface{}) []interface{} {
	l := app_root.Log()
	b, err := json.Marshal(r)
	if err != nil {
		l.Debug("Unable to marshal", zap.Error(err))
		return make([]interface{}, 0)
	}
	if !gjson.ValidBytes(b) {
		l.Debug("Invalid JSON sequence")
		return make([]interface{}, 0)
	}
	s := gjson.ParseBytes(b)
	cols := make([]interface{}, 0)

	for _, p := range z.header {
		cols = append(cols, s.Get(p).Value())
	}
	return cols
}

func (z *columnImpl) ValueStrings(r interface{}) []string {
	l := app_root.Log()
	b, err := json.Marshal(r)
	if err != nil {
		l.Debug("Unable to marshal", zap.Error(err))
		return make([]string, 0)
	}
	if !gjson.ValidBytes(b) {
		l.Debug("Invalid JSON sequence")
		return make([]string, 0)
	}
	s := gjson.ParseBytes(b)
	cols := make([]string, 0)

	for _, p := range z.header {
		cols = append(cols, s.Get(p).String())
	}
	return cols
}
