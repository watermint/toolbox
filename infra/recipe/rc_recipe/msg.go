package rc_recipe

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
	"strings"
)

func RecipeMessage(r Recipe, suffix string) app_msg.Message {
	path, name := Path(r)
	keyPath := make([]string, 0)
	keyPath = append(keyPath, "recipe")
	keyPath = append(keyPath, path...)
	keyPath = append(keyPath, name)
	keyPath = append(keyPath, suffix)
	key := strings.Join(keyPath, ".")
	return app_msg.M(key)
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
	return ut_reflect.Key(BasePackage, r)
}
