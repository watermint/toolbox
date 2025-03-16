package app_catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue_impl"
)

var (
	current = rc_catalogue_impl.NewEmptyCatalogue()
)

func Current() rc_catalogue.Catalogue {
	return current
}

func SetCurrent(c rc_catalogue.Catalogue) {
	current = c
}
