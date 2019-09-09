package app_report

import "github.com/watermint/toolbox/infra/ui/app_msg"

type Report interface {
	// Report data row
	Row(row interface{})

	// Report transaction result
	Success(input interface{}, result interface{})
	Failure(reason app_msg.Message, input interface{}, result interface{})
	Skip(reason app_msg.Message, input interface{}, result interface{})

	Flush()
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

const (
	reportPath = "reports"
)

var (
	msgSuccess     = app_msg.M("report.transaction.success")
	msgFailure     = app_msg.M("report.transaction.failure")
	msgSkip        = app_msg.M("report.transaction.skip")
	MsgInvalidData = app_msg.M("report.transaction.failure.invalid_data")
)
