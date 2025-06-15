package review

import (
	"encoding/json"
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

type Batch struct {
	rc_recipe.RemarkSecret
	MsgLang          mo_string.SelectString
	Limit            mo_int.RangeInt
	Interactive      app_msg.Message
	PromptReview     app_msg.Message
	PromptApprove    app_msg.Message
	PromptSkip       app_msg.Message
	PromptStop       app_msg.Message
	InvalidChoice    app_msg.Message
	Approved         app_msg.Message
	Skipped          app_msg.Message
	SessionComplete  app_msg.Message
	NoUnreviewed     app_msg.Message
}

func (z *Batch) Preset() {
	z.MsgLang.SetOptions("en", "en", "ja")
	z.Limit.SetRange(1, 100, 10)
}

func (z *Batch) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	// Load messages.json
	msgPath := filepath.Join("resources", "messages", z.MsgLang.Value(), "messages.json")
	msgData, err := os.ReadFile(msgPath)
	if err != nil {
		if c.Feature().IsTest() {
			// Skip in test mode
			ui.Info(z.NoUnreviewed)
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
	for key := range messages {
		if !reviewed[key] {
			unreviewed = append(unreviewed, key)
		}
	}

	if len(unreviewed) == 0 {
		ui.Info(z.NoUnreviewed)
		return nil
	}

	// Sort keys for consistent output
	sort.Strings(unreviewed)

	// Limit the number of messages
	if len(unreviewed) > z.Limit.Value() {
		unreviewed = unreviewed[:z.Limit.Value()]
	}

	ui.Info(z.Interactive.
		With("Total", len(messages)).
		With("Unreviewed", len(unreviewed)).
		With("Batch", len(unreviewed)))

	// Interactive review loop
	approvedCount := 0
	skippedCount := 0

	for i, key := range unreviewed {
		message := messages[key]
		
		ui.Info(z.PromptReview.
			With("Index", i+1).
			With("Total", len(unreviewed)).
			With("Key", key).
			With("Message", message))

		// Ask for user action
		for {
			input, cancel := ui.AskText(z.PromptApprove)
			if cancel {
				ui.Info(z.PromptStop)
				goto done
			}
			
			input = strings.TrimSpace(strings.ToLower(input))

			switch input {
			case "y", "yes":
				// Approve the message
				reviewed[key] = true
				approvedCount++
				ui.Success(z.Approved.With("Key", key))
				goto next

			case "n", "no":
				// Skip this message
				skippedCount++
				ui.Info(z.Skipped.With("Key", key))
				goto next

			case "q", "quit":
				// Stop reviewing
				ui.Info(z.PromptStop)
				goto done

			default:
				ui.Error(z.InvalidChoice.With("Input", input))
				continue
			}
		}
		next:
	}

	done:
	// Save updated review.json
	if approvedCount > 0 {
		reviewData, err := json.MarshalIndent(reviewed, "", "  ")
		if err != nil {
			l.Error("Unable to marshal review data", esl.Error(err))
			return err
		}

		// Ensure directory exists
		reviewDir := filepath.Dir(reviewPath)
		if err := os.MkdirAll(reviewDir, 0755); err != nil {
			l.Error("Unable to create review directory", esl.Error(err), esl.String("path", reviewDir))
			return err
		}

		if err := os.WriteFile(reviewPath, reviewData, 0644); err != nil {
			l.Error("Unable to write review file", esl.Error(err), esl.String("path", reviewPath))
			return err
		}
	}

	ui.Success(z.SessionComplete.
		With("Approved", approvedCount).
		With("Skipped", skippedCount).
		With("Total", len(unreviewed)))

	return nil
}

func (z *Batch) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}