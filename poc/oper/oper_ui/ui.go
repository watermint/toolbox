package oper_ui

import "github.com/watermint/toolbox/poc/oper/oper_msg"

type UI interface {
	// Tell message
	Tell(msg oper_msg.UIMessage)

	// Tell error message
	TellError(msg oper_msg.UIMessage)

	// Tell done
	TellDone(msg oper_msg.UIMessage)

	// Tell success
	TellSuccess(msg oper_msg.UIMessage)

	// Tell failure
	TellFailure(msg oper_msg.UIMessage)

	// Tell progress (text)
	TellProgress(msg oper_msg.UIMessage)

	// Ask retry with retry message. Returns true when
	// the user/client agreed retry
	AskRetry(msg oper_msg.UIMessage) bool

	// Ask continue with warning message. Returns true when
	// the user/client agreed to proceed
	AskWarn(msg oper_msg.UIMessage) bool

	// Ask a text. UI ask text as required option but,
	// a user/client can enter empty string.
	AskText(msg oper_msg.UIMessage) string
}
