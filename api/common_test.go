package api

import (
	"testing"
	"time"
)

func TestRebaseTimeForAPI(t *testing.T) {
	jst, err := time.LoadLocation("Japan")
	if err != nil {
		t.Error(err)
	}
	nowUtc := time.Now()
	nowJst := nowUtc.In(jst)
	nowRoundedUtc := nowUtc.Round(time.Second)

	if !RebaseTimeForAPI(nowJst).Equal(nowRoundedUtc) {
		t.Error("Invalid state")
	}
}
