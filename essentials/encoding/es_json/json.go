package es_json

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

var (
	ErrorInvalidJSONFormat    = errors.New("invalid json format")
	ErrorNotFound             = errors.New("data not found for the path")
	ErrorMissingRequired      = errors.New("missing required field")
	ErrorUnsupportedFieldType = errors.New("unsupported field type")
	ErrorNotAnArray           = errors.New("data is not an array")
)

const (
	TagPathName          = "path"
	TagPathValueRequired = "required"
	PathArrayFirst       = "0"
)

// Wrapper of gjson
type Json interface {
	// Raw JSON message
	Raw() json.RawMessage

	// True if the instance is null
	IsNull() bool

	// Returns an array value and true if this instance is an array.
	Array() (v []Json, t bool)

	// Run each entries
	ArrayEach(f func(e Json) error) error

	// Returns an object value and true if this instance is an object.
	Object() (v map[string]Json, t bool)

	// Returns an boolean value and true if this instance is true/false.
	Bool() (v bool, t bool)

	// Returns a number value and true if this instance is a number.
	Number() (v es_number.Number, t bool)

	// Returns a string value and true if this instance is a string.
	String() (v string, t bool)

	// Parse model with given type.
	Model(v interface{}) (err error)

	// Find value under the path. Returns nil & false if not found.
	Find(path string) (j Json, found bool)

	// Find then Model.
	FindModel(path string, v interface{}) (err error)

	// Find an array
	FindArray(path string) (v []Json, t bool)

	// Find an array
	FindArrayEach(path string, f func(e Json) error) error

	// Find an object
	FindObject(path string) (v map[string]Json, t bool)

	// Find Bool
	FindBool(path string) (v bool, t bool)

	// Returns a number value and true if this instance is a number.
	FindNumber(path string) (v es_number.Number, t bool)

	// Returns a string value and true if this instance is a string.
	FindString(path string) (v string, t bool)
}

func MustParseString(j string) Json {
	if j, err := ParseString(j); err != nil {
		return Null()
	} else {
		return j
	}
}

func MustParse(j []byte) Json {
	if j, err := Parse(j); err != nil {
		return Null()
	} else {
		return j
	}
}

func ParseString(j string) (Json, error) {
	return Parse([]byte(j))
}

func Parse(j []byte) (Json, error) {
	if !Validate(j) {
		return nil, ErrorInvalidJSONFormat
	}
	r := gjson.ParseBytes(j)
	return newWrapper(r), nil
}

func newWrapper(r gjson.Result) Json {
	return &wrapperImpl{
		r: r,
	}
}

type wrapperImpl struct {
	r gjson.Result
}

func (z wrapperImpl) FindArrayEach(path string, f func(e Json) error) error {
	if entries, found := z.FindArray(path); found {
		for _, e := range entries {
			if err := f(e); err != nil {
				return err
			}
		}
		return nil
	}
	return ErrorNotFound
}

func (z wrapperImpl) IsNull() bool {
	return z.r.Type == gjson.Null
}

func (z wrapperImpl) Raw() json.RawMessage {
	return json.RawMessage(z.r.Raw)
}

func (z wrapperImpl) Array() (v []Json, t bool) {
	if !z.r.IsArray() {
		return nil, false
	}
	v = make([]Json, 0)
	for _, e := range z.r.Array() {
		v = append(v, newWrapper(e))
	}
	return v, true
}

func (z wrapperImpl) ArrayEach(f func(e Json) error) error {
	if entries, t := z.Array(); !t {
		return ErrorNotAnArray
	} else {
		for _, e := range entries {
			if err := f(e); err != nil {
				return err
			}
		}
		return nil
	}
}

func (z wrapperImpl) Object() (v map[string]Json, t bool) {
	if !z.r.IsObject() {
		return nil, false
	}
	v = make(map[string]Json)
	for k, v0 := range z.r.Map() {
		v[k] = newWrapper(v0)
	}
	return v, true
}

func (z wrapperImpl) Bool() (v bool, t bool) {
	if z.r.Type != gjson.True && z.r.Type != gjson.False {
		return false, false
	}
	return z.r.Bool(), true
}

func (z wrapperImpl) Number() (v es_number.Number, t bool) {
	if z.r.Type != gjson.Number {
		return nil, false
	}
	return es_number.New(z.r.Raw), true
}

func (z wrapperImpl) String() (v string, t bool) {
	if z.r.Type != gjson.String {
		return "", false
	}
	return z.r.String(), true
}

func (z wrapperImpl) Model(v interface{}) (err error) {
	l := app_root.Log()

	vv := reflect.ValueOf(v).Elem()
	vt := vv.Type()

	l = l.With(zap.String("valueType", vt.Name()))

	for i := vt.NumField() - 1; i >= 0; i-- {
		vtf := vt.Field(i)
		vvf := vv.Field(i)

		if vtf.Name == "Raw" && vvf.Type().Kind() == reflect.TypeOf(json.RawMessage{}).Kind() {
			vvf.SetBytes(json.RawMessage(z.r.Raw))
			continue
		}

		p := vtf.Tag.Get(TagPathName)
		if p == "" {
			continue
		}
		pp := strings.Split(p, ",")
		path := pp[0]
		required := false
		if len(pp) > 1 && pp[1] == TagPathValueRequired {
			required = true
		}

		jv := z.r.Get(path)
		if !jv.Exists() {
			if required {
				l.Debug("Missing required field",
					zap.String("field", vtf.Name),
					zap.String("path", p))
				return ErrorMissingRequired
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
			l.Error("unexpected type found", zap.String("type.kind", vtf.Type.Kind().String()))
			return ErrorUnsupportedFieldType
		}
	}
	return nil
}

func (z wrapperImpl) Find(path string) (j Json, found bool) {
	r1 := z.r.Get(path)
	if !r1.Exists() {
		return nil, false
	}
	return newWrapper(r1), true
}

func (z wrapperImpl) FindModel(path string, v interface{}) (err error) {
	if x, found := z.Find(path); !found {
		return ErrorNotFound
	} else {
		return x.Model(v)
	}
}

func (z wrapperImpl) FindArray(path string) (v []Json, t bool) {
	if x, found := z.Find(path); !found {
		return nil, false
	} else {
		return x.Array()
	}
}

func (z wrapperImpl) FindObject(path string) (v map[string]Json, t bool) {
	if x, found := z.Find(path); !found {
		return nil, false
	} else {
		return x.Object()
	}
}

func (z wrapperImpl) FindBool(path string) (v bool, t bool) {
	if x, found := z.Find(path); !found {
		return false, false
	} else {
		return x.Bool()
	}
}

func (z wrapperImpl) FindNumber(path string) (v es_number.Number, t bool) {
	if x, found := z.Find(path); !found {
		return nil, false
	} else {
		return x.Number()
	}
}

func (z wrapperImpl) FindString(path string) (v string, t bool) {
	if x, found := z.Find(path); !found {
		return "", false
	} else {
		return x.String()
	}
}
