package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/catalogue"
	"github.com/watermint/toolbox/essentials/go/es_resource"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_golog"
	"github.com/watermint/toolbox/infra/control/app_bootstrap"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"log"
	"os"
)

func run(args []string, forTest bool) {
	bundle := es_resource.New(
		rice.MustFindBox("resources/templates"),
		rice.MustFindBox("resources/messages"),
		rice.MustFindBox("resources/web"),
		rice.MustFindBox("resources/keys"),
		rice.MustFindBox("resources/images"),
		rice.MustFindBox("resources/data"),
	)
	app_resource.SetBundle(bundle)
	app_catalogue.SetCurrent(catalogue.NewCatalogue())
	log.SetOutput(lgw_golog.NewLogWrapper(esl.Default()))

	b := app_bootstrap.NewBootstrap()
	b.Run(b.Parse(args[1:]...))
}

func main() {
	run(os.Args, false)
}
