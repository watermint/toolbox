package replay

import (
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"testing"
)

func TestApprove_Preset(t *testing.T) {
	a := &Approve{}
	a.Preset()
	// Preset doesn't do anything, but we test it for coverage
}

func TestApprove_Fields(t *testing.T) {
	// Test field initialization
	a := &Approve{
		Id:            "test-id",
		WorkspacePath: mo_string.NewOptional("test-workspace"),
		ReplayPath:    mo_string.NewOptional("test-replay"),
		Name:          mo_string.NewOptional("test-name"),
	}

	if a.Id != "test-id" {
		t.Error("Expected Id to be 'test-id'")
	}

	if !a.WorkspacePath.IsExists() || a.WorkspacePath.Value() != "test-workspace" {
		t.Error("Expected WorkspacePath to be set correctly")
	}

	if !a.ReplayPath.IsExists() || a.ReplayPath.Value() != "test-replay" {
		t.Error("Expected ReplayPath to be set correctly")
	}

	if !a.Name.IsExists() || a.Name.Value() != "test-name" {
		t.Error("Expected Name to be set correctly")
	}
}

func TestApprove_EmptyFields(t *testing.T) {
	// Test with empty optional fields - they will be nil when not set
	a := &Approve{
		Id: "empty-test",
	}

	// Just verify fields are properly accessible without panicking
	if a.Id != "empty-test" {
		t.Error("Expected Id to be 'empty-test'")
	}

	// Optional fields will be nil when not initialized
	if a.WorkspacePath != nil {
		t.Error("Expected WorkspacePath to be nil")
	}

	if a.ReplayPath != nil {
		t.Error("Expected ReplayPath to be nil")
	}

	if a.Name != nil {
		t.Error("Expected Name to be nil")
	}
}
