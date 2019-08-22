package app_recipe_group

import (
	"errors"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"reflect"
	"sort"
	"strings"
)

const (
	RecipeBasePackage = "github.com/watermint/toolbox/recipe"
)

func AppHeader(ui app_ui.UI) {
	ui.Header("run.app.header", app_msg.P("AppVersion", app.Version))
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

type Group struct {
	Name      string
	BasePkg   string
	Path      []string
	Recipes   map[string]app_recipe.Recipe
	SubGroups map[string]*Group
}

func NewGroup(path []string, name string) *Group {
	return &Group{
		Name:      name,
		BasePkg:   RecipeBasePackage,
		Path:      path,
		Recipes:   make(map[string]app_recipe.Recipe),
		SubGroups: make(map[string]*Group),
	}
}

func (z *Group) addToPath(fullPath []string, relPath []string, name string, r app_recipe.Recipe) {
	if len(relPath) > 0 {
		p0 := relPath[0]
		sg, ok := z.SubGroups[p0]
		if !ok {
			sg = NewGroup(fullPath, p0)
			z.SubGroups[p0] = sg
		}
		sg.addToPath(fullPath, relPath[1:], name, r)
	} else {
		z.Recipes[name] = r
	}
}

func (z *Group) Add(r app_recipe.Recipe) {
	path, name := RecipeInfo(z.BasePkg, r)

	z.addToPath(path, path, name, r)
}

func (z *Group) PrintUsage(ui app_ui.UI) {
	grpDesc := make([]string, 0)
	grpDesc = append(grpDesc, "recipe")
	grpDesc = append(grpDesc, z.Path...)
	grpDesc = append(grpDesc, "desc")

	AppHeader(ui)
	ui.Break()
	ui.Info(strings.Join(grpDesc, "."))
	ui.Break()

	ui.Header("run.group.header.usage")
	ui.Info(
		"run.group.usage",
		app_msg.P("Exec", os.Args[0]),
		app_msg.P("Group", strings.Join(z.Path, " ")),
	)
	ui.Break()

	ui.Header("run.group.header.available_commands")
	cmdTable := ui.InfoTable(false)
	for _, cmd := range z.AvailableCommands() {
		cmdTable.Row(app_msg.Raw(" "), app_msg.Raw(cmd), z.CommandDesc(cmd))
	}
	cmdTable.Flush()
}

func (z *Group) CommandDesc(cmd string) app_msg.Message {
	keyPath := make([]string, 0)
	keyPath = append(keyPath, "recipe")
	keyPath = append(keyPath, z.Path...)
	keyPath = append(keyPath, cmd)
	keyPath = append(keyPath, "desc")
	key := strings.Join(keyPath, ".")

	return app_msg.M(key)
}

func (z *Group) IsSecret() bool {
	for _, r := range z.Recipes {
		_, ok := r.(app_recipe.SecretRecipe)
		if !ok {
			return false
		}
	}
	return true
}

func (z *Group) AvailableCommands() (cmd []string) {
	cmd = make([]string, 0)
	for _, g := range z.SubGroups {
		if !g.IsSecret() {
			cmd = append(cmd, g.Name)
		}
	}
	for n, r := range z.Recipes {
		_, ok := r.(app_recipe.SecretRecipe)
		if !ok {
			cmd = append(cmd, n)
		}
	}
	sort.Strings(cmd)
	return
}

func (z *Group) Select(args []string) (name string, g *Group, r app_recipe.Recipe, remainder []string, err error) {
	if len(args) < 1 {
		return "", z, nil, args, nil
	}
	arg := args[0]
	for k, sg := range z.SubGroups {
		if arg == k {
			return sg.Select(args[1:])
		}
	}
	for k, sr := range z.Recipes {
		if arg == k {
			return k, z, sr, args[1:], nil
		}
	}
	return "", z, nil, args, errors.New("not found")
}
