package app_run

import (
	"github.com/watermint/toolbox/app86/app_recipe"
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

func recipes() []app_recipe.Recipe {
	return []app_recipe.Recipe{
		&recipe.License{},
		&dev.Version{},
		&member.Invite{},
		&teamfolder.List{},
	}
}

func catalogue() *Group {
	root := NewGroup([]string{}, "")
	for _, r := range recipes() {
		root.Add(r)
	}
	return root
}

func AppHeader() string {
	// TODO: Convert it into Message
	return "{{.AppName}} ({{.AppVersion}})\n" +
		"Licensed under MIT License. etc, etc"
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
