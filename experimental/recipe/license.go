package recipe

import (
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type License struct {
}

func (*License) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*License) Exec(k app_recipe.Kitchen) error {
	return nil
}
