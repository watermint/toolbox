package dbx_fs_copier_batch

import (
	"encoding/json"
	"testing"
)

func TestConstants(t *testing.T) {
	// Test that constants are defined and non-empty
	constants := []string{
		queueIdBlockCommit,
		queueIdBlockUpload,
		queueIdBlockBatch,
		queueIdBlockCheck,
	}
	
	for i, constant := range constants {
		if constant == "" {
			t.Errorf("Constant %d should not be empty", i)
		}
	}
	
	// Test specific values
	if queueIdBlockCommit != "upload_commit" {
		t.Errorf("Expected queueIdBlockCommit to be 'upload_commit', got %s", queueIdBlockCommit)
	}
	
	if queueIdBlockUpload != "upload_block" {
		t.Errorf("Expected queueIdBlockUpload to be 'upload_block', got %s", queueIdBlockUpload)
	}
	
	if queueIdBlockBatch != "upload_batch" {
		t.Errorf("Expected queueIdBlockBatch to be 'upload_batch', got %s", queueIdBlockBatch)
	}
	
	if queueIdBlockCheck != "upload_check" {
		t.Errorf("Expected queueIdBlockCheck to be 'upload_check', got %s", queueIdBlockCheck)
	}
}

func TestCommitInfo(t *testing.T) {
	info := CommitInfo{
		Path:           "/test/path",
		Mode:           "add",
		Autorename:     true,
		ClientModified: "2023-01-01T00:00:00Z",
		Mute:           false,
		StrictConflict: true,
	}
	
	// Test JSON marshaling
	data, err := json.Marshal(info)
	if err != nil {
		t.Errorf("Failed to marshal CommitInfo: %v", err)
	}
	
	// Test JSON unmarshaling
	var unmarshaled CommitInfo
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal CommitInfo: %v", err)
	}
	
	// Verify fields
	if unmarshaled.Path != info.Path {
		t.Errorf("Expected Path %s, got %s", info.Path, unmarshaled.Path)
	}
	if unmarshaled.Mode != info.Mode {
		t.Errorf("Expected Mode %s, got %s", info.Mode, unmarshaled.Mode)
	}
	if unmarshaled.Autorename != info.Autorename {
		t.Errorf("Expected Autorename %t, got %t", info.Autorename, unmarshaled.Autorename)
	}
}

func TestUploadCursor(t *testing.T) {
	cursor := UploadCursor{
		SessionId: "test-session",
		Offset:    1024,
	}
	
	// Test JSON marshaling
	data, err := json.Marshal(cursor)
	if err != nil {
		t.Errorf("Failed to marshal UploadCursor: %v", err)
	}
	
	// Test JSON unmarshaling
	var unmarshaled UploadCursor
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal UploadCursor: %v", err)
	}
	
	if unmarshaled.SessionId != cursor.SessionId {
		t.Errorf("Expected SessionId %s, got %s", cursor.SessionId, unmarshaled.SessionId)
	}
	if unmarshaled.Offset != cursor.Offset {
		t.Errorf("Expected Offset %d, got %d", cursor.Offset, unmarshaled.Offset)
	}
}

func TestSessionId(t *testing.T) {
	sessionId := SessionId{
		SessionId: "test-session-123",
	}
	
	// Test JSON marshaling
	data, err := json.Marshal(sessionId)
	if err != nil {
		t.Errorf("Failed to marshal SessionId: %v", err)
	}
	
	// Test JSON unmarshaling
	var unmarshaled SessionId
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal SessionId: %v", err)
	}
	
	if unmarshaled.SessionId != sessionId.SessionId {
		t.Errorf("Expected SessionId %s, got %s", sessionId.SessionId, unmarshaled.SessionId)
	}
}

func TestUploadAppend(t *testing.T) {
	upload := UploadAppend{
		Cursor: UploadCursor{
			SessionId: "test-session",
			Offset:    512,
		},
		Close: true,
	}
	
	// Test JSON marshaling
	data, err := json.Marshal(upload)
	if err != nil {
		t.Errorf("Failed to marshal UploadAppend: %v", err)
	}
	
	// Test JSON unmarshaling
	var unmarshaled UploadAppend
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal UploadAppend: %v", err)
	}
	
	if unmarshaled.Close != upload.Close {
		t.Errorf("Expected Close %t, got %t", upload.Close, unmarshaled.Close)
	}
	if unmarshaled.Cursor.SessionId != upload.Cursor.SessionId {
		t.Errorf("Expected Cursor.SessionId %s, got %s", upload.Cursor.SessionId, unmarshaled.Cursor.SessionId)
	}
}

func TestUploadFinish(t *testing.T) {
	finish := UploadFinish{
		Cursor: UploadCursor{
			SessionId: "test-session",
			Offset:    1024,
		},
		Commit: CommitInfo{
			Path: "/test/file.txt",
			Mode: "add",
		},
	}
	
	// Test JSON marshaling
	data, err := json.Marshal(finish)
	if err != nil {
		t.Errorf("Failed to marshal UploadFinish: %v", err)
	}
	
	// Test JSON unmarshaling
	var unmarshaled UploadFinish
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal UploadFinish: %v", err)
	}
	
	if unmarshaled.Commit.Path != finish.Commit.Path {
		t.Errorf("Expected Commit.Path %s, got %s", finish.Commit.Path, unmarshaled.Commit.Path)
	}
}

func TestUploadFinishBatch(t *testing.T) {
	batch := UploadFinishBatch{
		Entries: []UploadFinish{
			{
				Cursor: UploadCursor{SessionId: "session1", Offset: 100},
				Commit: CommitInfo{Path: "/file1.txt"},
			},
			{
				Cursor: UploadCursor{SessionId: "session2", Offset: 200},
				Commit: CommitInfo{Path: "/file2.txt"},
			},
		},
	}
	
	// Test JSON marshaling
	data, err := json.Marshal(batch)
	if err != nil {
		t.Errorf("Failed to marshal UploadFinishBatch: %v", err)
	}
	
	// Test JSON unmarshaling
	var unmarshaled UploadFinishBatch
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal UploadFinishBatch: %v", err)
	}
	
	if len(unmarshaled.Entries) != len(batch.Entries) {
		t.Errorf("Expected %d entries, got %d", len(batch.Entries), len(unmarshaled.Entries))
	}
}

func TestFinishBatch(t *testing.T) {
	batch := FinishBatch{
		Batch: []string{"session1", "session2", "session3"},
	}
	
	// Test JSON marshaling
	data, err := json.Marshal(batch)
	if err != nil {
		t.Errorf("Failed to marshal FinishBatch: %v", err)
	}
	
	// Test JSON unmarshaling
	var unmarshaled FinishBatch
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal FinishBatch: %v", err)
	}
	
	if len(unmarshaled.Batch) != len(batch.Batch) {
		t.Errorf("Expected %d batch items, got %d", len(batch.Batch), len(unmarshaled.Batch))
	}
	
	for i, item := range batch.Batch {
		if unmarshaled.Batch[i] != item {
			t.Errorf("Expected batch[%d] to be %s, got %s", i, item, unmarshaled.Batch[i])
		}
	}
}

func TestSessionCheck(t *testing.T) {
	check := SessionCheck{
		SessionId: "test-session-check",
		Path:      "/check/path",
	}
	
	// Test JSON marshaling
	data, err := json.Marshal(check)
	if err != nil {
		t.Errorf("Failed to marshal SessionCheck: %v", err)
	}
	
	// Test JSON unmarshaling
	var unmarshaled SessionCheck
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Errorf("Failed to unmarshal SessionCheck: %v", err)
	}
	
	if unmarshaled.SessionId != check.SessionId {
		t.Errorf("Expected SessionId %s, got %s", check.SessionId, unmarshaled.SessionId)
	}
	if unmarshaled.Path != check.Path {
		t.Errorf("Expected Path %s, got %s", check.Path, unmarshaled.Path)
	}
}