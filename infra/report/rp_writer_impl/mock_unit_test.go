package rp_writer_impl

import (
	"testing"
)

func TestNewMock(t *testing.T) {
	m := NewMock()
	if m == nil {
		t.Fatal("Expected non-nil mock")
	}
	
	if len(m.records) != 0 {
		t.Error("Expected empty records")
	}
	
	if m.isClosed {
		t.Error("Expected not closed initially")
	}
	
	if m.isOpened {
		t.Error("Expected not opened initially")
	}
}

func TestMock_Name(t *testing.T) {
	m := NewMock()
	if m.Name() != "" {
		t.Error("Expected empty name")
	}
}

func TestMock_IsOpened(t *testing.T) {
	m := NewMock()
	if m.IsOpened() {
		t.Error("Expected not opened initially")
	}
	
	// Simulate opening
	m.isOpened = true
	if !m.IsOpened() {
		t.Error("Expected opened after setting")
	}
}

func TestMock_IsClosed(t *testing.T) {
	m := NewMock()
	if m.IsClosed() {
		t.Error("Expected not closed initially")
	}
	
	// Close it
	m.Close()
	if !m.IsClosed() {
		t.Error("Expected closed after Close()")
	}
}

func TestMock_Records(t *testing.T) {
	m := NewMock()
	records := m.Records()
	if len(records) != 0 {
		t.Error("Expected no records initially")
	}
	
	// Add some records directly
	m.records = append(m.records, "test1", "test2")
	records = m.Records()
	if len(records) != 2 {
		t.Error("Expected 2 records")
	}
}

func TestMock_Row(t *testing.T) {
	m := NewMock()
	
	// Test panic when not opened
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when writing to non-opened mock")
		} else if r != ErrorMockTheWriterIsNotReady {
			t.Error("Expected ErrorMockTheWriterIsNotReady")
		}
	}()
	
	m.Row("test")
}

func TestMock_Row_AfterOpen(t *testing.T) {
	m := NewMock()
	
	// Open the mock
	err := m.Open(nil, nil)
	if err != nil {
		t.Error("Expected no error on open")
	}
	
	// Write some rows
	m.Row(&MockRecord{SKU: "ABC123", Quantity: 10})
	m.Row(&MockRecord{SKU: "DEF456", Quantity: 20})
	
	records := m.Records()
	if len(records) != 2 {
		t.Fatalf("Expected 2 records, got %d", len(records))
	}
	
	// Verify first record
	if rec1, ok := records[0].(*MockRecord); ok {
		if rec1.SKU != "ABC123" || rec1.Quantity != 10 {
			t.Error("First record mismatch")
		}
	} else {
		t.Error("Expected first record to be MockRecord")
	}
	
	// Close and try to write - should panic
	m.Close()
	
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when writing to closed mock")
		}
	}()
	
	m.Row(&MockRecord{SKU: "GHI789", Quantity: 30})
}

func TestMock_Open(t *testing.T) {
	m := NewMock()
	
	err := m.Open(nil, nil)
	if err != nil {
		t.Error("Expected no error on open")
	}
	
	if !m.isOpened {
		t.Error("Expected isOpened to be true after Open")
	}
}

func TestMock_Close(t *testing.T) {
	m := NewMock()
	
	m.Close()
	if !m.isClosed {
		t.Error("Expected isClosed to be true after Close")
	}
}

func TestMock_ConcurrentWrites(t *testing.T) {
	m := NewMock()
	m.Open(nil, nil)
	
	// Test concurrent writes
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			m.Row(&MockRecord{
				SKU:      "CONCURRENT",
				Quantity: n,
			})
			done <- true
		}(i)
	}
	
	// Wait for all writes
	for i := 0; i < 10; i++ {
		<-done
	}
	
	records := m.Records()
	if len(records) != 10 {
		t.Errorf("Expected 10 records from concurrent writes, got %d", len(records))
	}
}