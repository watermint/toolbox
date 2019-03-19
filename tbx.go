package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/cmd/cmd_root"
	"os"
)

func main() {
	bx := rice.MustFindBox("resources")
	ec, err := app.NewExecContext(bx)
	if err != nil {
		return
	}
	cmds := cmd_root.NewCommands()
	cmds.Exec(ec, os.Args)
}
