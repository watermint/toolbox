package app_ui

import "github.com/watermint/toolbox/app86/app_msg"

type Console struct {
}

func (z *Console) Info(key string, p ...app_msg.Param) {
	app_msg.M(key, p...)
}

func (z *Console) Error(key string, p ...app_msg.Param) {
	panic("implement me")
}

func (z *Console) AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool) {
	panic("implement me")
}

func (z *Console) AskText(key string, p ...app_msg.Param) (text string, cancel bool) {
	panic("implement me")
}
