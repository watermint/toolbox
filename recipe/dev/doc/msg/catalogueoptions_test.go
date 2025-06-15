package msg

import (
	"testing"

	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestCatalogueOptions_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &CatalogueOptions{})
}

func TestCatalogueOptions_DryRun(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &CatalogueOptions{DryRun: true})
}

func TestCatalogueOptions_GenerateOptionDescription(t *testing.T) {
	z := &CatalogueOptions{}
	
	tests := []struct {
		name      string
		fieldName string
		option    string
		expected  string
	}{
		{
			name:      "BasePath root",
			fieldName: "BasePath",
			option:    "root",
			expected:  "Full access to all folders with permissions",
		},
		{
			name:      "BasePath home",
			fieldName: "base_path",
			option:    "home",
			expected:  "Access limited to personal home folder",
		},
		{
			name:      "Visibility public",
			fieldName: "visibility",
			option:    "public",
			expected:  "Anyone with the link can access",
		},
		{
			name:      "Visibility team_only",
			fieldName: "new_visibility",
			option:    "team_only",
			expected:  "Only team members can access",
		},
		{
			name:      "AccessLevel editor",
			fieldName: "access_level",
			option:    "editor",
			expected:  "Can edit, comment, and share",
		},
		{
			name:      "AccessLevel viewer",
			fieldName: "AccessLevel",
			option:    "viewer",
			expected:  "Can view and comment",
		},
		{
			name:      "ManagementType company",
			fieldName: "management_type",
			option:    "company_managed",
			expected:  "Managed by company administrators",
		},
		{
			name:      "ManagementType user",
			fieldName: "ManagementType",
			option:    "user_managed",
			expected:  "Managed by individual users",
		},
		{
			name:      "Format HTML",
			fieldName: "format",
			option:    "html",
			expected:  "HTML format",
		},
		{
			name:      "Format markdown",
			fieldName: "Format",
			option:    "markdown",
			expected:  "Markdown format",
		},
		{
			name:      "Method block",
			fieldName: "method",
			option:    "block",
			expected:  "Block upload method (parallel chunks)",
		},
		{
			name:      "State open",
			fieldName: "state",
			option:    "open",
			expected:  "Open issues only",
		},
		{
			name:      "Generic field",
			fieldName: "some_field",
			option:    "some_option",
			expected:  "some field: some_option",
		},
		{
			name:      "Unknown visibility option",
			fieldName: "visibility",
			option:    "unknown",
			expected:  "Visibility option: unknown",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := z.generateOptionDescription(tt.fieldName, tt.option)
			if result != tt.expected {
				t.Errorf("generateOptionDescription(%q, %q) = %q, want %q", 
					tt.fieldName, tt.option, result, tt.expected)
			}
		})
	}
}


func TestCatalogueOptions_EmptyOptions(t *testing.T) {
	z := &CatalogueOptions{}
	
	// Test empty option handling
	result := z.generateOptionDescription("field", "")
	expected := "field: "
	if result != expected {
		t.Errorf("generateOptionDescription with empty option = %q, want %q", result, expected)
	}
}

func TestCatalogueOptions_FieldNameCaseInsensitive(t *testing.T) {
	z := &CatalogueOptions{}
	
	// Test that field names are case insensitive
	testCases := []struct {
		fieldName string
		option    string
		expected  string
	}{
		{"BASEPATH", "root", "Full access to all folders with permissions"},
		{"BasePath", "root", "Full access to all folders with permissions"},
		{"basepath", "root", "Full access to all folders with permissions"},
		{"base_path", "root", "Full access to all folders with permissions"},
		{"BASE_PATH", "root", "Full access to all folders with permissions"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.fieldName, func(t *testing.T) {
			result := z.generateOptionDescription(tc.fieldName, tc.option)
			if result != tc.expected {
				t.Errorf("generateOptionDescription(%q, %q) = %q, want %q",
					tc.fieldName, tc.option, result, tc.expected)
			}
		})
	}
}

func TestCatalogueOptions_AllFormats(t *testing.T) {
	z := &CatalogueOptions{}
	
	// Test all format options
	formats := []struct {
		option   string
		expected string
	}{
		{"html", "HTML format"},
		{"markdown", "Markdown format"},
		{"plain_text", "Plain text format"},
		{"pdf", "PDF document format"},
		{"unknown_format", "Format: unknown_format"},
	}
	
	for _, f := range formats {
		t.Run(f.option, func(t *testing.T) {
			result := z.generateOptionDescription("format", f.option)
			if result != f.expected {
				t.Errorf("generateOptionDescription(\"format\", %q) = %q, want %q",
					f.option, result, f.expected)
			}
		})
	}
}