package app_ui

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type UI interface {
	Header(key string, p ...app_msg.P)
	Info(key string, p ...app_msg.P)
	InfoTable(name string) Table
	Error(key string, p ...app_msg.P)
	Break()
	Text(key string, p ...app_msg.P) string
	TextOrEmpty(key string, p ...app_msg.P) string

	AskCont(key string, p ...app_msg.P) (cont bool, cancel bool)
	AskText(key string, p ...app_msg.P) (text string, cancel bool)
	AskSecure(key string, p ...app_msg.P) (secure string, cancel bool)

	OpenArtifact(path string)
	Success(key string, p ...app_msg.P)
	Failure(key string, p ...app_msg.P)

	IsConsole() bool
	IsWeb() bool
}

type Table interface {
	Header(h ...app_msg.Message)
	HeaderRaw(h ...string)
	Row(m ...app_msg.Message)
	RowRaw(m ...string)
	Flush()
}
