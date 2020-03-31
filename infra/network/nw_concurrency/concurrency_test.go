package nw_concurrency

import "testing"

func TestSetConcurrency(t *testing.T) {
	// Should accept negative
	SetConcurrency(-1)
	Start()
	End()

	// Should accept zero
	SetConcurrency(0)
	Start()
	End()

	// Should accept positive
	SetConcurrency(1)
	Start()
	End()

	SetConcurrency(2)
	Start()
	End()
}
