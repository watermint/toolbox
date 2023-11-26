package ethread

import (
	"sync"
	"testing"
)

func TestCurrentRoutineName(t *testing.T) {
	wg := sync.WaitGroup{}

	var n1, n2, n3 string
	n1 = CurrentRoutineName()
	wg.Add(2)
	go func() {
		n2 = CurrentRoutineName()
		wg.Done()
	}()
	go func() {
		n3 = CurrentRoutineName()
		wg.Done()
	}()
	wg.Wait()

	if n1 == n2 || n2 == n3 {
		t.Error(n1, n2, n3)
	}
}
