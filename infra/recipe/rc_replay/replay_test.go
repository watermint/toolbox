package rc_replay

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/essentials/network/nw_request"
	"testing"
)

func TestNew(t *testing.T) {
	logger := esl.Default()
	
	// Test with no options
	replay := New(logger)
	if replay == nil {
		t.Error("Expected non-nil replay instance")
	}
	
	// Test the replay implements the interface
	var _ Replay = replay
	
	// Test with options
	replayWithOpts := New(logger, ReportDiffs(true))
	if replayWithOpts == nil {
		t.Error("Expected non-nil replay instance with options")
	}
}

func TestReportDiffs(t *testing.T) {
	// Test ReportDiffs option function
	opt := ReportDiffs(true)
	if opt == nil {
		t.Error("Expected non-nil option function")
	}
	
	opts := Opts{}
	result := opt(opts)
	if !result.reportDiffs {
		t.Error("Expected reportDiffs to be true")
	}
	
	// Test with false
	optFalse := ReportDiffs(false)
	resultFalse := optFalse(opts)
	if resultFalse.reportDiffs {
		t.Error("Expected reportDiffs to be false")
	}
}

func TestOpts_Apply(t *testing.T) {
	opts := Opts{}
	
	// Test with no options
	result := opts.Apply([]Opt{})
	if result.reportDiffs {
		t.Error("Expected default reportDiffs to be false")
	}
	
	// Test with single option
	result = opts.Apply([]Opt{ReportDiffs(true)})
	if !result.reportDiffs {
		t.Error("Expected reportDiffs to be true")
	}
	
	// Test with multiple options
	result = opts.Apply([]Opt{ReportDiffs(true), ReportDiffs(false)})
	if result.reportDiffs {
		t.Error("Expected last option to override (reportDiffs should be false)")
	}
}

func TestCapture(t *testing.T) {
	// Test Capture struct creation
	req := nw_request.Req{RequestHash: "test-hash"}
	res := nw_capture.Res{ResponseCode: 200}
	
	capture := Capture{
		Req: req,
		Res: res,
	}
	
	if capture.Req.RequestHash != "test-hash" {
		t.Error("Expected request hash to be 'test-hash'")
	}
	
	if capture.Res.ResponseCode != 200 {
		t.Error("Expected response code to be 200")
	}
}

func TestPreserveLogFilePrefixes(t *testing.T) {
	// Test that PreserveLogFilePrefixes is properly defined
	if PreserveLogFilePrefixes == nil {
		t.Error("Expected PreserveLogFilePrefixes to be defined")
	}
	
	if len(PreserveLogFilePrefixes) == 0 {
		t.Error("Expected PreserveLogFilePrefixes to have at least one entry")
	}
	
	// Test that all entries are non-empty strings
	for i, prefix := range PreserveLogFilePrefixes {
		if prefix == "" {
			t.Errorf("Expected prefix at index %d to be non-empty", i)
		}
	}
}

func TestErrorReportDiffFound(t *testing.T) {
	// Test that ErrorReportDiffFound is properly defined
	if ErrorReportDiffFound == nil {
		t.Error("Expected ErrorReportDiffFound to be defined")
	}
	
	if ErrorReportDiffFound.Error() == "" {
		t.Error("Expected ErrorReportDiffFound to have a message")
	}
}

func TestReplayInterface(t *testing.T) {
	// Test that rpImpl implements Replay interface
	logger := esl.Default()
	replay := New(logger)
	
	// This will fail at compile time if rpImpl doesn't implement Replay
	var _ Replay = replay
}

func TestOpts_Defaults(t *testing.T) {
	// Test default values
	opts := Opts{}
	if opts.reportDiffs {
		t.Error("Expected default reportDiffs to be false")
	}
}