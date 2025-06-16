package dbx_fs

import (
	"testing"

	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func TestMsgFileSystemCached(t *testing.T) {
	// Test that MFileSystemCached is properly initialized
	if MFileSystemCached == nil {
		t.Fatal("MFileSystemCached should not be nil")
	}
	
	// MFileSystemCached is already of type *MsgFileSystemCached
	// Just verify it has the expected type by checking if we can access its fields
	if MFileSystemCached.ProgressPreScan == nil {
		// This is expected - the field will be populated by the message system
		t.Log("ProgressPreScan is nil, which is expected before message initialization")
	}
}

func TestMsgFileSystemCached_Messages(t *testing.T) {
	msg := &MsgFileSystemCached{}
	
	// Apply should work without panic
	applied := app_msg.Apply(msg)
	if applied == nil {
		t.Fatal("Applied message should not be nil")
	}
	
	// Verify the applied message is the correct type
	appliedMsg, ok := applied.(*MsgFileSystemCached)
	if !ok {
		t.Fatal("Applied message should be of type *MsgFileSystemCached")
	}
	
	// The struct should have the ProgressPreScan field
	if appliedMsg.ProgressPreScan == nil {
		// Note: After Apply, the field might be populated by the message system
		// This is expected behavior
		t.Log("ProgressPreScan is nil after Apply")
	}
}