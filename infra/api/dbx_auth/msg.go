package dbx_auth

import "github.com/watermint/toolbox/infra/ui/app_msg"

type MsgCcAuth struct {
	FailedOrCancelled app_msg.Message
	VerifyFailed      app_msg.Message
	GeneratedToken1   app_msg.Message
	GeneratedToken2   app_msg.Message
	OauthSeq1         app_msg.Message
	OauthSeq2         app_msg.Message
}

var (
	MCcAuth = app_msg.Apply(&MsgCcAuth{}).(*MsgCcAuth)
)
