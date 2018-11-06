package report_column

import (
	"errors"
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

func NewRow(row interface{}, logger *zap.Logger) Row {
	ri := &rowImpl{
		name:   RowName(row),
		logger: logger,
	}
	ri.header = ri.parseHeader(row)

	return ri
}

type Row interface {
	Name() string
	Header() []string
	Values(r interface{}) []string
}

type rowImpl struct {
	name   string
	header []string
	logger *zap.Logger
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

func (z *rowImpl) marshal(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.Ptr:
		return z.marshal(v.Elem())
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Int8:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Int16:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Int32:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Uint8:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Uint16:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Uint32:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.String:
		return v.String(), nil
	}
	return "", errors.New("unsupported type")
}

func (z *rowImpl) valueForPath(path string, value reflect.Value) string {
	if !value.IsValid() {
		return ""
	}
	if value.Type().Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if path == "" {
		if mv, err := z.marshal(value); err != nil {
			return ""
		} else {
			return mv
		}
	}

	paths := strings.Split(path, ".")
	p0 := paths[0]
	vt := value.Type()
	if _, ok := vt.FieldByName(p0); !ok {
		z.logger.Debug(
			"field not found",
			zap.String("path", path),
			zap.String("field", p0),
		)
		return ""
	}

	vf := value.FieldByName(p0)
	if !vf.IsValid() {
		z.logger.Debug(
			"field not found",
			zap.String("path", path),
			zap.String("field", p0),
		)
		return ""
	}
	if vf.Type().Kind() == reflect.Ptr {
		vf = vf.Elem()
	}
	if len(paths) > 1 {
		return z.valueForPath(strings.Join(paths[1:], "."), vf)
	}
	if mv, err := z.marshal(vf); err != nil {
		return ""
	} else {
		return mv
	}
}

func (z *rowImpl) Values(value interface{}) []string {
	vals := make([]string, 0)
	v := reflect.ValueOf(value)
	for _, c := range z.header {
		vals = append(vals, z.valueForPath(c, v))
	}
	return vals
}
