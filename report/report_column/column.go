package report_column

import (
	"fmt"
	"github.com/watermint/toolbox/app"
	"go.uber.org/zap"
	"reflect"
	"strconv"
	"strings"
)

func RowName(row interface{}) string {
	rowType := reflect.TypeOf(row)
	if rowType.Kind() == reflect.Ptr {
		rowType = rowType.Elem()
	}
	return rowType.Name()
}

func NewRow(row interface{}, ec *app.ExecContext) Row {
	ri := &rowImpl{
		name: RowName(row),
		ec:   ec,
	}
	ri.header = ri.parseHeader(row)

	return ri
}

type Row interface {
	Name() string
	Header() []string
	Values(r interface{}) []interface{}
	ValuesAsString(r interface{}) []string
}

type rowImpl struct {
	name   string
	header []string
	ec     *app.ExecContext
}

func (z *rowImpl) Name() string {
	return z.name
}
func (z *rowImpl) Header() []string {
	return z.header
}

func (z *rowImpl) typeOf(r interface{}) reflect.Type {
	rt := reflect.TypeOf(r)
	if rt.Kind() == reflect.Ptr {
		rt = reflect.ValueOf(r).Elem().Type()
	}
	return rt
}

func (z *rowImpl) supportedType(k reflect.Kind) bool {
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

func (z *rowImpl) parseHeader(row interface{}) []string {
	return z.headerFromType("", z.typeOf(row))
}

func (z *rowImpl) headerFromType(prefix string, rt reflect.Type) (cols []string) {
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

func (z *rowImpl) marshal(v reflect.Value) (interface{}, error) {
	switch v.Kind() {
	case reflect.Ptr:
		return z.marshal(v.Elem())
	default:
		return v.Interface(), nil
	}
}

func (z *rowImpl) valueForPathAsString(path string, value reflect.Value) string {
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

func (z *rowImpl) valueForPath(path string, value reflect.Value) (interface{}, bool) {
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
		z.ec.Log().Debug(
			"field not found",
			zap.String("path", path),
			zap.String("field", p0),
		)
		return nil, false
	}

	vf := value.FieldByName(p0)
	if !vf.IsValid() {
		z.ec.Log().Debug(
			"field not found",
			zap.String("path", path),
			zap.String("field", p0),
		)
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

func (z *rowImpl) ValuesAsString(value interface{}) []string {
	vals := make([]string, 0)
	v := reflect.ValueOf(value)
	for _, c := range z.header {
		vals = append(vals, z.valueForPathAsString(c, v))
	}
	return vals
}

func (z *rowImpl) Values(value interface{}) []interface{} {
	vals := make([]interface{}, 0)
	v := reflect.ValueOf(value)
	for _, c := range z.header {
		vp, _ := z.valueForPath(c, v)
		vals = append(vals, vp)
	}
	return vals
}
