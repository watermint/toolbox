package qtr_options

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
)

// SelectStringOption represents a SelectString option that needs a description
type SelectStringOption struct {
	RecipeName string
	FieldName  string
	Option     string
	Key        string
}

// VerifySelectStringOptions checks all recipes for SelectString fields and returns missing options
func VerifySelectStringOptions(c app_control.Control) ([]SelectStringOption, error) {
	l := c.Log()
	cat := app_catalogue.Current()
	
	var missingOptions []SelectStringOption
	
	l.Info("Verifying SelectString option descriptions")
	
	// Load messages file once
	messages := make(map[string]string)
	msgPath := filepath.Join("resources", "messages", "en", "messages.json")
	if msgData, err := os.ReadFile(msgPath); err != nil {
		l.Warn("Could not read messages file", esl.Error(err))
	} else {
		if err := json.Unmarshal(msgData, &messages); err != nil {
			l.Warn("Could not parse messages file", esl.Error(err))
		}
	}
	
	// Check all recipes for SelectString fields
	for _, r := range cat.Recipes() {
		spec := rc_spec.New(r)
		recipeName := spec.Name()
		
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
			
			// Extract options from type attributes
			if attrMap, ok := typeAttr.(map[string]interface{}); ok {
				if options, hasOptions := attrMap["options"]; hasOptions {
					if optionsList, ok := options.([]string); ok && len(optionsList) > 0 {
						fieldNameLow := es_case.ToLowerSnakeCase(valueName)
						
						// Check each option for a description message
						for _, option := range optionsList {
							optionKey := fmt.Sprintf("%s.flag.%s.options.%s", 
								recipeName, 
								fieldNameLow, 
								es_case.ToLowerSnakeCase(option))
							
							// Check if message exists
							if _, exists := messages[optionKey]; !exists {
								missingOptions = append(missingOptions, SelectStringOption{
									RecipeName: recipeName,
									FieldName:  valueName,
									Option:     option,
									Key:        optionKey,
								})
							}
						}
					}
				}
			}
		}
	}
	
	// Sort by key for consistent output
	sort.Slice(missingOptions, func(i, j int) bool {
		return missingOptions[i].Key < missingOptions[j].Key
	})
	
	return missingOptions, nil
}

// TouchSelectStringOptions marks all SelectString option messages as used
func TouchSelectStringOptions(c app_control.Control, touchFunc func(string)) {
	cat := app_catalogue.Current()
	
	// Check all recipes for SelectString fields
	for _, r := range cat.Recipes() {
		spec := rc_spec.New(r)
		recipeName := spec.Name()
		
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
			
			// Extract options from type attributes
			if attrMap, ok := typeAttr.(map[string]interface{}); ok {
				if options, hasOptions := attrMap["options"]; hasOptions {
					if optionsList, ok := options.([]string); ok && len(optionsList) > 0 {
						fieldNameLow := es_case.ToLowerSnakeCase(valueName)
						
						// Touch each option message to mark it as used
						for _, option := range optionsList {
							optionKey := fmt.Sprintf("%s.flag.%s.options.%s", 
								recipeName, 
								fieldNameLow, 
								es_case.ToLowerSnakeCase(option))
							
							touchFunc(optionKey)
						}
					}
				}
			}
		}
	}
}