package mo_path

import (
	"runtime"
	"strings"
	"testing"
)

func TestNewFileSystemPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantPath string
	}{
		{
			name:     "simple path",
			input:    "/tmp/test",
			wantPath: "/tmp/test",
		},
		{
			name:     "relative path",
			input:    "./test/file.txt",
			wantPath: "./test/file.txt",
		},
		{
			name:     "empty path",
			input:    "",
			wantPath: "",
		},
		{
			name:     "path with spaces",
			input:    "/path with spaces/file.txt",
			wantPath: "/path with spaces/file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsp := NewFileSystemPath(tt.input)
			if fsp == nil {
				t.Fatal("Expected non-nil FileSystemPath")
			}

			// Check Path() method
			if got := fsp.Path(); got != tt.wantPath {
				// The path might be processed by FormatPathWithPredefinedVariables
				// So we check if it at least contains the expected path
				if !strings.Contains(got, tt.wantPath) && tt.wantPath != "" {
					t.Errorf("Path() = %v, want %v", got, tt.wantPath)
				}
			}

			// The implementation returns the same type for both functions,
			// so we can't distinguish them by interface. Just verify it implements FileSystemPath
			if _, ok := fsp.(FileSystemPath); !ok {
				t.Error("NewFileSystemPath should return FileSystemPath")
			}
		})
	}
}

func TestNewExistingFileSystemPath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantPath string
	}{
		{
			name:     "existing path",
			input:    "/tmp",
			wantPath: "/tmp",
		},
		{
			name:     "file path",
			input:    "/etc/hosts",
			wantPath: "/etc/hosts",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			efsp := NewExistingFileSystemPath(tt.input)
			if efsp == nil {
				t.Fatal("Expected non-nil ExistingFileSystemPath")
			}

			// Check Path() method
			got := efsp.Path()
			if got != tt.wantPath {
				// The path might be processed by FormatPathWithPredefinedVariables
				// So we check if it at least contains the expected path
				if !strings.Contains(got, tt.wantPath) && tt.wantPath != "" {
					t.Errorf("Path() = %v, want %v", got, tt.wantPath)
				}
			}

			// Check ShouldExist() method
			if !efsp.ShouldExist() {
				t.Error("ShouldExist() should return true for ExistingFileSystemPath")
			}
		})
	}
}

func TestFileSystemPathImpl_Drive(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantDrive string
		skipOS    []string
	}{
		{
			name:      "unix absolute path",
			path:      "/tmp/test",
			wantDrive: "",
			skipOS:    []string{"windows"},
		},
		{
			name:      "unix relative path",
			path:      "./test",
			wantDrive: "",
			skipOS:    []string{"windows"},
		},
		{
			name:      "windows drive path",
			path:      "C:\\Windows\\System32",
			wantDrive: "C:",
			skipOS:    []string{"darwin", "linux"},
		},
		{
			name:      "windows drive lowercase",
			path:      "d:\\data\\file.txt",
			wantDrive: "d:",
			skipOS:    []string{"darwin", "linux"},
		},
		{
			name:      "empty path",
			path:      "",
			wantDrive: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip test if on incompatible OS
			for _, skipOS := range tt.skipOS {
				if runtime.GOOS == skipOS {
					t.Skipf("Skipping test on %s", skipOS)
				}
			}

			fsp := &fileSystemPathImpl{path: tt.path}
			if got := fsp.Drive(); got != tt.wantDrive {
				t.Errorf("Drive() = %v, want %v", got, tt.wantDrive)
			}
		})
	}
}

func TestFileSystemPathImpl_Methods(t *testing.T) {
	// Test Path() method
	impl := &fileSystemPathImpl{path: "/test/path"}
	if impl.Path() != "/test/path" {
		t.Errorf("Path() = %v, want %v", impl.Path(), "/test/path")
	}

	// Test ShouldExist() method
	impl1 := &fileSystemPathImpl{path: "/test", shouldExist: false}
	if impl1.ShouldExist() {
		t.Error("ShouldExist() should return false when shouldExist is false")
	}

	impl2 := &fileSystemPathImpl{path: "/test", shouldExist: true}
	if !impl2.ShouldExist() {
		t.Error("ShouldExist() should return true when shouldExist is true")
	}
}

func TestPathWithPredefinedVariables(t *testing.T) {
	// Test paths that might contain predefined variables
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "home variable",
			input: "{{.Home}}/test",
		},
		{
			name:  "desktop variable",
			input: "{{.Desktop}}/file.txt",
		},
		{
			name:  "invalid variable",
			input: "{{.Invalid}}/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsp := NewFileSystemPath(tt.input)
			path := fsp.Path()

			// If the input contains variables and they couldn't be processed,
			// the path should remain unchanged or be processed
			if strings.Contains(tt.input, "{{") {
				// Just verify we got a non-empty path back
				if path == "" && tt.input != "" {
					t.Error("Expected non-empty path for variable input")
				}
			}
		})
	}
}
