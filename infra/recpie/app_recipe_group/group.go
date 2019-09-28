package app_recipe_group

import (
	"bytes"
	"errors"
	"flag"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"sort"
	"strings"
)

func AppHeader(ui app_ui.UI) {
	ui.Header("run.app.header", app_msg.P{"AppVersion": app.Version})
	ui.Info("run.app.copyright")
	ui.Info("run.app.license")
	ui.Break()
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
	path, name := app_recipe.Path(r)

	z.addToPath(path, path, name, r)
}

func (z *Group) usageHeader(ui app_ui.UI, desc string) {
	AppHeader(ui)
	ui.Break()
	ui.Info(desc)
	ui.Break()
}

func (z *Group) PrintRecipeUsage(ui app_ui.UI, rcp app_recipe.Recipe, f *flag.FlagSet) {
	path, name := app_recipe.Path(rcp)
	z.usageHeader(ui, app_recipe.Desc(rcp).Key())

	ui.Header("run.recipe.header.usage")
	ui.Info(
		"run.recipe.usage",
		app_msg.P{
			"Exec":   os.Args[0],
			"Recipe": strings.Join(append(path, name), " "),
		},
	)

	ui.Break()
	ui.Header("run.recipe.header.available_flags")

	buf := new(bytes.Buffer)
	f.SetOutput(buf)
	f.PrintDefaults()
	ui.Info("raw", app_msg.P{"Raw": buf.String()})
	ui.Break()
}

func (z *Group) PrintUsage(ui app_ui.UI) {
	grpDesc := make([]string, 0)
	grpDesc = append(grpDesc, "recipe")
	grpDesc = append(grpDesc, z.Path...)
	grpDesc = append(grpDesc, "desc")

	z.usageHeader(ui, strings.Join(grpDesc, "."))

	ui.Header("run.group.header.usage")
	ui.Info(
		"run.group.usage",
		app_msg.P{
			"Exec":  os.Args[0],
			"Group": strings.Join(z.Path, " "),
		},
	)
	ui.Break()

	ui.Header("run.group.header.available_commands")
	cmdTable := ui.InfoTable("usage")
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
