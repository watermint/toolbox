package main

import (
	"github.com/watermint/toolbox/cmdlet/cmd_root"
	"github.com/watermint/toolbox/infra"
	"os"
)

func main() {
	ec := &infra.ExecContext{}
	cmds := cmd_root.NewCommands()
	cmds.Exec(ec, os.Args)
}
