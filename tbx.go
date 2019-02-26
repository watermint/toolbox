package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/cmd/cmd_root"
	"os"
)

func main() {
	bx := rice.MustFindBox("")
	ec := app.NewExecContext(bx)
	cmds := cmd_root.NewCommands()
	cmds.Exec(ec, os.Args)
}
