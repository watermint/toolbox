package mo_int

import (
	"testing"
)

func TestRangeInt_Value64(t *testing.T) {
	ri := NewRange()
	
	// Test with various values
	tests := []int64{
		0,
		-1,
		1,
		42,
		-42,
		1234567890,
		-1234567890,
	}
	
	for _, val := range tests {
		ri.SetValue(val)
		if ri.Value64() != val {
			t.Errorf("Value64() = %d, want %d", ri.Value64(), val)
		}
		
		// Also verify Value() returns truncated int
		if ri.Value() != int(val) {
			t.Errorf("Value() = %d, want %d", ri.Value(), int(val))
		}
	}
}

func TestRangeInt_SetRangeWithReversedValues(t *testing.T) {
	ri := NewRange()
	
	// Test that SetRange handles reversed min/max correctly
	ri.SetRange(10, 5, 7)
	
	min, max := ri.Range()
	if min != 5 || max != 10 {
		t.Errorf("Range() = (%d, %d), want (5, 10)", min, max)
	}
	
	if ri.Value64() != 7 {
		t.Errorf("Value64() = %d, want 7", ri.Value64())
	}
}

func TestRangeInt_BoundaryValues(t *testing.T) {
	ri := NewRange()
	
	// Test with minimum value at boundary
	ri.SetRange(0, 100, 0)
	if !ri.IsValid() {
		t.Error("Expected valid at minimum boundary")
	}
	
	// Test with maximum value at boundary
	ri.SetRange(0, 100, 100)
	if !ri.IsValid() {
		t.Error("Expected valid at maximum boundary")
	}
	
	// Test with value below range
	ri.SetValue(-1)
	if ri.IsValid() {
		t.Error("Expected invalid below range")
	}
	
	// Test with value above range
	ri.SetValue(101)
	if ri.IsValid() {
		t.Error("Expected invalid above range")
	}
}

func TestRangeInt_NegativeRange(t *testing.T) {
	ri := NewRange()
	
	// Test with negative range
	ri.SetRange(-100, -10, -50)
	
	min, max := ri.Range()
	if min != -100 || max != -10 {
		t.Errorf("Range() = (%d, %d), want (-100, -10)", min, max)
	}
	
	if ri.Value64() != -50 {
		t.Errorf("Value64() = %d, want -50", ri.Value64())
	}
	
	if !ri.IsValid() {
		t.Error("Expected valid in negative range")
	}
	
	// Test invalid values
	ri.SetValue(-101)
	if ri.IsValid() {
		t.Error("Expected invalid below negative range")
	}
	
	ri.SetValue(-9)
	if ri.IsValid() {
		t.Error("Expected invalid above negative range")
	}
}

func TestRangeInt_ZeroRange(t *testing.T) {
	ri := NewRange()
	
	// Test with same min and max
	ri.SetRange(42, 42, 42)
	
	min, max := ri.Range()
	if min != 42 || max != 42 {
		t.Errorf("Range() = (%d, %d), want (42, 42)", min, max)
	}
	
	if !ri.IsValid() {
		t.Error("Expected valid when value equals min/max")
	}
	
	// Any other value should be invalid
	ri.SetValue(41)
	if ri.IsValid() {
		t.Error("Expected invalid when value differs from single allowed value")
	}
	
	ri.SetValue(43)
	if ri.IsValid() {
		t.Error("Expected invalid when value differs from single allowed value")
	}
}