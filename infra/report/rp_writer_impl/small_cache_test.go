package rp_writer_impl

import (
	"testing"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_control"
)

type mockWriter struct {
	opened bool
	closed bool
	rows   []interface{}
}

func (m *mockWriter) Name() string {
	return "mock"
}

func (m *mockWriter) Row(r interface{}) {
	m.rows = append(m.rows, r)
}

func (m *mockWriter) Open(ctl interface{}, model interface{}, opts ...rp_model.ReportOpt) error {
	m.opened = true
	return nil
}

func (m *mockWriter) Close() {
	m.closed = true
}

func TestNewSmallCache(t *testing.T) {
	mock := &mockWriter{}
	writer := NewSmallCache("test", mock)
	
	if writer == nil {
		t.Error("Expected non-nil writer")
	}
	
	if writer.Name() != "test" {
		t.Errorf("Expected name 'test', got %s", writer.Name())
	}
}

func TestNewSmallCacheWithThreshold(t *testing.T) {
	mock := &mockWriter{}
	writer := NewSmallCacheWithThreshold("test", mock, 5)
	
	if writer == nil {
		t.Error("Expected non-nil writer")
	}
}

func TestSmallCache_BelowThreshold(t *testing.T) {
	mock := &mockWriter{}
	cache := NewSmallCacheWithThreshold("test", mock, 3)
	ctl := qt_control.WithFeature(false)
	
	// Open
	err := cache.Open(ctl, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// Add rows below threshold - should be cached
	cache.Row("row1")
	cache.Row("row2")
	
	// Mock writer should not have received rows yet
	if len(mock.rows) != 0 {
		t.Errorf("Expected 0 rows in mock, got %d", len(mock.rows))
	}
	
	// Close should flush cache
	cache.Close()
	
	if !mock.opened {
		t.Error("Expected mock to be opened")
	}
	
	if !mock.closed {
		t.Error("Expected mock to be closed")
	}
	
	if len(mock.rows) != 2 {
		t.Errorf("Expected 2 rows in mock, got %d", len(mock.rows))
	}
}

func TestSmallCache_ExceedThreshold(t *testing.T) {
	mock := &mockWriter{}
	cache := NewSmallCacheWithThreshold("test", mock, 2)
	ctl := qt_control.WithFeature(false)
	
	// Open
	err := cache.Open(ctl, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	
	// Add rows to exceed threshold
	cache.Row("row1")
	cache.Row("row2") // Still cached
	cache.Row("row3") // Should trigger flush and pass through
	
	// After exceeding threshold, cache should be flushed
	if len(mock.rows) != 3 {
		t.Errorf("Expected 3 rows in mock after threshold exceeded, got %d", len(mock.rows))
	}
	
	// Additional rows should pass through directly
	cache.Row("row4")
	
	if len(mock.rows) != 4 {
		t.Errorf("Expected 4 rows in mock, got %d", len(mock.rows))
	}
	
	cache.Close()
}

func TestSmallCache_RowWithoutOpen(t *testing.T) {
	mock := &mockWriter{}
	cache := NewSmallCacheWithThreshold("test", mock, 2)
	ctl := qt_control.WithFeature(false)
	
	// Set up cache with control and model but don't call Open explicitly
	sc := cache.(*smallCache)
	sc.ctl = ctl
	sc.model = "test"
	
	// Row should trigger open
	cache.Row("row1")
	
	if !mock.opened {
		t.Error("Expected mock to be opened after first row")
	}
}