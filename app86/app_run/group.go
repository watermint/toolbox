package app_run

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/app86/app_recipe"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

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

func (z *Group) Desc() string {
	key := "recipe." + strings.Join(z.Path, ".") + ".desc"
	//TODO: convert into txt resource
	return key
}

func (z *Group) Usage() string {
	// TODO: try lookup custom usage
	//key := strings.Join(z.Path, ".") + ".usage"

	usage := os.Args[0] + " " + strings.Join(z.Path, " ") + " [command]"

	return usage
}

func (z *Group) printIndented(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Println("  " + line)
	}
}

func (z *Group) PrintUsage() {
	fmt.Println(AppHeader())
	fmt.Println()
	z.printIndented(z.Desc())
	fmt.Println()

	usageHeader := "Usage:" // TODO: i18n
	fmt.Println(usageHeader)
	z.printIndented(z.Usage())
	fmt.Println()

	availableHeader := "Available Commands:" // TODO: i18n
	fmt.Println(availableHeader)
	aw := new(tabwriter.Writer)
	aw.Init(os.Stdout, 0, 2, 1, ' ', 0)
	for _, cmd := range z.AvailableCommands() {
		fmt.Fprintln(aw, "\t"+cmd+"\t"+z.commandDesc(cmd))
	}
	fmt.Fprintln(aw, "  \t")
	aw.Flush()

}

func (z *Group) commandDesc(cmd string) string {
	keyPath := make([]string, 0)
	keyPath = append(keyPath, "recipe")
	keyPath = append(keyPath, z.Path...)
	keyPath = append(keyPath, cmd)
	keyPath = append(keyPath, "desc")
	key := strings.Join(keyPath, ".")

	// TODO:
	return key
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

func (z *Group) Select(args []string) (g *Group, r app_recipe.Recipe, remainder []string, err error) {
	if len(args) < 1 {
		return z, nil, args, nil
	}
	arg := args[0]
	for k, sg := range z.SubGroups {
		if arg == k {
			return sg.Select(args[1:])
		}
	}
	for k, sr := range z.Recipes {
		if arg == k {
			return nil, sr, args[1:], nil
		}
	}
	return z, nil, args, errors.New("not found")
}
