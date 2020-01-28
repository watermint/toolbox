package rc_group

import (
	"flag"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgHeader struct {
	Header    app_msg.Message
	Copyright app_msg.Message
	License   app_msg.Message
}

var (
	MHeader = app_msg.Apply(&MsgHeader{}).(*MsgHeader)
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
	Recipes() map[string]rc_recipe.Spec
	SubGroups() map[string]Group
	Add(r rc_recipe.Spec)
	AddToPath(fullPath []string, relPath []string, name string, r rc_recipe.Spec)
	PrintRecipeUsage(ui app_ui.UI, spec rc_recipe.Spec, f *flag.FlagSet)
	PrintGroupUsage(ui app_ui.UI, exec, version string)
	GroupDesc() app_msg.Message
	CommandTitle(cmd string) app_msg.Message
	IsSecret() bool
	Select(args []string) (name string, g Group, r rc_recipe.Spec, remainder []string, err error)
}
