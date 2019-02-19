package dbx_api

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

func ParseModelJsonForTest(v interface{}, raw json.RawMessage) error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	return parseModelJson(logger, v, raw)
}

func parseModel(logger *zap.Logger, v interface{}, j gjson.Result) error {
	vv := reflect.ValueOf(v).Elem()
	vt := vv.Type()

	for i := vt.NumField() - 1; i >= 0; i-- {
		vtf := vt.Field(i)
		vvf := vv.Field(i)

		if vtf.Name == "Raw" && vvf.Type().Kind() == reflect.TypeOf(json.RawMessage{}).Kind() {
			vvf.SetBytes(json.RawMessage(j.Raw))
			continue
		}

		p := vtf.Tag.Get("path")
		if p == "" {
			continue
		}
		pp := strings.Split(p, ",")
		path := pp[0]
		required := false
		if len(pp) > 1 && pp[1] == "required" {
			required = true
		}

		jv := j.Get(path)
		if !jv.Exists() {
			if required {
				logger.Debug("missing required field", zap.String("field", vtf.Name))
				return errors.New("missing required field")
			}
			continue
		}

		switch vtf.Type.Kind() {
		case reflect.String:
			vvf.SetString(jv.String())
		case reflect.Int:
			vvf.SetInt(jv.Int())
		case reflect.Int8:
			vvf.SetInt(jv.Int())
		case reflect.Int16:
			vvf.SetInt(jv.Int())
		case reflect.Int32:
			vvf.SetInt(jv.Int())
		case reflect.Int64:
			vvf.SetInt(jv.Int())
		case reflect.Uint:
			vvf.SetUint(jv.Uint())
		case reflect.Uint8:
			vvf.SetUint(jv.Uint())
		case reflect.Uint16:
			vvf.SetUint(jv.Uint())
		case reflect.Uint32:
			vvf.SetUint(jv.Uint())
		case reflect.Uint64:
			vvf.SetUint(jv.Uint())
		case reflect.Bool:
			vvf.SetBool(jv.Bool())
		case reflect.Float32:
			vvf.SetFloat(jv.Float())
		case reflect.Float64:
			vvf.SetFloat(jv.Float())

		default:
			logger.Debug("unexpected type found", zap.String("type", vtf.Type.Kind().String()))
			return errors.New("unexpected type found")
		}
	}
	return nil
}

// `v` is pointer of the struct
func parseModelJson(logger *zap.Logger, v interface{}, raw json.RawMessage) error {
	if !gjson.ValidBytes(raw) {
		return errors.New("invalid json sequence")
	}

	j := gjson.ParseBytes(raw)

	return parseModel(logger, v, j)
}
