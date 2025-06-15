package app_ui

import (
	"bytes"
	"strings"
	"testing"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
)

func TestDefinitionList(t *testing.T) {
	mc := app_msg_container_impl.NewSingleWithMessagesForTest(map[string]string{
		"raw": "{{.Raw}}",
	})
	
	definitions := []Definition{
		{
			Term:        app_msg.Raw("-path"),
			Description: app_msg.Raw("Specify the file path. Default: /home"),
		},
		{
			Term:        app_msg.Raw("-recursive"),
			Description: app_msg.Raw("List files recursively. Options: true, false. Default: false"),
		},
		{
			Term:        app_msg.Raw("-base-path"),
			Description: app_msg.Raw("Choose the file path standard. Options: root, home. Default: root"),
		},
	}

	// Test Markdown output
	t.Run("Markdown", func(t *testing.T) {
		output := MakeMarkdown(mc, func(ui UI) {
			ui.DefinitionList(definitions)
		})
		
		// Check that markdown format is correct
		t.Logf("Markdown output:\n%s", output)
		if !strings.Contains(output, "**-path**") {
			t.Error("Expected markdown bold formatting for term")
		}
		if !strings.Contains(output, ": Specify the file path") {
			t.Error("Expected markdown definition format with colon")
		}
	})

	// Test Console output
	t.Run("Console", func(t *testing.T) {
		var buf bytes.Buffer
		lg := esl.Default()
		dg := es_dialogue.DenyAll()
		ui := NewConsole(mc, lg, &buf, dg)
		ui.DefinitionList(definitions)
		
		output := buf.String()
		t.Logf("Console output:\n%s", output)
		if !strings.Contains(output, "-path") {
			t.Error("Expected term in console output")
		}
		if !strings.Contains(output, "  Specify the file path") {
			t.Error("Expected indented description in console output")
		}
	})

	// Test Plain output
	t.Run("Plain", func(t *testing.T) {
		output := MakeConsoleDemo(mc, func(ui UI) {
			ui.DefinitionList(definitions)
		})
		
		t.Logf("Plain output:\n%s", output)
		if !strings.Contains(output, "-path") {
			t.Error("Expected term in plain output")
		}
		if !strings.Contains(output, "  Specify the file path") {
			t.Error("Expected indented description in plain output")
		}
	})
}