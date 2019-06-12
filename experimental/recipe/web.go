package recipe

import (
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type Web struct {
}

func (z *Web) Console() {
}

func (z *Web) Requirement() app_vo.ValueObject {
	panic("implement me")
}

func (z *Web) Exec(k app_recipe.Kitchen) error {
	panic("implement me")
}
