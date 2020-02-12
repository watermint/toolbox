package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_run"
	"github.com/watermint/toolbox/infra/control/app_workflow"
	"github.com/watermint/toolbox/infra/util/ut_ui"
	"os"
	"strings"
)

func runRunBook(b app_run.Bootstrap, path string, args []string) {
	_, com := b.ParseCommon(args, true)
	r, _ := b.Parse("job", "run", "-runbook-path", path)
	b.Run(r, com)
}

func run(args []string, forTest bool) {
	bx := rice.MustFindBox("resources")
	web := rice.MustFindBox("web")
	b := app_run.NewBootstrap(bx, web)

	switch {
	case len(args) <= 1:
		if path, _, found := app_workflow.DefaultRunBook(forTest); found {
			ut_ui.HideConsole()
			runRunBook(b, path, []string{})
		} else {
			b.Run(b.Parse(args[1:]...))
		}

	case strings.HasSuffix(strings.ToLower(args[1]), ".runbook"):
		path := args[1]
		if _, found := app_workflow.NewRunBook(path); found {
			runRunBook(b, path, args[2:])
		} else {
			fmt.Errorf("Unable to execute runbook: %s\n", path)
			os.Exit(app_control.FatalStartup)
		}

	default:
		b.Run(b.Parse(args[1:]...))
	}
}

func main() {
	run(os.Args, false)
}
