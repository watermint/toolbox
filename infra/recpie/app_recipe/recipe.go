package app_recipe

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"reflect"
	"strings"
)

const (
	BasePackage = "github.com/watermint/toolbox/recipe"
)

type Recipe interface {
	Requirement() app_vo.ValueObject
	Exec(k app_kitchen.Kitchen) error
	Test(c app_control.Control) error
}

// SecretRecipe will not be listed in available commands.
type SecretRecipe interface {
	Hidden()
}

// Console only recipe will not be listed in web console.
type ConsoleRecipe interface {
	Console()
}

func Desc(r Recipe) app_msg.Message {
	path, name := Path(r)
	keyPath := make([]string, 0)
	keyPath = append(keyPath, "recipe")
	keyPath = append(keyPath, path...)
	keyPath = append(keyPath, name)
	keyPath = append(keyPath, "desc")
	key := strings.Join(keyPath, ".")

	return app_msg.M(key)
}

func Path(r Recipe) (path []string, name string) {
	path = make([]string, 0)

	rt := reflect.ValueOf(r).Elem().Type()
	pkg := rt.PkgPath()
	pkg = strings.ReplaceAll(pkg, BasePackage, "")
	if strings.HasPrefix(pkg, "/") {
		pkg = pkg[1:]
	}
	if pkg != "" {
		path = append(path, strings.Split(pkg, "/")...)
	}
	return path, strings.ToLower(rt.Name())
}
