package ut_doc

import "github.com/watermint/toolbox/infra/ui/app_msg"

type MsgDoc struct {
	AuthExampleHeaderSeparator app_msg.Message
}

var (
	MDoc = app_msg.Apply(&MsgDoc{}).(*MsgDoc)
)
