package msg

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"os"
	"path/filepath"
	"sort"
)

type Translate struct {
	rc_recipe.RemarkSecret
	Key                mo_string.OptionalString
	ProcessingMessage  app_msg.Message
	AddedTranslation   app_msg.Message
	CompletedMessage   app_msg.Message
	NoMissingFound     app_msg.Message
	ErrorLoadingFiles  app_msg.Message
	SingleKeyNotFound  app_msg.Message
	CurrentValue       app_msg.Message
	PromptTranslation  app_msg.Message
}

func (z *Translate) Preset() {
}

func (z *Translate) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	// Load English messages
	enPath := filepath.Join("resources", "messages", "en", "messages.json")
	enData, err := os.ReadFile(enPath)
	if err != nil {
		ui.Error(z.ErrorLoadingFiles.With("Path", enPath).With("Error", err.Error()))
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
	if err != nil && !os.IsNotExist(err) {
		ui.Error(z.ErrorLoadingFiles.With("Path", jaPath).With("Error", err.Error()))
		return err
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

	// If specific key is provided, translate only that key
	if z.Key.IsExists() {
		key := z.Key.Value()
		enMessage, exists := enMessages[key]
		if !exists {
			ui.Error(z.SingleKeyNotFound.With("Key", key))
			return nil
		}

		// Show current values
		ui.Info(z.CurrentValue.With("Key", key).With("English", enMessage))
		if jaValue, jaExists := jaMessages[key]; jaExists {
			ui.Info(z.CurrentValue.With("Key", key).With("Japanese", jaValue))
		}

		// Show translation prompt
		ui.Info(z.PromptTranslation.With("Key", key))
		
		// For now, we'll return here since we need interactive input
		// In the future, this could be enhanced with interactive prompts
		return nil
	}

	// List missing translations
	missingKeys := make([]string, 0)
	for key := range enMessages {
		if _, exists := jaMessages[key]; !exists {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) == 0 {
		ui.Info(z.NoMissingFound)
		return nil
	}

	sort.Strings(missingKeys)

	// Show first 10 missing keys
	ui.Progress(z.ProcessingMessage.With("Count", len(missingKeys)))
	
	limit := 10
	if len(missingKeys) < limit {
		limit = len(missingKeys)
	}

	for i := 0; i < limit; i++ {
		key := missingKeys[i]
		enMessage := enMessages[key]
		ui.Info(z.CurrentValue.
			With("Key", key).
			With("English", enMessage))
	}

	if len(missingKeys) > limit {
		ui.Info(z.ProcessingMessage.With("Count", len(missingKeys)-limit))
	}

	return nil
}


func (z *Translate) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}