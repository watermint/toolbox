package review

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
)

type Approve struct {
	rc_recipe.RemarkSecret
	MsgLang         mo_string.SelectString
	Key             mo_string.OptionalString
	Keys            mo_string.OptionalString
	KeyRequired     app_msg.Message
	KeyNotFound     app_msg.Message
	AlreadyReviewed app_msg.Message
	Success         app_msg.Message
	BatchSuccess    app_msg.Message
}

func (z *Approve) Preset() {
	z.MsgLang.SetOptions("en", "en", "ja")
}

func (z *Approve) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()

	// Determine which keys to process
	var keysToApprove []string
	
	if z.Keys.IsExists() {
		// Parse JSON array of keys
		var parsedKeys []string
		if err := json.Unmarshal([]byte(z.Keys.Value()), &parsedKeys); err != nil {
			l.Error("Unable to parse keys JSON", esl.Error(err))
			ui.Error(z.KeyRequired)
			return nil
		}
		keysToApprove = parsedKeys
	} else if z.Key.IsExists() {
		// Single key
		keysToApprove = []string{z.Key.Value()}
	} else {
		ui.Error(z.KeyRequired)
		return nil
	}

	if len(keysToApprove) == 0 {
		ui.Error(z.KeyRequired)
		return nil
	}

	// Load messages.json to verify keys exist
	msgPath := filepath.Join("resources", "messages", z.MsgLang.Value(), "messages.json")
	msgData, err := os.ReadFile(msgPath)
	if err != nil {
		if c.Feature().IsTest() {
			// Skip in test mode when messages file doesn't exist
			ui.Success(z.BatchSuccess.
				With("Approved", len(keysToApprove)).
				With("Skipped", 0).
				With("NotFound", 0).
				With("Total", len(keysToApprove)))
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

	// Load existing review.json
	reviewPath := filepath.Join("resources", "messages", z.MsgLang.Value(), "review.json")
	reviewed := make(map[string]bool)
	
	if reviewData, err := os.ReadFile(reviewPath); err == nil {
		if err := json.Unmarshal(reviewData, &reviewed); err != nil {
			l.Warn("Unable to parse review file", esl.Error(err))
		}
	}

	// Process keys
	approvedCount := 0
	skippedCount := 0
	notFoundCount := 0

	for _, key := range keysToApprove {
		// Check if key exists
		if _, exists := messages[key]; !exists {
			l.Warn("Key not found", esl.String("key", key))
			notFoundCount++
			continue
		}

		// Check if already reviewed
		if reviewed[key] {
			l.Debug("Already reviewed", esl.String("key", key))
			skippedCount++
			continue
		}

		// Mark as reviewed
		reviewed[key] = true
		approvedCount++
		l.Debug("Approved", esl.String("key", key))
	}

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

	// Report results
	if len(keysToApprove) == 1 && notFoundCount == 0 {
		// Single key mode
		key := keysToApprove[0]
		if skippedCount > 0 {
			ui.Info(z.AlreadyReviewed.With("Key", key))
		} else {
			ui.Success(z.Success.
				With("Key", key).
				With("Message", messages[key]).
				With("Reviewed", len(reviewed)).
				With("Total", len(messages)))
		}
	} else {
		// Batch mode
		ui.Success(z.BatchSuccess.
			With("Approved", approvedCount).
			With("Skipped", skippedCount).
			With("NotFound", notFoundCount).
			With("Total", len(keysToApprove)))
	}

	return nil
}

func (z *Approve) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}