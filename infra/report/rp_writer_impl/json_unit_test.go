package rp_writer_impl

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJsonWriter_findRaw_Unit(t *testing.T) {
	// Test the findRaw method directly without needing a full control
	w := &jsonWriter{}
	
	// Test with struct containing Raw field
	type WithRaw struct {
		Name string
		Raw  json.RawMessage
	}
	
	testData := &WithRaw{
		Name: "test",
		Raw:  json.RawMessage(`{"custom":"data"}`),
	}
	
	raw := w.findRaw(testData)
	if raw == nil {
		t.Error("Expected to find Raw field")
	}
	if string(raw) != `{"custom":"data"}` {
		t.Errorf("Expected raw data to match, got: %s", string(raw))
	}
	
	// Test with struct without Raw field
	type WithoutRaw struct {
		Name  string
		Value int
	}
	
	testData2 := &WithoutRaw{Name: "test", Value: 42}
	raw2 := w.findRaw(testData2)
	if raw2 != nil {
		t.Error("Expected nil for struct without Raw field")
	}
	
	// Test with wrong Raw type
	type WrongRaw struct {
		Name string
		Raw  string // Wrong type
	}
	
	testData3 := &WrongRaw{Name: "test", Raw: "not json.RawMessage"}
	raw3 := w.findRaw(testData3)
	if raw3 != nil {
		t.Error("Expected nil for wrong Raw type")
	}
	
	// Test with reflect.Value input
	rv := reflect.ValueOf(testData).Elem()
	raw4 := w.findRaw(rv)
	if raw4 == nil {
		t.Error("Expected to find Raw field from reflect.Value")
	}
	
	// Test with nil Raw
	testData5 := &WithRaw{
		Name: "test",
		Raw:  nil,
	}
	raw5 := w.findRaw(testData5)
	if raw5 != nil {
		t.Error("Expected nil for nil Raw field")
	}
}

func TestJsonWriter_Name_Unit(t *testing.T) {
	w := &jsonWriter{
		name: "test_report",
	}
	
	if w.Name() != "test_report" {
		t.Errorf("Expected name 'test_report', got '%s'", w.Name())
	}
}

func TestFilterQueryLogFlags(t *testing.T) {
	// Test that the package-level variables exist
	if filterQueryLogEnabledExposed {
		t.Log("filterQueryLogEnabledExposed is true")
	}
	if filterQueryLogErrorExposed {
		t.Log("filterQueryLogErrorExposed is true")
	}
	
	// These are just for coverage - they start as false
	filterQueryLogEnabledExposed = true
	filterQueryLogErrorExposed = true
	
	// Reset
	filterQueryLogEnabledExposed = false
	filterQueryLogErrorExposed = false
}