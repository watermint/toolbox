package api_parser

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/log/esl"
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

	l := esl.Default().With(esl.String("valueType", vt.Name()))

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
				l.Error("Missing required field", esl.String("field", vtf.Name), esl.String("path", p))
				return errors.New("missing required field")
			}
			continue
		}

		switch vtf.Type.Kind() {
		case reflect.String:
			vvf.SetString(jv.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			vvf.SetInt(jv.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			vvf.SetUint(jv.Uint())
		case reflect.Bool:
			vvf.SetBool(jv.Bool())
		case reflect.Float32, reflect.Float64:
			vvf.SetFloat(jv.Float())

		default:
			l.Error("unexpected type found", esl.String("type.kind", vtf.Type.Kind().String()))
			return errors.New("unexpected type found")
		}
	}
	return nil
}

func ParseModelPathRaw(v interface{}, j json.RawMessage, path string) error {
	if !gjson.ValidBytes(j) {
		return errors.New("invalid json")
	}
	g := gjson.ParseBytes(j)
	p := g.Get(path)
	if !p.Exists() {
		return errors.New("unexpected data format")
	}
	return ParseModel(v, p)
}

func CombineRaw(raws map[string]json.RawMessage) json.RawMessage {
	b, err := json.Marshal(raws)
	if err != nil {
		esl.Default().Error("Unable to marshal", esl.Error(err))
		return json.RawMessage("{}")
	}
	return b
}
