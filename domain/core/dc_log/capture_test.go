package dc_log

import (
	"testing"
)

func TestErrorJobNotFound(t *testing.T) {
	// Test that ErrorJobNotFound is properly defined
	if ErrorJobNotFound == nil {
		t.Error("ErrorJobNotFound should not be nil")
	}
	
	if ErrorJobNotFound.Error() == "" {
		t.Error("ErrorJobNotFound should have a message")
	}
	
	expectedMessage := "job not found"
	if ErrorJobNotFound.Error() != expectedMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedMessage, ErrorJobNotFound.Error())
	}
}

func TestDefaultTimeIntervalSeconds(t *testing.T) {
	// Test that DefaultTimeIntervalSeconds has the expected value
	expectedValue := 3600
	if DefaultTimeIntervalSeconds != expectedValue {
		t.Errorf("Expected DefaultTimeIntervalSeconds to be %d, got %d", expectedValue, DefaultTimeIntervalSeconds)
	}
}

func TestConstants(t *testing.T) {
	// Test that constants are accessible and have reasonable values
	if DefaultTimeIntervalSeconds <= 0 {
		t.Error("DefaultTimeIntervalSeconds should be positive")
	}
	
	// 3600 seconds = 1 hour, which is a reasonable default
	if DefaultTimeIntervalSeconds != 3600 {
		t.Errorf("Expected DefaultTimeIntervalSeconds to be 3600, got %d", DefaultTimeIntervalSeconds)
	}
}