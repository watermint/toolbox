package app_ui

import (
	"fmt"
	"github.com/watermint/toolbox/infra/report/rp_artifact"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"go.uber.org/atomic"
)

type Syntax interface {
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

	// Ask to continue. This confirmation step may be skipped on some UI implementation.
	// If you want to ensure actual user acknowledge, please use AskCont instead.
	AskProceed(m app_msg.Message)

	// Ask continue
	AskCont(m app_msg.Message) (cont bool)

	// Ask for a text
	AskText(m app_msg.Message) (text string, cancel bool)

	// Ask for a credentials
	AskSecure(m app_msg.Message) (secure string, cancel bool)

	// Display success message
	Success(m app_msg.Message)

	// Display failure message
	Failure(m app_msg.Message)

	// Display progress
	Progress(m app_msg.Message)

	// Code block
	Code(code string)

	// Link to artifact
	Link(artifact rp_artifact.Artifact)

	// True when the syntax is for console
	IsConsole() bool

	// True when the syntax is for web
	IsWeb() bool

	// New UI with given message syntax
	WithContainerSyntax(mc app_msg_container.Container) Syntax

	// Message container of this UI
	Messages() app_msg_container.Container
}

type UI interface {
	Syntax

	// Test existence of the message key
	Exists(m app_msg.Message) bool

	// Dealing with table. Table will be automatically closed after the f finished.
	WithTable(name string, f func(t Table))

	// Compile text
	Text(m app_msg.Message) string

	// Compile text, returns empty string if the key is not found
	TextOrEmpty(m app_msg.Message) string

	// Unique identifier of this UI
	Id() string

	// New UI with given message container
	WithContainer(mc app_msg_container.Container) UI
}

type Table interface {
	Header(h ...app_msg.Message)
	HeaderRaw(h ...string)
	Row(m ...app_msg.Message)
	RowRaw(m ...string)
	Flush()
}

var (
	latestId atomic.Int64
)

func newId() string {
	return fmt.Sprintf("%d", latestId.Add(1))
}
