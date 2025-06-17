package uc_file_merge

import (
	"testing"
)

// Test the option functions
func TestMergeOptions(t *testing.T) {
	opts := &MergeOpts{}
	
	// Test DryRun
	opts = DryRun()(opts)
	if !opts.DryRun {
		t.Error("Expected DryRun to be true")
	}
	
	// Test WithinSameNamespace
	opts = &MergeOpts{}
	opts = WithinSameNamespace()(opts)
	if !opts.WithinSameNamespace {
		t.Error("Expected WithinSameNamespace to be true")
	}
	
	// Test ClearEmptyFolder
	opts = &MergeOpts{}
	opts = ClearEmptyFolder()(opts)
	if !opts.CleanEmptyFolder {
		t.Error("Expected CleanEmptyFolder to be true")
	}
}