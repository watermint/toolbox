package rp_model_impl

import "github.com/watermint/toolbox/infra/ui/app_msg"

type MsgTransactionReport struct {
	Success      app_msg.Message
	Failure      app_msg.Message
	Skip         app_msg.Message
	ErrorGeneral app_msg.Message
}

var (
	MTransactionReport = app_msg.Apply(&MsgTransactionReport{}).(*MsgTransactionReport)
)
