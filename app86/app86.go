package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app86/app_run"
	"os"
)

func main() {
	bx := rice.MustFindBox("resources")
	app_run.Run(os.Args[1:], bx)
}
