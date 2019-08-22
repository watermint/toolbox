package dev

import (
	"github.com/watermint/toolbox/infra/quality"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type Quality struct {
}

func (*Quality) Hidden() {
}

func (*Quality) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*Quality) Exec(k app_kitchen.Kitchen) error {
	quality.Suite(k.Control())
	return nil
}
