package dbx_error

import (
	"github.com/watermint/toolbox/infra/recipe/rc_error_handler"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgHandler struct {
	ErrorSummary    app_msg.Message
	ErrorBadRequest app_msg.Message
}

var (
	MHandler = app_msg.Apply(&MsgHandler{}).(*MsgHandler)
)

func NewHandler() rc_error_handler.ErrorHandler {
	return &handlerImpl{}
}

type handlerImpl struct {
}

func (z handlerImpl) Handle(ui app_ui.UI, e error) bool {
	if ebr, ok := e.(*ErrorBadRequest); ok {
		ui.Error(MHandler.ErrorBadRequest.With("Error", ebr.Reason))
		return true
	}
	de := NewErrors(e)
	if de == nil {
		return false
	}
	ui.Error(MHandler.ErrorSummary.With("Error", de.Summary()))
	return true
}
