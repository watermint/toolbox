package ut_compare

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"testing"
	"time"
)

func TestClone(t *testing.T) {
	// Test with various time values
	now := time.Now()
	loc, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name string
		time time.Time
	}{
		{"current time", now},
		{"zero time", time.Time{}},
		{"specific date", time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)},
		{"different timezone", time.Date(2023, 6, 15, 10, 0, 0, 0, loc)},
		{"unix epoch", time.Unix(0, 0)},
		{"with nanoseconds", time.Date(2023, 1, 1, 0, 0, 0, 999999999, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cloned := Clone(tt.time)

			// Verify all fields are equal
			if !cloned.Equal(tt.time) {
				t.Errorf("Clone() time not equal: got %v, want %v", cloned, tt.time)
			}

			// Verify it's a different instance (pointer comparison)
			if &cloned == &tt.time {
				t.Error("Clone() returned same instance")
			}

			// Verify all components match
			if cloned.Year() != tt.time.Year() ||
				cloned.Month() != tt.time.Month() ||
				cloned.Day() != tt.time.Day() ||
				cloned.Hour() != tt.time.Hour() ||
				cloned.Minute() != tt.time.Minute() ||
				cloned.Second() != tt.time.Second() ||
				cloned.Nanosecond() != tt.time.Nanosecond() ||
				cloned.Location().String() != tt.time.Location().String() {
				t.Error("Clone() components don't match")
			}
		})
	}
}

func TestClonePtr(t *testing.T) {
	// Test with nil
	var nilTime *time.Time
	cloned := ClonePtr(nilTime)
	if cloned != nil {
		t.Error("ClonePtr(nil) should return nil")
	}

	// Test with non-nil time
	now := time.Now()
	cloned = ClonePtr(&now)

	if cloned == nil {
		t.Fatal("ClonePtr() returned nil for non-nil input")
	}

	if !cloned.Equal(now) {
		t.Errorf("ClonePtr() time not equal: got %v, want %v", *cloned, now)
	}

	// Verify it's a different pointer
	if cloned == &now {
		t.Error("ClonePtr() returned same pointer")
	}

	// Test with zero time
	zeroTime := time.Time{}
	clonedZero := ClonePtr(&zeroTime)
	if clonedZero == nil {
		t.Fatal("ClonePtr() returned nil for zero time")
	}
	if !clonedZero.IsZero() {
		t.Error("ClonePtr() should preserve zero time")
	}
}

func TestEarliest_EdgeCases(t *testing.T) {
	// Test with single element
	single := time.Now()
	result := Earliest(single)
	if !result.Equal(single) {
		t.Error("Earliest() with single element should return that element")
	}

	// Test all elements are the same
	same := time.Now()
	result = Earliest(same, same, same)
	if !result.Equal(same) {
		t.Error("Earliest() with same elements should return that time")
	}
}

func TestLatest_EdgeCases(t *testing.T) {
	// Test with single element
	single := time.Now()
	result := Latest(single)
	if !result.Equal(single) {
		t.Error("Latest() with single element should return that element")
	}

	// Test all elements are the same
	same := time.Now()
	result = Latest(same, same, same)
	if !result.Equal(same) {
		t.Error("Latest() with same elements should return that time")
	}
}

func TestIsBetweenOptional_AllZero(t *testing.T) {
	// Test the case where both a and b are zero
	now := time.Now()
	zeroOpt := mo_time.NewOptional(time.Time{})

	// When both are zero, should always return true
	if !IsBetweenOptional(now, zeroOpt, zeroOpt) {
		t.Error("IsBetweenOptional() should return true when both bounds are zero")
	}

	// Test with different time values when both bounds are zero
	future := now.Add(100 * time.Hour)
	past := now.Add(-100 * time.Hour)

	if !IsBetweenOptional(future, zeroOpt, zeroOpt) {
		t.Error("IsBetweenOptional() should return true for any time when both bounds are zero")
	}

	if !IsBetweenOptional(past, zeroOpt, zeroOpt) {
		t.Error("IsBetweenOptional() should return true for any time when both bounds are zero")
	}
}
