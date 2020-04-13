package app_ui

import (
	"fmt"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

type UI interface {
	// Deprecated: use Header
	HeaderK(key string, p ...app_msg.P)

	// Header
	Header(m app_msg.Message)

	// Sub header
	SubHeader(m app_msg.Message)

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

	// Compile text, returns empty string if the key is not found
	TextOrEmpty(m app_msg.Message) string

	// Ask to continue. This confirmation step may be skipped on some UI implementation.
	// If you want to ensure actual user acknowledge, please use AskCont instead.
	AskProceed(m app_msg.Message)

	// Ask continue
	AskCont(m app_msg.Message) (cont bool, cancel bool)

	// Ask for a text
	AskText(m app_msg.Message) (text string, cancel bool)

	// Ask for a credentials
	AskSecure(m app_msg.Message) (secure string, cancel bool)

	OpenArtifact(path string, autoOpen bool)

	// Deprecated: use Success
	SuccessK(key string, p ...app_msg.P)

	// Deprecated: use Failure
	FailureK(key string, p ...app_msg.P)
	Success(m app_msg.Message)
	Failure(m app_msg.Message)
	Progress(m app_msg.Message)

	Code(code string)

	IsConsole() bool
	IsWeb() bool

	// Unique identifier of this UI
	Id() string
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

var (
	latestId atomic.Int64
)

func newId() string {
	return fmt.Sprintf("%d", latestId.Add(1))
}
