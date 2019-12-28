package rc_group

import (
	"bytes"
	"errors"
	"flag"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"sort"
	"strings"
)

type Catalogue struct {
	Recipes     []rc_recipe.Recipe
	Ingredients []rc_recipe.Recipe
	Messages    []interface{}
	RootGroup   *Group
}

type MsgHeader struct {
	Header    app_msg.Message
	Copyright app_msg.Message
	License   app_msg.Message
}

var (
	MHeader = app_msg.Apply(&MsgHeader{}).(*MsgHeader)
)

func AppHeader(ui app_ui.UI) {
	ui.Header(MHeader.Header.With("AppVersion", app.Version))
	ui.Info(MHeader.Copyright)
	ui.Info(MHeader.License)
	ui.Break()
}

type Group struct {
	Name      string
	BasePkg   string
	Path      []string
	Recipes   map[string]rc_recipe.Recipe
	SubGroups map[string]*Group
}

func NewGroup(path []string, name string) *Group {
	return &Group{
		Name:      name,
		Path:      path,
		Recipes:   make(map[string]rc_recipe.Recipe),
		SubGroups: make(map[string]*Group),
	}
}

func (z *Group) addToPath(fullPath []string, relPath []string, name string, r rc_recipe.Recipe) {
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

func (z *Group) Add(r rc_recipe.Recipe) {
	path, name := rc_recipe.Path(r)

	z.addToPath(path, path, name, r)
}

func (z *Group) usageHeader(ui app_ui.UI, desc string) {
	AppHeader(ui)
	ui.Break()
	ui.InfoK(desc)
	ui.Break()
}

func (z *Group) PrintRecipeUsage(ui app_ui.UI, spec rc_recipe.Spec, f *flag.FlagSet) {
	z.usageHeader(ui, spec.Title().Key())

	ui.HeaderK("run.recipe.header.usage")
	ui.InfoK(
		"run.recipe.usage",
		app_msg.P{
			"Exec":   os.Args[0],
			"Recipe": spec.CliPath(),
			"Args":   ui.TextOrEmptyK(spec.CliArgs().Key()),
		},
	)

	ui.Break()
	ui.HeaderK("run.recipe.header.available_flags")

	buf := new(bytes.Buffer)
	f.SetOutput(buf)
	f.PrintDefaults()
	ui.InfoK("raw", app_msg.P{"Raw": buf.String()})
	ui.Break()
}

func (z *Group) PrintUsage(ui app_ui.UI) {
	grpDesc := make([]string, 0)
	grpDesc = append(grpDesc, "recipe")
	grpDesc = append(grpDesc, z.Path...)
	grpDesc = append(grpDesc, "title")

	z.usageHeader(ui, strings.Join(grpDesc, "."))

	ui.HeaderK("run.group.header.usage")
	ui.InfoK(
		"run.group.usage",
		app_msg.P{
			"Exec":  os.Args[0],
			"Group": strings.Join(z.Path, " "),
		},
	)
	ui.Break()

	ui.HeaderK("run.group.header.available_commands")
	cmdTable := ui.InfoTable("usage")
	for _, cmd := range z.AvailableCommands() {
		cmdTable.Row(app_msg.Raw(" "), app_msg.Raw(cmd), z.CommandTitle(cmd))
	}
	cmdTable.Flush()
}

func (z *Group) CommandTitle(cmd string) app_msg.Message {
	keyPath := make([]string, 0)
	keyPath = append(keyPath, "recipe")
	keyPath = append(keyPath, z.Path...)
	keyPath = append(keyPath, cmd)
	keyPath = append(keyPath, "title")
	key := strings.Join(keyPath, ".")

	return app_msg.M(key)
}

func (z *Group) IsSecret() bool {
	for _, r := range z.Recipes {
		_, ok := r.(rc_recipe.SecretRecipe)
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
		_, ok := r.(rc_recipe.SecretRecipe)
		if !ok {
			cmd = append(cmd, n)
		}
	}
	sort.Strings(cmd)
	return
}

func (z *Group) Select(args []string) (name string, g *Group, r rc_recipe.Recipe, remainder []string, err error) {
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
