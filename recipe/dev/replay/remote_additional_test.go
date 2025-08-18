package replay

import (
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"testing"
)

func TestRemote_Preset(t *testing.T) {
	r := &Remote{}
	r.Preset()
	// Preset doesn't do anything, but we test it for coverage
}

func TestRemote_Fields(t *testing.T) {
	// Test field initialization with ReplayUrl
	r := &Remote{
		ReplayUrl: mo_string.NewOptional("https://example.com/replay.zip"),
	}

	if !r.ReplayUrl.IsExists() || r.ReplayUrl.Value() != "https://example.com/replay.zip" {
		t.Error("Expected ReplayUrl to be set correctly")
	}
}

func TestRemote_EmptyReplayUrl(t *testing.T) {
	// Test with empty ReplayUrl
	r := &Remote{}

	// ReplayUrl will be nil when not initialized
	if r.ReplayUrl != nil {
		t.Error("Expected ReplayUrl to be nil")
	}
}
