package rp_model_impl

import (
	"testing"

	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func TestNewSpec(t *testing.T) {
	// Test with simple model
	model := &TestModel{Name: "test", Value: 123}
	spec := newSpec("test-spec", model, nil)

	if spec.Name() != "test-spec" {
		t.Error("Expected spec name to be 'test-spec'")
	}

	if spec.Model() != model {
		t.Error("Expected spec model to match input model")
	}

	// Test columns
	cols := spec.Columns()
	if len(cols) == 0 {
		t.Error("Expected spec to have columns")
	}

	// Test with TransactionRow model
	txRow := &rp_model.TransactionRow{
		Input:  &TestModel{Name: "input", Value: 1},
		Result: &TestModel{Name: "result", Value: 2},
	}
	txSpec := newSpec("tx-spec", txRow, nil)

	txCols := txSpec.Columns()
	// Should have status, reason, and columns from input/result
	if len(txCols) < 2 {
		t.Error("Expected transaction spec to have at least status and reason columns")
	}

	// Check that status and reason are in columns
	hasStatus := false
	hasReason := false
	for _, col := range txCols {
		if col == "status" {
			hasStatus = true
		}
		if col == "reason" {
			hasReason = true
		}
	}
	if !hasStatus || !hasReason {
		t.Error("Expected transaction spec to have status and reason columns")
	}
}

func TestColumnSpec_Methods(t *testing.T) {
	model := &TestModel{Name: "test", Value: 456}
	spec := &ColumnSpec{
		name:  "col-spec",
		model: model,
		opts:  []rp_model.ReportOpt{},
		cols:  []string{"col1", "col2"},
		colDesc: map[string]app_msg.Message{
			"col1": app_msg.Raw("Column 1"),
			"col2": app_msg.Raw("Column 2"),
		},
	}

	// Test Name
	if spec.Name() != "col-spec" {
		t.Error("Expected name to be 'col-spec'")
	}

	// Test Model
	if spec.Model() != model {
		t.Error("Expected model to match")
	}

	// Test Columns
	cols := spec.Columns()
	if len(cols) != 2 {
		t.Error("Expected 2 columns")
	}

	// Test ColumnDesc
	desc1 := spec.ColumnDesc("col1")
	if desc1 == nil {
		t.Error("Expected column description for col1")
	}

	// Test ColumnDesc for unknown column
	descUnknown := spec.ColumnDesc("unknown")
	if descUnknown == nil {
		t.Error("Expected raw message for unknown column")
	}

	// Test Options
	opts := spec.Options()
	if len(opts) != 0 {
		t.Error("Expected empty options")
	}

	// Test Desc
	desc := spec.Desc()
	if desc == nil {
		t.Error("Expected description message")
	}
}

func TestColumnSpec_NilModel(t *testing.T) {
	// Test that Desc panics with nil model
	spec := &ColumnSpec{
		name:  "nil-spec",
		model: nil,
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for nil model")
		}
	}()

	// This should panic
	_ = spec.Desc()
}

func TestSpec_WithReportOpts(t *testing.T) {
	model := &TestModel{Name: "test", Value: 789}

	// Test with hidden columns
	opts := []rp_model.ReportOpt{
		func(o *rp_model.ReportOpts) *rp_model.ReportOpts {
			if o.HiddenColumns == nil {
				o.HiddenColumns = make(map[string]bool)
			}
			o.HiddenColumns["name"] = true
			return o
		},
	}

	spec := newSpec("opts-spec", model, opts)

	// Options should be preserved
	if len(spec.Options()) != 1 {
		t.Error("Expected options to be preserved")
	}
}
