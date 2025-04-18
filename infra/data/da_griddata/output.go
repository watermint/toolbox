package da_griddata

import (
	"io"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_multi"
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	OutputTypeAll = "all"
	OutputTypeCsv = "csv"
	// OutputTypeXlsx = "xlsx"
	OutputTypeJson = "json"
)

var (
	OutputTypes = []string{
		OutputTypeAll,
		OutputTypeCsv,
		// OutputTypeXlsx,
		OutputTypeJson,
	}
)

type GridDataOutput interface {
	mo_multi.MultiValue
	Row(column []interface{})
	SetFormatter(outType string, formatter GridDataFormatter)
	Open(c app_control.Control) error
	Close()
	Spec() GridDataOutputSpec
	Debug() interface{}
	FilePath() string
}

type GridDataWriter interface {
	Name() string
	Row(column []interface{})
	Open(c app_control.Control) error
	Close()
}

type PlainGridDataWriter interface {
	// FileSuffix returns the data type file suffix starting with a dot (e.g. `.json`, or `.csv`).
	FileSuffix() string

	// WriteRow writes a single row to the output.
	WriteRow(l esl.Logger, w io.Writer, formatter GridDataFormatter, row int, column []interface{}) error
}

type GridDataFormatter interface {
	Format(v interface{}, col, row int) interface{}
}

var (
	DefaultGridDataFormatter GridDataFormatter = &PlainGridDataFormatter{}
)

type PlainGridDataFormatter struct {
}

func (z PlainGridDataFormatter) Format(v interface{}, col, row int) interface{} {
	return v
}
