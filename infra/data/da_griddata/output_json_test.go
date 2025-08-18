package da_griddata

import (
    "bytes"
    "encoding/json"
    "strings"
    "testing"

    "github.com/watermint/toolbox/essentials/log/esl"
)

func TestNewJsonWriter(t *testing.T) {
	w := NewJsonWriter()
	if w == nil {
		t.Error("Expected non-nil JSON writer")
	}

	jw, ok := w.(*jsonWriter)
	if !ok {
		t.Error("Expected jsonWriter type")
	}
	if jw == nil {
		t.Error("Expected non-nil jsonWriter instance")
	}
}

func TestJsonWriter_FileSuffix(t *testing.T) {
	w := &jsonWriter{}

	suffix := w.FileSuffix()
	if suffix != ".json" {
		t.Errorf("Expected file suffix '.json', got '%s'", suffix)
	}
}

func TestJsonWriter_WriteRow(t *testing.T) {
	w := &jsonWriter{}
	l := esl.Default()
	formatter := &PlainGridDataFormatter{}

	tests := []struct {
		name   string
		row    int
		column []interface{}
	}{
		{
			name:   "string values",
			row:    0,
			column: []interface{}{"hello", "world", "test"},
		},
		{
			name:   "integer values",
			row:    1,
			column: []interface{}{1, 2, 3},
		},
		{
			name:   "float values",
			row:    2,
			column: []interface{}{1.5, 2.5, 3.5},
		},
		{
			name:   "mixed values",
			row:    3,
			column: []interface{}{"test", 123, 45.67, true},
		},
		{
			name:   "nil values",
			row:    4,
			column: []interface{}{nil, "test", nil},
		},
		{
			name:   "empty array",
			row:    5,
			column: []interface{}{},
		},
		{
			name:   "complex types",
			row:    6,
			column: []interface{}{map[string]interface{}{"key": "value"}, []int{1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := w.WriteRow(l, buf, formatter, tt.row, tt.column)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			result := strings.TrimSpace(buf.String())

			// Verify it's valid JSON
			var parsed []interface{}
			err = json.Unmarshal([]byte(result), &parsed)
			if err != nil {
				t.Errorf("Invalid JSON output: %v", err)
			}

			// Verify the number of elements
			if len(parsed) != len(tt.column) {
				t.Errorf("Expected %d elements, got %d", len(tt.column), len(parsed))
			}
		})
	}
}

func TestJsonWriter_WriteRow_Error(t *testing.T) {
	w := &jsonWriter{}
	l := esl.Default()
	formatter := &PlainGridDataFormatter{}

	// Test with writer that always fails
	errWriter := &errorJsonWriter{}
	err := w.WriteRow(l, errWriter, formatter, 0, []interface{}{"test"})
	if err == nil {
		t.Error("Expected error when writer fails")
	}
}

// Mock formatter that returns unmarshalable value
type badFormatter struct{}

func (b badFormatter) Format(data interface{}, col int, row int) interface{} {
	// Return a channel which cannot be marshaled to JSON
	return make(chan int)
}

func TestJsonWriter_WriteRow_MarshalError(t *testing.T) {
	w := &jsonWriter{}
	l := esl.Default()
	formatter := &badFormatter{}

	buf := &bytes.Buffer{}
	err := w.WriteRow(l, buf, formatter, 0, []interface{}{"test"})
	if err == nil {
		t.Error("Expected error when marshaling fails")
	}
}

// Mock writer that always returns an error
type errorJsonWriter struct{}

func (e *errorJsonWriter) Write(p []byte) (n int, err error) {
	return 0, bytes.ErrTooLarge
}
