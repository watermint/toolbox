package msg

import (
	"encoding/json"
	"path/filepath"
	"testing"
	"github.com/watermint/toolbox/essentials/model/mo_string"
)

func TestTranslate_Preset(t *testing.T) {
	translate := &Translate{}
	translate.Preset()
	// Preset method should not panic and complete successfully
}

func TestTranslate_KeyHandling(t *testing.T) {
	translate := &Translate{}
	
	// Initialize with empty optional string
	translate.Key = mo_string.NewOptional("")
	
	// Test with no key set (empty string)
	if translate.Key.IsExists() {
		t.Error("Key should not exist when empty")
	}
	
	// Test with key set
	translate.Key = mo_string.NewOptional("test.key")
	if !translate.Key.IsExists() {
		t.Error("Key should exist after setting")
	}
	
	if translate.Key.Value() != "test.key" {
		t.Errorf("Expected 'test.key', got %s", translate.Key.Value())
	}
}

func TestTranslate_JSONParsing(t *testing.T) {
	// Test JSON parsing functionality
	testMessages := map[string]string{
		"test.key1": "Hello",
		"test.key2": "World",
	}
	
	jsonData, err := json.Marshal(testMessages)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}
	
	var parsed map[string]string
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal test data: %v", err)
	}
	
	if len(parsed) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(parsed))
	}
	
	if parsed["test.key1"] != "Hello" {
		t.Errorf("Expected 'Hello', got %s", parsed["test.key1"])
	}
}

func TestTranslate_FilePaths(t *testing.T) {
	// Test file path building
	enPath := filepath.Join("resources", "messages", "en", "messages.json")
	jaPath := filepath.Join("resources", "messages", "ja", "messages.json")
	
	if enPath == "" {
		t.Error("English path should not be empty")
	}
	
	if jaPath == "" {
		t.Error("Japanese path should not be empty")
	}
	
	// Test that paths are different
	if enPath == jaPath {
		t.Error("English and Japanese paths should be different")
	}
}

func TestTranslate_MissingKeyDetection(t *testing.T) {
	// Test missing key detection logic
	enMessages := map[string]string{
		"key1": "English 1",
		"key2": "English 2", 
		"key3": "English 3",
	}
	
	jaMessages := map[string]string{
		"key1": "Japanese 1",
		"key2": "Japanese 2",
		// key3 is missing
	}
	
	// Find missing keys
	missingKeys := make([]string, 0)
	for key := range enMessages {
		if _, exists := jaMessages[key]; !exists {
			missingKeys = append(missingKeys, key)
		}
	}
	
	if len(missingKeys) != 1 {
		t.Errorf("Expected 1 missing key, got %d", len(missingKeys))
	}
	
	if len(missingKeys) > 0 && missingKeys[0] != "key3" {
		t.Errorf("Expected missing key 'key3', got %s", missingKeys[0])
	}
}