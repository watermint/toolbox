package uc_file_merge

import (
	"testing"
)

// Test message initialization
func TestMsgMerge(t *testing.T) {
	// Test that MMerge is initialized
	if MMerge == nil {
		t.Error("Expected MMerge to be initialized")
	}

	// Test that messages are accessible
	_ = MMerge.RemoveEmptyFolder
	_ = MMerge.RemoveDuplicatedFile
	_ = MMerge.RemoveOldContent
	_ = MMerge.MoveFile
}
