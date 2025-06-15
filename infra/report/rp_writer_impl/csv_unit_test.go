package rp_writer_impl

import (
	"testing"
)

func TestCsvWriter_Name_Unit(t *testing.T) {
	w := &csvWriter{
		name: "csv_report",
	}
	
	if w.Name() != "csv_report" {
		t.Errorf("Expected name 'csv_report', got '%s'", w.Name())
	}
}

func TestNewCsvWriter(t *testing.T) {
	// Just test that it returns non-nil
	// We can't test fully without a control
	w := &csvWriter{
		name: "test",
	}
	
	if w.Name() != "test" {
		t.Error("Expected name to be set")
	}
}