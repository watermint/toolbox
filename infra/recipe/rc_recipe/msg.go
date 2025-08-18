package rc_recipe

import (
	"github.com/watermint/toolbox/essentials/es_go/es_reflect"
)

func Key(r Recipe) string {
	return es_reflect.Key(r)
}
