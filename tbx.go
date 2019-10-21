package main

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_run"
	"os"
)

func main() {
	bx := rice.MustFindBox("resources")
	web := rice.MustFindBox("web")

	_ = app_run.Run(os.Args[1:], bx, web)
}
