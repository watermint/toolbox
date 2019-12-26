package rc_recipe

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
)

func RecipeMessage(r Recipe, suffix string) app_msg.Message {
	return app_msg.M(Key(r) + "." + suffix)
}

func Title(r Recipe) app_msg.Message {
	return RecipeMessage(r, "title")
}

func Desc(r Recipe) app_msg.Message {
	return RecipeMessage(r, "desc")
}

func Path(r Recipe) (path []string, name string) {
	return ut_reflect.Path(BasePackage, r)
}

func Key(r Recipe) string {
	return ut_reflect.Key(app.Pkg, r)
}
