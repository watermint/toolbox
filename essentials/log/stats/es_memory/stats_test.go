package es_memory

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"testing"
)

func TestDumpStats(t *testing.T) {
	l := esl.Default()
	DumpMemStats(l)
}
