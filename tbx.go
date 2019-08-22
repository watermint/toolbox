package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_run"
	app2 "github.com/watermint/toolbox/legacy/app"
	"github.com/watermint/toolbox/legacy/cmd/cmd_root"
	"os"
)

func main() {
	switch {
	case len(os.Args) > 1 && os.Args[1] == "experimental":
		bx := rice.MustFindBox("resources")
		web := rice.MustFindBox("web")

		app_run.Run(os.Args[2:], bx, web)

	default:
		bx := rice.MustFindBox("legacy/resources")
		ec, err := app2.NewExecContext(bx)
		if err != nil {
			return
		}
		cmds := cmd_root.NewCommands()
		cmds.Exec(ec, os.Args)
	}
}
