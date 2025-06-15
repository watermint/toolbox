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

type Delete struct {
	rc_recipe.RemarkSecret
	MsgLang     mo_string.SelectString
	Key         mo_string.OptionalString
	KeyRequired app_msg.Message
	KeyNotFound app_msg.Message
	Success     app_msg.Message
}

func (z *Delete) Preset() {
	z.MsgLang.SetOptions("en", "en", "ja")
}

func (z *Delete) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	if !z.Key.IsExists() {
		ui.Error(z.KeyRequired)
		return nil
	}

	key := z.Key.Value()

	// Load existing messages.json
	msgPath := filepath.Join("resources", "messages", z.MsgLang.Value(), "messages.json")
	messages := make(map[string]string)

	msgData, err := os.ReadFile(msgPath)
	if err != nil {
		if os.IsNotExist(err) || c.Feature().IsTest() {
			ui.Error(z.KeyNotFound.With("Key", key))
			return nil
		}
		l.Error("Unable to read messages file", esl.Error(err), esl.String("path", msgPath))
		return err
	}

	if err := json.Unmarshal(msgData, &messages); err != nil {
		l.Error("Unable to parse messages file", esl.Error(err))
		return err
	}

	// Check if key exists
	deletedMessage, exists := messages[key]
	if !exists {
		ui.Error(z.KeyNotFound.With("Key", key))
		return nil
	}

	// Delete the message
	delete(messages, key)

	// Sort keys for consistent output
	keys := make([]string, 0, len(messages))
	for k := range messages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create sorted map for output
	sortedMessages := make(map[string]string)
	for _, k := range keys {
		sortedMessages[k] = messages[k]
	}

	// Save updated messages.json
	msgData, err = json.MarshalIndent(sortedMessages, "", "  ")
	if err != nil {
		l.Error("Unable to marshal messages data", esl.Error(err))
		return err
	}

	if err := os.WriteFile(msgPath, msgData, 0644); err != nil {
		l.Error("Unable to write messages file", esl.Error(err), esl.String("path", msgPath))
		return err
	}

	ui.Success(z.Success.
		With("Key", key).
		With("Message", deletedMessage).
		With("Total", len(messages)))

	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}