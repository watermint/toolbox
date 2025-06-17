package replay

import (
	"testing"
	"github.com/watermint/toolbox/essentials/model/mo_string"
)

func TestRecipe_Preset(t *testing.T) {
	r := &Recipe{}
	r.Preset()
	// Preset doesn't do anything, but we test it for coverage
}

func TestRecipe_Fields(t *testing.T) {
	// Test field initialization
	r := &Recipe{
		Id:   "test-job-id",
		Path: mo_string.NewOptional("/test/path"),
	}
	
	if r.Id != "test-job-id" {
		t.Error("Expected Id to be 'test-job-id'")
	}
	
	if !r.Path.IsExists() || r.Path.Value() != "/test/path" {
		t.Error("Expected Path to be set correctly")
	}
}

func TestRecipe_EmptyPath(t *testing.T) {
	// Test with empty Path
	r := &Recipe{
		Id: "empty-path-test",
	}
	
	// Path will be nil when not initialized
	if r.Path != nil {
		t.Error("Expected Path to be nil")
	}
}

func TestErrorJobNotFound(t *testing.T) {
	// Test the error constant
	if ErrorJobNotFound.Error() != "job id not found" {
		t.Error("Expected ErrorJobNotFound to have correct message")
	}
}