package app_ui

import (
	"github.com/watermint/toolbox/app86/app_msg"
)

type UI interface {
	Info(key string, placeHolders ...app_msg.PlaceHolder)
	Error(key string, placeHolders ...app_msg.PlaceHolder)

	AskCont(key string, placeHolders ...app_msg.PlaceHolder) (cont bool, cancel bool)
	AskText(key string, placeHolders ...app_msg.PlaceHolder) (text string, cancel bool)
}
