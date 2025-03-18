package dc_command

import (
	"strings"

	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewUsage(media dc_index.MediaType, spec rc_recipe.Spec) dc_section.Section {
	return &Usage{
		media: media,
		spec:  spec,
	}
}

type Usage struct {
	spec                rc_recipe.Spec
	media               dc_index.MediaType
	Header              app_msg.Message
	Remarks             app_msg.Message
	HeaderRun           app_msg.Message
	UsageWin            app_msg.Message
	UsageOther          app_msg.Message
	RunCatalinaRemarks1 app_msg.Message
	RunCatalinaRemarks2 app_msg.Message
	RunCatalinaRemarks3 app_msg.Message
	HeaderOptions       app_msg.Message
	HeaderCommonOptions app_msg.Message
	TableOptionsOption  app_msg.Message
	TableOptionsDesc    app_msg.Message
	TableOptionsDefault app_msg.Message
}

func (z Usage) Title() app_msg.Message {
	return z.Header
}

func (z Usage) isFullVersion() bool {
	return z.media != dc_index.MediaKnowledge
}

func (z Usage) Body(ui app_ui.UI) {
	ui.Info(z.Remarks)

	z.bodyRun(ui)
	z.bodyOptions(ui)
	if z.isFullVersion() {
		z.bodyCommonOptions(ui)
	}
}

func (z Usage) bodyRun(ui app_ui.UI) {
	if z.isFullVersion() {
		cmdWin := strings.Join([]string{
			".\\" + app_definitions.ExecutableName + ".exe",
			z.spec.CliPath(),
			ui.TextOrEmpty(z.spec.CliArgs()),
		}, " ")

		cmdOther := strings.Join([]string{
			"$HOME/Desktop/" + app_definitions.ExecutableName,
			z.spec.CliPath(),
			ui.TextOrEmpty(z.spec.CliArgs()),
		}, " ")

		ui.SubHeader(z.HeaderRun)

		ui.Info(z.UsageWin)
		ui.Code("cd $HOME\\Desktop\n" + cmdWin)
		ui.Break()

		ui.Info(z.UsageOther)
		ui.Code(cmdOther)
	} else {
		cmdGeneral := strings.Join([]string{
			app_definitions.ExecutableName,
			z.spec.CliPath(),
			ui.TextOrEmpty(z.spec.CliArgs()),
		}, " ")
		ui.Code(cmdGeneral)
	}

	ui.Info(z.spec.CliNote())
	if z.isFullVersion() {
		ui.Break()
		ui.Info(z.RunCatalinaRemarks1)
		ui.Info(z.RunCatalinaRemarks2)
		ui.Break()
		ui.Info(z.RunCatalinaRemarks3)
		ui.Break()
	}
}

func (z Usage) bodyOptions(ui app_ui.UI) {
	z.bodyOptionsTable(ui, z.HeaderOptions, z.spec)
}

func (z Usage) bodyCommonOptions(ui app_ui.UI) {
	GenerateCommonOptionsDetail(ui, z.HeaderCommonOptions, z.TableOptionsOption, z.TableOptionsDesc, z.TableOptionsDefault)
}

func (z Usage) bodyOptionsTable(ui app_ui.UI, subHeader app_msg.Message, sv rc_recipe.SpecValue) {
	BodyOptionsTable(ui, subHeader, sv, z.TableOptionsOption, z.TableOptionsDesc, z.TableOptionsDefault)
}
