package mo_filter

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"testing"
)

func TestEmailFilter(t *testing.T) {
	f := NewEmailFilter()
	ef := f.(*emailFilterOpt)

	// Test initial state
	if ef.Enabled() {
		t.Error("Expected filter to be disabled when email is empty")
	}

	// Test Bind
	bound := ef.Bind()
	if _, ok := bound.(*string); !ok {
		t.Error("Expected Bind to return *string")
	}

	// Test NameSuffix
	if ef.NameSuffix() != "Email" {
		t.Errorf("Expected NameSuffix to be 'Email', got %s", ef.NameSuffix())
	}

	// Test Desc
	desc := ef.Desc()
	if desc == nil {
		t.Error("Expected Desc to return non-nil message")
	}

	// Set email
	ef.email = "test@example.com"

	// Test Enabled with email set
	if !ef.Enabled() {
		t.Error("Expected filter to be enabled when email is set")
	}

	// Test Capture
	captured := ef.Capture()
	if captured != "test@example.com" {
		t.Errorf("Expected Capture to return %s, got %v", "test@example.com", captured)
	}

	// Test Restore with valid JSON
	json := es_json.MustParseString(`"restored@example.com"`)
	err := ef.Restore(json)
	if err != nil {
		t.Errorf("Expected Restore to succeed, got error: %v", err)
	}
	if ef.email != "restored@example.com" {
		t.Errorf("Expected email to be restored@example.com, got %s", ef.email)
	}

	// Test Restore with invalid JSON (not a string)
	jsonInvalid := es_json.MustParseString(`123`)
	err = ef.Restore(jsonInvalid)
	if err != rc_recipe.ErrorValueRestoreFailed {
		t.Errorf("Expected Restore to fail with ErrorValueRestoreFailed, got %v", err)
	}
}

func TestEmailFilter_Accept(t *testing.T) {
	tests := []struct {
		name         string
		filterEmail  string
		input        interface{}
		wantAccept   bool
	}{
		// Direct email matches
		{
			name:        "exact match lowercase",
			filterEmail: "test@example.com",
			input:       "test@example.com",
			wantAccept:  true,
		},
		{
			name:        "case insensitive match",
			filterEmail: "Test@Example.com",
			input:       "test@example.com",
			wantAccept:  true,
		},
		{
			name:        "different email",
			filterEmail: "test@example.com",
			input:       "other@example.com",
			wantAccept:  false,
		},
		// Email with display name
		{
			name:        "email with display name - match address",
			filterEmail: "test@example.com",
			input:       "Test User <test@example.com>",
			wantAccept:  true,
		},
		{
			name:        "email with display name - match name",
			filterEmail: "Test User",
			input:       "Test User <test@example.com>",
			wantAccept:  true,
		},
		{
			name:        "email with display name - case insensitive address",
			filterEmail: "TEST@EXAMPLE.COM",
			input:       "Test User <test@example.com>",
			wantAccept:  true,
		},
		{
			name:        "email with quotes in display name",
			filterEmail: "test@example.com",
			input:       `"Test User" <test@example.com>`,
			wantAccept:  true,
		},
		// Invalid inputs
		{
			name:        "invalid email format",
			filterEmail: "test@example.com",
			input:       "not-an-email",
			wantAccept:  false,
		},
		{
			name:        "non-string input",
			filterEmail: "test@example.com",
			input:       123,
			wantAccept:  false,
		},
		{
			name:        "nil input",
			filterEmail: "test@example.com",
			input:       nil,
			wantAccept:  false,
		},
		{
			name:        "empty string input",
			filterEmail: "test@example.com",
			input:       "",
			wantAccept:  false,
		},
		// Edge cases
		{
			name:        "filter email is invalid format",
			filterEmail: "not-an-email",
			input:       "not-an-email",
			wantAccept:  true, // Direct string match
		},
		{
			name:        "empty filter matches empty",
			filterEmail: "",
			input:       "",
			wantAccept:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &emailFilterOpt{email: tt.filterEmail}
			if got := f.Accept(tt.input); got != tt.wantAccept {
				t.Errorf("Accept() = %v, want %v", got, tt.wantAccept)
			}
		})
	}
}