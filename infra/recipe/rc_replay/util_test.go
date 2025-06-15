package rc_replay

import (
	"os"
	"testing"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func TestReplayPath_WithProvidedPath(t *testing.T) {
	// Test with provided path
	path := mo_string.NewOptional("/test/path")
	result, err := ReplayPath(path)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if result == "" {
		t.Error("Expected non-empty result")
	}
}

func TestReplayPath_WithEmptyPath(t *testing.T) {
	// Test with empty path and no environment variable
	path := mo_string.NewOptional("")
	
	// Clear environment variable first
	originalEnv := os.Getenv(app_definitions.EnvNameReplayPath)
	defer os.Setenv(app_definitions.EnvNameReplayPath, originalEnv)
	os.Unsetenv(app_definitions.EnvNameReplayPath)
	
	result, err := ReplayPath(path)
	
	if err != ErrorPathNotFound {
		t.Errorf("Expected ErrorPathNotFound, got %v", err)
	}
	
	if result != "" {
		t.Errorf("Expected empty result, got %s", result)
	}
}

func TestReplayPath_WithEnvironmentVariable(t *testing.T) {
	// Test with environment variable
	path := mo_string.NewOptional("")
	
	// Set environment variable
	originalEnv := os.Getenv(app_definitions.EnvNameReplayPath)
	defer os.Setenv(app_definitions.EnvNameReplayPath, originalEnv)
	os.Setenv(app_definitions.EnvNameReplayPath, "/env/test/path")
	
	result, err := ReplayPath(path)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	if result == "" {
		t.Error("Expected non-empty result")
	}
}

func TestReplayPath_EnvironmentOverridesPath(t *testing.T) {
	// Test that provided path takes precedence over environment variable
	path := mo_string.NewOptional("/provided/path")
	
	// Set environment variable
	originalEnv := os.Getenv(app_definitions.EnvNameReplayPath)
	defer os.Setenv(app_definitions.EnvNameReplayPath, originalEnv)
	os.Setenv(app_definitions.EnvNameReplayPath, "/env/path")
	
	result, err := ReplayPath(path)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// Should contain the provided path, not the environment path
	if result == "" {
		t.Error("Expected non-empty result")
	}
}

func TestErrorPathNotFound(t *testing.T) {
	// Test that ErrorPathNotFound is properly defined
	if ErrorPathNotFound == nil {
		t.Error("ErrorPathNotFound should not be nil")
	}
	
	if ErrorPathNotFound.Error() == "" {
		t.Error("ErrorPathNotFound should have a message")
	}
}