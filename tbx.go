package main

import (
	"github.com/watermint/toolbox/catalogue"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_golog"
	"github.com/watermint/toolbox/infra/control/app_bootstrap"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/resources"
	"log"
	"os"
)

func run(args []string, forTest bool) {
	defer func() {
		if r := recover(); r != nil {
			if r == app_exit.Success {
				return
			} else {
				panic(r)
			}
		}
	}()

	app_exit.SetTestMode(forTest)

	bundle := resources.NewBundle()
	app_resource.SetBundle(bundle)
	app_catalogue.SetCurrent(catalogue.NewCatalogue())
	log.SetOutput(lgw_golog.NewLogWrapper(esl.Default()))

	b := app_bootstrap.NewBootstrap()
	b.Run(b.Parse(args[1:]...))
}

func main() {
	run(os.Args, false)
}
