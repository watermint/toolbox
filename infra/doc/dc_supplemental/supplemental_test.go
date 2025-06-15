package dc_supplemental

import (
	"testing"
	"github.com/watermint/toolbox/infra/doc/dc_index"
)

func TestDocs(t *testing.T) {
	// Test with markdown media type
	docs := Docs(dc_index.MediaTypeMarkdown)
	if docs == nil {
		t.Error("Expected non-nil docs")
	}
	
	// Should return multiple documents
	if len(docs) == 0 {
		t.Error("Expected at least one document")
	}
	
	// Test with HTML media type
	htmlDocs := Docs(dc_index.MediaTypeHtml)
	if htmlDocs == nil {
		t.Error("Expected non-nil docs for HTML")
	}
	
	if len(htmlDocs) == 0 {
		t.Error("Expected at least one document for HTML")
	}
	
	// Should return the same number of docs regardless of media type
	if len(docs) != len(htmlDocs) {
		t.Errorf("Expected same number of docs, got %d for markdown and %d for HTML", len(docs), len(htmlDocs))
	}
}