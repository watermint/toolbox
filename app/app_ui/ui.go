package app_ui

import "github.com/watermint/toolbox/app/app_msg"

type UI interface {
	// Tell message
	Tell(msg app_msg.UIMessage)

	// Tell error message
	TellError(msg app_msg.UIMessage)

	// Tell done
	TellDone(msg app_msg.UIMessage)

	// Tell success
	TellSuccess(msg app_msg.UIMessage)

	// Tell failure
	TellFailure(msg app_msg.UIMessage)

	// Ask retry with retry message. Returns true when
	// the user/client agreed retry
	AskRetry(msg app_msg.UIMessage) bool

	// Ask a text. UI ask text as required option but,
	// a user/client can enter empty string.
	AskText(msg app_msg.UIMessage) string
}
