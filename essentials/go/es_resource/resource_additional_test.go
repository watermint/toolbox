package es_resource

import (
	"embed"
	"testing"
)

//go:embed testdata
var testFS embed.FS

func TestNonTraversableResource(t *testing.T) {
	// Create a non-traversable resource
	res := NewNonTraversableResource("testdata", testFS)
	
	// Test reading an existing file
	data, err := res.Bytes("test.txt")
	if err != nil {
		t.Errorf("Expected to read test.txt, got error: %v", err)
	}
	if len(data) == 0 {
		t.Error("Expected non-empty data from test.txt")
	}
	
	// Test reading a non-existing file
	_, err = res.Bytes("nonexistent.txt")
	if err == nil {
		t.Error("Expected error when reading non-existent file")
	}
	
	// Test HttpFileSystem returns empty
	fs := res.HttpFileSystem()
	if fs == nil {
		t.Error("Expected non-nil http.FileSystem")
	}
	
	// Try to open a file through HttpFileSystem (should fail as it's empty)
	f, err := fs.Open("test.txt")
	if err == nil {
		t.Error("Expected error when opening file through empty filesystem")
		if f != nil {
			f.Close()
		}
	}
}

func TestBundleImpl_AllMethods(t *testing.T) {
	// Create test resources
	tpl := EmptyResource()
	msg := EmptyResource()
	web := EmptyResource()
	key := EmptyResource()
	img := EmptyResource()
	dat := EmptyResource()
	bld := EmptyResource()
	rel := EmptyResource()
	
	// Create bundle
	bundle := New(tpl, msg, web, key, img, dat, bld, rel)
	
	// Test all getter methods
	if bundle.Templates() != tpl {
		t.Error("Templates() should return the same resource")
	}
	if bundle.Messages() != msg {
		t.Error("Messages() should return the same resource")
	}
	if bundle.Web() != web {
		t.Error("Web() should return the same resource")
	}
	if bundle.Keys() != key {
		t.Error("Keys() should return the same resource")
	}
	if bundle.Images() != img {
		t.Error("Images() should return the same resource")
	}
	if bundle.Data() != dat {
		t.Error("Data() should return the same resource")
	}
	if bundle.Build() != bld {
		t.Error("Build() should return the same resource")
	}
	if bundle.Release() != rel {
		t.Error("Release() should return the same resource")
	}
}

func TestNewChainBundle(t *testing.T) {
	// Create test bundles
	bundle1 := EmptyBundle()
	bundle2 := EmptyBundle()
	
	langCodes := []string{"en", "ja"}
	
	// Create chain bundle
	chainBundle := NewChainBundle(langCodes, bundle1, bundle2)
	
	// Verify it returns non-nil resources
	if chainBundle.Templates() == nil {
		t.Error("Templates() should not be nil")
	}
	if chainBundle.Messages() == nil {
		t.Error("Messages() should not be nil")
	}
	if chainBundle.Web() == nil {
		t.Error("Web() should not be nil")
	}
	if chainBundle.Keys() == nil {
		t.Error("Keys() should not be nil")
	}
	if chainBundle.Images() == nil {
		t.Error("Images() should not be nil")
	}
	if chainBundle.Data() == nil {
		t.Error("Data() should not be nil")
	}
	if chainBundle.Build() == nil {
		t.Error("Build() should not be nil")
	}
	if chainBundle.Release() == nil {
		t.Error("Release() should not be nil")
	}
	
	// Test with single bundle
	singleChain := NewChainBundle([]string{"en"}, bundle1)
	if singleChain == nil {
		t.Error("Chain bundle should not be nil")
	}
	
	// Test with no language codes
	noLangChain := NewChainBundle([]string{}, bundle1, bundle2)
	if noLangChain == nil {
		t.Error("Chain bundle should not be nil even with no language codes")
	}
}

func TestNonTraversableResource_PathHandling(t *testing.T) {
	res := NewNonTraversableResource("testdata", testFS)
	
	// Test with different path separators
	tests := []struct {
		name     string
		path     string
		wantErr  bool
	}{
		{
			name:    "simple file",
			path:    "test.txt",
			wantErr: false,
		},
		{
			name:    "with backslash",
			path:    "test\\txt", // Will be converted to forward slash
			wantErr: true, // No such file
		},
		{
			name:    "empty path",
			path:    "",
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := res.Bytes(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bytes(%q) error = %v, wantErr %v", tt.path, err, tt.wantErr)
			}
		})
	}
}