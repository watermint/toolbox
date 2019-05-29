package app_run

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_ui"
	"github.com/watermint/toolbox/app86/recipe"
	"github.com/watermint/toolbox/app86/recipe/dev"
	"github.com/watermint/toolbox/app86/recipe/member"
	"github.com/watermint/toolbox/app86/recipe/teamfolder"
	"reflect"
	"strings"
)

const (
	RecipeBasePackage = "github.com/watermint/toolbox/app86/recipe"
)

func Recipes() []app_recipe.Recipe {
	return []app_recipe.Recipe{
		&recipe.License{},
		&dev.LongRunning{},
		&member.Invite{},
		&member.List{},
		&teamfolder.List{},
	}
}

func Catalogue() *Group {
	root := NewGroup([]string{}, "")
	for _, r := range Recipes() {
		root.Add(r)
	}
	return root
}

func AppHeader(ui app_ui.UI) {
	ui.Header("run.app.header", app_msg.P("AppVersion", app.AppVersion))
	ui.Info("run.app.copyright")
	ui.Info("run.app.license")
	ui.Break()
}

func RecipeInfo(basePkg string, r app_recipe.Recipe) (cmdPath []string, cmdName string) {
	cmdPath = make([]string, 0)

	rt := reflect.ValueOf(r).Elem().Type()
	pkg := rt.PkgPath()
	pkg = strings.ReplaceAll(pkg, basePkg, "")
	if strings.HasPrefix(pkg, "/") {
		pkg = pkg[1:]
	}
	if pkg != "" {
		cmdPath = append(cmdPath, strings.Split(pkg, "/")...)
	}
	name := rt.Name()

	return cmdPath, strings.ToLower(name)
}
