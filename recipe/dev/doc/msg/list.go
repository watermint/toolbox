package msg

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type List struct {
	rc_recipe.RemarkSecret
	MsgLang    mo_string.SelectString
	Prefix     mo_string.OptionalString
	NoMessages app_msg.Message
	Summary    app_msg.Message
}

func (z *List) Preset() {
	z.MsgLang.SetOptions("en", "en", "ja")
}

func (z *List) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	// Load messages.json
	msgPath := filepath.Join("resources", "messages", z.MsgLang.Value(), "messages.json")
	msgData, err := os.ReadFile(msgPath)
	if err != nil {
		if os.IsNotExist(err) || c.Feature().IsTest() {
			ui.Info(z.NoMessages)
			return nil
		}
		l.Error("Unable to read messages file", esl.Error(err), esl.String("path", msgPath))
		return err
	}

	var messages map[string]string
	if err := json.Unmarshal(msgData, &messages); err != nil {
		l.Error("Unable to parse messages file", esl.Error(err))
		return err
	}

	// Filter by prefix if provided
	filtered := make(map[string]string)
	prefix := ""
	if z.Prefix.IsExists() {
		prefix = z.Prefix.Value()
		for key, value := range messages {
			if strings.HasPrefix(key, prefix) {
				filtered[key] = value
			}
		}
	} else {
		filtered = messages
	}

	// Sort keys for consistent output
	keys := make([]string, 0, len(filtered))
	for k := range filtered {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Output results
	out := es_stdout.NewDefaultOut(c.Feature())
	defer out.Close()

	ui.Info(z.Summary.
		With("Total", len(messages)).
		With("Filtered", len(filtered)).
		With("Prefix", prefix))

	// Output as JSON for easy processing
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	result := make(map[string]interface{})
	for _, key := range keys {
		result[key] = filtered[key]
	}

	if err := encoder.Encode(result); err != nil {
		l.Error("Unable to encode output", esl.Error(err))
		return err
	}

	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}