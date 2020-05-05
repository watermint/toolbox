package app_catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue"
)

var (
	current rc_catalogue.Catalogue = rc_catalogue.NewEmptyCatalogue()
)

func Current() rc_catalogue.Catalogue {
	return current
}

func SetCurrent(c rc_catalogue.Catalogue) {
	current = c
}
