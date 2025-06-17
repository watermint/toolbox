package da_griddata

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
)

func TestNewConsoleWriter(t *testing.T) {
	formatter := &PlainGridDataFormatter{}
	pw := NewCsvWriter()
	
	w := NewConsoleWriter(formatter, pw)
	if w == nil {
		t.Error("Expected non-nil console writer")
	}
	
	cw, ok := w.(*consoleWriter)
	if !ok {
		t.Error("Expected consoleWriter type")
	}
	if cw == nil {
		t.Error("Expected non-nil consoleWriter instance")
	}
	if cw.formatter != formatter {
		t.Error("Expected formatter to be set")
	}
	if cw.pw != pw {
		t.Error("Expected plain writer to be set")
	}
}

func TestConsoleWriter_Name(t *testing.T) {
	w := &consoleWriter{
		name: "test-console",
	}
	
	name := w.Name()
	if name != "test-console" {
		t.Errorf("Expected name 'test-console', got '%s'", name)
	}
}

func TestConsoleWriter_Open(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		w := &consoleWriter{}
		
		err := w.Open(ctl)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		
		if w.ctl == nil {
			t.Error("Expected control to be set after Open")
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestConsoleWriter_Close(t *testing.T) {
	w := &consoleWriter{}
	
	// Should not panic
	w.Close()
}

func TestConsoleWriter_Row(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		// Create a mock plain writer to capture output
		mpw := &mockPlainWriter{
			suffix: ".test",
			writes: make([]string, 0),
		}
		
		formatter := &PlainGridDataFormatter{}
		w := &consoleWriter{
			ctl:       ctl,
			name:      "test",
			formatter: formatter,
			pw:        mpw,
			row:       0,
		}
		
		// Test writing rows
		testData := [][]interface{}{
			{"row1", "col2", "col3"},
			{1, 2, 3},
			{"mixed", 123, true},
		}
		
		for i, row := range testData {
			w.Row(row)
			
			// Verify row index incremented
			if w.row != i+1 {
				t.Errorf("Expected row index %d, got %d", i+1, w.row)
			}
		}
		
		// Verify we wrote the correct number of rows
		if len(mpw.writes) != len(testData) {
			t.Errorf("Expected %d writes, got %d", len(testData), len(mpw.writes))
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestConsoleWriter_ConcurrentRow(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		mpw := &mockPlainWriter{
			suffix: ".test",
			writes: make([]string, 0),
		}
		
		formatter := &PlainGridDataFormatter{}
		w := &consoleWriter{
			ctl:       ctl,
			name:      "test",
			formatter: formatter,
			pw:        mpw,
			row:       0,
		}
		
		// Test concurrent writes
		done := make(chan bool)
		for i := 0; i < 10; i++ {
			go func(idx int) {
				w.Row([]interface{}{"concurrent", idx})
				done <- true
			}(i)
		}
		
		// Wait for all goroutines
		for i := 0; i < 10; i++ {
			<-done
		}
		
		// Should have 10 rows written
		if w.row != 10 {
			t.Errorf("Expected row count 10, got %d", w.row)
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

// Mock plain writer for testing
type mockPlainWriter struct {
	suffix string
	writes []string
}

func (m *mockPlainWriter) FileSuffix() string {
	return m.suffix
}

func (m *mockPlainWriter) WriteRow(l esl.Logger, w io.Writer, formatter GridDataFormatter, row int, column []interface{}) error {
	buf := &bytes.Buffer{}
	for i, col := range column {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%v", formatter.Format(col, i, row)))
	}
	m.writes = append(m.writes, buf.String())
	_, err := w.Write(buf.Bytes())
	return err
}