package dev

import (
	"github.com/watermint/toolbox/app86/app_recipe"
	"github.com/watermint/toolbox/app86/app_vo"
)

type Version struct {
}

func (*Version) Hidden() {
}

func (*Version) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*Version) Exec(k app_recipe.Kitchen) error {
	return nil
}
