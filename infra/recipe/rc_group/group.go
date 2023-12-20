package rc_group

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgHeader struct {
	Header  app_msg.Message
	License app_msg.Message
}

var (
	MHeader = app_msg.Apply(&MsgHeader{}).(*MsgHeader)
)

func AppHeader(ui app_ui.UI, version string) {
	ui.Header(MHeader.Header.With("AppVersion", version).With("AppName", app_definitions.Name))
	ui.Info(app_msg.Raw(app_definitions.Copyright))
	ui.Info(MHeader.License)
	ui.Break()
}

func UsageHeader(ui app_ui.UI, desc app_msg.Message, version string) {
	AppHeader(ui, version)
	ui.Break()
	ui.Info(desc)
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
	PrintUsage(ui app_ui.UI, exec, version string)
	GroupDesc() app_msg.Message
	CommandTitle(cmd string) app_msg.Message
	IsSecret() bool
	Select(args []string) (g Group, r rc_recipe.Spec, remainder []string, err error)
}
