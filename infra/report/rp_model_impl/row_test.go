package rp_model_impl

import (
	"testing"
)

func TestRowReport_Basic(t *testing.T) {
	// Test NewRowReport
	report := NewRowReport("test-report")
	if report.name != "test-report" {
		t.Error("Expected report name to be 'test-report'")
	}
	if report.Rows() != 0 {
		t.Error("Expected initial row count to be 0")
	}

	// Test Fork with nil ctl
	forked := report.Fork(nil)
	if forked == nil {
		t.Error("Expected forked report to be non-nil")
	}
	if forked.Rows() != 0 {
		t.Error("Expected forked report to have 0 rows")
	}

	// Test SetModel
	model := &TestModel{Name: "test", Value: 1}
	report.SetModel(model)
	if report.model == nil {
		t.Error("Expected model to be set")
	}

	// Test Spec
	spec := report.Spec()
	if spec.Name() != "test-report" {
		t.Error("Expected spec name to match report name")
	}

	// Test that we can call Close without opening
	report.Close() // Should not panic
}


// Test model for testing
type TestModel struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}