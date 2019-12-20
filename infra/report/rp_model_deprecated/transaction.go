package rp_model_deprecated

import (
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func rowForFailure(ui app_ui.UI, err error, input interface{}) *rp_model.TransactionRow {
	reason := api_util.MsgFromError(err)
	if ui.TextOrEmpty(reason.Key()) == "" {
		summary := api_util.ErrorSummary(err)
		if summary == "" {
			summary = err.Error()
		}
		reason = app_msg.M("dbx.err.general_error", app_msg.P{"Error": summary})
	}
	return &rp_model.TransactionRow{
		Status: ui.Text(rp_model.MsgFailure.Key(), rp_model.MsgFailure.Params()...),
		Reason: ui.Text(reason.Key(), reason.Params()...),
		Input:  input,
		Result: nil,
	}
}
