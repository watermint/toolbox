package main

import (
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		return
	}
	run([]string{os.Args[0], "dev", "test", "echo", "-text", "Hey"}, true)
}

func TestReplayBundle(t *testing.T) {
	run([]string{os.Args[0], "dev", "replay", "bundle"}, true)
}
