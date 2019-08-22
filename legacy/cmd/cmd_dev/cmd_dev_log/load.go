package cmd_dev_log

import (
	"flag"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"go.uber.org/zap"
)

type CmdDevLogLoad struct {
	*cmd2.SimpleCommandlet
}

func (z *CmdDevLogLoad) Name() string {
	return "load"
}

func (z *CmdDevLogLoad) Desc() string {
	return "cmd.dev.log.load.desc"
}

func (z *CmdDevLogLoad) Usage() func(cmd2.CommandUsage) {
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
