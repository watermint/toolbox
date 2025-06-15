package review

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
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
	MsgLang mo_string.SelectString
	Limit   mo_int.RangeInt
	Prefix  mo_string.OptionalString
	Summary app_msg.Message
}

func (z *List) Preset() {
	z.MsgLang.SetOptions("en", "en", "ja")
	z.Limit.SetRange(1, 1000, 50)
}

func (z *List) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	// Load messages.json
	msgPath := filepath.Join("resources", "messages", z.MsgLang.Value(), "messages.json")
	msgData, err := os.ReadFile(msgPath)
	if err != nil {
		if c.Feature().IsTest() {
			// Skip in test mode when messages file doesn't exist
			ui.Info(z.Summary.With("Total", 0).With("Unreviewed", 0).With("Showing", 0))
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

	// Load review.json if exists
	reviewPath := filepath.Join("resources", "messages", z.MsgLang.Value(), "review.json")
	reviewed := make(map[string]bool)
	
	if reviewData, err := os.ReadFile(reviewPath); err == nil {
		if err := json.Unmarshal(reviewData, &reviewed); err != nil {
			l.Warn("Unable to parse review file", esl.Error(err))
		}
	}

	// Find unreviewed messages
	var unreviewed []string
	prefix := ""
	if z.Prefix.IsExists() {
		prefix = z.Prefix.Value()
	}
	
	for key := range messages {
		if !reviewed[key] {
			if prefix == "" || strings.HasPrefix(key, prefix) {
				unreviewed = append(unreviewed, key)
			}
		}
	}

	// Sort keys for consistent output
	sort.Strings(unreviewed)

	// Limit the number of messages
	if len(unreviewed) > z.Limit.Value() {
		unreviewed = unreviewed[:z.Limit.Value()]
	}

	// Output results
	out := es_stdout.NewDefaultOut(c.Feature())
	defer out.Close()

	ui.Info(z.Summary.
		With("Total", len(messages)).
		With("Unreviewed", len(unreviewed)).
		With("Showing", len(unreviewed)))

	// Output as JSON array for easy processing
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	result := make(map[string]interface{})
	for _, key := range unreviewed {
		result[key] = map[string]string{
			"key":     key,
			"message": messages[key],
		}
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