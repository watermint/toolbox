package dc_supplemental

import (
	"testing"
	"github.com/watermint/toolbox/infra/doc/dc_index"
)

func TestDocs(t *testing.T) {
	// Test with repository media type
	docs := Docs(dc_index.MediaRepository)
	if docs == nil {
		t.Error("Expected non-nil docs")
	}
	
	// Should return multiple documents
	if len(docs) == 0 {
		t.Error("Expected at least one document")
	}
	
	// Test with web media type
	webDocs := Docs(dc_index.MediaWeb)
	if webDocs == nil {
		t.Error("Expected non-nil docs for web")
	}
	
	if len(webDocs) == 0 {
		t.Error("Expected at least one document for web")
	}
	
	// Should return the same number of docs regardless of media type
	if len(docs) != len(webDocs) {
		t.Errorf("Expected same number of docs, got %d for repository and %d for web", len(docs), len(webDocs))
	}
}