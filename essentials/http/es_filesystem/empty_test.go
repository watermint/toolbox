package es_filesystem

import "testing"

func TestEmpty_Open(t *testing.T) {
	e := &Empty{}
	if f, err := e.Open("test"); err == nil || f != nil {
		t.Error(f, err)
	}
}
