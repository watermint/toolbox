package fd_file_impl

import (
	"encoding/csv"
	"errors"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"io"
	"os"
	"reflect"
	"strconv"
)

func NewRowFeed(name string) fd_file.RowFeed {
	return &RowFeed{
		name: name,
	}
}

type RowFeed struct {
	FilePath         string
	file             *os.File
	reader           *csv.Reader
	ctl              app_control.Control
	md               interface{}
	mt               reflect.Type
	name             string
	orderToFieldName map[int]string
	fieldNameToOrder map[string]int
	modelReady       bool
	mode             string
	headers          []string
	fields           []string
	colIndexToField  func(ci int, v reflect.Value, s string) error
}

func (z *RowFeed) ForkForTest(path string) fd_file.RowFeed {
	f := z.Fork()
	rf := f.(*RowFeed)
	rf.FilePath = path
	return rf
}

func (z *RowFeed) Fork() fd_file.RowFeed {
	rf := &RowFeed{
		FilePath: z.FilePath,
		name:     z.name,
		md:       z.md,
	}
	rf.applyModelInternal()
	return rf
}

func (z *RowFeed) Spec() fd_file.Spec {
	return newSpec(z)
}

func (z *RowFeed) SetFileName(filePath string) {
	z.FilePath = filePath
}

func (z *RowFeed) SetModel(m interface{}) {
	z.md = m
	z.applyModelInternal()
}

func (z *RowFeed) Model() interface{} {
	return z.md
}

func (z *RowFeed) applyModelInternal() {
	l := app_root.Log()
	if z.md == nil {
		l.Debug("No model defined")
		return
	}

	z.mt = reflect.TypeOf(z.md).Elem()
	z.orderToFieldName = make(map[int]string)
	z.fieldNameToOrder = make(map[string]int)
	z.fields = make([]string, 0)

	ord := 0

	appendField := func(f reflect.StructField) {
		z.fields = append(z.fields, strcase.ToSnake(f.Name))
		z.fieldNameToOrder[f.Name] = ord
		z.fieldNameToOrder[strcase.ToSnake(f.Name)] = ord
		z.fieldNameToOrder[strcase.ToCamel(f.Name)] = ord
		z.orderToFieldName[ord] = f.Name
		ord++
	}

	for i := 0; i < z.mt.NumField(); i++ {
		f := z.mt.Field(i)
		switch f.Type.Kind() {
		case reflect.Bool:
			appendField(f)
		case reflect.Int:
			appendField(f)
		case reflect.String:
			appendField(f)
		}
	}

	z.modelReady = true
}

func (z *RowFeed) ApplyModel(ctl app_control.Control) error {
	z.ctl = ctl
	l := ctl.Log()
	ui := ctl.UI()
	if z.md == nil {
		l.Error("No model defined")
		return errors.New("no model defined")
	}
	var err error
	z.file, err = os.Open(z.FilePath)
	if err != nil {
		ui.Error("flow.error.unable_to_read",
			app_msg.P{
				"Path":  z.FilePath,
				"Error": err,
			},
		)
		return err
	}
	z.reader = csv.NewReader(z.file)
	z.applyModelInternal()

	return nil
}

func (z *RowFeed) header(cols []string) (consumeLine bool, err error) {
	l := z.ctl.Log()
	l.Debug("Parse header", zap.Strings("cols", cols))

	z.headers = make([]string, len(cols))
	for i, col := range cols {
		z.headers[i] = strcase.ToSnake(col)
	}
	z.mode = "fieldName"
	for _, col := range cols {
		if _, ok := z.fieldNameToOrder[col]; !ok {
			z.mode = "order"
		}
	}
	l = l.With(zap.String("mode", z.mode))
	l.Debug("Feed injection mode")

	fieldSet := func(v reflect.Value, s string) error {
		switch v.Kind() {
		case reflect.Bool:
			b, err := strconv.ParseBool(s)
			if err != nil {
				l.Debug("Failed to parse field", zap.String("s", s), zap.Error(err))
				return err
			}
			v.SetBool(b)

		case reflect.Int:
			n, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				l.Debug("Failed to parse field", zap.String("s", s), zap.Error(err))
				return err
			}
			v.SetInt(n)

		case reflect.String:
			v.SetString(s)
		}
		return nil
	}

	switch z.mode {
	case "fieldName":
		z.colIndexToField = func(ci int, v reflect.Value, s string) error {
			if ci >= len(z.headers) {
				l.Debug("Column index out of range", zap.Int("ci", ci))
				return nil // ignore error
			}
			fieldName := z.headers[ci]
			f := v.Elem().FieldByName(strcase.ToCamel(fieldName))
			if !f.IsValid() || !f.CanSet() {
				l.Debug("Invalid column",
					zap.Bool("isValid", f.IsValid()),
					zap.Bool("canSet", f.CanSet()),
				)
				return errors.New("invalid field")
			}
			return fieldSet(f, s)
		}
		return true, nil

	case "order":
		z.colIndexToField = func(ci int, v reflect.Value, s string) error {
			fieldName, ok := z.orderToFieldName[ci]
			if !ok {
				l.Debug("Column for field not found", zap.Int("ci", ci))
				return errors.New("column for field not found")
			}
			f := v.Elem().FieldByName(fieldName)
			if !f.IsValid() || !f.CanSet() {
				l.Debug("Invalid column",
					zap.Bool("isValid", f.IsValid()),
					zap.Bool("canSet", f.CanSet()),
				)
				return errors.New("invalid field")
			}
			return fieldSet(f, s)
		}
		return false, nil

	default:
		return false, errors.New("unexpected row mode")
	}
}

func (z *RowFeed) row(cols []string) (m interface{}, err error) {
	rm := reflect.New(z.mt)
	for ci, col := range cols {
		if err := z.colIndexToField(ci, rm, col); err != nil {
			return nil, err
		}
	}
	return rm.Interface(), nil
}

func (z *RowFeed) EachRow(exec func(m interface{}, rowIndex int) error) error {
	ui := z.ctl.UI()

	if !z.modelReady {
		return errors.New("model is not ready")
	}
	defer z.file.Close()

	consumeRow := func(cols []string, rowIndex int) error {
		m, err := z.row(cols)
		if err != nil {
			return err
		}
		if err := exec(m, rowIndex); err != nil {
			return err
		}
		return nil
	}
	for ri := 0; ; ri++ {
		cols, err := z.reader.Read()
		switch {
		case err == io.EOF:
			return nil

		case err != nil:
			ui.Error("flow.error.unable_to_read",
				app_msg.P{
					"Path":  z.FilePath,
					"Error": err,
				},
			)
			return err

		case ri == 0:
			if consume, err := z.header(cols); err != nil {
				return err
			} else if !consume {
				if err := consumeRow(cols, ri); err != nil {
					return err
				}
			}

		default:
			if err := consumeRow(cols, ri); err != nil {
				return err
			}
		}
	}
}
