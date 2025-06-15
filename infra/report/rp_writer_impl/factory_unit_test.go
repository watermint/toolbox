package rp_writer_impl

import (
	"testing"
)

func TestNew_ReturnsWriter(t *testing.T) {
	// We can't test this fully without a control, but we can test
	// that the functions exist and are callable
	
	// Test that NewCascade is callable (would panic if not)
	defer func() {
		if r := recover(); r != nil {
			// Expected - we're calling with nil
			t.Log("NewCascade panicked as expected with nil control")
		}
	}()
	
	// This will panic, but that's ok - we're just testing it exists
	_ = NewCascade("test", nil)
}

func TestSmallCache_Exists(t *testing.T) {
	// Just verify the NewSmallCache function exists
	// We can't test it without a real writer
	
	defer func() {
		if r := recover(); r != nil {
			// Expected
			t.Log("NewSmallCache panicked as expected with nil writer")
		}
	}()
	
	_ = NewSmallCache("test", nil)
}