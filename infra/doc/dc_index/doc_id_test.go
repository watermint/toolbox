package dc_index

import (
	"strings"
	"testing"
	"github.com/watermint/toolbox/essentials/go/es_lang"
)

func TestGeneratedPath(t *testing.T) {
	// Test with default language
	lang := es_lang.Default
	result := GeneratedPath(lang, "test-doc")
	
	if !strings.Contains(result, "test-doc") {
		t.Errorf("Expected result to contain 'test-doc', got %s", result)
	}
	
	// Test with Japanese language
	jaLang := es_lang.Japanese
	result = GeneratedPath(jaLang, "test-doc")
	
	if !strings.Contains(result, "test-doc") {
		t.Errorf("Expected result to contain 'test-doc', got %s", result)
	}
	
	if !strings.Contains(result, jaLang.Suffix()) {
		t.Errorf("Expected result to contain language suffix, got %s", result)
	}
}

func TestNameOpts_Apply(t *testing.T) {
	// Test with no options
	opts := NameOpts{}
	result := opts.Apply([]NameOpt{})
	
	if result.CommandName != "" {
		t.Errorf("Expected empty CommandName, got %s", result.CommandName)
	}
	
	// Test with single option
	opts = NameOpts{}
	result = opts.Apply([]NameOpt{CommandName("test-command")})
	
	if result.CommandName != "test-command" {
		t.Errorf("Expected CommandName 'test-command', got %s", result.CommandName)
	}
	
	// Test with multiple options
	opts = NameOpts{}
	result = opts.Apply([]NameOpt{
		CommandName("test-command"),
		RefPath(true),
	})
	
	if result.CommandName != "test-command" {
		t.Errorf("Expected CommandName 'test-command', got %s", result.CommandName)
	}
	if !result.RefPath {
		t.Error("Expected RefPath to be true")
	}
}

func TestCommandName(t *testing.T) {
	opt := CommandName("my-command")
	opts := NameOpts{}
	result := opt(opts)
	
	if result.CommandName != "my-command" {
		t.Errorf("Expected CommandName 'my-command', got %s", result.CommandName)
	}
}

func TestRefPath(t *testing.T) {
	// Test enabling RefPath
	opt := RefPath(true)
	opts := NameOpts{}
	result := opt(opts)
	
	if !result.RefPath {
		t.Error("Expected RefPath to be true")
	}
	
	// Test disabling RefPath
	opt = RefPath(false)
	opts = NameOpts{}
	result = opt(opts)
	
	if result.RefPath {
		t.Error("Expected RefPath to be false")
	}
}

func TestWebDocPath(t *testing.T) {
	lang := es_lang.Default
	
	// Test WebCategoryHome without refPath
	result := WebDocPath(false, WebCategoryHome, "index", lang)
	expected := WebDocPathRoot + "index.md"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	
	// Test WebCategoryCommand without refPath
	result = WebDocPath(false, WebCategoryCommand, "test-cmd", lang)
	if !strings.Contains(result, "commands/") {
		t.Errorf("Expected result to contain 'commands/', got %s", result)
	}
	if !strings.Contains(result, "test-cmd") {
		t.Errorf("Expected result to contain 'test-cmd', got %s", result)
	}
	
	// Test WebCategoryGuide without refPath
	result = WebDocPath(false, WebCategoryGuide, "test-guide", lang)
	if !strings.Contains(result, "guides/") {
		t.Errorf("Expected result to contain 'guides/', got %s", result)
	}
	
	// Test WebCategoryKnowledge without refPath
	result = WebDocPath(false, WebCategoryKnowledge, "test-knowledge", lang)
	if !strings.Contains(result, "knowledge/") {
		t.Errorf("Expected result to contain 'knowledge/', got %s", result)
	}
	
	// Test WebCategoryContributor without refPath
	result = WebDocPath(false, WebCategoryContributor, "test-contrib", lang)
	if !strings.Contains(result, "contributor/") {
		t.Errorf("Expected result to contain 'contributor/', got %s", result)
	}
	
	// Test with refPath enabled
	result = WebDocPath(true, WebCategoryHome, "index", lang)
	if !strings.Contains(result, "{{ site.baseurl }}/") {
		t.Errorf("Expected result to contain baseurl template, got %s", result)
	}
	if !strings.HasSuffix(result, ".html") {
		t.Errorf("Expected result to end with .html, got %s", result)
	}
	
	// Test with empty name
	result = WebDocPath(false, WebCategoryHome, "", lang)
	if strings.HasSuffix(result, ".md") {
		t.Errorf("Expected no .md suffix for empty name, got %s", result)
	}
	
	// Test with Japanese language
	jaLang := es_lang.Japanese
	result = WebDocPath(false, WebCategoryHome, "test", jaLang)
	if !strings.Contains(result, jaLang.String()+"/") {
		t.Errorf("Expected result to contain language path, got %s", result)
	}
}

func TestWebDocPath_InvalidCategory(t *testing.T) {
	lang := es_lang.Default
	
	// Test with invalid category - should panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid category")
		}
	}()
	
	WebDocPath(false, WebCategory(999), "test", lang)
}

func TestDocName_Repository(t *testing.T) {
	lang := es_lang.Default
	
	// Test DocRootReadme
	result := DocName(MediaRepository, DocRootReadme, lang)
	if !strings.Contains(result, "README") {
		t.Errorf("Expected result to contain 'README', got %s", result)
	}
	if !strings.HasSuffix(result, ".md") {
		t.Errorf("Expected result to end with .md, got %s", result)
	}
	
	// Test DocRootLicense
	result = DocName(MediaRepository, DocRootLicense, lang)
	if !strings.Contains(result, "LICENSE") {
		t.Errorf("Expected result to contain 'LICENSE', got %s", result)
	}
	
	// Test DocRootBuild
	result = DocName(MediaRepository, DocRootBuild, lang)
	if !strings.Contains(result, "BUILD") {
		t.Errorf("Expected result to contain 'BUILD', got %s", result)
	}
	
	// Test DocRootContributing
	result = DocName(MediaRepository, DocRootContributing, lang)
	if !strings.Contains(result, "CONTRIBUTING") {
		t.Errorf("Expected result to contain 'CONTRIBUTING', got %s", result)
	}
	
	// Test DocRootCodeOfConduct
	result = DocName(MediaRepository, DocRootCodeOfConduct, lang)
	if !strings.Contains(result, "CODE_OF_CONDUCT") {
		t.Errorf("Expected result to contain 'CODE_OF_CONDUCT', got %s", result)
	}
}

func TestDocName_WithLanguageSuffix(t *testing.T) {
	jaLang := es_lang.Japanese
	
	result := DocName(MediaRepository, DocRootReadme, jaLang)
	if !strings.Contains(result, jaLang.Suffix()) {
		t.Errorf("Expected result to contain language suffix, got %s", result)
	}
}

func TestDocName_WithOptions(t *testing.T) {
	lang := es_lang.Default
	
	// Test with CommandName option
	result := DocName(MediaRepository, DocRootReadme, lang, CommandName("test-command"))
	// The function should handle the option without error
	if result == "" {
		t.Error("Expected non-empty result")
	}
	
	// Test with RefPath option
	result = DocName(MediaRepository, DocRootReadme, lang, RefPath(true))
	if result == "" {
		t.Error("Expected non-empty result")
	}
}

func TestConstants(t *testing.T) {
	// Test that WebDocPathRoot is defined
	if WebDocPathRoot == "" {
		t.Error("WebDocPathRoot should not be empty")
	}
	
	if WebDocPathRoot != "docs/" {
		t.Errorf("Expected WebDocPathRoot to be 'docs/', got %s", WebDocPathRoot)
	}
	
	// Test that AllMedia contains expected values
	if len(AllMedia) == 0 {
		t.Error("AllMedia should not be empty")
	}
	
	// Check that expected media types are present
	foundRepo := false
	foundWeb := false
	for _, media := range AllMedia {
		if media == MediaRepository {
			foundRepo = true
		}
		if media == MediaWeb {
			foundWeb = true
		}
	}
	
	if !foundRepo {
		t.Error("AllMedia should contain MediaRepository")
	}
	if !foundWeb {
		t.Error("AllMedia should contain MediaWeb")
	}
}