package es_sync

import (
	"errors"
	"testing"

	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
)

func TestSkipReasonConstants(t *testing.T) {
	if SkipSame != "same" {
		t.Errorf("Expected SkipSame to be 'same', got %s", SkipSame)
	}
	if SkipOld != "old" {
		t.Errorf("Expected SkipOld to be 'old', got %s", SkipOld)
	}
	if SkipExists != "exists" {
		t.Errorf("Expected SkipExists to be 'exists', got %s", SkipExists)
	}
	if SkipFilter != "filter" {
		t.Errorf("Expected SkipFilter to be 'filter', got %s", SkipFilter)
	}
}

func TestOpts_DefaultValues(t *testing.T) {
	opts := Opts{}
	
	if opts.SyncDelete() {
		t.Error("Expected SyncDelete to be false by default")
	}
	if opts.SyncOverwrite() {
		t.Error("Expected SyncOverwrite to be false by default")
	}
	if opts.SyncDontCompareTime() {
		t.Error("Expected SyncDontCompareTime to be false by default")
	}
	if opts.SyncDontCompareContent() {
		t.Error("Expected SyncDontCompareContent to be false by default")
	}
	if opts.OptimizeReduceCreateFolder() {
		t.Error("Expected OptimizeReduceCreateFolder to be false by default")
	}
	if opts.Progress() != nil {
		t.Error("Expected Progress to be nil by default")
	}
}

func TestSyncDelete(t *testing.T) {
	opts := Opts{}
	
	// Test enabling
	newOpts := SyncDelete(true)(opts)
	if !newOpts.SyncDelete() {
		t.Error("Expected SyncDelete to be true after enabling")
	}
	
	// Test disabling
	newOpts = SyncDelete(false)(newOpts)
	if newOpts.SyncDelete() {
		t.Error("Expected SyncDelete to be false after disabling")
	}
}

func TestSyncOverwrite(t *testing.T) {
	opts := Opts{}
	
	newOpts := SyncOverwrite(true)(opts)
	if !newOpts.SyncOverwrite() {
		t.Error("Expected SyncOverwrite to be true after enabling")
	}
	
	newOpts = SyncOverwrite(false)(newOpts)
	if newOpts.SyncOverwrite() {
		t.Error("Expected SyncOverwrite to be false after disabling")
	}
}

func TestSyncDontCompareTime(t *testing.T) {
	opts := Opts{}
	
	newOpts := SyncDontCompareTime(true)(opts)
	if !newOpts.SyncDontCompareTime() {
		t.Error("Expected SyncDontCompareTime to be true after enabling")
	}
}

func TestSyncDontCompareContent(t *testing.T) {
	opts := Opts{}
	
	newOpts := SyncDontCompareContent(true)(opts)
	if !newOpts.SyncDontCompareContent() {
		t.Error("Expected SyncDontCompareContent to be true after enabling")
	}
}

func TestOptimizePreventCreateFolder(t *testing.T) {
	opts := Opts{}
	
	newOpts := OptimizePreventCreateFolder(true)(opts)
	if !newOpts.OptimizeReduceCreateFolder() {
		t.Error("Expected OptimizeReduceCreateFolder to be true after enabling")
	}
}

func TestWithNameFilter(t *testing.T) {
	opts := Opts{}
	
	// Use nil filter for simplicity in testing
	_ = WithNameFilter(nil)(opts)
	
	// The option function should work without error
	// We can't directly test the filter since it's private
}

func TestOnCopySuccess(t *testing.T) {
	opts := Opts{}
	called := false
	
	listener := func(source es_filesystem.Entry, target es_filesystem.Entry) {
		called = true
	}
	
	newOpts := OnCopySuccess(listener)(opts)
	
	// Call with nil entries to test the listener mechanism
	newOpts.OnCopySuccess(nil, nil)
	
	if !called {
		t.Error("Expected OnCopySuccess listener to be called")
	}
}

func TestOnCopyFailure(t *testing.T) {
	opts := Opts{}
	called := false
	
	listener := func(source es_filesystem.Path, err es_filesystem.FileSystemError) {
		called = true
	}
	
	newOpts := OnCopyFailure(listener)(opts)
	
	sourcePath := es_filesystem_model.NewPath("/source.txt")
	mockError := es_filesystem_model.NewError(errors.New("test error"), es_filesystem_model.ErrorTypeOther)
	
	newOpts.OnCopyFailure(sourcePath, mockError)
	
	if !called {
		t.Error("Expected OnCopyFailure listener to be called")
	}
}

func TestOnDeleteSuccess(t *testing.T) {
	opts := Opts{}
	called := false
	
	listener := func(target es_filesystem.Path) {
		called = true
	}
	
	newOpts := OnDeleteSuccess(listener)(opts)
	
	targetPath := es_filesystem_model.NewPath("/target.txt")
	newOpts.OnDeleteSuccess(targetPath)
	
	if !called {
		t.Error("Expected OnDeleteSuccess listener to be called")
	}
}

func TestOnDeleteFailure(t *testing.T) {
	opts := Opts{}
	called := false
	
	listener := func(target es_filesystem.Path, err es_filesystem.FileSystemError) {
		called = true
	}
	
	newOpts := OnDeleteFailure(listener)(opts)
	
	targetPath := es_filesystem_model.NewPath("/target.txt")
	mockError := es_filesystem_model.NewError(errors.New("test error"), es_filesystem_model.ErrorTypeOther)
	
	newOpts.OnDeleteFailure(targetPath, mockError)
	
	if !called {
		t.Error("Expected OnDeleteFailure listener to be called")
	}
}

func TestOnCreateFolderSuccess(t *testing.T) {
	opts := Opts{}
	called := false
	
	listener := func(target es_filesystem.Path) {
		called = true
	}
	
	newOpts := OnCreateFolderSuccess(listener)(opts)
	
	targetPath := es_filesystem_model.NewPath("/newfolder")
	newOpts.OnCreateFolderSuccess(targetPath)
	
	if !called {
		t.Error("Expected OnCreateFolderSuccess listener to be called")
	}
}

func TestOnCreateFolderFailure(t *testing.T) {
	opts := Opts{}
	called := false
	
	listener := func(target es_filesystem.Path, err es_filesystem.FileSystemError) {
		called = true
	}
	
	newOpts := OnCreateFolderFailure(listener)(opts)
	
	targetPath := es_filesystem_model.NewPath("/newfolder")
	mockError := es_filesystem_model.NewError(errors.New("test error"), es_filesystem_model.ErrorTypeOther)
	
	newOpts.OnCreateFolderFailure(targetPath, mockError)
	
	if !called {
		t.Error("Expected OnCreateFolderFailure listener to be called")
	}
}

func TestOnSkip(t *testing.T) {
	opts := Opts{}
	called := false
	var receivedReason SkipReason
	
	listener := func(reason SkipReason, source es_filesystem.Entry, target es_filesystem.Path) {
		called = true
		receivedReason = reason
	}
	
	newOpts := OnSkip(listener)(opts)
	
	targetPath := es_filesystem_model.NewPath("/target.txt")
	
	newOpts.OnSkip(SkipSame, nil, targetPath)
	
	if !called {
		t.Error("Expected OnSkip listener to be called")
	}
	if receivedReason != SkipSame {
		t.Errorf("Expected reason to be SkipSame, got %s", receivedReason)
	}
}

func TestOpts_ListenersWithNil(t *testing.T) {
	opts := Opts{}
	
	// Test that calling listeners when they're nil doesn't panic
	sourcePath := es_filesystem_model.NewPath("/source.txt")
	targetPath := es_filesystem_model.NewPath("/target.txt")
	mockError := es_filesystem_model.NewError(errors.New("test error"), es_filesystem_model.ErrorTypeOther)
	
	// These should not panic
	opts.OnCopySuccess(nil, nil)
	opts.OnCopyFailure(sourcePath, mockError)
	opts.OnDeleteSuccess(targetPath)
	opts.OnDeleteFailure(targetPath, mockError)
	opts.OnCreateFolderSuccess(targetPath)
	opts.OnCreateFolderFailure(targetPath, mockError)
	opts.OnSkip(SkipSame, nil, targetPath)
}

func TestOpts_Apply_NoOptions(t *testing.T) {
	opts := Opts{}
	
	result := opts.Apply([]Opt{})
	
	// Should return the same opts
	if result.SyncDelete() != opts.SyncDelete() {
		t.Error("Apply with no options should return same opts")
	}
}

func TestOpts_Apply_SingleOption(t *testing.T) {
	opts := Opts{}
	
	result := opts.Apply([]Opt{SyncDelete(true)})
	
	if !result.SyncDelete() {
		t.Error("Expected SyncDelete to be true after applying single option")
	}
}

func TestOpts_Apply_MultipleOptions(t *testing.T) {
	opts := Opts{}
	
	result := opts.Apply([]Opt{
		SyncDelete(true),
		SyncOverwrite(true),
		SyncDontCompareTime(true),
	})
	
	if !result.SyncDelete() {
		t.Error("Expected SyncDelete to be true")
	}
	if !result.SyncOverwrite() {
		t.Error("Expected SyncOverwrite to be true")
	}
	if !result.SyncDontCompareTime() {
		t.Error("Expected SyncDontCompareTime to be true")
	}
}

func TestOpts_Apply_OptionsOverride(t *testing.T) {
	opts := Opts{}
	
	// Apply conflicting options - last one should win
	result := opts.Apply([]Opt{
		SyncDelete(true),
		SyncDelete(false),
	})
	
	if result.SyncDelete() {
		t.Error("Expected later option to override earlier one")
	}
}

func TestWithProgress(t *testing.T) {
	opts := Opts{}
	
	// Create a mock progress container (nil is valid)
	newOpts := WithProgress(nil)(opts)
	
	if newOpts.Progress() != nil {
		t.Error("Expected progress to be nil")
	}
	
	// We can't easily test with a real container without more dependencies,
	// but we can verify the option function works
}