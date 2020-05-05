package es_memory

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"testing"
)

func TestDumpStats(t *testing.T) {
	l := es_log.Default()
	DumpMemStats(l)
}
