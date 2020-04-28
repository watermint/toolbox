package rc_recipe

import (
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
)

func Key(r Recipe) string {
	return es_reflect.Key(app.Pkg, r)
}
