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

	// Build definitions for the option list
	definitions := make([]app_ui.Definition, 0, len(names))
	for _, k := range names {
		opt := fmt.Sprintf("-%s", es_case.ToLowerKebabCase(k))
		vd := sv.ValueDefault(k)
		vkd := sv.ValueCustomDefault(k)
		if ui.Exists(vkd) {
			vd = ui.Text(vkd)
		}

		// Build description with default value
		description := sv.ValueDesc(k)
		descParts := []string{ui.Text(description)}
		
		// Add available options for SelectString types with descriptions
		if val := sv.Value(k); val != nil {
			if _, typeAttr := val.Spec(); typeAttr != nil {
				if attrMap, ok := typeAttr.(map[string]interface{}); ok {
					if options, hasOptions := attrMap["options"]; hasOptions {
						if optionsList, ok := options.([]string); ok && len(optionsList) > 0 {
							// Build options with descriptions
							var optionDescs []string
							hasDescriptions := false
							for _, option := range optionsList {
								// Try to get option description from messages
								// The key pattern should match the field description pattern + ".options." + option
								baseDesc := sv.ValueDesc(k)
								if baseDesc != nil && baseDesc.Key() != "" {
									// First try with the full key
									optionKey := baseDesc.Key() + ".options." + option
									optionMsg := app_msg.CreateMessage(optionKey)
									if ui.Exists(optionMsg) {
										optionDescs = append(optionDescs, fmt.Sprintf("%s (%s)", option, ui.Text(optionMsg)))
										hasDescriptions = true
									} else {
										// Try with a shorter key pattern (last component of recipe path)
										keyParts := strings.Split(baseDesc.Key(), ".")
										if len(keyParts) > 2 {
											// Extract the last component (e.g., "list" from "citron.dropbox.file.list.flag.base_path")
											shortKey := keyParts[len(keyParts)-3] + ".flag." + es_case.ToLowerSnakeCase(k) + ".options." + option
											shortMsg := app_msg.CreateMessage(shortKey)
											if ui.Exists(shortMsg) {
												optionDescs = append(optionDescs, fmt.Sprintf("%s (%s)", option, ui.Text(shortMsg)))
												hasDescriptions = true
											} else {
												optionDescs = append(optionDescs, option)
											}
										} else {
											optionDescs = append(optionDescs, option)
										}
									}
								} else {
									optionDescs = append(optionDescs, option)
								}
							}
							
							// Format the options differently based on whether we have descriptions
							if hasDescriptions && len(optionsList) > 2 {
								// Use bullet list format for better readability when we have descriptions
								descParts = append(descParts, "Options:")
								for _, desc := range optionDescs {
									descParts = append(descParts, "  • " + desc)
								}
							} else {
								// Use inline format for simple options
								descParts = append(descParts, fmt.Sprintf("Options: %s", strings.Join(optionDescs, ", ")))
							}
						}
					}
				}
			}
		}
		
		// Add default value
		if vd != "" {
			descParts = append(descParts, fmt.Sprintf("Default: %s", vd))
		}
		
		definitions = append(definitions, app_ui.Definition{
			Term:        app_msg.Raw(opt),
			Description: app_msg.Raw(strings.Join(descParts, ". ")),
		})
	}
	
	ui.DefinitionList(definitions)
	ui.Break()
}

func GenerateCommonOptionsDetail(ui app_ui.UI, headerCommonOptions, tableOptionsOption, tableOptionsDesc, tableOptionsDefault app_msg.Message) {
	cv := rc_spec.NewCommonValue()
	BodyOptionsTable(ui, headerCommonOptions, cv, tableOptionsOption, tableOptionsDesc, tableOptionsDefault)
}
