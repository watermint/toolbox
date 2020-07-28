package sv_message

import "github.com/watermint/toolbox/infra/ui/app_msg"

type MsgProgress struct {
	ProgressRetrieve app_msg.Message
}

var (
	MProgress = app_msg.Apply(&MsgProgress{}).(*MsgProgress)
)
