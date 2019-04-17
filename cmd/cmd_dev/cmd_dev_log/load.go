package cmd_dev_log

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"go.uber.org/zap"
)

type CmdDevLogLoad struct {
	*cmd.SimpleCommandlet
}

func (z *CmdDevLogLoad) Name() string {
	return "load"
}

func (z *CmdDevLogLoad) Desc() string {
	return "cmd.dev.log.load.desc"
}

func (z *CmdDevLogLoad) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdDevLogLoad) FlagConfig(f *flag.FlagSet) {
}

func (z *CmdDevLogLoad) Exec(args []string) {
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			z.Log().Debug("Log load test", zap.Int("i", i), zap.Int("j", j))
		}
	}
}
