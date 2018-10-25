package cmdlet

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"os"
	"path/filepath"
	"testing"
)

func CmdTest(t *testing.T, g Commandlet, args []string) {
	tokensFilePath, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	for filepath.Base(tokensFilePath) != "toolbox" {
		tokensFilePath = filepath.Dir(tokensFilePath)
	}
	ec := &infra.ExecContext{}
	if ec.StartupForTest(tokensFilePath) != nil {
		return
	}

	// Finish tests if tokens file not available
	if !ec.IsTokensAvailable() {
		seelog.Info("Skip tests")
		return
	}

	seelog.Infof("Testing Args[%v]", args)
	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	ec.PrepareFlags(f)
	g.Init(nil)
	g.FlagConfig(f)
	g.Exec(ec, args)

	eq := ErrorQueue()
	if len(eq) > 0 {
		t.Error(eq[0].UserMessage())
	}
}
