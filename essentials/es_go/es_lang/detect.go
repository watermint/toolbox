package es_lang

import (
	"github.com/watermint/toolbox/essentials/i18n/es_locale"
)

// Detect & select language in select
func Detect(supported []Lang) Lang {
	bcp47 := es_locale.CurrentLocale().String()

	return Select(bcp47, supported)
}
