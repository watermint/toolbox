package ut_memory

import (
	"github.com/watermint/toolbox/infra/control/app_log"
	"testing"
)

func TestDumpStats(t *testing.T) {
	l := app_log.NewConsoleLogger(true)
	DumpStats(l)
}
