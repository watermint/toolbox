package test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_runtime"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type Resources struct {
}

func (z *Resources) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &app_vo.EmptyValueObject{}))
}

func (z *Resources) Hidden() {
}

func (z *Resources) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (z *Resources) Exec(k app_kitchen.Kitchen) error {
	qt_runtime.Suite(k.Control())
	return nil
}
