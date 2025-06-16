package dbx_fs

import (
	"testing"

	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestNewFileSystem(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_client_impl.NewMock("mock", ctl)
		fs := NewFileSystem(ctx)
		
		if fs == nil {
			t.Fatal("Expected non-nil filesystem")
		}
		
		dbxFs, ok := fs.(*dbxFs)
		if !ok {
			t.Fatal("Expected dbxFs type")
		}
		
		if dbxFs.ctx == nil {
			t.Error("Context not set")
		}
	})
}

func TestDbxFs_OperationalComplexity(t *testing.T) {
	fs := &dbxFs{}
	
	testCases := []struct {
		name       string
		numEntries int
		expected   int64
	}{
		{
			name:       "empty",
			numEntries: 0,
			expected:   1,
		},
		{
			name:       "small",
			numEntries: 100,
			expected:   1,
		},
		{
			name:       "below threshold",
			numEntries: ApiComplexityThreshold - 1,
			expected:   1,
		},
		{
			name:       "at threshold",
			numEntries: ApiComplexityThreshold,
			expected:   1,
		},
		{
			name:       "above threshold",
			numEntries: ApiComplexityThreshold + 1,
			expected:   int64(ApiComplexityThreshold + 1),
		},
		{
			name:       "way above threshold",
			numEntries: ApiComplexityThreshold * 2,
			expected:   int64(ApiComplexityThreshold * 2),
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entries := make([]es_filesystem.Entry, tc.numEntries)
			complexity := fs.OperationalComplexity(entries)
			if complexity != tc.expected {
				t.Errorf("Expected complexity %d, got %d", tc.expected, complexity)
			}
		})
	}
}

func TestDbxFs_FileSystemType(t *testing.T) {
	// Test the constant
	if FileSystemTypeDropbox != "dropbox" {
		t.Errorf("Expected FileSystemTypeDropbox to be 'dropbox', got '%s'", FileSystemTypeDropbox)
	}
}

func TestApiComplexityThreshold(t *testing.T) {
	// Test the constant value
	if ApiComplexityThreshold != 10_000 {
		t.Errorf("Expected ApiComplexityThreshold to be 10000, got %d", ApiComplexityThreshold)
	}
}

func TestDbxFs_EntryInterface(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_client_impl.NewMock("mock", ctl)
		fs := NewFileSystem(ctx)
		
		// Type assertion to ensure interface is implemented
		var _ es_filesystem.FileSystem = fs
	})
}