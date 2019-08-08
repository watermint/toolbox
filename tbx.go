package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/cmd/cmd_root"
	"github.com/watermint/toolbox/experimental/app_run"
	"os"
)

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "experimental":
		bx := rice.MustFindBox("experimental/resources")
		web := rice.MustFindBox("experimental/web")

		app_run.Run(os.Args[2:], bx, web)

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
