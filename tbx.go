package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/atbx/app_run"
	"github.com/watermint/toolbox/cmd/cmd_root"
	"os"
)

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "experimental":
		bx := rice.MustFindBox("atbx/resources")
		app_run.Run(os.Args[2:], bx)

	default:
		bx := rice.MustFindBox("resources")
		ec, err := app.NewExecContext(bx)
		if err != nil {
			return
		}
		cmds := cmd_root.NewCommands()
		cmds.Exec(ec, os.Args)
	}
}
