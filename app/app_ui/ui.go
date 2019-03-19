package app_ui

type UI interface {
	// Switch to debug mode
	DebugMode(debug bool)

	// Tell message
	Tell(msg UIMessage)

	// Tell error message
	TellError(msg UIMessage)

	// Tell success
	TellSuccess(msg UIMessage)

	// Tell failure
	TellFailure(msg UIMessage)

	// Ask retry with retry message. Returns true when
	// the user/client agreed retry
	AskRetry(msg UIMessage) bool

	// Ask a text. UI ask text as required option but,
	// a user/client can enter empty string.
	AskText(msg UIMessage) string

	// Ask confirm proceed operation.
	// Returns true when the user/client agreed proceed.
	AskConfirm(msg UIMessage) bool
}
