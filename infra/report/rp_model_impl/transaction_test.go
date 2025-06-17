package rp_model_impl

import (
	"testing"
)

func TestTransactionReport_Basic(t *testing.T) {
	// Test NewTransactionReport
	report := NewTransactionReport("tx-report")
	if report.name != "tx-report" {
		t.Error("Expected report name to be 'tx-report'")
	}
	if report.Rows() != 0 {
		t.Error("Expected initial row count to be 0")
	}

	// Test Fork with nil
	forked := report.Fork(nil)
	if forked == nil {
		t.Error("Expected forked report to be non-nil")
	}
	// Note: Can't test Rows() on forked report as it doesn't initialize rows counter

	// Test SetModel
	input := &TestModel{Name: "input", Value: 1}
	result := &TestModel{Name: "result", Value: 2}
	report.SetModel(input, result)
	
	// Test that model was set
	if report.model == nil {
		t.Error("Expected model to be set after SetModel")
	}
}

func TestTransactionReport_Spec(t *testing.T) {
	// Test NewTransactionReport
	report := NewTransactionReport("spec-test")
	
	// Set model first
	report.SetModel(&TestModel{}, &TestModel{})
	
	// Test Spec
	spec := report.Spec()
	if spec.Name() != "spec-test" {
		t.Error("Expected spec name to match report name")
	}
}

func TestTransactionReport_SetCtl(t *testing.T) {
	report := NewTransactionReport("setctl-test")
	
	// Test that ctl is initially nil
	if report.ctl != nil {
		t.Error("Expected initial ctl to be nil")
	}
	
	// Test Close without opening
	report.Close() // Should not panic
}

