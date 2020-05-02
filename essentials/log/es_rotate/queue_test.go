package es_rotate

import "testing"

func TestStartupShutdown(t *testing.T) {
	// Should not panic
	Shutdown()

	Startup()
	Shutdown()

	// Should able to restart
	Startup()
	Shutdown()
}
