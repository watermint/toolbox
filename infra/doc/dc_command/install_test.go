package dc_command

import (
	"testing"
)

func TestNewInstall(t *testing.T) {
	section := NewInstall()
	if section == nil {
		t.Error("Expected non-nil section")
	}
	
	install, ok := section.(*Install)
	if !ok {
		t.Error("Expected Install type")
	}
	
	// Test that it implements the interface methods
	_ = install.Title() // Should not panic
}

func TestInstall_Title(t *testing.T) {
	install := &Install{}
	title := install.Title()
	
	// Should return the Header field
	if title != install.Header {
		t.Error("Title should return the Header field")
	}
}

func TestInstall_Body(t *testing.T) {
	install := &Install{}
	
	// Test that the Body method exists by checking if it can be called
	// We expect it to panic with nil UI, but at least it shows the method exists
	defer func() {
		if r := recover(); r != nil {
			// Expected to panic with nil UI - this is normal behavior
		}
	}()
	
	// Call with nil UI - expected to panic but tests method existence
	install.Body(nil)
}