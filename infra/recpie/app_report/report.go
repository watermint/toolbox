package app_report

import "github.com/watermint/toolbox/infra/ui/app_msg"

type ReportOpt func(o *ReportOpts) *ReportOpts
type ReportOpts struct {
	HiddenColumns  map[string]bool
	ShowAllColumns bool
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

func ShowAllColumns(enabled bool) ReportOpt {
	return func(o *ReportOpts) *ReportOpts {
		o.ShowAllColumns = enabled
		return o
	}
}

func HideColumns(col []string) ReportOpt {
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
	Failure(reason app_msg.Message, input interface{}, result interface{})
	Skip(reason app_msg.Message, input interface{}, result interface{})

	Close()
}

func TransactionHeader(input interface{}, result interface{}) TransactionRow {
	return TransactionRow{
		Input:  input,
		Result: result,
	}
}

type TransactionRow struct {
	Status string
	Reason string
	Input  interface{}
	Result interface{}
}

var (
	msgSuccess     = app_msg.M("report.transaction.success")
	msgFailure     = app_msg.M("report.transaction.failure")
	msgSkip        = app_msg.M("report.transaction.skip")
	MsgInvalidData = app_msg.M("report.transaction.failure.invalid_data")
)
