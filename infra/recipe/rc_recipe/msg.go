package rc_recipe

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
)

func Key(r Recipe) string {
	return ut_reflect.Key(app.Pkg, r)
}
