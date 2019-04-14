package api_recpie_ctl

import "github.com/watermint/toolbox/app/app_recipe/api_recipe_msg"

type Controller interface {
	Shutdown()
	Fatal()
}

type UI interface {
	Info(key string, placeHolders ...api_recipe_msg.PlaceHolder)
	Error(key string, placeHolders ...api_recipe_msg.PlaceHolder)

	AskWarn(key string, placeHolders ...api_recipe_msg.PlaceHolder) (cont bool, cancel bool)
	AskText(key string, placeHolders ...api_recipe_msg.PlaceHolder) (text string, cancel bool)
}
