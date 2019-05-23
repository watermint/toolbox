package app_ui

import (
	"github.com/watermint/toolbox/app86/app_msg"
)

type UI interface {
	Info(key string, p ...app_msg.Param)
	Error(key string, p ...app_msg.Param)

	AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool)
	AskText(key string, p ...app_msg.Param) (text string, cancel bool)
}
