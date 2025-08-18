package nw_congestion

import (
	"sync"
	"testing"
	"time"
)

func TestGetReportInterval(t *testing.T) {
	// Test in non-production mode
	interval := getReportInterval()
	if interval != 10*time.Second {
		t.Errorf("Expected 10 seconds in non-production, got %v", interval)
	}
}

func TestWindowSettings(t *testing.T) {
	// Save original values
	origMax := maxCongestionWindow
	origInit := initCongestionWindow

	// Test SetMaxCongestionWindow with hard limit
	SetMaxCongestionWindow(20, false)
	if CurrentMaxCongestionWindow() != hardLimitCongestionWindow {
		t.Errorf("Expected max window to be limited to %d, got %d", hardLimitCongestionWindow, CurrentMaxCongestionWindow())
	}

	// Test SetMaxCongestionWindow ignoring hard limit
	SetMaxCongestionWindow(20, true)
	if CurrentMaxCongestionWindow() != 20 {
		t.Errorf("Expected max window to be 20, got %d", CurrentMaxCongestionWindow())
	}

	// Test SetInitCongestionWindow
	SetInitCongestionWindow(2)
	if CurrentInitCongestionWindow() != 2 {
		t.Errorf("Expected init window to be 2, got %d", CurrentInitCongestionWindow())
	}

	// Test CurrentMinCongestionWindow
	if CurrentMinCongestionWindow() != minCongestionWindow {
		t.Errorf("Expected min window to be %d, got %d", minCongestionWindow, CurrentMinCongestionWindow())
	}

	// Restore original values
	maxCongestionWindow = origMax
	initCongestionWindow = origInit
}

func TestCcImpl_Key(t *testing.T) {
	cc := NewControl().(*ccImpl)
	key := cc.key("hash", "endpoint")
	if key != "hash-endpoint" {
		t.Errorf("Expected key to be 'hash-endpoint', got %s", key)
	}
}

func TestBasicFlow(t *testing.T) {
	// Save original value
	origMax := maxCongestionWindow

	// Set a reasonable window for testing
	SetMaxCongestionWindow(4, true)

	cc := NewControl()
	hash := "test-hash"
	endpoint := "test-endpoint"

	// Test basic flow: Start -> EndSuccess
	cc.Start(hash, endpoint)
	cc.EndSuccess(hash, endpoint)

	// Test basic flow: Start -> EndTransportError
	cc.Start(hash, endpoint)
	cc.EndTransportError(hash, endpoint)

	// Test basic flow: Start -> EndRateLimit
	cc.Start(hash, endpoint)
	cc.EndRateLimit(hash, endpoint, time.Now().Add(5*time.Second))

	// Restore original value
	maxCongestionWindow = origMax
}

func TestConcurrentOperations(t *testing.T) {
	// Save original value
	origMax := maxCongestionWindow

	// Set a small window to test concurrency control
	SetMaxCongestionWindow(2, true)

	cc := NewControl()

	// Run concurrent operations
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			hash := "hash"
			endpoint := "endpoint"

			cc.Start(hash, endpoint)
			time.Sleep(50 * time.Millisecond)

			// Mix different end conditions
			switch i % 3 {
			case 0:
				cc.EndSuccess(hash, endpoint)
			case 1:
				cc.EndTransportError(hash, endpoint)
			case 2:
				cc.EndRateLimit(hash, endpoint, time.Now().Add(1*time.Minute))
			}
		}(i)
	}
	wg.Wait()

	// Restore original value
	maxCongestionWindow = origMax
}

func TestPackageLevelFunctions(t *testing.T) {
	hash := "test-hash"
	endpoint := "test-endpoint"

	// These should not panic
	Start(hash, endpoint)
	EndSuccess(hash, endpoint)

	Start(hash, endpoint)
	EndTransportError(hash, endpoint)

	Start(hash, endpoint)
	EndRateLimit(hash, endpoint, time.Now().Add(1*time.Minute))
}

func TestSignificantWait(t *testing.T) {
	cc := NewControl().(*ccImpl)

	// Test wait time less than threshold
	notSignificant := time.Now().Add(30 * time.Second)
	if cc.isSignificantWait(notSignificant) {
		t.Error("Expected wait to not be significant")
	}

	// Test wait time more than threshold
	significant := time.Now().Add(2 * time.Minute)
	if !cc.isSignificantWait(significant) {
		t.Error("Expected wait to be significant")
	}
}

func TestRateLimitWithSignificantWait(t *testing.T) {
	// Save original values
	origMax := maxCongestionWindow

	// Set window for testing
	SetMaxCongestionWindow(4, true)

	cc := NewControl()
	hash := "test-hash"
	endpoint := "test-endpoint"

	// Test with significant wait (should set window to minimum)
	cc.Start(hash, endpoint)
	cc.EndRateLimit(hash, endpoint, time.Now().Add(2*time.Minute))

	// Test with non-significant wait (should decrease window)
	cc.Start(hash, endpoint)
	cc.EndRateLimit(hash, endpoint, time.Now().Add(30*time.Second))

	// Restore original value
	maxCongestionWindow = origMax
}

func TestMonitorGoroutine(t *testing.T) {
	// Create a new control that will start monitor
	cc := NewControl()

	// Trigger monitor start by calling Start
	cc.Start("test", "test")

	// Give monitor time to start
	time.Sleep(100 * time.Millisecond)

	// Clean up by ending the operation
	cc.EndSuccess("test", "test")

	// The monitor goroutine will continue running, but that's expected
}
