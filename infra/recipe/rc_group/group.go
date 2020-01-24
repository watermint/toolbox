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

type Catalogue interface {
	Recipes() []rc_recipe.Recipe
	Ingredients() []rc_recipe.Recipe
	Messages() []interface{}
	RootGroup() Group
}

type catalogueImpl struct {
	recipes     []rc_recipe.Recipe
	ingredients []rc_recipe.Recipe
	messages    []interface{}
	root        Group
}

func (z *catalogueImpl) Recipes() []rc_recipe.Recipe {
	return z.recipes
}

func (z *catalogueImpl) Ingredients() []rc_recipe.Recipe {
	return z.ingredients
}

func (z *catalogueImpl) Messages() []interface{} {
	return z.messages
}

func (z *catalogueImpl) RootGroup() Group {
	return z.root
}

func NewCatalogue(recipes, ingredients []rc_recipe.Recipe, messages []interface{}) Catalogue {
	root := NewGroup([]string{}, "")
	for _, r := range recipes {
		root.Add(r)
	}

	return &catalogueImpl{
		recipes:     recipes,
		ingredients: ingredients,
		messages:    messages,
		root:        root,
	}
}

func NewEmptyCatalogue() Catalogue {
	return NewCatalogue([]rc_recipe.Recipe{}, []rc_recipe.Recipe{}, []interface{}{})
}

type MsgHeader struct {
	Header    app_msg.Message
	Copyright app_msg.Message
	License   app_msg.Message
}

type MsgGroup struct {
	RecipeHeaderUsage      app_msg.Message
	RecipeUsage            app_msg.Message
	RecipeAvailableFlags   app_msg.Message
	GroupHeaderUsage       app_msg.Message
	GroupUsage             app_msg.Message
	GroupAvailableCommands app_msg.Message
}

var (
	MHeader = app_msg.Apply(&MsgHeader{}).(*MsgHeader)
	MGroup  = app_msg.Apply(&MsgGroup{}).(*MsgGroup)
)

func AppHeader(ui app_ui.UI, version string) {
	ui.Header(MHeader.Header.With("AppVersion", version))
	ui.Info(MHeader.Copyright)
	ui.Info(MHeader.License)
	ui.Break()
}

type Group interface {
	Name() string
	BasePkg() string
	Path() []string
	Recipes() map[string]rc_recipe.Recipe
	SubGroups() map[string]Group
	Add(r rc_recipe.Recipe)
	AddToPath(fullPath []string, relPath []string, name string, r rc_recipe.Recipe)
	PrintRecipeUsage(ui app_ui.UI, spec rc_recipe.Spec, f *flag.FlagSet)
	PrintGroupUsage(ui app_ui.UI, exec, version string)
	GroupDesc() app_msg.Message
	CommandTitle(cmd string) app_msg.Message
	IsSecret() bool
	Select(args []string) (name string, g Group, r rc_recipe.Recipe, remainder []string, err error)
}

type GroupImpl struct {
	name      string
	basePkg   string
	path      []string
	recipes   map[string]rc_recipe.Recipe
	subGroups map[string]Group
}

func (z *GroupImpl) Name() string {
	return z.name
}

func (z *GroupImpl) BasePkg() string {
	return z.basePkg
}

func (z *GroupImpl) Path() []string {
	return z.path
}

func (z *GroupImpl) Recipes() map[string]rc_recipe.Recipe {
	return z.recipes
}

func (z *GroupImpl) SubGroups() map[string]Group {
	return z.subGroups
}

func (z *GroupImpl) GroupDesc() app_msg.Message {
	grpDesc := make([]string, 0)
	grpDesc = append(grpDesc, "recipe")
	grpDesc = append(grpDesc, z.path...)
	grpDesc = append(grpDesc, "title")

	return app_msg.M(strings.Join(grpDesc, "."))
}

func NewGroup(path []string, name string) Group {
	return &GroupImpl{
		name:      name,
		path:      path,
		recipes:   make(map[string]rc_recipe.Recipe),
		subGroups: make(map[string]Group),
	}
}

func (z *GroupImpl) AddToPath(fullPath []string, relPath []string, name string, r rc_recipe.Recipe) {
	if len(relPath) > 0 {
		p0 := relPath[0]
		sg, ok := z.subGroups[p0]
		if !ok {
			sg = NewGroup(fullPath, p0)
			z.subGroups[p0] = sg
		}
		sg.AddToPath(fullPath, relPath[1:], name, r)
	} else {
		z.recipes[name] = r
	}
}

func (z *GroupImpl) Add(r rc_recipe.Recipe) {
	path, name := rc_recipe.Path(r)

	z.AddToPath(path, path, name, r)
}

func (z *GroupImpl) usageHeader(ui app_ui.UI, desc app_msg.Message, version string) {
	AppHeader(ui, version)
	ui.Break()
	ui.Info(desc)
	ui.Break()
}

func (z *GroupImpl) PrintRecipeUsage(ui app_ui.UI, spec rc_recipe.Spec, f *flag.FlagSet) {
	z.usageHeader(ui, spec.Title(), app.Version)

	ui.Header(MGroup.RecipeHeaderUsage)
	ui.Info(MGroup.RecipeUsage.
		With("Exec", os.Args[0]).
		With("Recipe", spec.CliPath()).
		With("Args", ui.TextOrEmpty(spec.CliArgs())))

	ui.Break()
	ui.Header(MGroup.RecipeAvailableFlags)

	buf := new(bytes.Buffer)
	f.SetOutput(buf)
	f.PrintDefaults()
	ui.Info(app_msg.Raw(buf.String()))
	ui.Break()
}

func (z *GroupImpl) PrintGroupUsage(ui app_ui.UI, exec, version string) {
	z.usageHeader(ui, z.GroupDesc(), version)

	ui.Header(MGroup.GroupHeaderUsage)
	ui.Info(MGroup.GroupUsage.
		With("Exec", exec).
		With("Group", strings.Join(z.path, " ")))
	ui.Break()

	ui.Header(MGroup.GroupAvailableCommands)
	cmdTable := ui.InfoTable("usage")
	for _, cmd := range z.availableCommands() {
		cmdTable.Row(app_msg.Raw(" "), app_msg.Raw(cmd), z.CommandTitle(cmd))
	}
	cmdTable.Flush()
}

func (z *GroupImpl) CommandTitle(cmd string) app_msg.Message {
	keyPath := make([]string, 0)
	keyPath = append(keyPath, "recipe")
	keyPath = append(keyPath, z.path...)
	keyPath = append(keyPath, cmd)
	keyPath = append(keyPath, "title")
	key := strings.Join(keyPath, ".")

	return app_msg.M(key)
}

func (z *GroupImpl) IsSecret() bool {
	for _, r := range z.recipes {
		_, ok := r.(rc_recipe.SecretRecipe)
		if !ok {
			return false
		}
	}
	return true
}

func (z *GroupImpl) availableCommands() (cmd []string) {
	cmd = make([]string, 0)
	for _, g := range z.subGroups {
		if !g.IsSecret() {
			cmd = append(cmd, g.Name())
		}
	}
	for n, r := range z.recipes {
		_, ok := r.(rc_recipe.SecretRecipe)
		if !ok {
			cmd = append(cmd, n)
		}
	}
	sort.Strings(cmd)
	return
}

func (z *GroupImpl) Select(args []string) (name string, g Group, r rc_recipe.Recipe, remainder []string, err error) {
	if len(args) < 1 {
		return "", z, nil, args, nil
	}
	arg := args[0]
	for k, sg := range z.subGroups {
		if arg == k {
			return sg.Select(args[1:])
		}
	}
	for k, sr := range z.Recipes() {
		if arg == k {
			return k, z, sr, args[1:], nil
		}
	}
	return "", z, nil, args, errors.New("not found")
}
