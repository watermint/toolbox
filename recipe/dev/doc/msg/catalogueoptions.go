package msg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type CatalogueOptions struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkIrreversible
	DryRun bool
}

func (z *CatalogueOptions) Preset() {
}

func (z *CatalogueOptions) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	cat := app_catalogue.Current()

	l.Info("Generating option descriptions for all recipes in catalogue")

	totalGenerated := 0
	recipesProcessed := 0

	// Load existing messages
	msgPath := filepath.Join("resources", "messages", "en", "messages.json")
	messages := make(map[string]string)
	if msgData, err := os.ReadFile(msgPath); err == nil {
		if err := json.Unmarshal(msgData, &messages); err != nil {
			l.Warn("Could not parse messages file", esl.Error(err))
		}
	}

	// Process all recipes
	for _, r := range cat.Recipes() {
		spec := rc_spec.New(r)
		recipeName := spec.Name()
		hasSelectString := false

		// Get all value names (field names)
		for _, valueName := range spec.ValueNames() {
			value := spec.Value(valueName)
			if value == nil {
				continue
			}

			// Check if this is a SelectString type
			typeName, typeAttr := value.Spec()
			if typeName != "mo_string.SelectString" &&
				typeName != "github.com/watermint/toolbox/essentials/model/mo_string.selectStringInternal" &&
				typeName != "essentials.model.mo_string.select_string_internal" {
				continue
			}

			hasSelectString = true

			// Extract options from type attributes
			if attrMap, ok := typeAttr.(map[string]interface{}); ok {
				if options, hasOptions := attrMap["options"]; hasOptions {
					if optionsList, ok := options.([]string); ok && len(optionsList) > 0 {
						fieldNameLow := es_case.ToLowerSnakeCase(valueName)

						// Generate message for each option
						for _, option := range optionsList {
							optionKey := fmt.Sprintf("%s.flag.%s.options.%s",
								recipeName,
								fieldNameLow,
								es_case.ToLowerSnakeCase(option))

							// Check if message already exists
							if _, exists := messages[optionKey]; !exists {
								description := z.generateOptionDescription(valueName, option)

								if z.DryRun {
									ui.Info(app_msg.Raw(fmt.Sprintf("Would create message: %s = %s", optionKey, description)))
								} else {
									messages[optionKey] = description
									ui.Info(app_msg.Raw(fmt.Sprintf("Created option message: %s = %s", optionKey, description)))
								}
								totalGenerated++
							}
						}
					}
				}
			}
		}

		if hasSelectString {
			recipesProcessed++
		}
	}

	// Write back to file if not dry run
	if !z.DryRun && totalGenerated > 0 {
		data, err := json.MarshalIndent(messages, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal messages: %w", err)
		}

		if err := os.WriteFile(msgPath, data, 0644); err != nil {
			return fmt.Errorf("failed to write messages file: %w", err)
		}
	}

	l.Info("Option generation completed",
		esl.Int("total_generated", totalGenerated),
		esl.Int("recipes_processed", recipesProcessed))
	ui.Success(app_msg.Raw(fmt.Sprintf("Generated %d option descriptions for %d recipes", totalGenerated, recipesProcessed)))

	return nil
}

func (z *CatalogueOptions) generateOptionDescription(fieldName, option string) string {
	// Convert field name to lowercase for matching
	fieldNameLower := strings.ToLower(fieldName)

	// Handle common patterns
	switch fieldNameLower {
	case "basepath", "base_path":
		switch option {
		case "root":
			return "Full access to all folders with permissions"
		case "home":
			return "Access limited to personal home folder"
		default:
			return fmt.Sprintf("Base path option: %s", option)
		}
	case "visibility", "new_visibility":
		switch option {
		case "public":
			return "Anyone with the link can access"
		case "team_only":
			return "Only team members can access"
		case "password":
			return "Password protected access"
		case "team_and_password":
			return "Team members only with password"
		case "shared_folder_only":
			return "Only shared folder members can access"
		default:
			return fmt.Sprintf("Visibility option: %s", option)
		}
	case "accesslevel", "access_level":
		switch option {
		case "editor":
			return "Can edit, comment, and share"
		case "viewer":
			return "Can view and comment"
		case "viewer_no_comment":
			return "Can only view"
		case "owner":
			return "Full owner permissions"
		default:
			return fmt.Sprintf("Access level: %s", option)
		}
	case "managementtype", "management_type":
		switch option {
		case "company_managed":
			return "Managed by company administrators"
		case "user_managed":
			return "Managed by individual users"
		default:
			return fmt.Sprintf("Management type: %s", option)
		}
	case "format":
		switch option {
		case "html":
			return "HTML format"
		case "markdown":
			return "Markdown format"
		case "plain_text":
			return "Plain text format"
		case "pdf":
			return "PDF document format"
		default:
			return fmt.Sprintf("Format: %s", option)
		}
	case "method":
		switch option {
		case "block":
			return "Block upload method (parallel chunks)"
		case "sequential":
			return "Sequential upload method"
		default:
			return fmt.Sprintf("Method: %s", option)
		}
	case "state":
		switch option {
		case "open":
			return "Open issues only"
		case "closed":
			return "Closed issues only"
		case "all":
			return "All issues"
		default:
			return fmt.Sprintf("State: %s", option)
		}
	default:
		// Generic description
		return fmt.Sprintf("%s: %s", strings.ReplaceAll(fieldNameLower, "_", " "), option)
	}
}

func (z *CatalogueOptions) Test(c app_control.Control) error {
	return z.Exec(c)
}
