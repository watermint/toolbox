package nw_congestion

import (
	"testing"
	"time"
)

func TestCcImpl_IsSignificantWait(t *testing.T) {
	cc := NewControl().(*ccImpl)
	if x := cc.isSignificantWait(time.Now().Add(thresholdSignificantWait / 2)); x {
		t.Error(x)
	}
	if x := cc.isSignificantWait(time.Now().Add(thresholdSignificantWait * 2)); !x {
		t.Error(x)
	}
}
