package app_recipe

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
	"strings"
)

const (
	BasePackage = app.Pkg + "/recipe"
)

type Recipe interface {
	Requirement() app_vo.ValueObject
	Exec(k app_kitchen.Kitchen) error
	Test(c app_control.Control) error
	Reports() []rp_spec.ReportSpec
}

// SecretRecipe will not be listed in available commands.
type SecretRecipe interface {
	Hidden()
}

// Console only recipe will not be listed in web console.
type ConsoleRecipe interface {
	Console()
}

func Title(r Recipe) app_msg.Message {
	path, name := Path(r)
	keyPath := make([]string, 0)
	keyPath = append(keyPath, "recipe")
	keyPath = append(keyPath, path...)
	keyPath = append(keyPath, name)
	keyPath = append(keyPath, "title")
	key := strings.Join(keyPath, ".")

	return app_msg.M(key)
}

func Path(r Recipe) (path []string, name string) {
	return ut_reflect.Path(BasePackage, r)
}

func Key(r Recipe) string {
	return ut_reflect.Key(BasePackage, r)
}
