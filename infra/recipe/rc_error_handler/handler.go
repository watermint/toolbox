package rc_error_handler

import "github.com/watermint/toolbox/infra/ui/app_ui"

type ErrorHandler interface {
	// Handle an error. Returns true if the handler consumed the error.
	// Otherwise false. Caller (the framework) should pass to the error to the next handler.
	Handle(ui app_ui.UI, e error) bool
}
