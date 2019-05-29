package app_ui

import (
	"github.com/watermint/toolbox/app86/app_msg"
)

type UI interface {
	Header(key string, p ...app_msg.Param)
	Info(key string, p ...app_msg.Param)
	InfoTable(border bool) Table
	Error(key string, p ...app_msg.Param)
	Break()

	AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool)
	AskText(key string, p ...app_msg.Param) (text string, cancel bool)
}

type Table interface {
	Header(h ...app_msg.Message)
	HeaderRaw(h ...string)
	Row(m ...app_msg.Message)
	RowRaw(m ...string)
	Flush()
}
