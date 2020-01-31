package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/gonutz/w32"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_run"
	"os"
)

func main() {
	bx := rice.MustFindBox("resources")
	web := rice.MustFindBox("web")

	if rb, found := app_run.FindRunBook(false); found {
		if app.IsWindows() {
			cw := w32.GetConsoleWindow()
			if cw != 0 {
				_, pid := w32.GetWindowThreadProcessId(cw)
				if w32.GetCurrentProcessId() == pid {
					w32.ShowWindowAsync(cw, w32.SW_HIDE)
				}
			}
		}
		rb.Exec(bx, web)
	} else {
		app_run.Run(os.Args[1:], bx, web)
	}
}
