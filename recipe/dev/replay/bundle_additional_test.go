package replay

import (
	"testing"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

func TestBundle_Preset(t *testing.T) {
	b := &Bundle{}
	b.Preset()
	
	if b.Timeout != 60 {
		t.Errorf("Expected Timeout to be 60, got %d", b.Timeout)
	}
	
	if b.PeerName != app_definitions.PeerDeploy {
		t.Errorf("Expected PeerName to be %s, got %s", app_definitions.PeerDeploy, b.PeerName)
	}
	
	expectedPath := "/watermint-toolbox-logs/{{.Date}}-{{.Time}}/{{.Random}}"
	if b.ResultsPath.Path() != expectedPath {
		t.Errorf("Expected ResultsPath to be %s, got %s", expectedPath, b.ResultsPath.Path())
	}
}

func TestBundle_Fields(t *testing.T) {
	// Test field initialization
	b := &Bundle{
		ReplayPath:  mo_string.NewOptional("test-replay"),
		ResultsPath: mo_path.NewDropboxPath("/test/results"),
		PeerName:    "test-peer",
		Timeout:     120,
	}
	
	if !b.ReplayPath.IsExists() || b.ReplayPath.Value() != "test-replay" {
		t.Error("Expected ReplayPath to be set correctly")
	}
	
	if b.ResultsPath.Path() != "/test/results" {
		t.Error("Expected ResultsPath to be set correctly")
	}
	
	if b.PeerName != "test-peer" {
		t.Error("Expected PeerName to be 'test-peer'")
	}
	
	if b.Timeout != 120 {
		t.Error("Expected Timeout to be 120")
	}
}

func TestBundle_EmptyReplayPath(t *testing.T) {
	// Test with empty ReplayPath
	b := &Bundle{
		ResultsPath: mo_path.NewDropboxPath("/results"),
		PeerName:    "peer",
		Timeout:     30,
	}
	
	// ReplayPath will be nil when not initialized
	if b.ReplayPath != nil {
		t.Error("Expected ReplayPath to be nil")
	}
}