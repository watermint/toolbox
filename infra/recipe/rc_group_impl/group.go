package rc_group_impl

import (
	"errors"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sort"
	"strings"
)

type MsgGroup struct {
	GroupHeaderUsage          app_msg.Message
	GroupUsage                app_msg.Message
	GroupAvailableCommands    app_msg.Message
	HeaderAvailCommandCommand app_msg.Message
	HeaderAvailCommandDesc    app_msg.Message
	HeaderAvailCommandNote    app_msg.Message
}

var (
	MGroup = app_msg.Apply(&MsgGroup{}).(*MsgGroup)
)

func NewGroup() rc_group.Group {
	return newGroupWithPath([]string{}, "")
}

func newGroupWithPath(path []string, name string) rc_group.Group {
	return &groupImpl{
		name:      name,
		path:      path,
		recipes:   make(map[string]rc_recipe.Spec),
		subGroups: make(map[string]rc_group.Group),
	}
}

type groupImpl struct {
	name      string
	path      []string
	recipes   map[string]rc_recipe.Spec
	subGroups map[string]rc_group.Group
}

func (z *groupImpl) Name() string {
	return z.name
}

func (z *groupImpl) Path() []string {
	return z.path
}

func (z *groupImpl) Recipes() map[string]rc_recipe.Spec {
	return z.recipes
}

func (z *groupImpl) SubGroups() map[string]rc_group.Group {
	return z.subGroups
}

func (z *groupImpl) GroupDesc() app_msg.Message {
	grpDesc := make([]string, 0)
	grpDesc = append(grpDesc, "recipe")
	grpDesc = append(grpDesc, z.path...)
	grpDesc = append(grpDesc, "title")

	return app_msg.CreateMessage(strings.Join(grpDesc, "."))
}

func (z *groupImpl) AddToPath(fullPath []string, relPath []string, name string, r rc_recipe.Spec) {
	if len(relPath) > 0 {
		p0 := relPath[0]
		sg, ok := z.subGroups[p0]
		if !ok {
			sp := make([]string, 0)
			sp = append(sp, z.path...)
			sp = append(sp, p0)
			sg = newGroupWithPath(sp, p0)
			z.subGroups[p0] = sg
		}
		sg.AddToPath(fullPath, relPath[1:], name, r)
	} else {
		z.recipes[name] = r
	}
}

func (z *groupImpl) Add(r rc_recipe.Spec) {
	path, name := r.Path()

	z.AddToPath(path, path, name, r)
}

func (z *groupImpl) PrintUsage(ui app_ui.UI, exec, version string) {
	rc_group.UsageHeader(ui, z.GroupDesc(), version)

	ui.Header(MGroup.GroupHeaderUsage)
	ui.Info(MGroup.GroupUsage.
		With("Exec", exec).
		With("Group", strings.Join(z.path, " ")))
	ui.Break()

	ui.Header(MGroup.GroupAvailableCommands)
	cmdTable := ui.InfoTable("usage")
	cmdTable.Header(MGroup.HeaderAvailCommandCommand, MGroup.HeaderAvailCommandDesc, MGroup.HeaderAvailCommandNote)
	cmds, ca, specs := z.commandAnnotations(ui)
	for _, cmd := range cmds {
		ann := ca[cmd]
		spec := specs[cmd]
		if spec == nil {
			cmdTable.Row(app_msg.Raw(cmd), z.CommandTitle(cmd), app_msg.Raw(ann))
		} else {
			cmdTable.Row(app_msg.Raw(cmd), spec.Title(), app_msg.Raw(ann))
		}
	}
	cmdTable.Flush()
}

func (z *groupImpl) CommandTitle(cmd string) app_msg.Message {
	keyPath := make([]string, 0)
	keyPath = append(keyPath, "recipe")
	keyPath = append(keyPath, z.path...)
	keyPath = append(keyPath, cmd)
	keyPath = append(keyPath, "title")
	key := strings.Join(keyPath, ".")

	return app_msg.CreateMessage(key)
}

func (z *groupImpl) IsSecret() bool {
	for _, r := range z.recipes {
		if !r.IsSecret() {
			return false
		}
	}
	return true
}

func (z *groupImpl) commandAnnotations(ui app_ui.UI) (cmds []string, annotation map[string]string, spec map[string]rc_recipe.Spec) {
	cmds = make([]string, 0)
	annotation = make(map[string]string)
	spec = make(map[string]rc_recipe.Spec)
	for _, g := range z.subGroups {
		if !g.IsSecret() {
			cmds = append(cmds, g.Name())
		}
		annotation[g.Name()] = ""
	}
	for n, r := range z.recipes {
		if !r.IsSecret() {
			cmds = append(cmds, n)
		}
		annotation[n] = ui.TextOrEmpty(r.Remarks())
		spec[n] = r
	}
	sort.Strings(cmds)
	return
}

func (z *groupImpl) Select(args []string) (g rc_group.Group, r rc_recipe.Spec, remainder []string, err error) {
	if len(args) < 1 {
		return z, nil, args, nil
	}
	arg := args[0]
	for k, sg := range z.subGroups {
		if arg == k {
			return sg.Select(args[1:])
		}
	}
	for k, sr := range z.Recipes() {
		if arg == k {
			return z, sr, args[1:], nil
		}
	}
	return z, nil, args, errors.New("not found")
}
