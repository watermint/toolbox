package main

import (
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		return
	}
	run([]string{os.Args[0], "dev", "test", "echo", "-text", "Hey", "-debug"}, true)
}

func TestReplayBundle(t *testing.T) {
	run([]string{os.Args[0], "dev", "replay", "remote"}, true)
}
