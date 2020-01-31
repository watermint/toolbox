package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_run"
	"github.com/watermint/toolbox/infra/util/ut_ui"
	"os"
)

func main() {
	bx := rice.MustFindBox("resources")
	web := rice.MustFindBox("web")

	if rb, found := app_run.FindRunBook(false); found {
		rb.Exec(bx, web)
	} else {
		app_run.Run(os.Args[1:], bx, web)
	}
}
