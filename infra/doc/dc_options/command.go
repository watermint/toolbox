package dc_options

import (
	"strings"
	
	"github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgDoc struct {
	HeaderOption      app_msg.Message
	HeaderDescription app_msg.Message
	HeaderDefault     app_msg.Message
	HeaderOptions     app_msg.Message
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
		MDoc.HeaderOptions,
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

		// Get available options for SelectString types
		availableOptions := ""
		if val := spec.Value(k); val != nil {
			if _, typeAttr := val.Spec(); typeAttr != nil {
				if attrMap, ok := typeAttr.(map[string]interface{}); ok {
					if options, hasOptions := attrMap["options"]; hasOptions {
						if optionsList, ok := options.([]string); ok && len(optionsList) > 0 {
							availableOptions = strings.Join(optionsList, ", ")
						}
					}
				}
			}
		}

		mt.Row(
			MDoc.BodyOption.With("Option", es_case.ToLowerKebabCase(k)),
			spec.ValueDesc(k),
			app_msg.Raw(vd),
			app_msg.Raw(availableOptions),
		)
	}

	mt.Flush()
}
