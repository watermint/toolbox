package rp_column_impl

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

var (
	ErrorInvalidRowDataFormat = errors.New("invalid row data format")
)

func Headers(r interface{}, isHidden func(name string) bool) (headers []string, err error) {
	l := app_root.Log()
	b, err := json.Marshal(r)
	if err != nil {
		l.Debug("Unable to marshal", zap.Error(err))
		return nil, err
	}
	if !gjson.ValidBytes(b) {
		l.Debug("Invalid JSON sequence")
		return nil, ErrorInvalidRowDataFormat
	}
	s := gjson.ParseBytes(b)
	headers = make([]string, 0)

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
				if !isHidden(colName) {
					headers = append(headers, colName)
				}
				return true
			}
		})
	}

	parse(s, "")
	return headers, nil
}

func Parse(r interface{}) (s gjson.Result, err error) {
	l := app_root.Log()
	b, err := json.Marshal(r)
	if err != nil {
		l.Debug("Unable to marshal", zap.Error(err))
		return gjson.Parse("{}"), ErrorInvalidRowDataFormat
	}
	if !gjson.ValidBytes(b) {
		l.Debug("Invalid JSON sequence")
		return gjson.Parse("{}"), ErrorInvalidRowDataFormat
	}
	s = gjson.ParseBytes(b)
	return s, nil
}
