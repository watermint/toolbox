package app_ui

import "github.com/watermint/toolbox/app86/app_msg"

type Mock struct {
}

func (*Mock) Info(key string, p ...app_msg.Param) {
}

func (*Mock) Error(key string, p ...app_msg.Param) {
}

// always cancel process
func (*Mock) AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool) {
	return false, true
}

// always cancel
func (*Mock) AskText(key string, p ...app_msg.Param) (text string, cancel bool) {
	return "", true
}
