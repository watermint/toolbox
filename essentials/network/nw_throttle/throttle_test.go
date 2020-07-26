package nw_throttle

import "testing"

func TestThrottle(t *testing.T) {
	Throttle("test", "https://www.example.com/test", func() {
	})
}
