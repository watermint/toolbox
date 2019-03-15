package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_doc"
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
	if len(os.Args) > 1 && os.Args[1] == "-markdown" {
		d := app_doc.CmdDoc{ExecContext: ec}
		d.Init()
		d.Parse(cmds.RootCommand())
		d.Markdown()
		return
	}
	cmds.Exec(ec, os.Args)
}
