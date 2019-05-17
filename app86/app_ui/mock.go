package app_ui

import "github.com/watermint/toolbox/app86/app_msg"

type MockUI struct {
}

func (*MockUI) Info(key string, placeHolders ...app_msg.Param) {
	panic("implement me")
}

func (*MockUI) Error(key string, placeHolders ...app_msg.Param) {
	panic("implement me")
}

func (*MockUI) AskCont(key string, placeHolders ...app_msg.Param) (cont bool, cancel bool) {
	panic("implement me")
}

func (*MockUI) AskText(key string, placeHolders ...app_msg.Param) (text string, cancel bool) {
	panic("implement me")
}
