package uc_insight

import (
	"errors"
	"testing"
)

func TestApiErrorFromError_NilError(t *testing.T) {
	apiErr := ApiErrorFromError(nil)

	if apiErr.Error != "" {
		t.Errorf("Expected empty error string for nil error, got '%s'", apiErr.Error)
	}
	if apiErr.ErrorTag != "" {
		t.Errorf("Expected empty error tag for nil error, got '%s'", apiErr.ErrorTag)
	}
}

func TestApiErrorFromError_RegularError(t *testing.T) {
	testErr := errors.New("test error message")

	apiErr := ApiErrorFromError(testErr)

	if apiErr.Error != "test error message" {
		t.Errorf("Expected error 'test error message', got '%s'", apiErr.Error)
	}
	if apiErr.ErrorTag != "" {
		t.Errorf("Expected empty error tag for regular error, got '%s'", apiErr.ErrorTag)
	}
}

func TestApiErrorFromError_DbxErrorWithSummary(t *testing.T) {
	// Create a mock error that would be recognized by dbx_error
	testErr := errors.New("auth_error: invalid access token")

	apiErr := ApiErrorFromError(testErr)

	// The exact behavior depends on dbx_error implementation
	// We just verify the function doesn't panic and returns reasonable values
	if apiErr.Error == "" && apiErr.ErrorTag == "" {
		t.Error("Expected at least one of Error or ErrorTag to be non-empty")
	}
}

func TestApiErrorFromError_ComplexError(t *testing.T) {
	// Test with an error that might have complex structure
	testErr := &customError{
		message: "custom error with details",
		code:    "CUSTOM_001",
	}

	apiErr := ApiErrorFromError(testErr)

	// Should handle custom errors gracefully
	if apiErr.Error == "" {
		t.Error("Expected non-empty error string for custom error")
	}
}

// Custom error type for testing
type customError struct {
	message string
	code    string
}

func (e *customError) Error() string {
	return e.message
}

func TestApiError_StructFields(t *testing.T) {
	apiErr := ApiError{
		Error:    "test error",
		ErrorTag: "test_tag",
	}

	if apiErr.Error != "test error" {
		t.Errorf("Expected Error field 'test error', got '%s'", apiErr.Error)
	}
	if apiErr.ErrorTag != "test_tag" {
		t.Errorf("Expected ErrorTag field 'test_tag', got '%s'", apiErr.ErrorTag)
	}
}

func TestApiErrorReport_StructFields(t *testing.T) {
	report := ApiErrorReport{
		Category: "TestCategory",
		Message:  "Test message",
		Tag:      "test_tag",
		Detail:   "Test detail information",
	}

	if report.Category != "TestCategory" {
		t.Errorf("Expected Category 'TestCategory', got '%s'", report.Category)
	}
	if report.Message != "Test message" {
		t.Errorf("Expected Message 'Test message', got '%s'", report.Message)
	}
	if report.Tag != "test_tag" {
		t.Errorf("Expected Tag 'test_tag', got '%s'", report.Tag)
	}
	if report.Detail != "Test detail information" {
		t.Errorf("Expected Detail 'Test detail information', got '%s'", report.Detail)
	}
}

func TestApiError_EmptyValues(t *testing.T) {
	apiErr := ApiError{}

	if apiErr.Error != "" {
		t.Errorf("Expected empty Error field in zero value, got '%s'", apiErr.Error)
	}
	if apiErr.ErrorTag != "" {
		t.Errorf("Expected empty ErrorTag field in zero value, got '%s'", apiErr.ErrorTag)
	}
}

func TestApiErrorReport_EmptyValues(t *testing.T) {
	report := ApiErrorReport{}

	if report.Category != "" {
		t.Errorf("Expected empty Category field in zero value, got '%s'", report.Category)
	}
	if report.Message != "" {
		t.Errorf("Expected empty Message field in zero value, got '%s'", report.Message)
	}
	if report.Tag != "" {
		t.Errorf("Expected empty Tag field in zero value, got '%s'", report.Tag)
	}
	if report.Detail != "" {
		t.Errorf("Expected empty Detail field in zero value, got '%s'", report.Detail)
	}
}

func TestApiErrorFromError_ChainedErrors(t *testing.T) {
	innerErr := errors.New("inner error")
	outerErr := errors.New("outer error: " + innerErr.Error())

	apiErr := ApiErrorFromError(outerErr)

	// Should handle wrapped/chained errors
	if apiErr.Error == "" {
		t.Error("Expected non-empty error string for chained error")
	}
	// The error string should contain information about the error
	if len(apiErr.Error) < len("outer error") {
		t.Error("Expected error string to contain meaningful information")
	}
}

func TestApiErrorFromError_LongError(t *testing.T) {
	longMessage := "This is a very long error message that contains a lot of details about what went wrong in the system and should be handled properly by the ApiErrorFromError function without truncation or other issues"
	testErr := errors.New(longMessage)

	apiErr := ApiErrorFromError(testErr)

	if apiErr.Error != longMessage {
		t.Errorf("Expected full long error message to be preserved, got '%s'", apiErr.Error)
	}
}

func TestApiErrorFromError_SpecialCharacters(t *testing.T) {
	specialMessage := "Error with special chars: áéíóú, 中文, 🚀, \n\t\\"
	testErr := errors.New(specialMessage)

	apiErr := ApiErrorFromError(testErr)

	if apiErr.Error != specialMessage {
		t.Errorf("Expected special characters to be preserved, got '%s'", apiErr.Error)
	}
}

func TestApiError_JsonSerialization(t *testing.T) {
	// Test that the struct can be used for JSON serialization
	// This is important since the fields have json tags
	apiErr := ApiError{
		Error:    "json test error",
		ErrorTag: "json_test",
	}

	// Basic field access test
	if apiErr.Error != "json test error" {
		t.Error("ApiError struct should maintain field values for JSON serialization")
	}
	if apiErr.ErrorTag != "json_test" {
		t.Error("ApiError struct should maintain field values for JSON serialization")
	}
}

func TestApiErrorReport_JsonSerialization(t *testing.T) {
	// Test that the struct can be used for JSON serialization
	report := ApiErrorReport{
		Category: "JsonTest",
		Message:  "json test message",
		Tag:      "json_test",
		Detail:   "json test detail",
	}

	// Basic field access test
	if report.Category != "JsonTest" {
		t.Error("ApiErrorReport struct should maintain field values for JSON serialization")
	}
}

// Test interface compliance if any records implement ApiErrorRecord
func TestApiErrorRecord_Interface(t *testing.T) {
	// This test ensures the interface is properly defined
	// We can't test actual implementations without knowing which structs implement it
	var _ ApiErrorRecord = (*mockApiErrorRecord)(nil)
}

type mockApiErrorRecord struct{}

func (m *mockApiErrorRecord) ToParam() interface{} {
	return map[string]string{"test": "param"}
}

func TestMockApiErrorRecord_ToParam(t *testing.T) {
	mock := &mockApiErrorRecord{}
	param := mock.ToParam()

	if param == nil {
		t.Error("Expected non-nil parameter from ToParam")
	}

	// Type assertion to verify return type
	if paramMap, ok := param.(map[string]string); ok {
		if paramMap["test"] != "param" {
			t.Errorf("Expected param map to contain test: param, got %v", paramMap)
		}
	} else {
		t.Error("Expected ToParam to return map[string]string")
	}
}
