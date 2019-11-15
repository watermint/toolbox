package rp_model_impl

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
	"reflect"
	"strconv"
	"strings"
)

type columnReflectImpl struct {
	header []string
	ctl    app_control.Control
}

func (z *columnReflectImpl) Header() []string {
	return z.header
}

func (z *columnReflectImpl) typeOf(r interface{}) reflect.Type {
	rt := reflect.TypeOf(r)
	if rt.Kind() == reflect.Ptr {
		rt = reflect.ValueOf(r).Elem().Type()
	}
	return rt
}

func (z *columnReflectImpl) supportedType(k reflect.Kind) bool {
	switch k {
	case reflect.Array:
		return false
	case reflect.Chan:
		return false
	case reflect.Func:
		return false
	case reflect.Map:
		return false
	case reflect.Slice:
		return false
	case reflect.UnsafePointer:
		return false
	case reflect.Uintptr:
		return false
	}
	return true
}

func (z *columnReflectImpl) Parse(row interface{}) []string {
	return z.headerFromType("", z.typeOf(row))
}

func (z *columnReflectImpl) headerFromType(prefix string, rt reflect.Type) (cols []string) {
	cols = make([]string, 0)
	if rt.Kind() == reflect.Struct {
		n := rt.NumField()
		for i := 0; i < n; i++ {
			rf := rt.Field(i)
			rfk := rf.Type.Kind()
			rft := rf.Type
			if rfk == reflect.Ptr {
				rfk = rf.Type.Elem().Kind()
				rft = rf.Type.Elem()
			}
			if rfk == reflect.Struct {
				cols = append(cols, z.headerFromType(prefix+rf.Name+".", rft)...)
			} else if z.supportedType(rfk) {
				cols = append(cols, prefix+rf.Name)
			}
		}
	} else if z.supportedType(rt.Kind()) {
		cols = append(cols, prefix+"")
	}
	return
}

func (z *columnReflectImpl) marshal(v reflect.Value) (interface{}, error) {
	switch v.Kind() {
	case reflect.Ptr:
		return z.marshal(v.Elem())
	case reflect.String:
		return v.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint(), nil
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	default:
		return nil, errors.New("unsupported kind")
	}
}

func (z *columnReflectImpl) valueForPathAsString(path string, value reflect.Value) string {
	pathValue, e := z.valueForPath(path, value)
	if !e {
		return ""
	}
	switch v := pathValue.(type) {
	case bool:
		return strconv.FormatBool(v)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (z *columnReflectImpl) valueForPath(path string, value reflect.Value) (interface{}, bool) {
	l := z.ctl.Log().With(zap.String("path", path))

	if !value.IsValid() {
		return nil, false
	}
	if value.Type().Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if path == "" {
		return nil, false
	}

	paths := strings.Split(path, ".")
	p0 := paths[0]
	vt := value.Type()
	if _, ok := vt.FieldByName(p0); !ok {
		l.Debug("field not found", zap.String("field", p0))
		return nil, false
	}

	vf := value.FieldByName(p0)
	if !vf.IsValid() {
		l.Debug("field not found", zap.String("field", p0))
		return nil, false
	}
	if vf.Type().Kind() == reflect.Ptr {
		vf = vf.Elem()
	}
	if len(paths) > 1 {
		return z.valueForPath(strings.Join(paths[1:], "."), vf)
	}
	mv, err := z.marshal(vf)
	if err != nil {
		return nil, false
	}
	return mv, true
}

func (z *columnReflectImpl) ValuesAsString(value interface{}) []string {
	vals := make([]string, 0)
	v := reflect.ValueOf(value)
	for _, c := range z.header {
		vals = append(vals, z.valueForPathAsString(c, v))
	}
	return vals
}

func (z *columnReflectImpl) Values(value interface{}) []interface{} {
	vals := make([]interface{}, 0)
	v := reflect.ValueOf(value)
	for _, c := range z.header {
		vp, _ := z.valueForPath(c, v)
		vals = append(vals, vp)
	}
	return vals
}
