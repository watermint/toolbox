package app_opt

import "testing"

func TestDefault(t *testing.T) {
	com := Default()
	if com.Concurrency < 1 {
		t.Error(com.Concurrency)
	}
}
