package model2

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"reflect"
	"strings"
)

type Profile struct {
	Raw json.RawMessage

	// for processing, csv export, etc
	Email        string `path:"email,required"`
	TeamMemberId string `path:"team_member_id"`
	Status       string `path:"status.\\.tag"`
}

// `v` is pointer of the struct
func Parse(v interface{}, raw json.RawMessage) bool {
	if !gjson.ValidBytes(raw) {
		return false
	}

	json := gjson.ParseBytes(raw)

	vv := reflect.ValueOf(v).Elem()
	vt := vv.Type()

	for i := vt.NumField() - 1; i >= 0; i-- {
		vtf := vt.Field(i)
		vvf := vv.Field(i)
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

		jv := json.Get(path)
		if !jv.Exists() {
			if required {
				// omit error
				return false
			}
			continue
		}

		switch vtf.Type.Kind() {
		case reflect.String:
			vvf.SetString(jv.String())
		case reflect.Int:
			vvf.SetInt(jv.Int())
		// TODO: other int types
		case reflect.Bool:
			vvf.SetBool(jv.Bool())
		case reflect.Float32:
			vvf.SetFloat(jv.Float())
		case reflect.Float64:
			vvf.SetFloat(jv.Float())

		default:
			// TODO: error
		}
	}
	return true
}
