package app_ui

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type UI interface {
	// Deprecated: use Header
	HeaderK(key string, p ...app_msg.P)

	// Header
	Header(m app_msg.Message)
	// Deprecated: use Info
	InfoK(key string, p ...app_msg.P)

	// Info
	Info(m app_msg.Message)

	// Create information table
	InfoTable(name string) Table
	// Deprecated: use Error
	ErrorK(key string, p ...app_msg.P)

	// Error
	Error(m app_msg.Message)

	// Break
	Break()

	// Deprecated: use Text
	TextK(key string, p ...app_msg.P) string

	// Compile text
	Text(m app_msg.Message) string

	// Deprecated: use TextOrEmpty
	TextOrEmptyK(key string, p ...app_msg.P) string

	// Compile text, returns empty string if the key is not found
	TextOrEmpty(m app_msg.Message) string

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
