package dev

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type Quality struct {
}

func (z *Quality) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &app_vo.EmptyValueObject{}))
}

func (z *Quality) Hidden() {
}

func (z *Quality) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (z *Quality) Exec(k app_kitchen.Kitchen) error {
	quality.Suite(k.Control())
	return nil
}
