package app_ui

import "github.com/watermint/toolbox/app86/app_msg"

type Quiet struct {
}

func (*Quiet) Break() {
}

func (*Quiet) Header(key string, p ...app_msg.Param) {
}

func (*Quiet) InfoTable(border bool) Table {
	panic("mock")
}

func (*Quiet) Info(key string, p ...app_msg.Param) {
}

func (*Quiet) Error(key string, p ...app_msg.Param) {
}

// always cancel process
func (*Quiet) AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool) {
	return false, true
}

// always cancel
func (*Quiet) AskText(key string, p ...app_msg.Param) (text string, cancel bool) {
	return "", true
}
