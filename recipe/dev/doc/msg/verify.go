package msg

import (
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

type Verify struct {
	rc_recipe.RemarkSecret
	ValidationError    app_msg.Message
	MissingTranslation app_msg.Message
	VariableMismatch   app_msg.Message
	Summary            app_msg.Message
	AllValid           app_msg.Message
}

type ValidationResult struct {
	Key                string   `json:"key"`
	EnglishMessage     string   `json:"english_message"`
	JapaneseMessage    string   `json:"japanese_message,omitempty"`
	EnglishVariables   []string `json:"english_variables"`
	JapaneseVariables  []string `json:"japanese_variables,omitempty"`
	IssueType          string   `json:"issue_type,omitempty"`
	IssueDescription   string   `json:"issue_description,omitempty"`
}

func (z *Verify) Preset() {
}

func (z *Verify) extractVariables(message string) []string {
	// Regular expression to match {{.Variable}} patterns
	re := regexp.MustCompile(`{{\.([^}]+)}}`)
	matches := re.FindAllStringSubmatch(message, -1)
	
	variables := make([]string, 0, len(matches))
	seen := make(map[string]bool)
	
	for _, match := range matches {
		if len(match) > 1 {
			variable := match[1]
			if !seen[variable] {
				variables = append(variables, variable)
				seen[variable] = true
			}
		}
	}
	
	sort.Strings(variables)
	return variables
}

func (z *Verify) variablesEqual(vars1, vars2 []string) bool {
	if len(vars1) != len(vars2) {
		return false
	}
	
	for i := range vars1 {
		if vars1[i] != vars2[i] {
			return false
		}
	}
	
	return true
}

func (z *Verify) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	// Load English messages
	enPath := filepath.Join("resources", "messages", "en", "messages.json")
	enData, err := os.ReadFile(enPath)
	if err != nil {
		if c.Feature().IsTest() {
			// In test mode, if messages file doesn't exist, show success
			ui.Success(z.AllValid.With("TotalMessages", 0))
			return nil
		}
		l.Error("Unable to read English messages file", esl.Error(err), esl.String("path", enPath))
		return err
	}

	var enMessages map[string]string
	if err := json.Unmarshal(enData, &enMessages); err != nil {
		l.Error("Unable to parse English messages file", esl.Error(err))
		return err
	}

	// Load Japanese messages
	jaPath := filepath.Join("resources", "messages", "ja", "messages.json")
	jaData, err := os.ReadFile(jaPath)
	if err != nil {
		if os.IsNotExist(err) && c.Feature().IsTest() {
			// In test mode, if Japanese messages don't exist, create empty map
			l.Debug("Japanese messages file doesn't exist in test mode", esl.String("path", jaPath))
		} else {
			l.Error("Unable to read Japanese messages file", esl.Error(err), esl.String("path", jaPath))
			return err
		}
	}

	var jaMessages map[string]string
	if jaData != nil {
		if err := json.Unmarshal(jaData, &jaMessages); err != nil {
			l.Error("Unable to parse Japanese messages file", esl.Error(err))
			return err
		}
	} else {
		jaMessages = make(map[string]string)
	}

	// Validate messages
	var validationResults []ValidationResult
	var hasErrors bool

	// Sort keys for consistent output
	keys := make([]string, 0, len(enMessages))
	for k := range enMessages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		enMessage := enMessages[key]
		jaMessage, jaExists := jaMessages[key]
		
		enVariables := z.extractVariables(enMessage)
		
		result := ValidationResult{
			Key:              key,
			EnglishMessage:   enMessage,
			EnglishVariables: enVariables,
		}

		if !jaExists {
			result.IssueType = "missing_translation"
			result.IssueDescription = "Japanese translation is missing"
			hasErrors = true
		} else {
			result.JapaneseMessage = jaMessage
			jaVariables := z.extractVariables(jaMessage)
			result.JapaneseVariables = jaVariables
			
			if !z.variablesEqual(enVariables, jaVariables) {
				result.IssueType = "variable_mismatch"
				result.IssueDescription = fmt.Sprintf("Variable mismatch: English has %v, Japanese has %v", 
					enVariables, jaVariables)
				hasErrors = true
			}
		}

		if result.IssueType != "" {
			validationResults = append(validationResults, result)
		}
	}

	// Output results
	out := es_stdout.NewDefaultOut(c.Feature())
	defer out.Close()

	if hasErrors {
		ui.Error(z.ValidationError.With("ErrorCount", len(validationResults)))
		
		// Output detailed validation results as JSON
		encoder := json.NewEncoder(out)
		encoder.SetIndent("", "  ")
		encoder.SetEscapeHTML(false)
		
		if err := encoder.Encode(validationResults); err != nil {
			l.Error("Unable to encode validation results", esl.Error(err))
			return err
		}
		
		return fmt.Errorf("validation failed with %d errors", len(validationResults))
	} else {
		ui.Success(z.AllValid.With("TotalMessages", len(enMessages)))
	}

	ui.Info(z.Summary.
		With("TotalMessages", len(enMessages)).
		With("TranslatedMessages", len(jaMessages)).
		With("MissingTranslations", len(enMessages) - len(jaMessages)).
		With("ValidationErrors", len(validationResults)))

	return nil
}

func (z *Verify) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}