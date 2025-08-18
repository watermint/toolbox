package rc_replay

import (
	"os"
	"testing"

	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func TestReplayPath_ErrorCases(t *testing.T) {
	// Clear any existing env var
	originalEnv := os.Getenv(app_definitions.EnvNameReplayPath)
	os.Unsetenv(app_definitions.EnvNameReplayPath)
	defer func() {
		if originalEnv != "" {
			os.Setenv(app_definitions.EnvNameReplayPath, originalEnv)
		}
	}()

	// Test with empty optional string
	emptyOpt := mo_string.NewOptional("")
	_, err := ReplayPath(emptyOpt)
	if err != ErrorPathNotFound {
		t.Errorf("Expected ErrorPathNotFound, got %v", err)
	}
}

func TestReplayPath_PathFormatting(t *testing.T) {
	// Test path with predefined variables
	testCases := []struct {
		name     string
		input    string
		hasError bool
	}{
		{
			name:     "simple path",
			input:    "/tmp/replay",
			hasError: false,
		},
		{
			name:     "path with home",
			input:    "~/replay",
			hasError: false,
		},
		{
			name:     "relative path",
			input:    "./replay",
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			opt := mo_string.NewOptional(tc.input)
			path, err := ReplayPath(opt)

			if tc.hasError && err == nil {
				t.Error("Expected error but got nil")
			}
			if !tc.hasError && err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			if !tc.hasError && path == "" {
				t.Error("Expected non-empty path")
			}
		})
	}
}

func TestReplayPath_ComplexScenarios(t *testing.T) {
	// Save original env
	originalEnv := os.Getenv(app_definitions.EnvNameReplayPath)
	defer func() {
		if originalEnv != "" {
			os.Setenv(app_definitions.EnvNameReplayPath, originalEnv)
		} else {
			os.Unsetenv(app_definitions.EnvNameReplayPath)
		}
	}()

	// Test with both path and env var set - path should take precedence
	os.Setenv(app_definitions.EnvNameReplayPath, "/env/replay")
	opt := mo_string.NewOptional("/direct/replay")
	path, err := ReplayPath(opt)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if path != "/direct/replay" {
		t.Errorf("Expected direct path to take precedence, got %s", path)
	}
}

func TestErrorPathNotFound_Properties(t *testing.T) {
	// Verify error message
	expectedMsg := "replay path not found"
	if ErrorPathNotFound.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, ErrorPathNotFound.Error())
	}
}
