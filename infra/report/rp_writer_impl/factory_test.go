package rp_writer_impl

import (
	"testing"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
)

func TestNew(t *testing.T) {
	// Test in normal mode
	ctl := qt_control.WithFeature(false)
	writer := New("test", ctl)
	if writer == nil {
		t.Error("Expected non-nil writer")
	}

	// Test in test mode
	testCtl := qt_control.WithFeature(true)
	testWriter := New("test", testCtl)
	if testWriter == nil {
		t.Error("Expected non-nil writer in test mode")
	}
}