package app_ui

import (
	"fmt"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

type UI interface {
	// Header
	Header(m app_msg.Message)

	// Sub header
	SubHeader(m app_msg.Message)

	// Info
	Info(m app_msg.Message)

	// Create information table
	InfoTable(name string) Table

	// Error
	Error(m app_msg.Message)

	// Break
	Break()

	// Test existence of the message key
	Exists(m app_msg.Message) bool

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
