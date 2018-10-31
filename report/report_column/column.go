package report_column

import (
	"errors"
	"go.uber.org/zap"
	"reflect"
	"strconv"
)

type ColumnType struct {
	FieldIndex int
	Tag        string
	Type       reflect.Type
}

type Column struct {
	Row        int
	Column     int
	ColumnName string
	Value      string
}

const (
	ColumnTag = "column"
)

type ColumnMarshaller struct {
	firstColumn []ColumnType
	rowIndex    int
	Logger      *zap.Logger
}

func (c *ColumnMarshaller) IsFirstRow() bool {
	return c.rowIndex == 0
}

func (c *ColumnMarshaller) ColumnTypes(row interface{}) []ColumnType {
	rt := reflect.TypeOf(row)

	if rt.Kind() == reflect.Ptr {
		rt = reflect.Indirect(reflect.ValueOf(row)).Type()
	}

	cols := make([]ColumnType, 0)
	for i := 0; i < rt.NumField(); i++ {
		rf := rt.Field(i)
		ft := rf.Tag.Get(ColumnTag)

		if ft == "" {
			continue
		}

		col := ColumnType{
			FieldIndex: i,
			Tag:        ft,
			Type:       rf.Type,
		}
		cols = append(cols, col)
	}
	return cols
}

func (c *ColumnMarshaller) Row(row interface{}) (cols []Column, err error) {
	rt := reflect.TypeOf(row)
	rv := reflect.ValueOf(row)

	if rt.Kind() == reflect.Ptr {
		p := reflect.Indirect(reflect.ValueOf(row))
		rt0 := rt
		rt = reflect.TypeOf(p)
		rv = reflect.ValueOf(p)

		c.Logger.Debug("indirect",
			zap.Int("row", c.rowIndex),
			zap.Any("kind", rt.Kind()),
			zap.String("rt0", rt0.Name()),
			zap.String("rt", rt.Name()),
		)
	}

	if c.rowIndex == 0 {
		c.firstColumn = c.ColumnTypes(row)
	} else {
		colTypes := c.ColumnTypes(row)
		if !reflect.DeepEqual(c.firstColumn, colTypes) {
			c.Logger.Warn(
				"incompatible row found",
				zap.Any("expected", c.firstColumn),
				zap.Any("found", colTypes),
			)
			return nil, errors.New("incompatible row found")
		}
	}

	if len(c.firstColumn) < 1 {
		c.Logger.Warn("No column found")
	}

	cols = make([]Column, len(c.firstColumn))

	for i := 0; i < rv.NumField(); i++ {
		rvf := rv.Field(i)
		c.Logger.Debug("fieldName",
			zap.Int("i", i),
			zap.Any("rt.kind", rt.Kind()),
			zap.Any("rvf.Kind", rvf.Kind()),
			zap.String("typeName", rvf.Type().Name()),
		)
	}

	for i, ct := range c.firstColumn {
		rvf := rv.Field(ct.FieldIndex)
		c.Logger.Debug("fieldName",
			zap.Int("i", i),
			zap.String("typeName", rvf.Type().Name()),
		)
		v, err := c.MarshalColumn(rvf)
		if err != nil {
			c.Logger.Warn(
				"Unable to marshal data into columns",
				zap.Int("row", c.rowIndex),
				zap.Int("col", i),
				zap.String("col_name", ct.Tag),
				zap.Error(err),
			)
			v = ""
		}

		cols[i] = Column{
			Row:        c.rowIndex,
			Column:     i,
			ColumnName: ct.Tag,
			Value:      v,
		}
	}

	c.rowIndex++
	return
}

func (c *ColumnMarshaller) MarshalColumn(v reflect.Value) (string, error) {
	c.Logger.Debug("Marshal col", zap.String("type", v.Type().Name()))
	switch v.Kind() {
	case reflect.Ptr:
		return c.MarshalColumn(reflect.Indirect(v))
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

	c.Logger.Error(
		"Unsupported column type found",
		zap.Any("kind", v.Kind()),
	)
	return "", errors.New("unsupported column type found")
}
