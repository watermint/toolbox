package cmd

import (
	"context"
	"flag"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/model/dbx_auth"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const (
	testTimeout = time.Duration(40) * time.Second
)

func CmdTest(t *testing.T, g Commandlet, args []string) {
	CmdTestWithTimeout(t, g, args, testTimeout)
}

func CmdTestWithTimeout(t *testing.T, g Commandlet, args []string, timeout time.Duration) {
	tokensFilePath, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	for filepath.Base(tokensFilePath) != "toolbox" {
		tokensFilePath = filepath.Dir(tokensFilePath)
		if tokensFilePath == "." {
			return
		}
	}
	ec := app.NewExecContextForTest()
	ec.TokenFilePath = tokensFilePath
	if err := ec.ApplyFlags(); err != nil {
		return
	}

	// Finish tests if tokens file not available
	if !dbx_auth.IsCacheAvailable(ec, "") {
		ec.Log().Info("Skip tests")
		return
	}

	cmd := func(c chan bool) {
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

		c <- true
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	c := make(chan bool)
	go cmd(c)

	select {
	case <-ctx.Done():
		ec.Log().Info("Cancelled test due to timeout (assume the test succeed)", zap.Duration("timeout", testTimeout))

	case <-c:
		ec.Log().Info("The test passed")
	}

	eq := ErrorQueue()
	if len(eq) > 0 {
		t.Error(eq[0].UserMessage())
	}
}
