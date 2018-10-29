package cmdlet

import (
	"flag"
	"github.com/watermint/toolbox/infra"
	"go.uber.org/zap"
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
		if tokensFilePath == "." {
			return
		}
	}
	ec := infra.NewExecContext()
	ec.TokenFilePath = tokensFilePath
	if err := ec.ApplyFlags(); err != nil {
		return
	}

	// Finish tests if tokens file not available
	if !ec.IsTokensAvailable() {
		ec.Log().Info("Skip tests")
		return
	}

	ec.Log().Info(
		"Testing",
		zap.Strings("args", args),
	)
	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	ec.PrepareFlags(f)
	g.Init(nil)
	g.FlagConfig(f)
	ec.ApplyFlags()
	defer ec.Shutdown()
	g.Setup(ec)
	g.Exec(args)

	eq := ErrorQueue()
	if len(eq) > 0 {
		t.Error(eq[0].UserMessage())
	}
}
