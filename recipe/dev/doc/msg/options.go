package msg

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Options struct {
	rc_recipe.RemarkSecret
	TargetPath      string
	DryRun          bool
	ScanStarted     app_msg.Message
	ScanCompleted   app_msg.Message
	MessageCreated  app_msg.Message
	GenerationError app_msg.Message
}

type SelectStringField struct {
	RecipePath    string
	StructName    string
	FieldName     string
	FieldNameLow  string
	OptionValues  []string
	MessagePrefix string
}

func (z *Options) Preset() {
	z.TargetPath = "citron"
	z.DryRun = false
}

func (z *Options) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	ui.Info(z.ScanStarted.With("Path", z.TargetPath))

	// Find all SelectString fields in the codebase
	fields, err := z.scanSelectStringFields(c, z.TargetPath)
	if err != nil {
		l.Error("Failed to scan SelectString fields", esl.Error(err))
		return err
	}

	ui.Info(z.ScanCompleted.With("Count", len(fields)))

	// Generate option description messages
	for _, field := range fields {
		if err := z.generateOptionMessages(c, field); err != nil {
			l.Error("Failed to generate messages for field", 
				esl.Error(err), 
				esl.String("recipe", field.RecipePath),
				esl.String("field", field.FieldName))
			ui.Error(z.GenerationError.With("Field", field.FieldName).With("Recipe", field.RecipePath))
			continue
		}
	}

	return nil
}

func (z *Options) scanSelectStringFields(c app_control.Control, targetPath string) ([]SelectStringField, error) {
	l := c.Log()
	var fields []SelectStringField

	err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		// Parse the Go file
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			l.Debug("Failed to parse file", esl.Error(err), esl.String("path", path))
			return nil
		}

		// Find SelectString fields
		ast.Inspect(node, func(n ast.Node) bool {
			if typeSpec, ok := n.(*ast.TypeSpec); ok {
				if structType, ok := typeSpec.Type.(*ast.StructType); ok {
					structName := typeSpec.Name.Name
					packagePath := z.getPackagePath(path)
					
					for _, field := range structType.Fields.List {
						if z.isSelectStringField(field) && len(field.Names) > 0 {
							fieldName := field.Names[0].Name
							
							// Get message prefix from package path
							messagePrefix := z.buildMessagePrefix(packagePath, structName)
							
							selectField := SelectStringField{
								RecipePath:    path,
								StructName:    structName,
								FieldName:     fieldName,
								FieldNameLow:  es_case.ToLowerSnakeCase(fieldName),
								MessagePrefix: messagePrefix,
							}

							// Try to find SetOptions call to get option values
							options := z.findSetOptionsCall(node, fieldName)
							selectField.OptionValues = options

							fields = append(fields, selectField)
						}
					}
				}
			}
			return true
		})

		return nil
	})

	return fields, err
}

func (z *Options) isSelectStringField(field *ast.Field) bool {
	if selectorExpr, ok := field.Type.(*ast.SelectorExpr); ok {
		if ident, ok := selectorExpr.X.(*ast.Ident); ok {
			return ident.Name == "mo_string" && selectorExpr.Sel.Name == "SelectString"
		}
	}
	return false
}

func (z *Options) getPackagePath(filePath string) string {
	// Convert file path to package path
	// e.g., "citron/dropbox/team/sharedlink/update/visibility.go" -> "citron.dropbox.team.sharedlink.update.visibility"
	path := strings.TrimSuffix(filePath, ".go")
	path = strings.ReplaceAll(path, "/", ".")
	return path
}

func (z *Options) buildMessagePrefix(packagePath, structName string) string {
	// The package path already includes the struct name as the filename
	// e.g., "citron.dropbox.team.sharedlink.update.visibility" with struct "Visibility"
	// should result in "citron.dropbox.team.sharedlink.update.visibility"
	return packagePath
}

func (z *Options) findSetOptionsCall(node *ast.File, fieldName string) []string {
	var options []string

	ast.Inspect(node, func(n ast.Node) bool {
		if callExpr, ok := n.(*ast.CallExpr); ok {
			if selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				if selectorExpr.Sel.Name == "SetOptions" {
					// Check if this is for our field
					if ident, ok := selectorExpr.X.(*ast.SelectorExpr); ok {
						if baseIdent, ok := ident.X.(*ast.Ident); ok && baseIdent.Name == "z" {
							if ident.Sel.Name == fieldName {
								// Extract string literals from arguments
								for _, arg := range callExpr.Args {
									if basicLit, ok := arg.(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
										// Remove quotes
										value := strings.Trim(basicLit.Value, `"`)
										options = append(options, value)
									}
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	return options
}

func (z *Options) generateOptionMessages(c app_control.Control, field SelectStringField) error {
	ui := c.UI()
	l := c.Log()

	// Remove duplicates from option values
	uniqueOptions := make(map[string]bool)
	var options []string
	for _, option := range field.OptionValues {
		if !uniqueOptions[option] {
			uniqueOptions[option] = true
			options = append(options, option)
		}
	}

	// Generate option description messages for each unique option value
	for _, option := range options {
		optionKey := fmt.Sprintf("%s.flag.%s.options.%s", 
			field.MessagePrefix, 
			field.FieldNameLow, 
			es_case.ToLowerSnakeCase(option))

		// Generate a descriptive message for the option
		optionDescription := z.generateOptionDescription(field.FieldNameLow, option)

		if z.DryRun {
			ui.Info(app_msg.Raw(fmt.Sprintf("Would create: %s = %s", optionKey, optionDescription)))
		} else {
			// Create the message using the existing add command logic
			if err := z.createMessage(c, optionKey, optionDescription); err != nil {
				l.Error("Failed to create option message", 
					esl.Error(err), 
					esl.String("key", optionKey))
				return err
			}
			ui.Info(z.MessageCreated.With("Key", optionKey).With("Value", optionDescription))
		}
	}

	return nil
}

func (z *Options) generateOptionDescription(fieldName, option string) string {
	// Generate contextual descriptions based on field name and option
	switch fieldName {
	case "base_path":
		switch option {
		case "root":
			return "Access all folders with permissions (includes team folders and member folders)"
		case "home":
			return "Access personal folder only (convenient for personal file operations)"
		default:
			return fmt.Sprintf("Base path option: %s", option)
		}
	case "new_visibility", "visibility":
		switch option {
		case "public":
			return "Anyone with the link can access (public sharing)"
		case "team_only":
			return "Only team members can access (restricted to team)"
		case "password":
			return "Password protected access (requires password)"
		default:
			return fmt.Sprintf("Visibility option: %s", option)
		}
	case "format":
		switch option {
		case "png":
			return "PNG image format (raster, supports transparency)"
		case "jpg", "jpeg":
			return "JPEG image format (raster, smaller file size)"
		case "svg":
			return "SVG vector format (scalable, smaller file size)"
		case "pdf":
			return "PDF document format (high quality, printable)"
		default:
			return fmt.Sprintf("Format option: %s", option)
		}
	case "budget_memory":
		switch option {
		case "low":
			return "Reduce memory usage (may impact performance)"
		case "normal":
			return "Standard memory usage (balanced performance)"
		default:
			return fmt.Sprintf("Memory budget option: %s", option)
		}
	case "budget_storage":
		switch option {
		case "low":
			return "Minimize storage usage (limited logs)"
		case "normal":
			return "Standard storage usage (normal logging)"
		case "unlimited":
			return "No storage limits (full logging)"
		default:
			return fmt.Sprintf("Storage budget option: %s", option)
		}
	case "output":
		switch option {
		case "text":
			return "Plain text output (human readable)"
		case "markdown":
			return "Markdown formatted output (documentation friendly)"
		case "json":
			return "JSON formatted output (machine readable)"
		case "none":
			return "No output (silent mode)"
		default:
			return fmt.Sprintf("Output format option: %s", option)
		}
	case "lang", "language":
		switch option {
		case "en":
			return "English language"
		case "ja":
			return "Japanese language (日本語)"
		case "auto":
			return "Auto-detect language from system settings"
		default:
			return fmt.Sprintf("Language option: %s", option)
		}
	case "retain_job_data":
		switch option {
		case "default":
			return "Standard job data retention policy"
		case "on_error":
			return "Retain job data only when errors occur"
		case "none":
			return "Do not retain job data (minimal storage)"
		default:
			return fmt.Sprintf("Job data retention option: %s", option)
		}
	default:
		// Generic description
		return fmt.Sprintf("%s option: %s", strings.ReplaceAll(fieldName, "_", " "), option)
	}
}

func (z *Options) createMessage(c app_control.Control, key, value string) error {
	// Directly create the message by updating the JSON file
	msgPath := filepath.Join("resources", "messages", "en", "messages.json")
	messages := make(map[string]string)

	// Load existing messages
	if msgData, err := os.ReadFile(msgPath); err == nil {
		if err := json.Unmarshal(msgData, &messages); err != nil {
			return fmt.Errorf("failed to parse messages file: %w", err)
		}
	}

	// Add the new message
	messages[key] = value

	// Write back to file
	data, err := json.MarshalIndent(messages, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal messages: %w", err)
	}

	if err := os.WriteFile(msgPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write messages file: %w", err)
	}

	return nil
}

func (z *Options) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}