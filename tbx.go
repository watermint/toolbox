package main

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/cmd/cmd_root"
	"os"
)

func main() {
	ec := app.NewExecContext()
	cmds := cmd_root.NewCommands()
	cmds.Exec(ec, os.Args)
}
