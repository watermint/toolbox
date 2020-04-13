package dbx_conn_impl

import "github.com/watermint/toolbox/infra/ui/app_msg"

type MsgConnect struct {
	VerifySuccess app_msg.Message
}

var (
	MConnect = app_msg.Apply(&MsgConnect{}).(*MsgConnect)
)
