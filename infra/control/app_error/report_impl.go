package app_error

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_model_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type MsgErrorReport struct {
	ErrorOneOrMoreOperationErrors app_msg.Message
}

var (
	MErrorReport = app_msg.Apply(&MsgErrorReport{}).(*MsgErrorReport)
)

func NewErrorReport(lg esl.Logger, wb app_workspace.Bundle, ui app_ui.UI) ErrorReport {
	return &errorReportImpl{
		lg: lg,
		ui: ui,
		wb: wb,
		rr: rp_model_impl.NewRowReport(ErrorReportName),
	}
}

type errorReportImpl struct {
	lg esl.Logger
	ui app_ui.UI
	wb app_workspace.Bundle
	rr *rp_model_impl.RowReport
}

func (z *errorReportImpl) Up(ctl app_control.Control) error {
	z.rr.SetCtl(ctl)
	z.rr.SetModel(&ErrorReportRow{})
	return z.rr.Open(rp_model.NoConsoleOutput())
}

func (z *errorReportImpl) Down() {
	if x := z.rr.Rows(); x > 0 {
		z.ui.Error(MErrorReport.ErrorOneOrMoreOperationErrors.With("Errors", x).With("ReportPath", z.wb.Workspace().Report()))
	}
	z.rr.Close()
}

func (z *errorReportImpl) ErrorHandler(err error, mouldId, batchId string, p interface{}) {
	if err == qt_errors.ErrorMock {
		return
	}
	d, em := json.Marshal(p)
	if em != nil {
		z.lg.Debug("Unable to marshal", esl.Error(err))
		z.rr.Row(&ErrorReportRow{
			OperationName: mouldId,
			BatchId:       batchId,
			Data:          nil,
			Error:         err.Error(),
		})
	} else {
		z.rr.Row(&ErrorReportRow{
			OperationName: mouldId,
			BatchId:       batchId,
			Data:          d,
			Error:         err.Error(),
		})
	}
}
