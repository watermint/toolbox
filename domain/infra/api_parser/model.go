package api_parser

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"reflect"
	"strings"
)

func ParseModelString(v interface{}, j string) error {
	if !gjson.Valid(j) {
		return errors.New("invalid json")
	}
	g := gjson.Parse(j)
	return ParseModel(v, g)
}

func ParseModelRaw(v interface{}, j json.RawMessage) error {
	if !gjson.ValidBytes(j) {
		return errors.New("invalid json")
	}
	g := gjson.ParseBytes(j)
	return ParseModel(v, g)
}

func ParseModel(v interface{}, j gjson.Result) error {
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
			return errors.New("unexpected type found")
		}
	}
	return nil
}
