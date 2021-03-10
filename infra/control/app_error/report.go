package app_error

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	ErrorReportName = "operation_errors"
)

type ErrorReportRow struct {
	OperationName string          `json:"operation_name"`
	BatchId       string          `json:"batch_id"`
	Data          json.RawMessage `json:"data"`
	Error         string          `json:"error"`
}

type ErrorReport interface {
	Up(ctl app_control.Control) error
	Down()
	ErrorListener(err error, mouldId, batchId string, p interface{})
}
