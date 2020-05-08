package main

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/catalogue"
	"github.com/watermint/toolbox/essentials/go/es_resource"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_golog"
	"github.com/watermint/toolbox/essentials/terminal/es_window"
	"github.com/watermint/toolbox/infra/control/app_bootstrap"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/control/app_workflow"
	"log"
	"os"
	"strings"
)

func runRunBook(b app_bootstrap.Bootstrap, path string, args []string) {
	_, com := b.ParseCommon(args, true)
	r, _ := b.Parse("job", "run", "-runbook-path", path)
	b.Run(r, com)
}

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

	switch {
	case len(args) <= 1:
		if path, _, found := app_workflow.DefaultRunBook(forTest); found {
			es_window.HideConsole()
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
			app_exit.Abort(app_exit.FatalStartup)
		}

	default:
		b.Run(b.Parse(args[1:]...))
	}
}

func main() {
	run(os.Args, false)
}
