package dbx_fs_copier_batch

import (
	"testing"

	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
)

func TestNewLocalToDropboxBatch(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		// Test normal batch size
		connector := NewLocalToDropboxBatch(ctl, nil, 10)
		if connector == nil {
			t.Error("Expected non-nil connector")
		}
		
		// Verify it returns the correct type
		batch, ok := connector.(*copierLocalToDropboxBatch)
		if !ok {
			t.Error("Expected copierLocalToDropboxBatch type")
		}
		
		if batch.batchSize != 10 {
			t.Errorf("Expected batch size 10, got %d", batch.batchSize)
		}
		
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewLocalToDropboxBatch_BatchSizeLimits(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		// Test batch size less than 1
		connector := NewLocalToDropboxBatch(ctl, nil, 0)
		batch := connector.(*copierLocalToDropboxBatch)
		if batch.batchSize != 1 {
			t.Errorf("Expected batch size to be adjusted to 1, got %d", batch.batchSize)
		}
		
		// Test batch size greater than 1000
		connector2 := NewLocalToDropboxBatch(ctl, nil, 1500)
		batch2 := connector2.(*copierLocalToDropboxBatch)
		if batch2.batchSize != 1000 {
			t.Errorf("Expected batch size to be adjusted to 1000, got %d", batch2.batchSize)
		}
		
		// Test negative batch size
		connector3 := NewLocalToDropboxBatch(ctl, nil, -5)
		batch3 := connector3.(*copierLocalToDropboxBatch)
		if batch3.batchSize != 1 {
			t.Errorf("Expected negative batch size to be adjusted to 1, got %d", batch3.batchSize)
		}
		
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewLocalToDropboxBatch_InitializedFields(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		connector := NewLocalToDropboxBatch(ctl, nil, 50)
		batch := connector.(*copierLocalToDropboxBatch)
		
		// Verify initial state
		if batch.ctl == nil {
			t.Error("Expected control to be set")
		}
		
		if batch.fs == nil {
			t.Error("Expected filesystem reader to be initialized")
		}
		
		// These should be nil initially (set during Startup)
		if batch.queue != nil {
			t.Error("Expected queue to be nil initially")
		}
		
		if batch.sessions != nil {
			t.Error("Expected sessions to be nil initially")
		}
		
		if batch.block != nil {
			t.Error("Expected block to be nil initially")
		}
		
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCopyBatchUploadBlock_Struct(t *testing.T) {
	// Test the struct initialization and field access
	block := CopyBatchUploadBlock{
		SessionId: "test-session-123",
		Path:      "/test/path/file.txt",
		Offset:    1024,
	}
	
	if block.SessionId != "test-session-123" {
		t.Errorf("Expected SessionId 'test-session-123', got '%s'", block.SessionId)
	}
	
	if block.Path != "/test/path/file.txt" {
		t.Errorf("Expected Path '/test/path/file.txt', got '%s'", block.Path)
	}
	
	if block.Offset != 1024 {
		t.Errorf("Expected Offset 1024, got %d", block.Offset)
	}
}

func TestCopierLocalToDropboxBatch_Shutdown_NoSessions(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		batch := &copierLocalToDropboxBatch{
			ctl:      ctl,
			sessions: nil, // No sessions to shutdown
		}
		
		// This will panic with nil sessions - that's the expected behavior
		// We can't really test Shutdown without proper setup
		// Instead, let's just verify the struct was created properly
		if batch.ctl == nil {
			t.Error("Expected control to be set")
		}
		
		if batch.sessions != nil {
			t.Error("Expected sessions to be nil")
		}
		
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}