package rp_model

import (
	"errors"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_column"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type ReportOpt func(o *ReportOpts) *ReportOpts
type ReportOpts struct {
	HiddenColumns   map[string]bool
	ShowAllColumns  bool
	ShowReportTitle bool
	ReportSuffix    string
	NoConsoleOutput bool
	ColumnModel     rp_column.Column
}

func (z *ReportOpts) IsHiddenColumn(name string) bool {
	if z.ShowAllColumns {
		return false
	}
	if z.HiddenColumns == nil {
		return false
	}
	_, ok := z.HiddenColumns[name]
	return ok
}

func ShowReportTitle() ReportOpt {
	return func(o *ReportOpts) *ReportOpts {
		o.ShowReportTitle = true
		return o
	}
}

func NoConsoleOutput() ReportOpt {
	return func(o *ReportOpts) *ReportOpts {
		o.NoConsoleOutput = true
		return o
	}
}
func Suffix(suffix string) ReportOpt {
	return func(o *ReportOpts) *ReportOpts {
		o.ReportSuffix = suffix
		return o
	}
}
func ColumnModel(model rp_column.Column) ReportOpt {
	return func(o *ReportOpts) *ReportOpts {
		o.ColumnModel = model
		return o
	}
}
func ShowAllColumns(enabled bool) ReportOpt {
	return func(o *ReportOpts) *ReportOpts {
		o.ShowAllColumns = enabled
		return o
	}
}

func HiddenColumns(col ...string) ReportOpt {
	return func(o *ReportOpts) *ReportOpts {
		if o.HiddenColumns == nil {
			o.HiddenColumns = make(map[string]bool)
		}
		for _, c := range col {
			o.HiddenColumns[c] = true
		}
		return o
	}
}

type Report interface {
	Open(opts ...ReportOpt) error
	Spec() Spec
	Rows() int64

	// Close report, close should not raise exception when the report already closed.
	Close()
}

type RowReport interface {
	Report
	OpenNew(opts ...ReportOpt) (r RowReport, err error)
	Row(row interface{})
	SetModel(row interface{}, opts ...ReportOpt)
}

type TransactionReport interface {
	Report
	OpenNew(opts ...ReportOpt) (r TransactionReport, err error)
	Success(input interface{}, result interface{})
	Failure(err error, input interface{})
	Skip(reason app_msg.Message, input interface{})
	SetModel(input interface{}, result interface{}, opts ...ReportOpt)
}

type Spec interface {
	Name() string
	Model() interface{}
	Desc() app_msg.Message
	Columns() []string
	ColumnDesc(col string) app_msg.Message
	Options() []ReportOpt

	Doc(ui app_ui.UI) *dc_recipe.Report
}

type TransactionRow struct {
	Status    string      `json:"status"`
	StatusTag string      `json:"status_tag"`
	Reason    string      `json:"reason"`
	Input     interface{} `json:"input"`
	Result    interface{} `json:"result"`
}

const (
	StatusTagSuccess = "success"
	StatusTagFailure = "failure"
	StatusTagSkip    = "skip"
)

var (
	ErrorWriterIsNotReady = errors.New("writer is not ready")
)
