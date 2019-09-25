package app_file_impl

import (
	"encoding/csv"
	"errors"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"io"
	"os"
	"reflect"
	"strconv"
)

func NewData() app_file.Data {
	return &CsvData{}
}

type CsvData struct {
	FilePath         string
	file             *os.File
	reader           *csv.Reader
	ctl              app_control.Control
	md               interface{}
	mt               reflect.Type
	orderToFieldName map[int]string
	fieldNameToOrder map[string]int
	modelReady       bool
	mode             string
	headers          []string
	colIndexToField  func(ci int, v reflect.Value, s string) error
}

func (z *CsvData) Model(ctl app_control.Control, m interface{}) (err error) {
	ui := ctl.UI()
	z.ctl = ctl
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
	z.md = m
	z.mt = reflect.TypeOf(m).Elem()
	z.orderToFieldName = make(map[int]string)
	z.fieldNameToOrder = make(map[string]int)

	ord := 0

	appendField := func(f reflect.StructField) {
		z.fieldNameToOrder[f.Name] = ord
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
	return nil
}

func (z *CsvData) header(cols []string) error {
	l := z.ctl.Log()
	l.Debug("Parse header", zap.Strings("cols", cols))

	z.headers = make([]string, len(cols))
	for i, col := range cols {
		z.headers[i] = strcase.ToCamel(col)
	}
	z.mode = "fieldName"
	for _, col := range cols {
		if _, ok := z.fieldNameToOrder[col]; !ok {
			z.mode = "order"
		}
	}
	l = l.With(zap.String("mode", z.mode))
	l.Debug("Data injection mode")

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
			f := v.Elem().FieldByName(fieldName)
			if !f.IsValid() || !f.CanSet() {
				l.Debug("Invalid column",
					zap.Bool("isZero", f.IsZero()),
					zap.Bool("isValid", f.IsValid()),
					zap.Bool("canSet", f.CanSet()),
				)
				return errors.New("invalid field")
			}
			return fieldSet(f, s)
		}

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
					zap.Bool("isZero", f.IsZero()),
					zap.Bool("isValid", f.IsValid()),
					zap.Bool("canSet", f.CanSet()),
				)
				return errors.New("invalid field")
			}
			return fieldSet(f, s)
		}
	}
	return nil
}

func (z *CsvData) row(cols []string) (m interface{}, err error) {
	rm := reflect.New(z.mt)
	for ci, col := range cols {
		if err := z.colIndexToField(ci, rm, col); err != nil {
			return nil, err
		}
	}
	return rm.Interface(), nil
}

func (z *CsvData) EachRow(exec func(m interface{}, rowIndex int) error) error {
	ui := z.ctl.UI()

	if !z.modelReady {
		return errors.New("model is not ready")
	}
	defer z.file.Close()
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
			if err := z.header(cols); err != nil {
				return err
			}

		default:
			m, err := z.row(cols)
			if err != nil {
				return err
			}
			if err := exec(m, ri); err != nil {
				return err
			}
		}
	}
}
