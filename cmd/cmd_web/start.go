package cmd_web

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/poc/webui"
)

type CmdWebStart struct {
	*cmd.SimpleCommandlet
	Web webui.WebUI
}

func (z *CmdWebStart) Name() string {
	return "start"
}

func (z *CmdWebStart) Desc() string {
	return "Start web console"
}

func (CmdWebStart) Usage() string {
	return ""
}

func (z *CmdWebStart) FlagConfig(f *flag.FlagSet) {
	z.Web.FlagConfig(f)
}

func (z *CmdWebStart) Exec(args []string) {
	z.Web.Logger = z.Log()
	z.Web.Start()
}
