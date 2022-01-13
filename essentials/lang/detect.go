package lang

import (
	"github.com/watermint/essentials/ei18n/elocale"
)

// Detect & select language in select
func Detect(supported []Lang) Lang {
	bcp47 := elocale.CurrentLocale().String()

	return Select(bcp47, supported)
}
