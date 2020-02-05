package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_run"
	"github.com/watermint/toolbox/infra/util/ut_ui"
	"os"
	"strings"
)

func main() {
	bx := rice.MustFindBox("resources")
	web := rice.MustFindBox("web")

	switch {
	case len(os.Args) <= 1:
		if rb, found := app_run.DefaultRunBook(false); found {
			ut_ui.HideConsole()
			rb.Exec(bx, web)
		} else {
			app_run.Run(os.Args[1:], bx, web)
		}

	case strings.HasSuffix(strings.ToLower(os.Args[1]), ".runbook"):
		if rb, found := app_run.NewRunBook(os.Args[1]); found {
			ut_ui.HideConsole()
			rb.Exec(bx, web)
		} else {
			fmt.Errorf("Unable to execute runbook\n")
			os.Exit(app_control.FatalStartup)
		}

	default:
		app_run.Run(os.Args[1:], bx, web)
	}
}
