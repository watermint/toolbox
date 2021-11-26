package dc_options

import (
	"github.com/watermint/essentials/estring/ecase"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgDoc struct {
	HeaderOption      app_msg.Message
	HeaderDescription app_msg.Message
	HeaderDefault     app_msg.Message
	BodyOption        app_msg.Message
}

var (
	MDoc = app_msg.Apply(&MsgDoc{}).(*MsgDoc)
)

func PrintOptionsTable(ui app_ui.UI, spec rc_recipe.SpecValue) {
	mt := ui.InfoTable("")

	mt.Header(
		MDoc.HeaderOption,
		MDoc.HeaderDescription,
		MDoc.HeaderDefault,
	)

	if len(spec.ValueNames()) < 1 {
		return
	}

	for _, k := range spec.ValueNames() {
		vd := spec.ValueDefault(k)
		vkd := spec.ValueCustomDefault(k)
		if ui.Exists(vkd) {
			vd = ui.Text(vkd)
		}

		mt.Row(
			MDoc.BodyOption.With("Option", ecase.ToLowerKebabCase(k)),
			spec.ValueDesc(k),
			app_msg.Raw(vd),
		)
	}

	mt.Flush()
}
