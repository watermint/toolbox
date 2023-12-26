package rc_recipe

import (
	"github.com/watermint/toolbox/essentials/go/es_reflect"
)

func Key(r Recipe) string {
	return es_reflect.Key(r)
}
