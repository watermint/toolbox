package diag

import "testing"

func TestNewDiagnosticsInfra(t *testing.T) {
	d := NewDiagnosticsInfra()
	if d.NumCpu < 1 {
		t.Errorf("No CPU found: %d", d.NumCpu)
	}
}

func TestNewDiagnosticsRuntime(t *testing.T) {
	d := NewDiagnosticsRuntime()
	if d.PID < 1 {
		t.Errorf("Invalid PID: %d", d.PID)
	}
}
