package app_apikey

import "github.com/watermint/toolbox/infra/ui/app_msg"

type MsgApiKey struct {
	AskBasicAuthKey app_msg.Message
}

var (
	MApiKey = app_msg.Apply(&MsgApiKey{}).(*MsgApiKey)
)
