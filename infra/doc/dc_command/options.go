package dc_command

import (
	"fmt"
	"strings"

	"github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func BodyOptionsTable(ui app_ui.UI, subHeader app_msg.Message, sv rc_recipe.SpecValue, tableOptionsOption, tableOptionsDesc, tableOptionsDefault app_msg.Message) {
	names := sv.ValueNames()
	if len(names) < 1 {
		return
	}
	ui.SubHeader(subHeader)
	ui.Break()

	ui.WithTable("Options", func(t app_ui.Table) {
		t.Header(tableOptionsOption, tableOptionsDesc, tableOptionsDefault)
		for _, k := range names {
			opt := fmt.Sprintf("`-%s`", es_case.ToLowerKebabCase(k))
			vd := sv.ValueDefault(k)
			vkd := sv.ValueCustomDefault(k)
			if ui.Exists(vkd) {
				vd = ui.Text(vkd)
			}

			// Get available options for SelectString types and add to description
			description := sv.ValueDesc(k)
			if val := sv.Value(k); val != nil {
				if _, typeAttr := val.Spec(); typeAttr != nil {
					if attrMap, ok := typeAttr.(map[string]interface{}); ok {
						if options, hasOptions := attrMap["options"]; hasOptions {
							if optionsList, ok := options.([]string); ok && len(optionsList) > 0 {
								optionsText := strings.Join(optionsList, ", ")
								description = app_msg.Raw(fmt.Sprintf("%s (Options: %s)", ui.Text(description), optionsText))
							}
						}
					}
				}
			}

			t.Row(app_msg.Raw(opt), description, app_msg.Raw(vd))
		}
	})
	ui.Break()
}

func GenerateCommonOptionsDetail(ui app_ui.UI, headerCommonOptions, tableOptionsOption, tableOptionsDesc, tableOptionsDefault app_msg.Message) {
	cv := rc_spec.NewCommonValue()
	BodyOptionsTable(ui, headerCommonOptions, cv, tableOptionsOption, tableOptionsDesc, tableOptionsDefault)
}
