package es_lang

import (
	"github.com/watermint/toolbox/essentials/islet/ei18n/elocale"
)

// Detect & select language in select
func Detect(supported []Lang) Lang {
	bcp47 := elocale.CurrentLocale().String()

	return Select(bcp47, supported)
}
