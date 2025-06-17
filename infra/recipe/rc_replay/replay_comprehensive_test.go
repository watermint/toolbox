package rc_replay

import (
	"encoding/json"
	"testing"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/essentials/network/nw_request"
)

func TestPreserveLogFilePrefixes_Values(t *testing.T) {
	// Test expected prefixes
	// Note: These are examples, actual prefixes are defined in the package
	
	// Verify some common prefixes are included
	prefixMap := make(map[string]bool)
	for _, prefix := range PreserveLogFilePrefixes {
		prefixMap[prefix] = true
	}
	
	// Check that capture is included (most important for replay)
	if !prefixMap["capture"] {
		t.Error("Expected 'capture' to be in PreserveLogFilePrefixes")
	}
	
	// The actual prefixes are defined by constants in other packages
	// Just verify we have some prefixes
	if len(PreserveLogFilePrefixes) < 2 {
		t.Error("Expected at least 2 prefixes in PreserveLogFilePrefixes")
	}
}

func TestReplayImpl_Structure(t *testing.T) {
	// Test rpImpl structure
	logger := esl.Default()
	
	impl := &rpImpl{
		logger: logger,
		opt:    Opts{reportDiffs: true},
	}
	
	if impl.logger == nil {
		t.Error("Expected logger to be set")
	}
	
	if !impl.opt.reportDiffs {
		t.Error("Expected reportDiffs to be true")
	}
}

func TestCapture_JSON(t *testing.T) {
	// Test JSON marshaling/unmarshaling of Capture
	capture := Capture{
		Req: nw_request.Req{
			RequestHash: "hash123",
			// Additional fields would go here
		},
		Res: nw_capture.Res{
			ResponseCode: 200,
			// Additional fields would go here
		},
	}
	
	// Marshal to JSON
	data, err := json.Marshal(capture)
	if err != nil {
		t.Fatalf("Failed to marshal capture: %v", err)
	}
	
	// Unmarshal back
	var decoded Capture
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal capture: %v", err)
	}
	
	// Verify fields
	if decoded.Req.RequestHash != capture.Req.RequestHash {
		t.Errorf("Expected request hash %s, got %s", capture.Req.RequestHash, decoded.Req.RequestHash)
	}
	if decoded.Res.ResponseCode != capture.Res.ResponseCode {
		t.Errorf("Expected response code %d, got %d", capture.Res.ResponseCode, decoded.Res.ResponseCode)
	}
}

func TestNew_WithMultipleOptions(t *testing.T) {
	logger := esl.Default()
	
	// Test with multiple options
	replay := New(logger, 
		ReportDiffs(true),
		ReportDiffs(false), // Second call should override
	)
	
	if replay == nil {
		t.Fatal("Expected non-nil replay")
	}
	
	// Verify it's the correct type
	impl, ok := replay.(*rpImpl)
	if !ok {
		t.Fatal("Expected replay to be *rpImpl")
	}
	
	// Last option should win
	if impl.opt.reportDiffs {
		t.Error("Expected reportDiffs to be false (last option should win)")
	}
}

func TestOpt_Function(t *testing.T) {
	// Test that Opt function type works correctly
	customOpt := func(o Opts) Opts {
		o.reportDiffs = true
		return o
	}
	
	opts := Opts{}
	result := customOpt(opts)
	
	if !result.reportDiffs {
		t.Error("Expected custom option to set reportDiffs to true")
	}
}

