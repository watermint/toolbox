package cmd_web

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/webui"
)

type CmdWebStart struct {
	*cmd.SimpleCommandlet
}

func (c *CmdWebStart) Name() string {
	return "start"
}

func (c *CmdWebStart) Desc() string {
	return "Start web console"
}

func (CmdWebStart) Usage() string {
	return ""
}

func (c *CmdWebStart) FlagConfig(f *flag.FlagSet) {
}

func (c *CmdWebStart) Exec(args []string) {
	webui.Start()
}
