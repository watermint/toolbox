package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app86/app_run"
	"github.com/watermint/toolbox/app86/app_ui"
	"os"
)

func main() {
	bx := rice.MustFindBox("resources")
	mc := app_run.NewContainer(bx)
	ui := app_ui.NewConsole(mc)
	c := app_run.Catalogue()

	g, _, _, err := c.Select(os.Args[1:])

	switch {
	case err != nil:
		if g != nil {
			g.PrintUsage(ui)
		} else {
			c.PrintUsage(ui)
		}

	case g != nil:
		g.PrintUsage(ui)
	}
}
