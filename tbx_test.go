package main

import (
	"os"
	"testing"

	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

func TestEcho(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		t.Skipped()
		return
	}
	run([]string{os.Args[0], "dev", "test", "echo", "-text", "Hey", "-debug"}, true)
}

func TestReplayBundle(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		t.Skipped()
		return
	}
	run([]string{os.Args[0], "dev", "replay", "remote"}, true)
}

func TestRun_PanicAndSuccess(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != nil && r != 0 { // app_exit.Success is 0
				t.Logf("Recovered panic: %v", r)
			}
		}
	}()
	// This should not panic (simulate success)
	run([]string{os.Args[0]}, true)
}
