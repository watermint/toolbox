package oper_ui

import (
	"github.com/watermint/toolbox/poc/oper/oper_i18n"
	"os"
)

type UITable interface {
	Header() []oper_i18n.UIMessage
	Row(row int) []oper_i18n.UIMessage
	RowCount() int
}

type UI interface {
	// Tell message
	Tell(msg oper_i18n.UIMessage)

	// Tell error message
	TellError(msg oper_i18n.UIMessage)

	// Tell table messages
	TellTable(tbl UITable)

	// Tell done
	TellDone(msg oper_i18n.UIMessage)

	// Tell success
	TellSuccess(msg oper_i18n.UIMessage)

	// Tell failure
	TellFailure(msg oper_i18n.UIMessage)

	// Tell progress (text)
	TellProgress(msg oper_i18n.UIMessage)

	// Ask retry with retry message. Returns true when
	// the user/client agreed retry
	AskRetry(msg oper_i18n.UIMessage) bool

	// Ask continue with warning message. Returns true when
	// the user/client agreed to proceed
	AskWarn(msg oper_i18n.UIMessage) bool

	// Ask options. Returns selected option key.
	AskOptions(title oper_i18n.UIMessage, opts map[string]oper_i18n.UIMessage) string

	// Ask a file for input. Returns file if the user choose file.
	// Returns nil when the user canceled selection or file not found.
	AskInputFile(msg oper_i18n.UIMessage) *os.File

	// Ask a file for output.
	AskOutputFile(msg oper_i18n.UIMessage, filename string, tmpFilePath string)

	// Ask a text. UI ask text as required option but,
	// a user/client can enter empty string.
	AskText(msg oper_i18n.UIMessage) string
}
