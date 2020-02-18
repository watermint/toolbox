package app_ui

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
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

	// Test existence of the message key
	Exists(m app_msg.Message) bool

	// Deprecated: use Text
	TextK(key string, p ...app_msg.P) string

	// Compile text
	Text(m app_msg.Message) string

	// Deprecated: use TextOrEmpty
	TextOrEmptyK(key string, p ...app_msg.P) string

	// Compile text, returns empty string if the key is not found
	TextOrEmpty(m app_msg.Message) string

	// Deprecated: use AskCont
	AskContK(key string, p ...app_msg.P) (cont bool, cancel bool)

	// Ask continue
	AskCont(m app_msg.Message) (cont bool, cancel bool)

	// Deprecated: use AskText
	AskTextK(key string, p ...app_msg.P) (text string, cancel bool)

	// Ask for a text
	AskText(m app_msg.Message) (text string, cancel bool)

	// Deprecated: use AskSecure
	AskSecureK(key string, p ...app_msg.P) (secure string, cancel bool)

	// Ask for a credentials
	AskSecure(m app_msg.Message) (secure string, cancel bool)

	OpenArtifact(path string, autoOpen bool)

	// Deprecated: use Success
	SuccessK(key string, p ...app_msg.P)

	// Deprecated: use Failure
	FailureK(key string, p ...app_msg.P)
	Success(m app_msg.Message)
	Failure(m app_msg.Message)

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

type UILog interface {
	SetLogger(l *zap.Logger)
}
