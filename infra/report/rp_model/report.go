package rp_model

import "github.com/watermint/toolbox/infra/ui/app_msg"

type ReportOpt func(o *ReportOpts) *ReportOpts
type ReportOpts struct {
	HiddenColumns  map[string]bool
	ShowAllColumns bool
	ReportSuffix   string
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

func Suffix(suffix string) ReportOpt {
	return func(o *ReportOpts) *ReportOpts {
		o.ReportSuffix = suffix
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
		o.HiddenColumns = make(map[string]bool)
		for _, c := range col {
			o.HiddenColumns[c] = true
		}
		return o
	}
}

type Report interface {
	// Report data row
	Row(row interface{})

	// Report transaction result
	Success(input interface{}, result interface{})
	Failure(err error, input interface{})
	Skip(reason app_msg.Message, input interface{}, result interface{})

	Close()
}

func TransactionHeader(input interface{}, result interface{}) *TransactionRow {
	return &TransactionRow{
		Input:  input,
		Result: result,
	}
}

type TransactionRow struct {
	Status string      `json:"status"`
	Reason string      `json:"reason"`
	Input  interface{} `json:"input"`
	Result interface{} `json:"result"`
}

var (
	MsgSuccess     = app_msg.M("report.transaction.success")
	MsgFailure     = app_msg.M("report.transaction.failure")
	MsgSkip        = app_msg.M("report.transaction.skip")
	MsgInvalidData = app_msg.M("report.transaction.failure.invalid_data")
)

type NotFound struct {
	Id string
}

func (z *NotFound) Error() string {
	return "entry not found for id: " + z.Id
}

type InvalidData struct {
}

func (z *InvalidData) Error() string {
	return "invalid data"
}
