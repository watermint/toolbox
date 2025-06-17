package em_file_random

import (
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	opts := Default()
	
	// Check default values
	if opts.fileSizeRangeMax != 2048 {
		t.Errorf("Expected fileSizeRangeMax to be 2048, got %d", opts.fileSizeRangeMax)
	}
	if opts.fileSizeRangeMin != 0 {
		t.Errorf("Expected fileSizeRangeMin to be 0, got %d", opts.fileSizeRangeMin)
	}
	if opts.maxFilesInFolder != 1<<15 {
		t.Errorf("Expected maxFilesInFolder to be %d, got %d", 1<<15, opts.maxFilesInFolder)
	}
	if opts.maxFoldersInFolder != 64 {
		t.Errorf("Expected maxFoldersInFolder to be 64, got %d", opts.maxFoldersInFolder)
	}
	if opts.numFiles != 1000 {
		t.Errorf("Expected numFiles to be 1000, got %d", opts.numFiles)
	}
	if opts.depthRangeMax != 8 {
		t.Errorf("Expected depthRangeMax to be 8, got %d", opts.depthRangeMax)
	}
	
	// Check date ranges are reasonable
	now := time.Now()
	if opts.fileDateRangeMax.After(now.Add(time.Minute)) {
		t.Error("Expected fileDateRangeMax to be around now")
	}
	expectedMin := now.Add(-2 * 365 * 24 * time.Hour)
	if opts.fileDateRangeMin.Before(expectedMin.Add(-time.Hour)) || opts.fileDateRangeMin.After(expectedMin.Add(time.Hour)) {
		t.Error("Expected fileDateRangeMin to be around 2 years ago")
	}
	
	// Seed should be non-zero
	if opts.seed == 0 {
		t.Error("Expected seed to be non-zero")
	}
}

func TestOpts_Apply(t *testing.T) {
	// Test with no options
	opts := Default()
	result := opts.Apply([]Opt{})
	if result.numFiles != opts.numFiles {
		t.Error("Apply with no options should return unchanged opts")
	}
	
	// Test with single option
	opt1 := NumFiles(500)
	result = opts.Apply([]Opt{opt1})
	if result.numFiles != 500 {
		t.Errorf("Expected numFiles to be 500, got %d", result.numFiles)
	}
	
	// Test with multiple options
	opt2 := Depth(10)
	opt3 := Seed(12345)
	result = opts.Apply([]Opt{opt1, opt2, opt3})
	if result.numFiles != 500 {
		t.Errorf("Expected numFiles to be 500, got %d", result.numFiles)
	}
	if result.depthRangeMax != 10 {
		t.Errorf("Expected depthRangeMax to be 10, got %d", result.depthRangeMax)
	}
	if result.seed != 12345 {
		t.Errorf("Expected seed to be 12345, got %d", result.seed)
	}
}

func TestFileSize(t *testing.T) {
	opts := Default()
	modified := FileSize(100, 5000)(opts)
	
	if modified.fileSizeRangeMin != 100 {
		t.Errorf("Expected fileSizeRangeMin to be 100, got %d", modified.fileSizeRangeMin)
	}
	if modified.fileSizeRangeMax != 5000 {
		t.Errorf("Expected fileSizeRangeMax to be 5000, got %d", modified.fileSizeRangeMax)
	}
	
	// Other fields should remain unchanged
	if modified.numFiles != opts.numFiles {
		t.Error("FileSize option should not modify numFiles")
	}
}

func TestFileDate(t *testing.T) {
	opts := Default()
	minDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
	
	modified := FileDate(minDate, maxDate)(opts)
	
	if !modified.fileDateRangeMin.Equal(minDate) {
		t.Errorf("Expected fileDateRangeMin to be %v, got %v", minDate, modified.fileDateRangeMin)
	}
	if !modified.fileDateRangeMax.Equal(maxDate) {
		t.Errorf("Expected fileDateRangeMax to be %v, got %v", maxDate, modified.fileDateRangeMax)
	}
}

func TestDepth(t *testing.T) {
	opts := Default()
	modified := Depth(15)(opts)
	
	if modified.depthRangeMax != 15 {
		t.Errorf("Expected depthRangeMax to be 15, got %d", modified.depthRangeMax)
	}
	
	// Test with zero depth
	modified = Depth(0)(opts)
	if modified.depthRangeMax != 0 {
		t.Errorf("Expected depthRangeMax to be 0, got %d", modified.depthRangeMax)
	}
}

func TestNumDescendant(t *testing.T) {
	opts := Default()
	modified := NumDescendant(100, 20)(opts)
	
	if modified.maxFilesInFolder != 100 {
		t.Errorf("Expected maxFilesInFolder to be 100, got %d", modified.maxFilesInFolder)
	}
	if modified.maxFoldersInFolder != 20 {
		t.Errorf("Expected maxFoldersInFolder to be 20, got %d", modified.maxFoldersInFolder)
	}
}

func TestNumFiles(t *testing.T) {
	opts := Default()
	modified := NumFiles(2500)(opts)
	
	if modified.numFiles != 2500 {
		t.Errorf("Expected numFiles to be 2500, got %d", modified.numFiles)
	}
}

func TestSeed(t *testing.T) {
	opts := Default()
	modified := Seed(9876543210)(opts)
	
	if modified.seed != 9876543210 {
		t.Errorf("Expected seed to be 9876543210, got %d", modified.seed)
	}
}

func TestChainedOptions(t *testing.T) {
	// Test that options can be chained together
	opts := Default().Apply([]Opt{
		NumFiles(3000),
		Depth(12),
		FileSize(512, 4096),
		Seed(11111),
		NumDescendant(50, 10),
	})
	
	if opts.numFiles != 3000 {
		t.Errorf("Expected numFiles to be 3000, got %d", opts.numFiles)
	}
	if opts.depthRangeMax != 12 {
		t.Errorf("Expected depthRangeMax to be 12, got %d", opts.depthRangeMax)
	}
	if opts.fileSizeRangeMin != 512 {
		t.Errorf("Expected fileSizeRangeMin to be 512, got %d", opts.fileSizeRangeMin)
	}
	if opts.fileSizeRangeMax != 4096 {
		t.Errorf("Expected fileSizeRangeMax to be 4096, got %d", opts.fileSizeRangeMax)
	}
	if opts.seed != 11111 {
		t.Errorf("Expected seed to be 11111, got %d", opts.seed)
	}
	if opts.maxFilesInFolder != 50 {
		t.Errorf("Expected maxFilesInFolder to be 50, got %d", opts.maxFilesInFolder)
	}
	if opts.maxFoldersInFolder != 10 {
		t.Errorf("Expected maxFoldersInFolder to be 10, got %d", opts.maxFoldersInFolder)
	}
}