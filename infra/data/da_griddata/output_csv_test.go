package da_griddata

import (
	"bytes"
	"strings"
	"testing"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func TestNewCsvWriter(t *testing.T) {
	w := NewCsvWriter()
	if w == nil {
		t.Error("Expected non-nil CSV writer")
	}
	
	cw, ok := w.(*csvWriter)
	if !ok {
		t.Error("Expected csvWriter type")
	}
	if cw == nil {
		t.Error("Expected non-nil csvWriter instance")
	}
}

func TestCsvWriter_FileSuffix(t *testing.T) {
	w := &csvWriter{}
	
	suffix := w.FileSuffix()
	if suffix != ".csv" {
		t.Errorf("Expected file suffix '.csv', got '%s'", suffix)
	}
}

func TestCsvWriter_WriteRow(t *testing.T) {
	w := &csvWriter{}
	l := esl.Default()
	formatter := &PlainGridDataFormatter{}
	
	tests := []struct {
		name     string
		row      int
		column   []interface{}
		expected string
	}{
		{
			name:     "string values",
			row:      0,
			column:   []interface{}{"hello", "world", "test"},
			expected: "hello,world,test",
		},
		{
			name:     "integer values",
			row:      1,
			column:   []interface{}{1, 2, 3},
			expected: "1,2,3",
		},
		{
			name:     "float values",
			row:      2,
			column:   []interface{}{1.5, 2.5, 3.5},
			expected: "1.500000,2.500000,3.500000",
		},
		{
			name:     "mixed values",
			row:      3,
			column:   []interface{}{"test", 123, 45.67, true},
			expected: "test,123,45.670000,true",
		},
		{
			name:     "values with quotes",
			row:      4,
			column:   []interface{}{"hello \"world\"", "test,value", "line\nbreak"},
			expected: "\"hello \"\"world\"\"\",\"test,value\",\"line\nbreak\"",
		},
		{
			name:     "empty values",
			row:      5,
			column:   []interface{}{"", "", ""},
			expected: ",,", // CSV with empty values
		},
		{
			name:     "nil values",
			row:      6,
			column:   []interface{}{nil, "test", nil},
			expected: "<nil>,test,<nil>",
		},
		{
			name:     "various integer types",
			row:      7,
			column:   []interface{}{int8(8), int16(16), int32(32), int64(64)},
			expected: "8,16,32,64",
		},
		{
			name:     "various unsigned types",
			row:      8,
			column:   []interface{}{uint(1), uint8(8), uint16(16), uint32(32), uint64(64)},
			expected: "1,8,16,32,64",
		},
		{
			name:     "various float types",
			row:      9,
			column:   []interface{}{float32(1.5), float64(2.5)},
			expected: "1.500000,2.500000",
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
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestCsvWriter_WriteRow_Error(t *testing.T) {
	w := &csvWriter{}
	l := esl.Default()
	formatter := &PlainGridDataFormatter{}
	
	// Test with writer that always fails
	errWriter := &errorWriter{}
	err := w.WriteRow(l, errWriter, formatter, 0, []interface{}{"test"})
	if err == nil {
		t.Error("Expected error when writer fails")
	}
}

// Mock writer that always returns an error
type errorWriter struct{}

func (e *errorWriter) Write(p []byte) (n int, err error) {
	return 0, bytes.ErrTooLarge
}