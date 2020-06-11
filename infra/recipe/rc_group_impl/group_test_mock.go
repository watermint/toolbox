package rc_group_impl

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MockSpec struct {
	name string
	path []string
}

func (z MockSpec) ValueNames() []string {
	panic("implement me")
}

func (z MockSpec) ValueDesc(name string) app_msg.Message {
	panic("implement me")
}

func (z MockSpec) ValueDefault(name string) interface{} {
	panic("implement me")
}

func (z MockSpec) Value(name string) rc_recipe.Value {
	panic("implement me")
}

func (z MockSpec) ValueCustomDefault(name string) app_msg.MessageOptional {
	panic("implement me")
}

func (z MockSpec) SetFlags(f *flag.FlagSet, ui app_ui.UI) {
	panic("implement me")
}

func (z MockSpec) Name() string {
	return z.name
}

func (z MockSpec) Title() app_msg.Message {
	panic("implement me")
}

func (z MockSpec) Desc() app_msg.MessageOptional {
	panic("implement me")
}

func (z MockSpec) Remarks() app_msg.MessageOptional {
	panic("implement me")
}

func (z MockSpec) Path() (path []string, name string) {
	return z.path, z.name
}

func (z MockSpec) SpecId() string {
	panic("implement me")
}

func (z MockSpec) CliPath() string {
	panic("implement me")
}

func (z MockSpec) CliArgs() app_msg.MessageOptional {
	panic("implement me")
}

func (z MockSpec) CliNote() app_msg.MessageOptional {
	panic("implement me")
}

func (z MockSpec) Reports() []rp_model.Spec {
	panic("implement me")
}

func (z MockSpec) Feeds() map[string]fd_file.Spec {
	panic("implement me")
}

func (z MockSpec) Messages() []app_msg.Message {
	panic("implement me")
}

func (z MockSpec) Services() []string {
	panic("implement me")
}

func (z MockSpec) ConnUsePersonal() bool {
	panic("implement me")
}

func (z MockSpec) ConnUseBusiness() bool {
	panic("implement me")
}

func (z MockSpec) ConnScopes() []string {
	panic("implement me")
}

func (z MockSpec) ConnScopeMap() map[string]string {
	panic("implement me")
}

func (z MockSpec) SpinUp(ctl app_control.Control, custom func(r rc_recipe.Recipe)) (rcp rc_recipe.Recipe, err error) {
	panic("implement me")
}

func (z MockSpec) Debug() map[string]interface{} {
	panic("implement me")
}

func (z MockSpec) SpinDown(ctl app_control.Control) error {
	panic("implement me")
}

func (z MockSpec) IsSecret() bool {
	return false
}

func (z MockSpec) IsConsole() bool {
	return false
}

func (z MockSpec) IsExperimental() bool {
	return false
}

func (z MockSpec) IsIrreversible() bool {
	return false
}

func (z MockSpec) IsTransient() bool {
	return false
}

func (z MockSpec) PrintUsage(ui app_ui.UI) {
	panic("implement me")
}

func (z MockSpec) New() rc_recipe.Spec {
	panic("implement me")
}

func (z MockSpec) Doc(ui app_ui.UI) *dc_recipe.Recipe {
	panic("implement me")
}
